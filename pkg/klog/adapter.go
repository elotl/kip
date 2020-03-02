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

package klog

import (
	"flag"
	"fmt"

	"github.com/virtual-kubelet/virtual-kubelet/log"
	"k8s.io/klog"
)

const (
	errorKey = "error"
)

type KlogAdapter struct {
	fields         log.Fields
	extraArgsStr   string
	extraFormatStr string
}

func NewKlogAdapter() *KlogAdapter {
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	return &KlogAdapter{}
}

func (g *KlogAdapter) update() {
	g.extraArgsStr = ""
	for k, v := range g.fields {
		g.extraArgsStr += fmt.Sprintf(" %s=%v", k, v)
	}
	g.extraFormatStr = ""
	for k, v := range g.fields {
		g.extraFormatStr += fmt.Sprintf(" %s=%v", k, v)
	}
}

func (g *KlogAdapter) getArgs(args ...interface{}) []interface{} {
	if len(g.extraArgsStr) > 0 {
		return append(args, g.extraArgsStr)
	}
	return args
}

func (g *KlogAdapter) getFormat(format string) string {
	return format + g.extraFormatStr
}

func (g *KlogAdapter) Debug(args ...interface{}) {
	args = g.getArgs(args...)
	if klog.V(4) {
		klog.InfoDepth(1, args...)
	}
}

func (g *KlogAdapter) Debugf(format string, args ...interface{}) {
	format = g.getFormat(format)
	if klog.V(4) {
		klog.InfoDepth(1, fmt.Sprintf(format, args...))
	}
}

func (g *KlogAdapter) Info(args ...interface{}) {
	args = g.getArgs(args...)
	klog.InfoDepth(1, args...)
}

func (g *KlogAdapter) Infof(format string, args ...interface{}) {
	format = g.getFormat(format)
	klog.InfoDepth(1, fmt.Sprintf(format, args...))
}

func (g *KlogAdapter) Warn(args ...interface{}) {
	args = g.getArgs(args...)
	klog.WarningDepth(1, args...)
}

func (g *KlogAdapter) Warnf(format string, args ...interface{}) {
	format = g.getFormat(format)
	klog.WarningDepth(1, fmt.Sprintf(format, args...))
}

func (g *KlogAdapter) Error(args ...interface{}) {
	args = g.getArgs(args...)
	klog.ErrorDepth(1, args...)
}

func (g *KlogAdapter) Errorf(format string, args ...interface{}) {
	format = g.getFormat(format)
	klog.ErrorDepth(1, fmt.Sprintf(format, args...))
}

func (g *KlogAdapter) Fatal(args ...interface{}) {
	args = g.getArgs(args...)
	klog.FatalDepth(1, args...)
}

func (g *KlogAdapter) Fatalf(format string, args ...interface{}) {
	format = g.getFormat(format)
	klog.FatalDepth(1, fmt.Sprintf(format, args...))
}

func (g *KlogAdapter) WithField(key string, value interface{}) log.Logger {
	logger := &KlogAdapter{
		fields: map[string]interface{}{
			key: value,
		},
	}
	logger.update()
	return logger
}

func (g *KlogAdapter) WithFields(fields log.Fields) log.Logger {
	logger := &KlogAdapter{
		fields: make(map[string]interface{}),
	}
	for k, v := range fields {
		logger.fields[k] = v
	}
	logger.update()
	return logger
}

func (g *KlogAdapter) WithError(err error) log.Logger {
	logger := &KlogAdapter{
		fields: map[string]interface{}{
			errorKey: err,
		},
	}
	logger.update()
	return logger
}
