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

## Refresh data
To override KIP instance data, run `make update-pricing-data` from kip root dir.
