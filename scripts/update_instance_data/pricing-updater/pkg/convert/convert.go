package convert

import (
	"fmt"

	"github.com/gobwas/glob"
)

var (
	customBaselines = map[string]float64{
		"t2.2xlarge": 1.359,
		"t2.xlarge":  0.9,
		"t2.large":   0.6,
		"t2.medium":  0.4,
		"t2.small":   0.2,
		"t2.micro":   0.1,
		"t2.nano":    0.05,

		"t3.2xlarge": 3.2,
		"t3.xlarge":  1.6,
		"t3.large":   0.6,
		"t3.medium":  0.4,
		"t3.small":   0.4,
		"t3.micro":   0.2,
		"t3.nano":    0.1,

		"t3a.2xlarge": 3.2,
		"t3a.xlarge":  1.6,
		"t3a.large":   0.6,
		"t3a.medium":  0.4,
		"t3a.small":   0.4,
		"t3a.micro":   0.2,
		"t3a.nano":    0.1,
	}
	// https://cloud.google.com/compute/docs/machine-types#cpu-bursting
	gceBurstableTypes        = []string{"f1-micro", "g1-small", "e2-micro", "e2-small", "e2-medium"}
	unsupportedInstanceTypes = []glob.Glob{
		glob.MustCompile("a1.*"),
	}
)

func isUnsupportedInstanceType(instanceType string) bool {
	for _, g := range unsupportedInstanceTypes {
		if g.Match(instanceType) {
			return true
		}
	}
	return false
}

func CloudInfoRespToKipFormat(resp CloudinfoResponse) ([]TargetInstanceInfo, error) {
	var regionPricing []TargetInstanceInfo
	for _, product := range resp.Products {
		if isUnsupportedInstanceType(product.Type) {
			continue
		}
		generation := "current"
		//if !product.CurrentGeneration {
		//	// TODO - do we want to support previous generation instances?
		//  // this actually
		//	continue
		//}
		regionPricing = append(regionPricing, TargetInstanceInfo{
			Baseline:      getBaseline(product),
			Generation:    generation,
			OnDemandPrice: product.OnDemandPrice,
			SpotPrice:     getLowestSpotPrice(product),
			Memory:        product.Memory,
			InstanceType:  product.Type,
			Burstable:     getBurstable(product),
			CPU:           product.Cpu,
			GPU:           product.GPU,
		})
	}
	return regionPricing, nil
}

func getBaseline(instance InstanceInfo) float64 {
	if baseline, ok := customBaselines[instance.Type]; ok {
		return baseline
	}
	return float64(instance.Cpu)
}

func getBurstable(instance InstanceInfo) bool {
	burstable := instance.Burstable
	if !instance.Burstable {
		burstable = false
	}
	for _, instanceName := range gceBurstableTypes {
		if instance.Type == instanceName {
			burstable = true
			fmt.Printf("set %s as burstable\n", instance.Type)
		}
	}
	return burstable
}

func getLowestSpotPrice(instance InstanceInfo) float64 {
	lowest := instance.OnDemandPrice
	for _, zoneInfo := range instance.SpotPrices {
		if zoneInfo.Price < lowest {
			lowest = zoneInfo.Price
		}
	}
	return lowest
}
