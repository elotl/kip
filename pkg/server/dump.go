package server

import (
	"bytes"
	"fmt"
	"runtime/pprof"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"golang.org/x/net/context"
	"k8s.io/klog"
)

func (s InstanceProvider) dumpController(name string) ([]byte, error) {
	if name == "all" {
		buf := bytes.Buffer{}
		ctlByName := s.controllerManager.GetAllControllers()
		i := 0
		buf.Write([]byte("\n{\n"))
		for name, ctl := range ctlByName {
			header := fmt.Sprintf("\"%s\":", name)
			_, err := buf.Write([]byte(header))
			if err != nil {
				continue
			}
			_, _ = buf.Write(ctl.Dump())
			if i != len(ctlByName)-1 {
				buf.Write([]byte(","))
			}
			buf.Write([]byte("\n"))
			i++
		}
		buf.Write([]byte("}\n"))
		return buf.Bytes(), nil
	} else {
		ctl, exists := s.controllerManager.GetController(name)
		if !exists {
			return nil, fmt.Errorf("Asked to dump unknown controller: %s", name)
		}
		return ctl.Dump(), nil
	}

}

func dumpStack() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := pprof.Lookup("goroutine").WriteTo(buf, 1)
	if err != nil {
		return nil, fmt.Errorf("Could not dump goroutine stacks: %s", err)
	}
	return buf.Bytes(), nil
}

func (s InstanceProvider) Dump(context context.Context, request *clientapi.DumpRequest) (*clientapi.APIReply, error) {
	klog.V(2).Infof("Dump request for: %s", request.Kind)
	kind := string(request.Kind)
	var b []byte
	var err error
	if kind == "stack" {
		b, err = dumpStack()
	} else if kind == "all" {
		b, err = dumpStack()
		if err == nil {
			b2, err2 := s.dumpController("all")
			if err2 == nil {
				b = append(b, b2...)
			}
		}
	} else {
		b, err = s.dumpController(kind)
	}
	if err != nil {
		return errToAPIReply(err), nil
	}
	if b == nil {
		b = []byte("{}")
	}
	reply := clientapi.APIReply{
		Status: 200,
		Body:   b,
	}
	return &reply, nil
}
