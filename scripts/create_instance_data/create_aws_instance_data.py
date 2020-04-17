from cStringIO import StringIO
import json
from collections import defaultdict
import os
from pprint import pprint

import awspricing
import boto3

from util import (
    get_milpadir,
    parse_args,
    upload,
    write_go,
    dumpjson,
)
#
# This processes the data found at https://ec2instances.info/. The raw data
# can be re-downloaded via:
#
# curl https://raw.githubusercontent.com/powdahound/ec2instances.info/master/www/instances.json > aws_instances.json
#

regions = [
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ca-central-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2",
]

region_name_mapping = {
    'ap-northeast-1': 'Asia Pacific (Tokyo)',
    'ap-northeast-2': 'Asia Pacific (Seoul)',
    'ap-northeast-3': 'Asia Pacific (Osaka-Local)',
    'ap-south-1': 'Asia Pacific (Mumbai)',
    'ap-southeast-1': 'Asia Pacific (Singapore)',
    'ap-southeast-2': 'Asia Pacific (Sydney)',
    'ca-central-1': 'Canada (Central)',
    'cn-north-1': 'China (Beijing)',
    'cn-northwest-1': 'China (Ningxia)',
    'eu-central-1': 'EU (Frankfurt)',
    'eu-north-1': 'EU (Stockholm)',
    'eu-west-1': 'EU (Ireland)',
    'eu-west-2': 'EU (London)',
    'eu-west-3': 'EU (Paris)',
    'sa-east-1': 'South America (Sao Paulo)',
    'us-east-1': 'US East (N. Virginia)',
    'us-east-2': 'US East (Ohio)',
    'us-gov-east-1': 'AWS GovCloud (US-East)',
    'us-gov-west-1': 'AWS GovCloud (US)',
    'us-west-1': 'US West (N. California)',
    'us-west-2': 'US West (Oregon)',
}

HOURS_IN_MONTH = 30 * 24

def get_instance_data(raw_data):
    print('computing instance data')
    ec2_offer = awspricing.offer('AmazonEC2')
    simple_instance_data = defaultdict(list)
    for instance in raw_data:
        baseline = instance['vCPU']
        if instance['base_performance'] is not None:
            baseline = instance['base_performance']
        if baseline == "N/A":
            continue
        instance_info = {
            'instanceType': instance['instance_type'],
            'gpu': instance['GPU'],
            'memory': instance['memory'],
            'cpu': instance['vCPU'],
            'burstable': instance["burst_minutes"] is not None,
            'baseline': baseline,
            'generation': instance['generation'],
        }
        if instance["generation"] != 'current':
            continue
        for region in regions:
            try:
                price = ec2_offer.ondemand_hourly(
                  instance_type=instance['instance_type'],
                  operating_system='Linux',
                  region=region,
                )
                instance_info['price'] = float(price)
                print instance_info
                simple_instance_data[region].append(instance_info)
            except ValueError:
                continue
    return simple_instance_data


def get_raw_instance_data():
    print("reading data")
    milpadir = get_milpadir()
    filename = os.path.join(
        milpadir, "scripts/create_instance_data/aws_instances.json")
    raw_data = json.load(open(filename))
    return raw_data


def make_filter(*args):
    filters = []
    for field, value in args:
        filters.append({
            "Field": field,
            "Value": value,
            "Type": "TERM_MATCH",
        })
    return filters


def get_price_from_product_response(response):
    pricelist_json = response['PriceList'][0]
    pricelist = json.loads(pricelist_json)
    return pricelist['terms']['OnDemand'].values()[0]['priceDimensions'].values()[0]['pricePerUnit']['USD']


def get_storage_pricing():
    storage_by_region = {}
    pricing = boto3.client('pricing')
    for region in regions:
        region_name = region_name_mapping[region]
        filters = make_filter(
            ("volumeType", "General Purpose"),
            ("location", region_name))
        response = pricing.get_products(
            ServiceCode='AmazonEC2', Filters=filters)
        try:
            price = get_price_from_product_response(response)
            storage_by_region[region] = {
                # this name/key must match the product name in our
                # usage records
                'gp2': {'price': float(price) / HOURS_IN_MONTH},
            }
        except (IndexError, KeyError):
            print("No ebs pricing available in", region)
            continue
    pprint(storage_by_region)
    return storage_by_region


def get_elb_pricing():
    elb_by_region = {}
    pricing = boto3.client('pricing')
    for region in regions:
        region_name = region_name_mapping[region]
        filters = make_filter(
            ("productFamily", "Load Balancer"),
            ("groupDescription", "Standard Elastic Load Balancer"),
            ("location", region_name))
        response = pricing.get_products(
            ServiceCode='AmazonEC2', Filters=filters)
        try:
            price = get_price_from_product_response(response)
            elb_by_region[region] = {
                # this name/key must match the product name in our
                # usage records
                'ELB-Classic': {'price': float(price)},
            }
        except (IndexError, KeyError):
            print("No ELB pricing available in", region)
            continue
    pprint(elb_by_region)
    return elb_by_region


def update_instance_data(args):
    raw_data = get_raw_instance_data()
    simple_instance_data = get_instance_data(raw_data)
    jsonfp = dumpjson(simple_instance_data)
    if args.upload:
        upload('aws_instance_data.json', jsonfp)
    else:
        write_go('aws', jsonfp)


def update_network_data(args):
    elb_pricing = get_elb_pricing()
    jsonfp = dumpjson(elb_pricing)
    if args.upload:
        upload('aws_network_data.json', jsonfp)


def update_storage_data(args):
    pricing = get_storage_pricing()
    jsonfp = dumpjson(pricing)
    if args.upload:
        upload('aws_storage_data.json', jsonfp)


if __name__ == '__main__':
    args = parse_args()
    update_instance_data(args)
    update_network_data(args)
    update_storage_data(args)
