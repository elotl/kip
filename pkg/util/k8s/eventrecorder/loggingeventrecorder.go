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
