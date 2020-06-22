from cStringIO import StringIO
import json
import re
import os
import sys
#
# To install azure requirements:
#
#     pip install azure-mgmt-commerce azure-mgmt-compute
#
from azure.common.credentials import ServicePrincipalCredentials
from azure.mgmt.compute import ComputeManagementClient
from azure.mgmt.commerce import UsageManagementClient

from util import (
    parse_args,
    upload,
    write_go,
    dumpjson,
)

LOCATIONS = {
    'eastasia': 'East Asia',
    'southeastasia': 'Southeast Asia',
    'centralus': 'Central US',
    'eastus': 'East US',
    'eastus2': 'East US 2',
    'westus': 'West US',
    'northcentralus': 'North Central US',
    'southcentralus': 'South Central US',
    'northeurope': 'North Europe',
    'westeurope': 'West Europe',
    'japanwest': 'Japan West',
    'japaneast': 'Japan East',
    'brazilsouth': 'Brazil South',
    'australiaeast': 'Australia East',
    'australiasoutheast': 'Australia Southeast',
    'southindia': 'South India',
    'centralindia': 'Central India',
    'westindia': 'West India',
    'canadacentral': 'Canada Central',
    'canadaeast': 'Canada East',
    'uksouth': 'UK South',
    'ukwest': 'UK West',
    'westcentralus': 'West Central US',
    'westus2': 'West US 2',
    'koreacentral': 'Korea Central',
    'koreasouth': 'Korea South',
    'francecentral': 'France Central',
    'francesouth': 'France South',
    'australiacentral': 'Australia Central',
    'australiacentral2': 'Australia Central 2',
    ###
    'AP East': 'eastasia',
    'AP Southeast': 'southeastasia',
    'AU Central': 'australiacentral',
    'AU Central 2': 'australiacentral2',
    'AU East': 'australiaeast',
    'AU Southeast': 'australiasoutheast',
    'BR South': 'brazilsouth',
    'CA Central': 'canadacentral',
    'CA East': 'canadaeast',
    'EU North': 'northeurope',
    'EU West': 'westeurope',
    'FR Central': 'francecentral',
    'FR South': 'francesouth',
    'IN Central': 'centralindia',
    'IN South': 'southindia',
    'IN West': 'westindia',
    'JA East': 'japaneast',
    'JA West': 'japanwest',
    'KR Central': 'koreacentral',
    'KR South': 'koreasouth',
    'UK South': 'uksouth',
    'UK West': 'ukwest',
    'US Central': 'centralus',
    'US East': 'eastus',
    'US East 2': 'eastus2',
    'US North Central': 'northcentralus',
    'US South Central': 'southcentralus',
    'US West': 'westus',
    'US West 2': 'westus2',
    'US West Central': 'westcentralus',
}

#
# The SKU API does not provide any information on baseline CPU performance for
# burstable instances. This data is from:
#
# https://docs.microsoft.com/en-us/azure/virtual-machines/windows/b-series-burstable
#
BURSTABLE_INSTANCES = {
    'Standard_B1s': 0.1,
    'Standard_B1ms': 0.2,
    'Standard_B2s': 0.4,
    'Standard_B2ms': 0.6,
    'Standard_B4ms': 0.9,
    'Standard_B8ms': 1.35,
}


def get_credentials():
    subscription_id = os.environ['AZURE_SUBSCRIPTION_ID']
    credentials = ServicePrincipalCredentials(
        client_id=os.environ['AZURE_CLIENT_ID'],
        secret=os.environ['AZURE_CLIENT_SECRET'],
        tenant=os.environ['AZURE_TENANT_ID']
    )
    return credentials, subscription_id


def get_pricing():
    print("Getting pricing for azure instances")
    credentials, subscription_id = get_credentials()
    commerce_client = UsageManagementClient(credentials, subscription_id)
    rate = commerce_client.rate_card.get(
        "OfferDurableId eq 'MS-AZR-0003P' and " +  # Pay-as-you-go
        "Currency eq 'USD' and " +
        "Locale eq 'en-US' and " +
        "RegionInfo eq 'US'"
    )
    rates = {}
    for r in rate.meters:
        if r.meter_category != "Virtual Machines":
            continue
        if r.meter_region not in LOCATIONS:
            continue
        name = r.meter_name
        if ('Low Priority' in name or
            'Compute Hours' in name or
            'Windows' in r.meter_sub_category):
            continue
        name = name.replace(' ', '_', -1)
        names = [name]
        if '/' in name:
            names = name.split('/')
        loc = LOCATIONS[r.meter_region]
        for n in names:
            d = rates.get(n, {})
            d[loc] = r.meter_rates['0']
            rates[n] = d
    compute_client = ComputeManagementClient(credentials, subscription_id)
    resource_skus = compute_client.resource_skus.list()
    result = {}
    for rs in resource_skus:
        # E.g. name: Standard_DS12-1_v2 size: DS12-1_v2
        if not rs.name.startswith('Standard_'):
            continue
        if rs.size is None:
            continue
        if not rs.capabilities:
            continue
        if not rs.location_info:
            continue
        is_burstable = False
        baseline = 1.0
        memory = 0
        vcpus = 0
        m = re.search(r'([a-zA-Z0-9]+)(-[0-9]+)?(_v[0-9]+)?(_Promo)?', rs.size)
        if not m:
            continue
        name = m.group(1)
        if m.group(3) is not None:
            name = name + m.group(3)
        for c in rs.capabilities:
            if c.name == 'vCPUs':
                vcpus = int(c.value)
            if c.name == 'vCPUsAvailable' and vcpus == 0:
                vcpus = int(c.value)
            if c.name == 'MemoryGB':
                memory = float(c.value)
        if rs.name in BURSTABLE_INSTANCES:
            is_burstable = True
            baseline = BURSTABLE_INSTANCES[rs.name]
        for l in rs.location_info:
            loc = l.location.lower()
            if name not in rates:
                print >> sys.stderr, \
                        "Warning: can't find rate for", name
                continue
            if loc not in rates[name]:
                print >> sys.stderr, \
                        "Warning: can't find rate for", name, "in", loc
                continue
            if loc not in LOCATIONS:
                print >> sys.stderr, \
                        "Warning: unknown location", loc
                continue
            location = LOCATIONS[loc]
            if location not in result:
                result[location] = []
            result[location].append({
                'baseline': baseline,
                'burstable': is_burstable,
                'generation': 'current',
                'cpu': vcpus,
                'memory': memory,
                'instanceType': rs.name,
                'price': rates[name][loc],
            })
    return result


def update_instance_data(args):
    azure_pricing = get_pricing()
    jsonfp = dumpjson(azure_pricing)
    if args.upload:
        upload('azure_instance_data.json', jsonfp)
    else:
        write_go('azure', jsonfp, dumpjson({}))


def update_network_data(args):
    pass


def update_storage_data(args):
    pass

if __name__ == '__main__':
    args = parse_args()
    update_instance_data(args)
