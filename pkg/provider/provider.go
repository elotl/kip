package provider

// type InstanceProvider struct {
// 	nodeName           string
// 	operatingSystem    string
// 	internalIP         string
// 	daemonEndpointPort int32
// 	config             Config
// 	startTime          time.Time
// 	notifier           func(*v1.Pod)
// }

// func NewProvider(providerConfigPath, nodeName, operatingSystem string, internalIP string, daemonEndpointPort int32) (*InstanceProvider, error) {
// 	config, err := loadConfig(providerConfigPath, nodeName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	provider := InstanceProvider{
// 		nodeName:           nodeName,
// 		operatingSystem:    operatingSystem,
// 		internalIP:         internalIP,
// 		daemonEndpointPort: daemonEndpointPort,
// 		//pods:               make(map[string]*v1.Pod),
// 		config:    config,
// 		startTime: time.Now(),
// 	}
// 	return &provider, nil
// }
