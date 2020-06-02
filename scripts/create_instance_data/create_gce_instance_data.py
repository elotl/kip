import argparse
from pprint import pprint

import googleapiclient.discovery

from util import (
    write_go,
    dumpjson,
)

nano = 10**-9
custom_baselines = {
    'f1-micro': 0.2,  # 0.60GB ram
    'g1-small': 0.5,  # 1.70GB ram
}


def gce_to_kip_memory(memory_mb):
    gb = memory_mb / 1024.0
    return round(gb, 2)


def compute_machine_price(pricing, family, cpus, memory):
    family_prices = pricing[family]
    if type(family_prices) == float:
        return family_prices
    cpu_price = family_prices['cpu']
    ram_price = family_prices['ram']
    return cpus * cpu_price + memory * ram_price


def make_instance_data(machine, pricing, gpus):
    name = machine['name']
    family = name.split('-')[0]
    burstable = machine['isSharedCpu']
    cpus = machine['guestCpus']
    baseline = cpus
    if name in custom_baselines:
        baseline = custom_baselines[name]
    memory = gce_to_kip_memory(machine['memoryMb'])
    price = compute_machine_price(pricing, family, cpus, memory)
    maxGPUs = 0
    for n in gpus.values():
        if n > maxGPUs:
            maxGPUs = n
    return {
        "baseline": baseline,
        "generation": "current",
        "price": price,
        "memory": memory,
        "instanceType": name,
        "burstable": burstable,
        "gpu": maxGPUs,
        "supportedGPUTypes": gpus,
        "cpu": cpus,
    }


def get_available_gpus(zone, machine, gpus):
    '''
    The number of GPUs that can be attached to instance depends on the instance
    type and the zone.
    '''
    name = machine['name']
    family = name.split('-')[0]
    if family != "n1":
        return {}
    if zone not in gpus:
        return {}
    return gpus[zone]


def cleanup_single_region_prices(region_prices):
    '''
    for now, this just copies prices from the m1 instance type over to
    the m2 instance type.
    '''
    if 'm1' in region_prices:
        m1_vals = region_prices['m1']
        region_prices['m2'] = m1_vals
    return region_prices


def cleanup_prices(prices):
    '''
    The billing API isn't returning prices for us-west3 and us-west4
    but n1 and e families are available there.  Add those regions
    to our price list by copying prices from us-west2 since those
    prices appear to be the closest.
    '''
    missing_west_prices = {}
    for family in ('f1', 'g1', 'e2', 'n1'):
        missing_west_prices[family] = prices['us-west2'][family]
    if not prices['us-west3']:
        prices['us-west3'] = missing_west_prices
    if not prices['us-west4']:
        prices['us-west4'] = missing_west_prices
    return prices


def create_pricing_map(regions, skus):
    prices_by_region = {}
    for region in regions:
        region_prices = {}
        for sku in skus:
            if not sku_we_want(sku, region):
                continue
            desc = sku['description'].lower()
            group = sku['category']['resourceGroup'].lower()
            if 'custom' in desc:
                continue
            family = get_instance_family(desc, group)
            if family is None:
                continue
            price = get_price(sku)
            if price <= 0:
                # if the price doesn't exist, then we assume that
                # instance type isn't supported. E.g N2D Ram in Osaka
                continue
            if is_all_in_one_pricing(group):
                region_prices[family] = price
            else:
                resource_type = get_resource_type(desc, group)
                if family not in region_prices:
                    region_prices[family] = {}
                region_prices[family][resource_type] = price
        region_prices = cleanup_single_region_prices(region_prices)
        prices_by_region[region] = region_prices
    prices_by_region = cleanup_prices(prices_by_region)
    return prices_by_region


def is_all_in_one_pricing(group):
    return group in ('g1small', 'f1micro')


def get_instance_family(desc, group):
    if group == 'g1small':
        return 'g1'
    elif group == 'f1micro':
        return 'f1'
    if 'compute optimized' in desc:
        return 'c2'
    elif 'e2' in desc:
        return 'e2'
    elif 'memory' in desc and 'optimized' in desc:
        return 'm1'
    elif 'n1 ' in desc:
        return 'n1'
    elif 'n2 ' in desc:
        return 'n2'
    elif 'n2d ' in desc:
        return 'n2d'


def get_resource_type(desc, group):
    if group == 'cpu':
        return 'cpu'
    elif group == 'ram':
        return 'ram'
    elif 'core' in desc:
        return 'cpu'
    elif 'ram' in desc:
        return 'ram'
    return None


