# Pricing updater

This app gets pricing data for given region and provider from cloudinfo, converts it to format that KIP understands and creates ConfigMap, which KIP can use.

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
