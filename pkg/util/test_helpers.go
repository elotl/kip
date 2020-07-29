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

func GCEEnvVarsSet() bool {
	if os.Getenv("GCE_CLIENT_EMAIL") == "" ||
		os.Getenv("GCE_PRIVATE_KEY") == "" {
		return false
	}
	return true
}
