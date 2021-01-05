# Pricing updater

This tool gets pricing data for given region and provider from cloudinfo, converts it to format that KIP understands and creates ConfigMap (or dumps to json), which KIP can use.


## Supported providers
[x] AWS
[x] Azure


Warning: GCE is not supported, because cloudinfo doesn't have data about custom instances.

## Data format:
```json
{
    "<region>": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "spotPrice": 0.002,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        ...
    ]
}
```
You may also check examples of valid and invalid json data files in [instance selector tests](../../../pkg/util/instanceselector/testdata) used for unit testing.

## Updating instance data locally
To override KIP instance data, run `make update-pricing-data` from kip root dir.
This will generate new literals in `kip/pkg/util/instanceselect/<provider>_instance_data.go`.
You can commit those updates to repository later.


## Updating instance data periodically
Pricing-updater also allow you to run it inside kubernetes cluster, alongside KIP, to allow operating on "live" pricing data.
Steps:
1. Run `pricing-updater` as a kubernetes CronJob (check CronJob example & needed RBAC in [examples](manifests)).
2. `pricing-updater` will run every X seconds and create ConfigMap from instance data json file. You can specify desired ConfigMap name (`pricing-updater -update-configmap -configmap-name live-pricing-data)
3. You need to add volume mount to kip with created ConfigMap as a source.
4. Pass `instance-data-path <mounted-path>` to kip command.
