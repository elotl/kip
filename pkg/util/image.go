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
	"fmt"
	"strings"
)

// Most registry URLs don't have a leading scheme (e.g. http://).
// This means that we can't use url.Parse to find the host name from
// the image's path.  We'll use some heuristics to get the server and
// image path.
func ParseImageSpec(image string) (string, string, error) {
	server := ""
	imageRepo := image
	var err error
	parts := strings.Split(image, "/")
	// ECS: ACCOUNT.dkr.ecr.REGION.amazonaws.com/imagename:tag
	if len(parts) == 1 {
		// all good
	} else if len(parts) >= 2 && strings.HasSuffix(parts[0], "amazonaws.com") {
		server = parts[0]
		imageRepo = strings.Join(parts[1:], "/")
	} else if len(parts) == 2 {
		//assumes docker registry: user/registry:tag
		server = ""
	} else if len(parts) >= 3 {
		// everything else: server.tld/user/registry:tag
		server = parts[0]
		imageRepo = strings.Join(parts[1:], "/")
	} else {
		err = fmt.Errorf("Unknown image registry format")
	}
	return server, imageRepo, err
}
