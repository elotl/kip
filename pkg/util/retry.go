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
	"time"
)

var sleepDelay = time.Duration(1 * time.Second)

// This got written after a few Old Rasputins.  I highly recommend the beer but
// not as much as I recommend reading this code with a critical eye
func Retry(timeout time.Duration, f func() error, isRetryable func(error) bool) error {
	start := time.Now()
	timeoutTime := start.Add(timeout)
	var err error
	for {
		err = f()
		if err == nil {
			return nil
		}
		if !isRetryable(err) {
			return err
		}

		if time.Since(start) > timeout {
			break
		}

		sleepDelay = time.Duration(float64(sleepDelay) * 1.2)
		now := time.Now()
		nextWakeup := now.Add(sleepDelay)
		if nextWakeup.After(timeoutTime) {
			sleepDelay = timeoutTime.Sub(now)
		}
		time.Sleep(sleepDelay)
	}
	return WrapError(err, "Timed out retrying, last error")
}
