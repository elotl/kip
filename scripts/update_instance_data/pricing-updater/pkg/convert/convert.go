package convert

var (
	customBaselines = map[string]float64{
		"t2.2xlarge": 1.359,
		"t2.xlarge": 0.9,
		"t2.large": 0.6,
		"t2.medium": 0.4,
		"t2.small": 0.2,
		"t2.micro": 0.1,
		"t2.nano": 0.05,

		"t3.2xlarge": 3.2,
		"t3.xlarge": 1.6,
		"t3.large": 0.6,
		"t3.medium": 0.4,
		"t3.small": 0.4,
		"t3.micro": 0.2,
		"t3.nano": 0.1,
	}
)

func CloudInfoRespToKipFormat(resp CloudinfoResponse) ([]TargetInstanceInfo, error)  {
	var regionPricing []TargetInstanceInfo
	for _, product := range resp.Products {
		burstable := product.Burstable
		if !product.Burstable {
			burstable = false
		}
		generation := "previous"
		if product.CurrentGeneration {
			generation = "current"
		}
		regionPricing = append(regionPricing, TargetInstanceInfo{
			Baseline:      getBaseline(product),
			Generation:    generation,
			OnDemandPrice: product.OnDemandPrice,
			SpotPrice:     getLowestSpotPrice(product),
			Memory:        product.Memory,
			InstanceType:  product.Type,
			Burstable:     burstable,
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

func getLowestSpotPrice(instance InstanceInfo) float64 {
	lowest := instance.OnDemandPrice
	for _, zoneInfo := range instance.SpotPrices {
		if zoneInfo.Price < lowest {
			lowest = zoneInfo.Price
		}
	}
	return lowest
}