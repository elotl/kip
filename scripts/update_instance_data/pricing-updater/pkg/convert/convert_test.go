package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloudInfoRespToKipFormat(t *testing.T) {
	testCases := []struct {
		name         string
		response     CloudinfoResponse
		region       string
		expectedData map[string][]TargetInstanceInfo
	}{
		{
			name: "happy path",
			response: CloudinfoResponse{Products: []InstanceInfo{
				{
					Type:          "dummy",
					OnDemandPrice: 0.01,
					SpotPrices: []ZonePriceInfo{
						{
							"us-east-1a",
							0.002,
						},
						{
							"us-east-1b",
							0.003,
						},
					},
					Cpu:               1,
					Memory:            2,
					CurrentGeneration: true,
					Burstable: false,
				},
			}},
			region: "us-east-1",
			expectedData: map[string][]TargetInstanceInfo{
				"us-east-1": {
					{
						Baseline:      1,
						Generation:    "current",
						OnDemandPrice: 0.01,
						SpotPrice:     0.002,
						Memory:        2,
						InstanceType:  "dummy",
						Burstable:     false,
						CPU:           1,
						GPU:           0,
					},
				},
			},
		},
		{
			name: "custom baseline instance type",
			response: CloudinfoResponse{Products: []InstanceInfo{
				{
					Type:          "t3.micro", // we have static list of instances with baseline != vCpu
					OnDemandPrice: 0.01,
					SpotPrices: []ZonePriceInfo{
						{
							"us-east-1a",
							0.002,
						},
						{
							"us-east-1b",
							0.003,
						},
					},
					Cpu:               1,
					Memory:            2,
					CurrentGeneration: true,
					Burstable: true,
				},
			}},
			region: "us-east-1",
			expectedData: map[string][]TargetInstanceInfo{
				"us-east-1": {
					{
						Baseline:      0.2,
						Generation:    "current",
						OnDemandPrice: 0.01,
						SpotPrice:     0.002,
						Memory:        2,
						InstanceType:  "t3.micro",
						Burstable:     true,
						CPU:           1,
						GPU:           0,
					},
				},
			},
		},
		{
			name: "no spot price",
			response: CloudinfoResponse{Products: []InstanceInfo{
				{
					Type:          "dummy",
					OnDemandPrice: 0.01,
					SpotPrices: []ZonePriceInfo{}, // in such case, spotPrice == onDemand.
					Cpu:               1,
					Memory:            2,
					CurrentGeneration: true,
					Burstable: false,
				},
			}},
			region: "us-east-1",
			expectedData: map[string][]TargetInstanceInfo{
				"us-east-1": {
					{
						Baseline:      1,
						Generation:    "current",
						OnDemandPrice: 0.01,
						SpotPrice:     0.01,
						Memory:        2,
						InstanceType:  "dummy",
						Burstable:     false,
						CPU:           1,
						GPU:           0,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			regionData, err := CloudInfoRespToKipFormat(testCase.response)
			got := map[string][]TargetInstanceInfo{testCase.region: regionData}
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedData, got)
		})
	}
}
