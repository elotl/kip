package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/elotl/pricing-updater/pkg/cloudinfo"
	"github.com/elotl/pricing-updater/pkg/convert"
	"github.com/elotl/pricing-updater/pkg/store"
	"log"
	"os"
)

func scrapeAndConvert(provider, region string, skipRegionValidation bool) ([]convert.TargetInstanceInfo, error) {
	url, err := cloudinfo.ValidateURL(provider, region, skipRegionValidation)
	if err != nil {
		return nil, fmt.Errorf("cannot build url: %v", err)
	}
	var respStruct convert.CloudinfoResponse
	respBody, err := cloudinfo.GetResponseBody(url)

	if err != nil {
		return nil, fmt.Errorf("cannot get response: %v", err)
	}
	err = json.Unmarshal(respBody, &respStruct)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal response: %v", err)
	}
	outputData, err := convert.CloudInfoRespToKipFormat(respStruct)
	if err != nil {
		return nil, fmt.Errorf("cannot convert data to KIP format: %v", err)
	}
	return outputData, nil
}

func scrapeAllProvidersAndRegions() {
	for _, provider := range cloudinfo.SupportedProviders {
		providerRegions, err := cloudinfo.GetSupportedRegions(provider)
		if err != nil {
			log.Fatalf("failed getting supported regions for %s provider: %v", provider, err)
		}
		providerData := make(map[string][]convert.TargetInstanceInfo, 0)
		for _, region := range providerRegions {
			regionData, err := scrapeAndConvert(provider, region, true)
			if err != nil {
				log.Printf("error getting region %s data for provider %s: %v", region, provider, err)
			}
			providerData[region] = regionData
		}
		err = store.SaveAllRegions(providerData, provider)
		if err != nil {
			log.Fatalf("cannot save instance data for provider %s : %v", provider, err)
		}
	}
	log.Println("successfully saved data for all providers.")
	os.Exit(0)
}

func main() {
	// those flags can be used if pricing-updater runs as a CronJob in cluster
	provider := flag.String("provider", convert.ProviderAWS, fmt.Sprintf("provider name. Supported: %v", cloudinfo.SupportedProviders))
	region := flag.String("region", "", "region for a given provider")
	configmap := flag.String("configmap", "kip-instance-data", "target config map name which will be created and populated with instance pricing data")
	namespace := flag.String("namespace", "default", "target config map namespace")
	// this flag is used if pricing-updater need to generate .go files with instance data on KIP build.
	scrapeAll := flag.Bool("scrape-all", false, "setting this flag will scrape all supported providers and regions (and ignore provider and region flags)")

	flag.Parse()
	if scrapeAll != nil && *scrapeAll {
		scrapeAllProvidersAndRegions()
	}
	regionData, err := scrapeAndConvert(*provider, *region, false)
	if err != nil {
		log.Fatal(err)
	}
	outputData := map[string][]convert.TargetInstanceInfo{*region: regionData}
	fileName, err := store.SaveToJson(outputData)
	if err != nil {
		log.Fatalf("cannot save data to json: %v", err)
	}
	log.Printf("Data saved to: %s", fileName)
	err = store.CreateConfigMap(*configmap, *namespace, outputData)
	if err != nil {
		log.Fatal(err)
	}
}
