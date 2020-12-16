package convert

const (
	ProviderAWS   = "amazon"
	ProviderGCE   = "google"
	ProviderAzure = "azure"
)

type TargetInstanceInfo struct {
	Baseline      float64 `json:"baseline"`
	Generation    string  `json:"generation"`
	OnDemandPrice float64 `json:"price"`
	SpotPrice     float64 `json:"spotPrice,omitempty"`
	Memory        float64 `json:"memory"`
	InstanceType  string  `json:"instanceType"`
	Burstable     bool    `json:"burstable"`
	CPU           int     `json:"cpu"`
	GPU           int     `json:"gpu"`
}

type ZonePriceInfo struct {
	Zone  string  `json:"zone"`
	Price float64 `json:"price"`
}

type InstanceInfo struct {
	Type              string          `json:"type"`
	OnDemandPrice     float64         `json:"onDemandPrice"`
	SpotPrices        []ZonePriceInfo `json:"spotPrice"`
	Cpu               int             `json:"cpusPerVm"`
	Memory            float64         `json:"memPerVm"` // in GiB
	GPU               int             `json:"gpusPerVm"`
	CurrentGeneration bool            `json:"currentGen"`
	Burstable         bool            `json:"burst,omitempty"`
}

type CloudinfoResponse struct {
	Products []InstanceInfo
}

type RegionInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type RegionResp []RegionInfo
