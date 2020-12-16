package cloudinfo

import (
	"encoding/json"
	"fmt"
	"github.com/elotl/pricing-updater/pkg/convert"
	"io/ioutil"
	"net/http"
)

const (
	RegionsUrlPattern = "https://banzaicloud.com/cloudinfo/api/v1/providers/%s/services/compute/regions/"
	UrlPattern = "https://banzaicloud.com/cloudinfo/api/v1/providers/%s/services/compute/regions/%s/products"
)

var (
	SupportedProviders = []string{
		convert.ProviderAWS,
		//convert.ProviderGCE,
		convert.ProviderAzure,
	}
)


func GetResponseBody(url string) ([]byte, error) {
	httpClient := http.Client{}
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot get response: %v", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %v", err)
	}
	return body, nil
}

func GetSupportedRegions(provider string) (convert.RegionResp, error) {
	url := fmt.Sprintf(RegionsUrlPattern, provider)
	var respStruct convert.RegionResp
	respBody, err := GetResponseBody(url)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respBody, &respStruct)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal response: %v", err)
	}
	return respStruct, nil

}

func ValidateURL(provider, region string, skipRegionValidation bool) (string, error) {
	// validate provider
	providerIsValid := false
	for _, providerName := range SupportedProviders {
		if providerName == provider {
			providerIsValid = true
			break
		}
	}
	if !providerIsValid {
		return "", fmt.Errorf("%s provider is not supported. Supported providers: %s", provider, SupportedProviders)
	}
	if skipRegionValidation {
		return fmt.Sprintf(UrlPattern, provider, region), nil
	}
	supportedRegions, err := GetSupportedRegions(provider)
	if err != nil {
		return "", err
	}
	// validate region
	regionIsValid := false
	for _, regionName := range supportedRegions {
		if regionName.Id == region || regionName.Name == region {
			regionIsValid = true
			break
		}
	}
	if !regionIsValid {
		return "", fmt.Errorf("%s region is not supported. Supported regions for %s provider: %v", region, provider, supportedRegions)
	}
	return fmt.Sprintf(UrlPattern, provider, region), nil
}
