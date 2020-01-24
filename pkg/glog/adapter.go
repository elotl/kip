package glog

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/virtual-kubelet/virtual-kubelet/log"
)

const (
	errorKey = "error"
)

type GlogAdapter struct {
	fields         log.Fields
	extraArgsStr   string
	extraFormatStr string
}

func NewGlogAdapter() *GlogAdapter {
	return &GlogAdapter{}
}

func (g *GlogAdapter) update() {
	g.extraArgsStr = ""
	for k, v := range g.fields {
		g.extraArgsStr += fmt.Sprintf(" %s=%v", k, v)
	}
	g.extraFormatStr = ""
	for k, v := range g.fields {
		g.extraFormatStr += fmt.Sprintf(" %s=%v", k, v)
	}
}

func (g *GlogAdapter) getArgs(args ...interface{}) []interface{} {
	if len(g.extraArgsStr) > 0 {
		return append(args, g.extraArgsStr)
	}
	return args
}

func (g *GlogAdapter) getFormat(format string) string {
	return format + g.extraFormatStr
}

func (g *GlogAdapter) Debug(args ...interface{}) {
	args = g.getArgs(args...)
	if glog.V(4) {
		glog.InfoDepth(1, args...)
	}
}

func (g *GlogAdapter) Debugf(format string, args ...interface{}) {
	format = g.getFormat(format)
	if glog.V(4) {
		glog.InfoDepth(1, fmt.Sprintf(format, args...))
	}
}

func (g *GlogAdapter) Info(args ...interface{}) {
	args = g.getArgs(args...)
	glog.InfoDepth(1, args...)
}

func (g *GlogAdapter) Infof(format string, args ...interface{}) {
	format = g.getFormat(format)
	glog.InfoDepth(1, fmt.Sprintf(format, args...))
}

func (g *GlogAdapter) Warn(args ...interface{}) {
	args = g.getArgs(args...)
	glog.WarningDepth(1, args...)
}

func (g *GlogAdapter) Warnf(format string, args ...interface{}) {
	format = g.getFormat(format)
	glog.WarningDepth(1, fmt.Sprintf(format, args...))
}

func (g *GlogAdapter) Error(args ...interface{}) {
	args = g.getArgs(args...)
	glog.ErrorDepth(1, args...)
}

func (g *GlogAdapter) Errorf(format string, args ...interface{}) {
	format = g.getFormat(format)
	glog.ErrorDepth(1, fmt.Sprintf(format, args...))
}

func (g *GlogAdapter) Fatal(args ...interface{}) {
	args = g.getArgs(args...)
	glog.FatalDepth(1, args...)
}

func (g *GlogAdapter) Fatalf(format string, args ...interface{}) {
	format = g.getFormat(format)
	glog.FatalDepth(1, fmt.Sprintf(format, args...))
}

func (g *GlogAdapter) WithField(key string, value interface{}) log.Logger {
	logger := &GlogAdapter{
		fields: map[string]interface{}{
			key: value,
		},
	}
	logger.update()
	return logger
}

func (g *GlogAdapter) WithFields(fields log.Fields) log.Logger {
	logger := &GlogAdapter{
		fields: make(map[string]interface{}),
	}
	for k, v := range fields {
		logger.fields[k] = v
	}
	logger.update()
	return logger
}

func (g *GlogAdapter) WithError(err error) log.Logger {
	logger := &GlogAdapter{
		fields: map[string]interface{}{
			errorKey: err,
		},
	}
	logger.update()
	return logger
}
