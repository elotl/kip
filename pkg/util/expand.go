package util

import (
	"github.com/elotl/cloud-instance-provider/pkg/api"
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
