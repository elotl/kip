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

package azure

import "github.com/Azure/go-autorest/autorest"

func isNotFoundError(err error) bool {
	if detailedError, ok := err.(autorest.DetailedError); ok {
		intSC, ok := detailedError.StatusCode.(int)
		return ok && intSC == 404
	}
	return false
}
