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
	"github.com/elotl/kip/pkg/api"
	"k8s.io/kubernetes/third_party/forked/golang/expansion"
)

func EnvVarsToMap(envs []api.EnvVar) map[string]string {
	result := map[string]string{}
	for _, env := range envs {
		result[env.Name] = env.Value
	}
	return result
}

func ExpandCommandAndArgs(spec api.PodSpec) api.PodSpec {
	for i, unit := range spec.Units {
		mapping := expansion.MappingFuncFor(EnvVarsToMap(unit.Env))
		if len(unit.Command) != 0 {
			command := make([]string, len(unit.Command))
			for j, cmd := range unit.Command {
				command[j] = expansion.Expand(cmd, mapping)
			}
			spec.Units[i].Command = command
		}
		if len(unit.Args) != 0 {
			args := make([]string, len(unit.Args))
			for j, arg := range unit.Args {
				args[j] = expansion.Expand(arg, mapping)
			}
			spec.Units[i].Args = args
		}
	}
	return spec
}
