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

	"k8s.io/klog"
)

type WrappedError struct {
	msg   string
	cause error
}

func (e WrappedError) Error() string {
	return e.msg
}

func (e WrappedError) Cause() error {
	return e.cause
}

func WrapError(err error, format string, args ...interface{}) error {
	s := fmt.Sprintf(format, args...)
	var msg string
	if err != nil {
		if s != "" {
			msg = s + ": " + err.Error()
		} else {
			msg = err.Error()
		}
	} else {
		klog.Errorln("WrapError: nil error:", s)
		msg = s
	}
	if we, ok := err.(WrappedError); ok {
		we.msg = msg
		return we
	} else {
		return WrappedError{
			msg:   msg,
			cause: err,
		}
	}
}
