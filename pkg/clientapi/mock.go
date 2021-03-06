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

package clientapi

import (
	"bytes"
	"strings"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util/yaml"
)

type MockKipClient struct {
	GetVersioner func(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
	Creator      func(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*APIReply, error)
	Updater      func(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*APIReply, error)
	Getter       func(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*APIReply, error)
	Deleter      func(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*APIReply, error)
	GetLogser    func(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*APIReply, error)
	Dumper       func(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*APIReply, error)
	Deployer     func(ctx context.Context, opts ...grpc.CallOption) (Kip_DeployClient, error)
	StreamLogser func(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (Kip_StreamLogsClient, error)
	Execer       func(ctx context.Context, opts ...grpc.CallOption) (Kip_ExecClient, error)
	Attacher     func(ctx context.Context, opts ...grpc.CallOption) (Kip_AttachClient, error)
	Leader       func(ctx context.Context, in *IsLeaderRequest, opts ...grpc.CallOption) (*IsLeaderReply, error)
}

func (m MockKipClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	return m.GetVersioner(ctx, in, opts...)
}

func (m MockKipClient) IsLeader(ctx context.Context, in *IsLeaderRequest, opts ...grpc.CallOption) (*IsLeaderReply, error) {
	return m.Leader(ctx, in, opts...)
}

func (m MockKipClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Creator(ctx, in, opts...)
}

func (m MockKipClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Updater(ctx, in, opts...)
}

func (m MockKipClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Getter(ctx, in, opts...)
}

func (m MockKipClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Deleter(ctx, in, opts...)
}

func (m MockKipClient) GetLogs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.GetLogser(ctx, in, opts...)
}

func (m MockKipClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Dumper(ctx, in, opts...)
}

func (m MockKipClient) Deploy(ctx context.Context, opts ...grpc.CallOption) (Kip_DeployClient, error) {
	return m.Deployer(ctx, opts...)
}

func (m MockKipClient) StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (Kip_StreamLogsClient, error) {
	return m.StreamLogser(ctx, in, opts...)
}

func (m MockKipClient) Exec(ctx context.Context, opts ...grpc.CallOption) (Kip_ExecClient, error) {
	return m.Execer(ctx, opts...)
}

func (m MockKipClient) Attach(ctx context.Context, opts ...grpc.CallOption) (Kip_AttachClient, error) {
	return m.Attacher(ctx, opts...)
}

func errorReply(msg string) *APIReply {
	return &APIReply{
		Status: 400,
		Body:   []byte(msg),
	}
}

func versionAndKind(m []byte) (string, string, error) {
	var typeMeta api.TypeMeta
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(m), 8000)
	err := decoder.Decode(&typeMeta)
	if err != nil {
		return "", "", err
	}
	return typeMeta.APIVersion, strings.ToLower(typeMeta.Kind), nil
}

func NewMockKipClient() MockKipClient {
	encoder := api.IndentingJsonCodec{}
	// Currently only pods are supported.
	podstore := make(map[string]*api.Pod)
	cli := MockKipClient{}
	cli.Creator = func(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*APIReply, error) {
		_, objectKind, err := versionAndKind(in.Manifest)
		if err != nil {
			return errorReply("Error determining manifest kind"), nil
		}
		if objectKind != "pod" {
			return errorReply("Unsupported object kind"), nil
		}
		pod := api.NewPod()
		decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(in.Manifest), 8000)
		err = decoder.Decode(pod)
		if err != nil {
			return errorReply("Error loading manifest"), nil
		}
		body, err := encoder.Marshal(pod)
		if err != nil {
			return errorReply("Error creating reply object"), nil
		}
		podstore[pod.Name] = pod
		reply := APIReply{
			Status: 201,
			Body:   body,
		}
		return &reply, nil
	}
	cli.Updater = func(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*APIReply, error) {
		req := &CreateRequest{
			Manifest: in.Manifest,
		}
		return cli.Creator(ctx, req, opts...)
	}
	cli.Getter = func(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*APIReply, error) {
		name := string(in.Name)
		var body []byte
		if name != "" {
			var err error
			pod, exists := podstore[name]
			if !exists {
				return errorReply("No such object"), nil
			}
			body, err = encoder.Marshal(pod)
			if err != nil {
				return errorReply("Error creating reply object"), nil
			}
		} else {
			var err error
			podlist := api.NewPodList()
			for _, pod := range podstore {
				podlist.Items = append(podlist.Items, pod)
			}
			body, err = encoder.Marshal(podlist)
			if err != nil {
				return errorReply("Error creating reply object"), nil
			}
		}
		reply := APIReply{
			Status: 200,
			Body:   body,
		}
		return &reply, nil
	}
	cli.Deleter = func(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*APIReply, error) {
		name := string(in.Name)
		_, exists := podstore[name]
		if !exists {
			return errorReply("No such object"), nil
		}
		delete(podstore, name)
		reply := APIReply{
			Status: 200,
			Body:   []byte(""),
		}
		return &reply, nil
	}
	cli.Leader = func(ctx context.Context, in *IsLeaderRequest, opts ...grpc.CallOption) (*IsLeaderReply, error) {
		return &IsLeaderReply{IsLeader: true}, nil
	}
	return cli
}
