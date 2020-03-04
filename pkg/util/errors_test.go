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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	e := fmt.Errorf("1")
	we1 := WrapError(e, "we1")
	assert.Equal(t, "we1: 1", we1.Error())
	innerError := we1.(WrappedError).Cause()
	assert.Equal(t, "1", innerError.Error())

	// make sure that having no message doesn't make strange formatting
	we2 := WrapError(e, "")
	assert.Equal(t, "1", we2.Error())
}