def get_price(sku):
    pricing_info = sku['pricingInfo']
    assert len(pricing_info) == 1
    tiered_rates = pricing_info[0]['pricingExpression']['tieredRates']
    assert len(tiered_rates) > 0
    max_price = 0.0
    # there can be multiple prices listed, so far I've only seen the
    # additional prices be $0.0 For now, we'll take the highest priced
    # entry.
    for rate in tiered_rates:
        unit_price = rate['unitPrice']
        if 'currencyCode' not in unit_price:
            continue
        if unit_price['currencyCode'] != 'USD':
            msg = "Non USD price in SKUs: {} - {}".format(
                rate['currencyCode'], sku['name'])
            raise Exception(msg)
        units = int(unit_price['units'])
        nanos = int(unit_price['nanos'])
        price = units + nanos * nano
        if price > max_price:
            max_price = price
    return round(max_price, 8)


def get_all_regions(client, project):
    regions = []
    request = client.regions().list(project=project)
    while request is not None:
        response = request.execute()
        for region in response['items']:
            regions.append(region['name'])
            request = client.regions().list_next(
                previous_request=request, previous_response=response)
    return regions


def list_all_machine_types(client, project, zones):
    machines_by_zone = {}
    for zone in zones:
        print('getting machines in zone %s' % zone)
        request = client.machineTypes().list(project=project, zone=zone)
        machine_types = []
        while request is not None:
            response = request.execute()
            for machine_type in response['items']:
                machine_types.append(machine_type)
            request = client.machineTypes().list_next(
                previous_request=request, previous_response=response)
        machines_by_zone[zone] = machine_types
    return machines_by_zone


def sku_we_want(sku, region):
    if sku['category']['resourceFamily'] != 'Compute':
        return False
    if sku['category']['usageType'] != 'OnDemand':
        return False
    if ('global' not in sku['serviceRegions'] and
        region not in sku['serviceRegions']):
        return False
    if sku['category']['resourceGroup'] in (
            'PdSnapshotEgress', 'SecurityPolicy', 'GPU'):
        return False
    return True


def zone_to_region(zone):
    return zone[:-2]


def get_all_zones(client, project):
    print('getting all zones')
    zones = []
    request = client.zones().list(project=project)
    while request is not None:
        response = request.execute()
        for region in response['items']:
            zones.append(region['name'])
            request = client.zones().list_next(
                previous_request=request, previous_response=response)
    return zones


def get_supported_gpus(client, project, zones):
    gpus = {}
    for zone in zones:
        print('getting supported gpus in zone %s' % zone)
        gpus[zone] = {}
        request = client.acceleratorTypes().list(project=project, zone=zone)
        while request is not None:
            response = request.execute()
            if not 'items' in response:
                request = client.acceleratorTypes().list_next(
                    previous_request=request, previous_response=response)
                continue
            for gpu in response['items']:
                gpuType = gpu['name']
                gpus[zone][gpuType] = gpu['maximumCardsPerInstance']
                request = client.acceleratorTypes().list_next(
                    previous_request=request, previous_response=response)
    return gpus


def get_service(client, display_name):
    request = client.services().list()
    while request is not None:
        resp = request.execute()
        for svc in resp['services']:
            if svc['displayName'] == display_name:
                pprint(svc)
                return svc
        if resp.get('nextPageToken'):
            request = client.services().list(
                pageToken=resp.get('nextPageToken'))


def get_skus(client, parent):
    request = client.services().skus().list(parent=parent)
    skus = []
    while request is not None:
        response = request.execute()
        for sku in response['skus']:
            skus.append(sku)
        request = client.services().skus().list_next(
            previous_request=request, previous_response=response)
    return skus


def get_compute_skus():
    client = googleapiclient.discovery.build('cloudbilling', 'v1')
    print('getting services')
    ce_svc = get_service(client, 'Compute Engine')
    print('getting skus')
    skus = get_skus(client, ce_svc['name'])
    return skus


def get_instance_data(project):
    skus = get_compute_skus()
    client = googleapiclient.discovery.build('compute', 'v1')
    zones = get_all_zones(client, project)
    supportedGPUTypes = get_supported_gpus(client, project, zones)
    regions = set([zone_to_region(z) for z in zones])
    machines = list_all_machine_types(client, project, zones)
    pricing_by_region = create_pricing_map(regions, skus)
    instance_data = {}
    for zone, machines in machines.iteritems():
        region = zone_to_region(zone)
        zone_data = []
        pricing = pricing_by_region[region]
        for machine in machines:
            gpus = get_available_gpus(zone, machine, supportedGPUTypes)
            zone_data.append(make_instance_data(machine, pricing, gpus))
        instance_data[zone] = zone_data
    return instance_data


if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument('--project_id', help='Google Cloud project ID.',
                        default='elotl-kip')
    args = parser.parse_args()
    instance_data = get_instance_data(args.project_id)
    jsonfp = dumpjson(instance_data)
    write_go('gce', jsonfp)
