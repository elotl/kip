from cStringIO import StringIO
import json
import argparse
import os

import boto3


def get_kipdir():
    gopath = os.getenv("GOPATH")
    if not gopath:
        raise Exception("GOPATH must be set")
    kipdir = os.path.join(gopath, "src/github.com/elotl/kip")
    return kipdir


def parse_args():
    parser = argparse.ArgumentParser(
        usage='''Examples
# create instance data in .go source file
python create_instance_data.py

# create instance data locally and upload to s3
python create_instance_data.py --s3
        ''')
    parser.add_argument('--upload', action="store_true", default=False)
    args = parser.parse_args()
    return args


def upload(key, jsonfp):
    '''cloudname should be one of aws, azure or gce'''
    print("uploading", key)
    s3 = boto3.client('s3')
    bucket_name = 'elotl-cloud-data'
    s3.upload_fileobj(jsonfp, bucket_name, key)
    s3.put_object_acl(ACL='public-read', Bucket=bucket_name, Key=key)


def write_go(cloudname, jsonfp, custom_jsonfp):
    '''cloudname should be one of aws, azure or gce'''
    print("Writing go files")
    kipdir = get_kipdir()
    filepath = "pkg/util/instanceselector/{}_instance_data.go".format(
        cloudname)
    outfile = os.path.join(kipdir, filepath)
    with open(outfile, "w") as fp:
        headerpath = os.path.join(kipdir, "scripts/boilerplate.go.txt")
        with open(headerpath) as headerfp:
            header = headerfp.read()
            fp.write(header)
        fp.write("""package instanceselector

const {}InstanceJson = `
""".format(cloudname))
        fp.write(jsonfp.getvalue())
        fp.write("\n`\n\n")
        fp.write("""const {}CustomInstanceJson = `
""".format(cloudname))
        fp.write(custom_jsonfp.getvalue())
        fp.write("\n`\n")


def dumpjson(data):
    '''dumps data to a json file-like object'''
    jsonfp = StringIO()
    json.dump(data, jsonfp, indent=4, separators=(',', ': '))
    jsonfp.seek(0)
    return jsonfp
