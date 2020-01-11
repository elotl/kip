package util

import (
	"io/ioutil"
	"os"
)

func MakeTempFile(prefix string) (*os.File, func()) {
	tempFile, err := ioutil.TempFile("", prefix)
	if err != nil {
		panic(err)
	}
	return tempFile, func() { os.Remove(tempFile.Name()) }
}

func MakeTempFileName(prefix string) (string, func()) {
	f, closer := MakeTempFile(prefix)
	return f.Name(), closer
}

func AWSEnvVarsSet() bool {
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" ||
		os.Getenv("AWS_SECRET_ACCESS_KEY") == "" ||
		os.Getenv("AWS_REGION") == "" {
		return false
	}
	return true
}

func AzureEnvVarsSet() bool {
	if os.Getenv("AZURE_TENANT_ID") == "" ||
		os.Getenv("AZURE_CLIENT_ID") == "" ||
		os.Getenv("AZURE_CLIENT_SECRET") == "" {
		return false
	}
	return true
}
