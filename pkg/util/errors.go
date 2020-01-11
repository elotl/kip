package util

import (
	"fmt"

	"github.com/golang/glog"
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
		glog.Errorln("WrapError: nil error:", s)
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
