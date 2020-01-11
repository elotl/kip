package instanceselector

const awsInstanceJson = `
{
    "us-east-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.186,
            "memory": 16.0,
            "instanceType": "z1d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.372,
            "memory": 32.0,
            "instanceType": "z1d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.744,
            "memory": 64.0,
            "instanceType": "z1d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 12,
            "generation": "current",
            "price": 1.116,
            "memory": 96.0,
            "instanceType": "z1d.3xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 12
        },
        {
            "baseline": 24,
            "generation": "current",
            "price": 2.232,
            "memory": 192.0,
            "instanceType": "z1d.6xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 24
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 4.464,
            "memory": 384.0,
            "instanceType": "z1d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.65,
            "memory": 122.0,
            "instanceType": "f1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.3,
            "memory": 244.0,
            "instanceType": "f1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.2,
            "memory": 976.0,
            "instanceType": "f1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.468,
            "memory": 32.0,
            "instanceType": "h1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.936,
            "memory": 64.0,
            "instanceType": "h1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 1.872,
            "memory": 128.0,
            "instanceType": "h1.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.744,
            "memory": 256.0,
            "instanceType": "h1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "us-west-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.186,
            "memory": 16.0,
            "instanceType": "z1d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.372,
            "memory": 32.0,
            "instanceType": "z1d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.744,
            "memory": 64.0,
            "instanceType": "z1d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 12,
            "generation": "current",
            "price": 1.116,
            "memory": 96.0,
            "instanceType": "z1d.3xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 12
        },
        {
            "baseline": 24,
            "generation": "current",
            "price": 2.232,
            "memory": 192.0,
            "instanceType": "z1d.6xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 24
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 4.464,
            "memory": 384.0,
            "instanceType": "z1d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.65,
            "memory": 122.0,
            "instanceType": "f1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.3,
            "memory": 244.0,
            "instanceType": "f1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.2,
            "memory": 976.0,
            "instanceType": "f1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "ap-northeast-2": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "ap-northeast-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.186,
            "memory": 16.0,
            "instanceType": "z1d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.372,
            "memory": 32.0,
            "instanceType": "z1d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.744,
            "memory": 64.0,
            "instanceType": "z1d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 12,
            "generation": "current",
            "price": 1.116,
            "memory": 96.0,
            "instanceType": "z1d.3xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 12
        },
        {
            "baseline": 24,
            "generation": "current",
            "price": 2.232,
            "memory": 192.0,
            "instanceType": "z1d.6xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 24
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 4.464,
            "memory": 384.0,
            "instanceType": "z1d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "sa-east-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        }
    ],
    "ap-southeast-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.186,
            "memory": 16.0,
            "instanceType": "z1d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.372,
            "memory": 32.0,
            "instanceType": "z1d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.744,
            "memory": 64.0,
            "instanceType": "z1d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 12,
            "generation": "current",
            "price": 1.116,
            "memory": 96.0,
            "instanceType": "z1d.3xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 12
        },
        {
            "baseline": 24,
            "generation": "current",
            "price": 2.232,
            "memory": 192.0,
            "instanceType": "z1d.6xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 24
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 4.464,
            "memory": 384.0,
            "instanceType": "z1d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "ca-central-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "ap-southeast-2": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "us-west-2": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.186,
            "memory": 16.0,
            "instanceType": "z1d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.372,
            "memory": 32.0,
            "instanceType": "z1d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.744,
            "memory": 64.0,
            "instanceType": "z1d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 12,
            "generation": "current",
            "price": 1.116,
            "memory": 96.0,
            "instanceType": "z1d.3xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 12
        },
        {
            "baseline": 24,
            "generation": "current",
            "price": 2.232,
            "memory": 192.0,
            "instanceType": "z1d.6xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 24
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 4.464,
            "memory": 384.0,
            "instanceType": "z1d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.65,
            "memory": 122.0,
            "instanceType": "f1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.3,
            "memory": 244.0,
            "instanceType": "f1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.2,
            "memory": 976.0,
            "instanceType": "f1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.468,
            "memory": 32.0,
            "instanceType": "h1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.936,
            "memory": 64.0,
            "instanceType": "h1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 1.872,
            "memory": 128.0,
            "instanceType": "h1.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.744,
            "memory": 256.0,
            "instanceType": "h1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "us-east-2": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.468,
            "memory": 32.0,
            "instanceType": "h1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.936,
            "memory": 64.0,
            "instanceType": "h1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 1.872,
            "memory": 128.0,
            "instanceType": "h1.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.744,
            "memory": 256.0,
            "instanceType": "h1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "ap-south-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "eu-central-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "eu-west-1": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.834,
            "memory": 122.0,
            "instanceType": "x1e.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.668,
            "memory": 244.0,
            "instanceType": "x1e.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.336,
            "memory": 488.0,
            "instanceType": "x1e.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 6.672,
            "memory": 976.0,
            "instanceType": "x1e.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.344,
            "memory": 1952.0,
            "instanceType": "x1e.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 26.688,
            "memory": 3904.0,
            "instanceType": "x1e.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.186,
            "memory": 16.0,
            "instanceType": "z1d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.372,
            "memory": 32.0,
            "instanceType": "z1d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.744,
            "memory": 64.0,
            "instanceType": "z1d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 12,
            "generation": "current",
            "price": 1.116,
            "memory": 96.0,
            "instanceType": "z1d.3xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 12
        },
        {
            "baseline": 24,
            "generation": "current",
            "price": 2.232,
            "memory": 192.0,
            "instanceType": "z1d.6xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 24
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 4.464,
            "memory": 384.0,
            "instanceType": "z1d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.9,
            "memory": 61.0,
            "instanceType": "p2.xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 4
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 7.2,
            "memory": 488.0,
            "instanceType": "p2.8xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 14.4,
            "memory": 732.0,
            "instanceType": "p2.16xlarge",
            "burstable": false,
            "gpu": 16,
            "cpu": 64
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.14,
            "memory": 122.0,
            "instanceType": "g3.4xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.28,
            "memory": 244.0,
            "instanceType": "g3.8xlarge",
            "burstable": false,
            "gpu": 2,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.56,
            "memory": 488.0,
            "instanceType": "g3.16xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.65,
            "memory": 122.0,
            "instanceType": "f1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 3.3,
            "memory": 244.0,
            "instanceType": "f1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 13.2,
            "memory": 976.0,
            "instanceType": "f1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.468,
            "memory": 32.0,
            "instanceType": "h1.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.936,
            "memory": 64.0,
            "instanceType": "h1.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 1.872,
            "memory": 128.0,
            "instanceType": "h1.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.744,
            "memory": 256.0,
            "instanceType": "h1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ],
    "eu-west-2": [
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0052,
            "memory": 0.5,
            "instanceType": "t3.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0104,
            "memory": 1.0,
            "instanceType": "t3.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0208,
            "memory": 2.0,
            "instanceType": "t3.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "memory": 4.0,
            "instanceType": "t3.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "memory": 8.0,
            "instanceType": "t3.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 1.6,
            "generation": "current",
            "price": 0.1664,
            "memory": 16.0,
            "instanceType": "t3.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 3.2,
            "generation": "current",
            "price": 0.3328,
            "memory": 32.0,
            "instanceType": "t3.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 0.05,
            "generation": "current",
            "price": 0.0058,
            "memory": 0.5,
            "instanceType": "t2.nano",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "memory": 1.0,
            "instanceType": "t2.micro",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "memory": 2.0,
            "instanceType": "t2.small",
            "burstable": true,
            "gpu": 0,
            "cpu": 1
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "memory": 4.0,
            "instanceType": "t2.medium",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "memory": 8.0,
            "instanceType": "t2.large",
            "burstable": true,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.1856,
            "memory": 16.0,
            "instanceType": "t2.xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 1.3599999999999999,
            "generation": "current",
            "price": 0.3712,
            "memory": 32.0,
            "instanceType": "t2.2xlarge",
            "burstable": true,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 8.0,
            "instanceType": "m5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 16.0,
            "instanceType": "m5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 32.0,
            "instanceType": "m5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 64.0,
            "instanceType": "m5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.304,
            "memory": 192.0,
            "instanceType": "m5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 4.608,
            "memory": 384.0,
            "instanceType": "m5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.113,
            "memory": 8.0,
            "instanceType": "m5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.226,
            "memory": 16.0,
            "instanceType": "m5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.452,
            "memory": 32.0,
            "instanceType": "m5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.904,
            "memory": 64.0,
            "instanceType": "m5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 2.712,
            "memory": 192.0,
            "instanceType": "m5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 5.424,
            "memory": 384.0,
            "instanceType": "m5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 8.0,
            "instanceType": "m4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.2,
            "memory": 16.0,
            "instanceType": "m4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.4,
            "memory": 32.0,
            "instanceType": "m4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.8,
            "memory": 64.0,
            "instanceType": "m4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 40,
            "generation": "current",
            "price": 2.0,
            "memory": 160.0,
            "instanceType": "m4.10xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 40
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 3.2,
            "memory": 256.0,
            "instanceType": "m4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.085,
            "memory": 4.0,
            "instanceType": "c5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.17,
            "memory": 8.0,
            "instanceType": "c5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.34,
            "memory": 16.0,
            "instanceType": "c5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.68,
            "memory": 32.0,
            "instanceType": "c5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.53,
            "memory": 72.0,
            "instanceType": "c5.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.06,
            "memory": 144.0,
            "instanceType": "c5.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.096,
            "memory": 4.0,
            "instanceType": "c5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.192,
            "memory": 8.0,
            "instanceType": "c5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.384,
            "memory": 16.0,
            "instanceType": "c5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.768,
            "memory": 32.0,
            "instanceType": "c5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.728,
            "memory": 72.0,
            "instanceType": "c5d.9xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 72,
            "generation": "current",
            "price": 3.456,
            "memory": 144.0,
            "instanceType": "c5d.18xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 72
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.1,
            "memory": 3.75,
            "instanceType": "c4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.199,
            "memory": 7.5,
            "instanceType": "c4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.398,
            "memory": 15.0,
            "instanceType": "c4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 0.796,
            "memory": 30.0,
            "instanceType": "c4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 1.591,
            "memory": 60.0,
            "instanceType": "c4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.126,
            "memory": 16.0,
            "instanceType": "r5.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.252,
            "memory": 32.0,
            "instanceType": "r5.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.504,
            "memory": 64.0,
            "instanceType": "r5.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.008,
            "memory": 128.0,
            "instanceType": "r5.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.024,
            "memory": 384.0,
            "instanceType": "r5.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.048,
            "memory": 768.0,
            "instanceType": "r5.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.144,
            "memory": 16.0,
            "instanceType": "r5d.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.288,
            "memory": 32.0,
            "instanceType": "r5d.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.576,
            "memory": 64.0,
            "instanceType": "r5d.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.152,
            "memory": 128.0,
            "instanceType": "r5d.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 48,
            "generation": "current",
            "price": 3.456,
            "memory": 384.0,
            "instanceType": "r5d.12xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 48
        },
        {
            "baseline": 96,
            "generation": "current",
            "price": 6.912,
            "memory": 768.0,
            "instanceType": "r5d.24xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 96
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.133,
            "memory": 15.25,
            "instanceType": "r4.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.266,
            "memory": 30.5,
            "instanceType": "r4.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.532,
            "memory": 61.0,
            "instanceType": "r4.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.064,
            "memory": 122.0,
            "instanceType": "r4.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.128,
            "memory": 244.0,
            "instanceType": "r4.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.256,
            "memory": 488.0,
            "instanceType": "r4.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 6.669,
            "memory": 976.0,
            "instanceType": "x1.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 128,
            "generation": "current",
            "price": 13.338,
            "memory": 1952.0,
            "instanceType": "x1.32xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 128
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 3.06,
            "memory": 61.0,
            "instanceType": "p3.2xlarge",
            "burstable": false,
            "gpu": 1,
            "cpu": 8
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 12.24,
            "memory": 244.0,
            "instanceType": "p3.8xlarge",
            "burstable": false,
            "gpu": 4,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 24.48,
            "memory": 488.0,
            "instanceType": "p3.16xlarge",
            "burstable": false,
            "gpu": 8,
            "cpu": 64
        },
        {
            "baseline": 2,
            "generation": "current",
            "price": 0.156,
            "memory": 15.25,
            "instanceType": "i3.large",
            "burstable": false,
            "gpu": 0,
            "cpu": 2
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.312,
            "memory": 30.5,
            "instanceType": "i3.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 0.624,
            "memory": 61.0,
            "instanceType": "i3.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 1.248,
            "memory": 122.0,
            "instanceType": "i3.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 32,
            "generation": "current",
            "price": 2.496,
            "memory": 244.0,
            "instanceType": "i3.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 32
        },
        {
            "baseline": 64,
            "generation": "current",
            "price": 4.992,
            "memory": 488.0,
            "instanceType": "i3.16xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 64
        },
        {
            "baseline": 4,
            "generation": "current",
            "price": 0.69,
            "memory": 30.5,
            "instanceType": "d2.xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 4
        },
        {
            "baseline": 8,
            "generation": "current",
            "price": 1.38,
            "memory": 61.0,
            "instanceType": "d2.2xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 8
        },
        {
            "baseline": 16,
            "generation": "current",
            "price": 2.76,
            "memory": 122.0,
            "instanceType": "d2.4xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 16
        },
        {
            "baseline": 36,
            "generation": "current",
            "price": 5.52,
            "memory": 244.0,
            "instanceType": "d2.8xlarge",
            "burstable": false,
            "gpu": 0,
            "cpu": 36
        }
    ]
}
`