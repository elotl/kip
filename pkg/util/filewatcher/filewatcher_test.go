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

package filewatcher

import (
	"testing"
	"time"

	"github.com/elotl/kip/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestStatPeriodRefresh(t *testing.T) {
	t.Parallel()
	f, closer := util.MakeTempFile("filewatcher")
	defer closer()

	c1 := "contents1\n"
	_, err := f.Write([]byte(c1))
	assert.NoError(t, err)
	fw := New(f.Name())
	assert.Equal(t, c1, fw.Contents())
	c2 := "more data\n"
	// Filesystem time is only accurate to 1s
	time.Sleep(1250 * time.Millisecond)
	_, err = f.Write([]byte(c2))
	assert.NoError(t, err)
	assert.Equal(t, c1, fw.Contents())
	fw.CheckPeriod = 1 * time.Nanosecond
	assert.Equal(t, c1+c2, fw.Contents())
}
