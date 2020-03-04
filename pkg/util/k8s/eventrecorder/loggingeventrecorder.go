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

package eventrecorder

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
)

type LoggingEventRecorder struct {
	level klog.Level
}

func NewLoggingEventRecorder(loglevel int) record.EventRecorder {
	return &LoggingEventRecorder{level: klog.Level(loglevel)}
}

func (l *LoggingEventRecorder) Event(object runtime.Object, eventtype, reason, message string) {
	klog.V(l.level).Infof("event %q reason: %q msg: %q on %+v",
		eventtype, reason, message, object)
}

func (l *LoggingEventRecorder) Eventf(object runtime.Object, eventtype, reason, messageFmt string, args ...interface{}) {
	message := fmt.Sprintf(messageFmt, args...)
	klog.V(l.level).Infof("event %q reason: %q msg: %q on %+v",
		eventtype, reason, message, object)
}

func (l *LoggingEventRecorder) PastEventf(object runtime.Object, timestamp v1.Time, eventtype, reason, messageFmt string, args ...interface{}) {
	message := fmt.Sprintf(messageFmt, args...)
	klog.V(l.level).Infof("event %q reason: %q msg: %q on %+v timestamp %v",
		eventtype, reason, message, object, timestamp)
}

func (l *LoggingEventRecorder) AnnotatedEventf(object runtime.Object, annotations map[string]string, eventtype, reason, messageFmt string, args ...interface{}) {
	message := fmt.Sprintf(messageFmt, args...)
	klog.V(l.level).Infof("event %q reason: %q msg: %q on %+v %v",
		eventtype, reason, message, object, annotations)
}
