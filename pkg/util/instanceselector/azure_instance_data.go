/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package instanceselector

const azureInstanceJson = `
{
    "France Central": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.022,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.026,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.132,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.224,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.188,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.088,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.343,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.202,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.404,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.808,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.178,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.786,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.467,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.026,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.104,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.052,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.208,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.416,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.088,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.202,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.404,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.808,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.112,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.224,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.448,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.896,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.792,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.584,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.112,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.224,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.448,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.896,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.792,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.584,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.488,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.488,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.488,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.488,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.202,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.404,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.808,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.616,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.232,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.636,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "East US 2": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.24,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.48,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.495,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.524,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0208,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0104,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.166,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.333,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.133,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.266,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.532,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.064,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.128,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.133,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.266,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.532,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.064,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.128,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.495,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.067,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.134,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.268,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.536,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.173,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.347,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.694,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.387,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.12,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.464,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.24,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.55,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.1,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.82,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.55,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.1,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.82,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.82,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.82,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.343,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.686,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.373,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.746,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.202,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.404,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.807,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.067,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.134,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.268,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.536,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.173,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.347,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.694,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.387,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.9,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.8,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.6,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.338,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.677,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.353,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.706,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.045,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1716,
            "burstable": false,
            "instanceType": "Standard_L8s_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3432,
            "burstable": false,
            "instanceType": "Standard_L16s_v2",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.24,
            "burstable": false,
            "instanceType": "Standard_L80s_v2",
            "memory": 640.0,
            "cpu": 80
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.5365,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.073,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.873,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.146,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.707,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.415,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.337,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.669,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 26.688,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.338,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        }
    ],
    "Central India": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.17,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.34,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.68,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.36,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.72,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.1165,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.2331,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.27522,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.4661,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.08598,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.1731,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.24,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.187,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 36.765,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 18.374,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.524,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.52,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.04,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.047,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.333,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.206,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.699,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.433,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0267,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0133,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.107,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.053,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.214,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.428,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.049,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.198,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.395,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.79,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.049,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.198,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.395,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.79,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.105,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.36,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.105,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.36,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.137,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.274,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.548,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.096,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.192,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.937,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.137,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.274,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.548,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.096,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.192,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.937,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.895,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.718,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.436,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.872,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.895,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.814,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.628,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 25.5816,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 23.256,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        }
    ],
    "Australia Central": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.03,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.08875,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1225,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.355,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3475,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.71,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.695,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.39,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.105,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.493,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.081,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.162,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.325,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.65,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.3,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0625,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.18625,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1325,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.39,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2775,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.81875,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.58375,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.139,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.277,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.554,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.108,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.217,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.434,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.988,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.04,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.02,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.16,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.08,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.32,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.64,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.105,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.081,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.162,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.325,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.65,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.3,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15625,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3125,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.625,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.25,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.5,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.0,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15625,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3125,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.625,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.25,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.5,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.0,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.39875,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7975,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.59625,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.1925,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.434,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.39875,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7975,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.59625,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.1925,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.434,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.785,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.57,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.31,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.14,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.061,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.123,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 18.737,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.088,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 48.372,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 24.175,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        }
    ],
    "Central US": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.025,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.188,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.376,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.853,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.438,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.27,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.568,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.76,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.025,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0998,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0499,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.2,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.399,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.438,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.52,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.76,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.52,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.341,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.199,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.2,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.341,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.199,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.2,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.853,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.204,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.816,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.632,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.264,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.672,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "West US": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0248,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0124,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0992,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0496,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.198,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.397,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.07,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.279,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.559,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.117,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.852,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.279,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.559,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.117,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.062,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.124,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.498,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.996,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.117,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.468,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.936,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.872,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.081,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.188,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.376,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.07,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.279,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.559,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.117,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.852,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.279,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.559,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.117,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.062,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.124,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.498,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.996,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.117,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.468,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.936,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.872,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.744,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.744,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.148,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.296,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.593,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.186,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.48,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.371,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.032,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.032,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.148,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.296,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.593,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.186,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.48,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.371,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.032,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.032,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.61,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.22,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.44,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.88,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.69,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.61,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.22,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.44,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.88,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.88,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.88,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.69,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.69,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.69,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.344,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.688,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.376,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.752,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.975,
            "burstable": false,
            "instanceType": "Standard_A8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.95,
            "burstable": false,
            "instanceType": "Standard_A9",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.78,
            "burstable": false,
            "instanceType": "Standard_A10",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_A11",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.971,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.941,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.301,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.601,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.136,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.861,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.767,
            "burstable": false,
            "instanceType": "Standard_NV6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.534,
            "burstable": false,
            "instanceType": "Standard_NV12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.069,
            "burstable": false,
            "instanceType": "Standard_NV24s_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.212,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.424,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.848,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.696,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.392,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.816,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "UK West": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0235,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.094,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.047,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.188,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.376,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.088,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.059,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.237,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.475,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.95,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.088,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.404,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.059,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.237,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.475,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.95,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.333,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.699,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.467,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.116,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.928,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.856,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.712,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.116,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.928,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.856,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.712,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.022,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.264,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.188,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.343,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.343,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.202,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.405,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.809,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.618,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.237,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.641,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.936,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.872,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.36141,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.744,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.16719,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.33555,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.024,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.403,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 33.6269,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 16.806,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        }
    ],
    "UK South": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0236,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0118,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0944,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0472,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.189,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.378,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.026,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.132,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.224,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.188,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.088,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.405,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.343,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.405,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.059,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.237,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.475,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.95,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.088,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.405,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.343,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.351,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.702,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.405,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.937,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.874,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.059,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.237,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.475,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.95,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.178,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.786,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.467,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.116,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.928,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.856,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.116,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.928,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.856,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.507,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.003,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.006,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.712,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.712,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.264,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.283,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.566,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.133,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.36141,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.744,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.16719,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.33555,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.024,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.403,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 33.6269,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 16.806,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.952,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.902,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.274,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.549,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.092,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.803,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.202,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.404,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.808,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.616,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.232,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.636,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.936,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.872,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.825,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.65,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 16.83,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 15.3,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.439,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.878,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.757,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.514,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.027,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.439,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.878,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.757,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.514,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.514,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.514,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.027,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.027,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.027,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.362,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.724,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.448,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.896,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        }
    ],
    "Brazil South": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.024,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.041,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.291,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.582,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.164,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.095,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.19,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.38,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.76,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.939,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.086,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.171,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.343,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.685,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.349,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.171,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.343,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.685,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.072,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.287,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.574,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.148,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.061,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.208,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.437,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.27,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.917,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.567,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.159,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.318,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.636,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.272,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.544,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0336,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0168,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.134,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0672,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.269,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.538,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.086,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.171,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.343,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.685,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.171,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.343,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.685,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.072,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.287,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.574,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.148,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.088,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.159,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.318,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.636,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.272,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.544,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.088,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.8,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.758,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.764,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.764,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.47,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.94,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.879,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.8,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.758,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.764,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.764,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.349,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.131,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.262,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.524,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.048,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.096,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.192,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.716,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "Korea Central": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0281,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.014,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.112,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0562,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.225,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.449,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.083,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.165,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.33,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.66,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.321,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.165,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.33,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.66,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.321,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.051,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.204,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.816,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.083,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.165,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.33,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.66,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.321,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.165,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.33,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.66,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.321,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.051,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.204,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.816,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.049,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.133,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.278,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.45,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.123,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.246,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.492,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.984,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.968,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.936,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.123,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.246,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.492,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.984,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.968,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.936,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.108,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.224,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.172,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.343,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.687,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.456,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "Korea South": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.026,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.104,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.052,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.207,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.415,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.074,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.189,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.179,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.179,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.359,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.359,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.359,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.718,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.718,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.718,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.435,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.435,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.435,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.794,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.189,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.179,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.359,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.718,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.435,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.051,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.204,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.816,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.221,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.443,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.886,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.772,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.074,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.189,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.179,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.359,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.718,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.435,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.297,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.594,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.189,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.179,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.359,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.718,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.435,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.051,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.204,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.816,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.044,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.092,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.251,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.526,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.405,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.221,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.443,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.886,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.772,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.544,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.544,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.288,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.576,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.152,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.304,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.778,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.881,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.288,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.576,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.152,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.304,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.778,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.881,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.029,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.238,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.155,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.404,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.309,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.618,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.794,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.44,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.44,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0961,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.537,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.074,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.458,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "South Central US": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.352,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.064,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.509,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.662,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.509,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.438,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.124,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.248,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.495,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.025,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0998,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0499,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.2,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.399,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.064,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.509,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.662,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.509,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.438,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.76,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.067,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.134,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.268,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.536,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.173,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.347,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.694,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.387,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.76,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.15,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.3,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.6,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.067,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.134,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.268,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.536,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.173,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.347,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.694,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.387,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.254,
            "burstable": false,
            "instanceType": "Standard_HB60rs",
            "memory": 223.52,
            "cpu": 60
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.204,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.816,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.632,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.264,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.672,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.366,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.732,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.8104,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.464,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.08,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.16,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.32,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.52,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.52,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.341,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.199,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.199,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.341,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.199,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.199,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.749,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.172,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.343,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.924,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.577,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_A8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.75,
            "burstable": false,
            "instanceType": "Standard_A9",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7,
            "burstable": false,
            "instanceType": "Standard_A10",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.4,
            "burstable": false,
            "instanceType": "Standard_A11",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.484,
            "burstable": false,
            "instanceType": "Standard_NC6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.968,
            "burstable": false,
            "instanceType": "Standard_NC12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.93,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.936,
            "burstable": false,
            "instanceType": "Standard_NC24s_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.684,
            "burstable": false,
            "instanceType": "Standard_NV6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.368,
            "burstable": false,
            "instanceType": "Standard_NV12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.736,
            "burstable": false,
            "instanceType": "Standard_NV24s_v2",
            "memory": 448.0,
            "cpu": 24
        }
    ],
    "East US": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0207,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0104,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.166,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.333,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.853,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.023,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.079,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.352,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.853,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.238,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.475,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.26,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.016,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.26,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.016,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.904,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.807,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.211,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.422,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.988,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.664,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.14,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.28,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.56,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.07,
            "burstable": false,
            "instanceType": "Standard_NC6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.14,
            "burstable": false,
            "instanceType": "Standard_NC12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.108,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.28,
            "burstable": false,
            "instanceType": "Standard_NC24s_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.9,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.8,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.6,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.338,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.677,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.353,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.706,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.045,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.07,
            "burstable": false,
            "instanceType": "Standard_ND6s",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.14,
            "burstable": false,
            "instanceType": "Standard_ND12s",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.108,
            "burstable": false,
            "instanceType": "Standard_ND24rs",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.28,
            "burstable": false,
            "instanceType": "Standard_ND24s",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.198,
            "burstable": false,
            "instanceType": "Standard_DC2s",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.395,
            "burstable": false,
            "instanceType": "Standard_DC4s",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.12,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.464,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.24,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_L8s_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_L16s_v2",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.24,
            "burstable": false,
            "instanceType": "Standard_L80s_v2",
            "memory": 640.0,
            "cpu": 80
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.975,
            "burstable": false,
            "instanceType": "Standard_A8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.95,
            "burstable": false,
            "instanceType": "Standard_A9",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.78,
            "burstable": false,
            "instanceType": "Standard_A10",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_A11",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.5365,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.073,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.873,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.146,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.707,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.415,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.337,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.669,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 26.688,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.338,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        }
    ],
    "Canada Central": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.023,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0116,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0928,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0464,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.185,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.371,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.076,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.152,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.305,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.61,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.22,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.152,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.305,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.61,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.22,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.052,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.104,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.207,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.415,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.83,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.076,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.152,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.305,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.61,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.22,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.152,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.305,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.61,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.22,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.052,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.104,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.207,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.415,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.83,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.047,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.353,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.206,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.742,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.433,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.444,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.888,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.776,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.552,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.444,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.888,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.776,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.552,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.292,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.168,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.336,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.292,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.168,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.336,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.093,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.186,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.488,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.976,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.348,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.024,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.061,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.121,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.264,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.486,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.056,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.056,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.112,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.224,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.448,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.528,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.056,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.112,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.224,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.224,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.224,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.448,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.448,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.448,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.344,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.688,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.376,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.752,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.69015,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.3803,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.1603,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.7606,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.9777,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.9565,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.3707,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.3359,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 29.3568,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.6718,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.672,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.344,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 16.157,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.688,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        }
    ],
    "South India": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.016,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.033,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.118,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.236,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.234,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.418,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.468,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.936,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.895,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.121,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.483,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.966,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.047,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.301,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.206,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.633,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.433,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0294,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0145,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.118,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.058,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.236,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.471,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.121,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.483,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.966,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.135,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.271,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.541,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.082,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.164,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.328,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.135,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.271,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.541,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.082,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.164,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.328,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.167,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.334,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.669,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.338,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.67,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.675,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.801,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.801,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.167,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.334,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.669,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.338,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.67,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.675,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.801,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.801,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0935,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.187,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.748,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.496,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.992,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.366,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.895,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.3278,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.6556,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.61998,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.3112,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.41082,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.8229,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 15.664,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.106,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 40.441,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 20.211,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        }
    ],
    "West Europe": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.078,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.24,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.27,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.408,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.54,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.08,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.336,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.672,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.221,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.442,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.885,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.769,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.041,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.124,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.087,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.546,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.383,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.336,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.672,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.221,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.442,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.885,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.769,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.068,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.136,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.272,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.544,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.087,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.19,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.759,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.518,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.897,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.136,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.272,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.544,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.087,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.19,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.759,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.518,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.227,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.454,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.909,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.024,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.012,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.096,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.048,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.192,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.384,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.068,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.136,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.272,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.544,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.087,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.19,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.19,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.759,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.759,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.759,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.518,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.518,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.518,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.897,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.136,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.272,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.544,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.087,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.19,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.759,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.518,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.227,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.454,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.909,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.24,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.48,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.96,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.92,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.24,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.48,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.96,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.92,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.84,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.37,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.73,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.46,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.84,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.097,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.194,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.388,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.776,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.552,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.104,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.492,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.978,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.956,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 17.5032,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 15.912,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.4,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.8,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.6,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.99,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.4,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.8,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.6,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.6,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.6,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.99,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.99,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.99,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.488,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.976,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.166,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.333,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.666,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.032,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.065,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.383,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.767,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.271,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.043,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.1511,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.3022,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.4476,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.6044,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2484,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.498,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.472,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.337,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 37.3632,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 18.674,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.682,
            "burstable": false,
            "instanceType": "Standard_ND6s",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.366,
            "burstable": false,
            "instanceType": "Standard_ND12s",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.804,
            "burstable": false,
            "instanceType": "Standard_ND24rs",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.732,
            "burstable": false,
            "instanceType": "Standard_ND24s",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.682,
            "burstable": false,
            "instanceType": "Standard_NC6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.366,
            "burstable": false,
            "instanceType": "Standard_NC12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.804,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.732,
            "burstable": false,
            "instanceType": "Standard_NC24s_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.257,
            "burstable": false,
            "instanceType": "Standard_DC2s",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.514,
            "burstable": false,
            "instanceType": "Standard_DC4s",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.121,
            "burstable": false,
            "instanceType": "Standard_A8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.243,
            "burstable": false,
            "instanceType": "Standard_A9",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.897,
            "burstable": false,
            "instanceType": "Standard_A10",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.794,
            "burstable": false,
            "instanceType": "Standard_A11",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_L8s_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_L16s_v2",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.44,
            "burstable": false,
            "instanceType": "Standard_L80s_v2",
            "memory": 640.0,
            "cpu": 80
        }
    ],
    "Australia Central 2": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.04,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.02,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.16,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.08,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.32,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.64,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.105,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.081,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.162,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.325,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.65,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.3,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.105,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.841,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.681,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49875,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.997,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.081,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.162,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.325,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.65,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.3,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0625,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.18625,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1325,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.39,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2775,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.81875,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.58375,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15625,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3125,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.625,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.25,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.5,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.0,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15625,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3125,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.625,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.25,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.5,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.0,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.39875,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7975,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.59625,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.1925,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.434,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.39875,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7975,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.59625,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.1925,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.434,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.03,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.04,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1775,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.355,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3475,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.58,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.695,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.39,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.493,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.139,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.277,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.554,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.108,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.217,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.434,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.988,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "Canada East": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0233,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0932,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0466,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.186,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.373,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.07,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.28,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.56,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.12,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.463,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.28,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.56,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.12,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.463,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.218,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.437,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.874,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.07,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.28,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.56,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.12,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.463,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.463,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.463,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.14,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.28,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.56,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.12,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.463,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.218,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.437,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.874,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.27,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.568,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.444,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.888,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.776,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.552,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.444,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.888,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.776,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.552,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.292,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.168,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.336,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.292,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.168,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.46,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.336,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.974,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.026,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.087,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.194,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.446,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.484,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.968,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.829,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.69015,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.3803,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.1603,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.7606,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.9777,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.9565,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.3707,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.3359,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 29.3568,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.6718,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0927,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.965,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.335,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.829,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.484,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.968,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.936,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.872,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.744,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.484,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.968,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.936,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.872,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.872,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.872,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.744,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.744,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.744,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.344,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.688,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.376,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.752,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        }
    ],
    "West India": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.52,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.04,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.895,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.439,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.878,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.047,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.333,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.206,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.699,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.433,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.895,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.169,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.337,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.675,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.35,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.189,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.379,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.758,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.516,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.439,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.878,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.03,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0147,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.119,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.059,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.238,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.476,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.123,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.246,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.492,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.968,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.936,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.123,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.246,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.492,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.968,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.936,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.152,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.52,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.432,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.368,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.368,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.152,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.52,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.432,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.368,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.368,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.17,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.34,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.68,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.36,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.72,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "France South": [
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0338,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0169,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.135,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0676,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.27,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.541,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.228,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.456,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.912,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.824,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.304,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.608,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.217,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.434,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.228,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.456,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.912,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.824,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.304,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.608,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.217,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.434,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.131,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.262,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.525,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.05,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.228,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.456,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.912,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.824,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.304,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.304,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.608,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.608,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.608,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.217,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.217,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.217,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.434,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.434,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.434,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.228,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.456,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.912,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.824,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.304,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.608,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.217,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.434,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.131,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.262,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.525,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.05,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.065,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2314,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1378,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4862,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2886,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0218,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.6071,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1456,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.291,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.582,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.1648,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.3296,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.6592,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1456,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.291,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.582,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.1648,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.3296,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.6592,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2028,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4056,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.8112,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6224,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2448,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.835,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.835,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2028,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4056,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.8112,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6224,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2448,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.835,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.835,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.026,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0338,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1313,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3432,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.3861,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5837,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.7722,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.5444,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.042,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.131,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.262,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.524,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.047,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.094,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.189,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.712,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "Japan West": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.182,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.365,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.729,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.459,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.182,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.365,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.729,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.459,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.063,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.0171,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.032,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2196,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.258,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.516,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.032,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.054,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.178,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.113,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.238,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.786,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.258,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.516,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.032,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.064,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.128,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.182,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.365,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.729,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.459,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.182,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.365,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.729,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.459,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.063,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.03,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.016,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.12,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0599,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.24,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.479,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.258,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.516,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.032,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.064,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.128,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.117,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.235,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.469,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.938,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.877,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.754,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.223,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2279,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4559,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.4476,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.9117,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2484,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.498,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.99,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.671,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 38.701,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 19.342,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        }
    ],
    "East Asia": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.038,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.24,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.294,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.48,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.588,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.176,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.107,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.857,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.714,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.294,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.857,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.714,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.072,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.289,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.578,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.155,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.178,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.786,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.467,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.313,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.625,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.25,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.5,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.113,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.225,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.451,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.902,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.483,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.966,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.932,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.034,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.017,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.135,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.068,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.27,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.54,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.107,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.857,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.714,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.294,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.857,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.714,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.072,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.144,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.289,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.578,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.155,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.156,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.313,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.625,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.25,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.5,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.0,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.0,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.8,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.625,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.625,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.8,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.625,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.625,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.113,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.225,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.451,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.902,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.483,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.966,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.932,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.784875,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.569875,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.3095,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.139625,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0605,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.1225,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 18.7375,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.08875,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 48.37625,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 24.1775,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.122,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.245,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.979,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.958,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.917,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.406,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "Southeast Asia": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.59,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.18,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.36,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.095,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.27,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.48,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.54,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.08,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.079,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.316,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.631,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.263,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.382,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.765,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.53,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.912,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.316,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.631,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.263,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.382,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.765,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.53,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.058,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.115,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.231,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.462,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.924,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.045,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.134,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.095,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.281,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.198,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.589,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.417,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.125,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.196,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.392,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.784,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0264,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0132,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.106,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0528,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.211,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.422,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.079,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.316,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.631,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.263,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.382,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.382,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.382,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.765,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.765,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.765,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.53,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.53,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.53,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.158,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.316,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.631,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.263,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.382,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.765,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.53,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.058,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.115,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.231,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.462,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.924,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.125,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.609,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.912,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.814,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 11.628,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 25.5816,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 23.256,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.196,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.392,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.784,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2279,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4559,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.4476,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.9117,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2484,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.498,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.99,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.671,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 38.701,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 19.342,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.038,
            "burstable": false,
            "instanceType": "Standard_NC6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.077,
            "burstable": false,
            "instanceType": "Standard_NC12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.37,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.153,
            "burstable": false,
            "instanceType": "Standard_NC24s_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_L8s_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.748,
            "burstable": false,
            "instanceType": "Standard_L16s_v2",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.48,
            "burstable": false,
            "instanceType": "Standard_L80s_v2",
            "memory": 640.0,
            "cpu": 80
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.66,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.32,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.64,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.28,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.39,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.66,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.32,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.64,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.28,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.28,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.28,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.39,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.39,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.39,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.748,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.496,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.992,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.098,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.196,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.392,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.784,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.568,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.136,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.528,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.038,
            "burstable": false,
            "instanceType": "Standard_ND6s",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.077,
            "burstable": false,
            "instanceType": "Standard_ND12s",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.37,
            "burstable": false,
            "instanceType": "Standard_ND24rs",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.153,
            "burstable": false,
            "instanceType": "Standard_ND24s",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.925,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.848,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.238,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.476,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.033,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.724,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        }
    ],
    "Australia Southeast": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.029,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.071,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.142,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.284,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.278,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.556,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.112,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.093,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.186,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.178,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.786,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.467,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2279,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4559,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.4476,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.9117,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2484,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.498,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.99,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.671,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 38.701,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 19.341,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.093,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.186,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.032,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.016,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.128,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.064,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.256,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.512,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.078,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15576,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.311,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.623,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.246,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15576,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.311,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.623,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.246,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.065,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.13,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.521,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.042,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.078,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15576,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.311,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.623,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.246,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.15576,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.311,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.623,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.246,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.065,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.13,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.521,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.042,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.125,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.125,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.319,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.638,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.277,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.554,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.319,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.638,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.277,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.554,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.223,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.445,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.891,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.782,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.563,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.009,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "Australia East": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.638,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.276,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.552,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.104,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.208,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.638,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.276,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.552,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.104,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.104,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.104,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.208,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.208,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.208,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.374,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.748,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.496,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.992,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0264,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0132,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.106,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0528,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.211,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.422,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.336,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.673,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.345,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.336,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.673,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.345,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.065,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.13,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.521,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.042,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.084,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.336,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.673,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.345,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.168,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.336,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.673,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.345,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.399,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.798,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.596,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.065,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.13,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.26,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.521,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.042,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.106,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.655,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.467,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.125,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.125,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.0,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.0,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.319,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.638,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.277,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.554,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.319,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.638,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.277,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.554,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.347,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2279,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4559,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.4476,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.9117,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2484,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.498,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.99,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.671,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 38.6976,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 19.341,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.024,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.032,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.142,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.232,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.278,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.464,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.556,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.112,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.093,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.186,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.185,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.369,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.587,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.175,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.606,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.492,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.995,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.59,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.18,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.36,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.093,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.186,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.372,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.744,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.21,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.42,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.84,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.68,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.24178,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.48356,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 23.063832,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 20.96712,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.55,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.11,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.22,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.111,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.222,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.444,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.888,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.776,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.552,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.996,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        }
    ],
    "West US 2": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.17,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.34,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.68,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.36,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.72,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0208,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0104,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.166,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.333,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.023,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.405,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.495,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.495,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.114,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.458,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.916,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.149,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.299,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.598,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.196,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.036,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.076,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.208,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.159,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.437,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.333,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.873,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.146,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.707,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 5.415,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 10.337,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.669,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 26.688,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.338,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.26,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.016,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.26,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.016,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.07,
            "burstable": false,
            "instanceType": "Standard_NC6s_v2",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.14,
            "burstable": false,
            "instanceType": "Standard_NC12s_v2",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.108,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.28,
            "burstable": false,
            "instanceType": "Standard_NC24s_v2",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.591,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.066,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.132,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.75,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.345,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_NC6s_v3",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.12,
            "burstable": false,
            "instanceType": "Standard_NC12s_v3",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 13.464,
            "burstable": false,
            "instanceType": "Standard_NC24rs_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.24,
            "burstable": false,
            "instanceType": "Standard_NC24s_v3",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.584,
            "burstable": false,
            "instanceType": "Standard_HC44rs",
            "memory": 327.83,
            "cpu": 44
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.9,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.8,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.6,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.981,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.961,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.922,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.843,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.49,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.981,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.961,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.922,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.922,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.922,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.843,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.843,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.843,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.312,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.624,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.248,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.496,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.093,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.185,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.37,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.5365,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.073,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.07,
            "burstable": false,
            "instanceType": "Standard_ND6s",
            "memory": 112.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.14,
            "burstable": false,
            "instanceType": "Standard_ND12s",
            "memory": 224.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.108,
            "burstable": false,
            "instanceType": "Standard_ND24rs",
            "memory": 448.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.28,
            "burstable": false,
            "instanceType": "Standard_ND24s",
            "memory": 448.0,
            "cpu": 24
        }
    ],
    "North Europe": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.132,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.263,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.527,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.053,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.132,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.263,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.527,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.053,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.113,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.226,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.452,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.905,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.025,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.12,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.188,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.248,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.376,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.496,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.992,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.041,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.139,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.087,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.291,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.183,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.611,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.383,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.107,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.856,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.712,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.424,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.141,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.282,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.564,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.128,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.41,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.256,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.061,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.061,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0227,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0113,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.091,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.045,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.182,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.364,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.066,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.132,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.263,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.527,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.053,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.853,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.132,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.263,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.527,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.053,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.371,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.057,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.113,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.226,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.452,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.905,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.107,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.856,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.712,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.424,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.141,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.282,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.564,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.128,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.41,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.256,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.061,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.061,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.853,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.292,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.21,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.42,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.84,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.8438,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.6876,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.18903,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 7.3752,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.00477,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.01065,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 12.405,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.003,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 32.0256,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 16.006,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.972,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.944,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.888,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.292,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.584,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.096,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.192,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.384,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.768,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.536,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.072,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.456,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.975,
            "burstable": false,
            "instanceType": "Standard_A8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.95,
            "burstable": false,
            "instanceType": "Standard_A9",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.78,
            "burstable": false,
            "instanceType": "Standard_A10",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_A11",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.971,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.941,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.301,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.601,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.136,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.861,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        }
    ],
    "North Central US": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.02,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.06,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.24,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.25,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.48,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.0,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.852,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.043,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.091,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.27,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.191,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.568,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.8,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.26,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.016,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.904,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.808,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.211,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.423,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.989,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.665,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0208,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0104,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0832,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0416,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.166,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.333,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.073,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.146,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.293,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.585,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.185,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.37,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.741,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.482,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.05,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.099,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.199,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.398,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.796,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.1,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.2,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.4,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.8,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.26,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.016,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.629,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.077,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.154,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.308,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.616,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.193,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.386,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.771,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.542,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.14,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.28,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.56,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.085,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.17,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.34,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.68,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.36,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.72,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.06,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.9,
            "burstable": false,
            "instanceType": "Standard_NC6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.8,
            "burstable": false,
            "instanceType": "Standard_NC12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.6,
            "burstable": false,
            "instanceType": "Standard_NC24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.852,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.975,
            "burstable": false,
            "instanceType": "Standard_A8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.95,
            "burstable": false,
            "instanceType": "Standard_A9",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.78,
            "burstable": false,
            "instanceType": "Standard_A10",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.56,
            "burstable": false,
            "instanceType": "Standard_A11",
            "memory": 112.0,
            "cpu": 16
        }
    ],
    "Japan East": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.58,
            "burstable": false,
            "instanceType": "Standard_NV6",
            "memory": 56.0,
            "cpu": 6
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.16,
            "burstable": false,
            "instanceType": "Standard_NV12",
            "memory": 112.0,
            "cpu": 12
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.32,
            "burstable": false,
            "instanceType": "Standard_NV24",
            "memory": 224.0,
            "cpu": 24
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.022,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.032,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.324,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.281,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.648,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.562,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.124,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.205,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.409,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.818,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.636,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.294,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.205,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.409,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.818,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.636,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.063,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.054,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.153,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.113,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.322,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.238,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.677,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.5,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.258,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.516,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.032,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.064,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.2279,
            "burstable": false,
            "instanceType": "Standard_M8ms",
            "memory": 218.75,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.4559,
            "burstable": false,
            "instanceType": "Standard_M16ms",
            "memory": 437.5,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.4476,
            "burstable": false,
            "instanceType": "Standard_M32ls",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 8.9117,
            "burstable": false,
            "instanceType": "Standard_M32ms",
            "memory": 875.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.2484,
            "burstable": false,
            "instanceType": "Standard_M32ts",
            "memory": 192.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 6.498,
            "burstable": false,
            "instanceType": "Standard_M64ls",
            "memory": 512.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 14.99,
            "burstable": false,
            "instanceType": "Standard_M64ms",
            "memory": 1750.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.671,
            "burstable": false,
            "instanceType": "Standard_M64s",
            "memory": 1000.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 38.701,
            "burstable": false,
            "instanceType": "Standard_M128ms",
            "memory": 3800.0,
            "cpu": 128
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 19.342,
            "burstable": false,
            "instanceType": "Standard_M128s",
            "memory": 2000.0,
            "cpu": 128
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.0272,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.0136,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.109,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0544,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.218,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.435,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.102,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.205,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.409,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.818,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.636,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.294,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.205,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.409,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.818,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.636,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.229,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.459,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.918,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.835,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.063,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.126,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.252,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.504,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.008,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.129,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.258,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.516,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.032,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.064,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.145,
            "burstable": false,
            "instanceType": "Standard_H8",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.291,
            "burstable": false,
            "instanceType": "Standard_H16",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.535,
            "burstable": false,
            "instanceType": "Standard_H8m",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.07,
            "burstable": false,
            "instanceType": "Standard_H16m",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.52,
            "burstable": false,
            "instanceType": "Standard_H16r",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.377,
            "burstable": false,
            "instanceType": "Standard_H16mr",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.107,
            "burstable": false,
            "instanceType": "Standard_F2s_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.214,
            "burstable": false,
            "instanceType": "Standard_F4s_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.428,
            "burstable": false,
            "instanceType": "Standard_F8s_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.856,
            "burstable": false,
            "instanceType": "Standard_F16s_v2",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.712,
            "burstable": false,
            "instanceType": "Standard_F32s_v2",
            "memory": 64.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.424,
            "burstable": false,
            "instanceType": "Standard_F64s_v2",
            "memory": 128.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 3.852,
            "burstable": false,
            "instanceType": "Standard_F72s_v2",
            "memory": 144.0,
            "cpu": 72
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.128,
            "burstable": false,
            "instanceType": "Standard_D64_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.128,
            "burstable": false,
            "instanceType": "Standard_D64s_v3",
            "memory": 256.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64i_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.16,
            "burstable": false,
            "instanceType": "Standard_E2s_v3",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.32,
            "burstable": false,
            "instanceType": "Standard_E4s_v3",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.64,
            "burstable": false,
            "instanceType": "Standard_E8s_v3",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.28,
            "burstable": false,
            "instanceType": "Standard_E16s_v3",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.6,
            "burstable": false,
            "instanceType": "Standard_E20s_v3",
            "memory": 160.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.56,
            "burstable": false,
            "instanceType": "Standard_E32s_v3",
            "memory": 256.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64is_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.376,
            "burstable": false,
            "instanceType": "Standard_E64s_v3",
            "memory": 432.0,
            "cpu": 64
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.221,
            "burstable": false,
            "instanceType": "Standard_D2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.442,
            "burstable": false,
            "instanceType": "Standard_D3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.883,
            "burstable": false,
            "instanceType": "Standard_D4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_D11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.483,
            "burstable": false,
            "instanceType": "Standard_D12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.966,
            "burstable": false,
            "instanceType": "Standard_D13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.932,
            "burstable": false,
            "instanceType": "Standard_D14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_DS1",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.221,
            "burstable": false,
            "instanceType": "Standard_DS2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.442,
            "burstable": false,
            "instanceType": "Standard_DS3",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.883,
            "burstable": false,
            "instanceType": "Standard_DS4",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.242,
            "burstable": false,
            "instanceType": "Standard_DS11",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.483,
            "burstable": false,
            "instanceType": "Standard_DS12",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.966,
            "burstable": false,
            "instanceType": "Standard_DS13",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.932,
            "burstable": false,
            "instanceType": "Standard_DS14",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.576,
            "burstable": false,
            "instanceType": "Standard_G1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.151,
            "burstable": false,
            "instanceType": "Standard_G2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.302,
            "burstable": false,
            "instanceType": "Standard_G3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.602,
            "burstable": false,
            "instanceType": "Standard_G4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.205,
            "burstable": false,
            "instanceType": "Standard_G5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.576,
            "burstable": false,
            "instanceType": "Standard_GS1",
            "memory": 28.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.151,
            "burstable": false,
            "instanceType": "Standard_GS2",
            "memory": 56.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.302,
            "burstable": false,
            "instanceType": "Standard_GS3",
            "memory": 112.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.602,
            "burstable": false,
            "instanceType": "Standard_GS4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.602,
            "burstable": false,
            "instanceType": "Standard_GS4-4",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 4.602,
            "burstable": false,
            "instanceType": "Standard_GS4-8",
            "memory": 224.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.205,
            "burstable": false,
            "instanceType": "Standard_GS5",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.205,
            "burstable": false,
            "instanceType": "Standard_GS5-8",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 9.205,
            "burstable": false,
            "instanceType": "Standard_GS5-16",
            "memory": 448.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.366,
            "burstable": false,
            "instanceType": "Standard_L4s",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.732,
            "burstable": false,
            "instanceType": "Standard_L8s",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.464,
            "burstable": false,
            "instanceType": "Standard_L16s",
            "memory": 128.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 2.928,
            "burstable": false,
            "instanceType": "Standard_L32s",
            "memory": 256.0,
            "cpu": 32
        }
    ],
    "West Central US": [
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.064,
            "burstable": false,
            "instanceType": "Standard_DS1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_DS2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_DS3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.509,
            "burstable": false,
            "instanceType": "Standard_DS4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_DS5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_DS11-1_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_DS11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12-1_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12-2_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13-2_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13-4_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14-4_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14-8_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.662,
            "burstable": false,
            "instanceType": "Standard_DS15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_DS2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_DS3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.509,
            "burstable": false,
            "instanceType": "Standard_DS4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_DS5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_DS11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_DS12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_DS13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.33,
            "burstable": false,
            "instanceType": "Standard_DS14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1s",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4s",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.438,
            "burstable": false,
            "instanceType": "Standard_F8s",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_F16s",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 0.2,
            "generation": "current",
            "price": 0.025,
            "burstable": true,
            "instanceType": "Standard_B1ms",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 0.1,
            "generation": "current",
            "price": 0.013,
            "burstable": true,
            "instanceType": "Standard_B1s",
            "memory": 1.0,
            "cpu": 1
        },
        {
            "baseline": 0.6,
            "generation": "current",
            "price": 0.0998,
            "burstable": true,
            "instanceType": "Standard_B2ms",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 0.4,
            "generation": "current",
            "price": 0.0499,
            "burstable": true,
            "instanceType": "Standard_B2s",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 0.9,
            "generation": "current",
            "price": 0.2,
            "burstable": true,
            "instanceType": "Standard_B4ms",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.35,
            "generation": "current",
            "price": 0.399,
            "burstable": true,
            "instanceType": "Standard_B8ms",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D2s_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_D4s_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D8s_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_D16s_v3",
            "memory": 64.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.76,
            "burstable": false,
            "instanceType": "Standard_D32s_v3",
            "memory": 128.0,
            "cpu": 32
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.018,
            "burstable": false,
            "instanceType": "Standard_A0",
            "memory": 0.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.051,
            "burstable": false,
            "instanceType": "Standard_A1",
            "memory": 1.75,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.101,
            "burstable": false,
            "instanceType": "Standard_A2",
            "memory": 3.5,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.176,
            "burstable": false,
            "instanceType": "Standard_A3",
            "memory": 7.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_A5",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.352,
            "burstable": false,
            "instanceType": "Standard_A4",
            "memory": 14.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_A6",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_A7",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.064,
            "burstable": false,
            "instanceType": "Standard_D1_v2",
            "memory": 3.5,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_D2_v2",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_D3_v2",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D4_v2",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_D5_v2",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_D11_v2",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_D12_v2",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_D13_v2",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D14_v2",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.662,
            "burstable": false,
            "instanceType": "Standard_D15_v2",
            "memory": 140.0,
            "cpu": 20
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.127,
            "burstable": false,
            "instanceType": "Standard_D2_v2_Promo",
            "memory": 7.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.254,
            "burstable": false,
            "instanceType": "Standard_D3_v2_Promo",
            "memory": 14.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D4_v2_Promo",
            "memory": 28.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.017,
            "burstable": false,
            "instanceType": "Standard_D5_v2_Promo",
            "memory": 56.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.166,
            "burstable": false,
            "instanceType": "Standard_D11_v2_Promo",
            "memory": 14.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.332,
            "burstable": false,
            "instanceType": "Standard_D12_v2_Promo",
            "memory": 28.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.665,
            "burstable": false,
            "instanceType": "Standard_D13_v2_Promo",
            "memory": 56.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.17,
            "burstable": false,
            "instanceType": "Standard_D14_v2_Promo",
            "memory": 112.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.055,
            "burstable": false,
            "instanceType": "Standard_F1",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.109,
            "burstable": false,
            "instanceType": "Standard_F2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.219,
            "burstable": false,
            "instanceType": "Standard_F4",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.438,
            "burstable": false,
            "instanceType": "Standard_F8",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.875,
            "burstable": false,
            "instanceType": "Standard_F16",
            "memory": 32.0,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.04,
            "burstable": false,
            "instanceType": "Standard_A1_v2",
            "memory": 2.0,
            "cpu": 1
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.119,
            "burstable": false,
            "instanceType": "Standard_A2m_v2",
            "memory": 16.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.083,
            "burstable": false,
            "instanceType": "Standard_A2_v2",
            "memory": 4.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.249,
            "burstable": false,
            "instanceType": "Standard_A4m_v2",
            "memory": 32.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.175,
            "burstable": false,
            "instanceType": "Standard_A4_v2",
            "memory": 8.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.524,
            "burstable": false,
            "instanceType": "Standard_A8m_v2",
            "memory": 64.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.367,
            "burstable": false,
            "instanceType": "Standard_A8_v2",
            "memory": 16.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.11,
            "burstable": false,
            "instanceType": "Standard_D2_v3",
            "memory": 8.0,
            "cpu": 2
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.22,
            "burstable": false,
            "instanceType": "Standard_D4_v3",
            "memory": 16.0,
            "cpu": 4
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.44,
            "burstable": false,
            "instanceType": "Standard_D8_v3",
            "memory": 32.0,
            "cpu": 8
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 0.88,
            "burstable": false,
            "instanceType": "Standard_D16_v3",
            "memory": 64.1,
            "cpu": 16
        },
        {
            "baseline": 1.0,
            "generation": "current",
            "price": 1.76,
            "burstable": false,
            "instanceType": "Standard_D32_v3",
            "memory": 128.0,
            "cpu": 32
        }
    ]
}
`