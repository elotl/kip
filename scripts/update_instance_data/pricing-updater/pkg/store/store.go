package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elotl/pricing-updater/pkg/convert"
	"io/ioutil"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

var OutputFileNames = map[string]string{
	convert.ProviderAWS: "awsInstanceJson.json",
	convert.ProviderGCE: "gceInstanceJson.json",
	convert.ProviderAzure: "azureInstanceJson.json",
}

func SaveToJson(data map[string][]convert.TargetInstanceInfo) (string, error) {
	f, err := ioutil.TempFile("/tmp", "kip-instance-data-*.json")
	if err != nil {
		return "", err
	}
	err = dumpToJson(data, f)
	return f.Name(), err
}

func dumpToJson(data map[string][]convert.TargetInstanceInfo, file *os.File) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	err = file.Close()
	return err
}

func SaveAllRegions(data map[string][]convert.TargetInstanceInfo, provider string) error {
	f, err := os.Create(OutputFileNames[provider])
	if err != nil {
		return err
	}
	err = dumpToJson(data, f)
	return err
}

func CreateConfigMap(name, namespace string, data map[string][]convert.TargetInstanceInfo) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return fmt.Errorf("cannot load kube config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	configMapClient := clientset.CoreV1().ConfigMaps(namespace)

	marshalledData, err := json.Marshal(data)
	configMap := v1.ConfigMap{
		Data: map[string]string{
			"instance-data.json": string(marshalledData),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	_, err = configMapClient.Create(context.TODO(), &configMap, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("cannot create ConfigMap: %v", err)
	}
	return nil
}
