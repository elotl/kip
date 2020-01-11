package clientapi

import (
	"bytes"
	"strings"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util/yaml"
)

type MockMilpaClient struct {
	GetVersioner        func(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
	Creator             func(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*APIReply, error)
	Updater             func(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*APIReply, error)
	Getter              func(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*APIReply, error)
	Deleter             func(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*APIReply, error)
	GetLogser           func(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*APIReply, error)
	Dumper              func(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*APIReply, error)
	SetupIPForwardinger func(ctx context.Context, in *SetupIPForwardingRequest, opts ...grpc.CallOption) (*APIReply, error)
	Deployer            func(ctx context.Context, opts ...grpc.CallOption) (Milpa_DeployClient, error)
	StreamLogser        func(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (Milpa_StreamLogsClient, error)
	PortForwarder       func(ctx context.Context, opts ...grpc.CallOption) (Milpa_PortForwardClient, error)
	Execer              func(ctx context.Context, opts ...grpc.CallOption) (Milpa_ExecClient, error)
	Attacher            func(ctx context.Context, opts ...grpc.CallOption) (Milpa_AttachClient, error)
	Leader              func(ctx context.Context, in *IsLeaderRequest, opts ...grpc.CallOption) (*IsLeaderReply, error)
}

func (m MockMilpaClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	return m.GetVersioner(ctx, in, opts...)
}

func (m MockMilpaClient) IsLeader(ctx context.Context, in *IsLeaderRequest, opts ...grpc.CallOption) (*IsLeaderReply, error) {
	return m.Leader(ctx, in, opts...)
}

func (m MockMilpaClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Creator(ctx, in, opts...)
}

func (m MockMilpaClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Updater(ctx, in, opts...)
}

func (m MockMilpaClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Getter(ctx, in, opts...)
}

func (m MockMilpaClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Deleter(ctx, in, opts...)
}

func (m MockMilpaClient) GetLogs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.GetLogser(ctx, in, opts...)
}

func (m MockMilpaClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.Dumper(ctx, in, opts...)
}

func (m MockMilpaClient) SetupIPForwarding(ctx context.Context, in *SetupIPForwardingRequest, opts ...grpc.CallOption) (*APIReply, error) {
	return m.SetupIPForwardinger(ctx, in, opts...)
}

func (m MockMilpaClient) Deploy(ctx context.Context, opts ...grpc.CallOption) (Milpa_DeployClient, error) {
	return m.Deployer(ctx, opts...)
}

func (m MockMilpaClient) StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (Milpa_StreamLogsClient, error) {
	return m.StreamLogser(ctx, in, opts...)
}

func (m MockMilpaClient) PortForward(ctx context.Context, opts ...grpc.CallOption) (Milpa_PortForwardClient, error) {
	return m.PortForwarder(ctx, opts...)
}

func (m MockMilpaClient) Exec(ctx context.Context, opts ...grpc.CallOption) (Milpa_ExecClient, error) {
	return m.Execer(ctx, opts...)
}

func (m MockMilpaClient) Attach(ctx context.Context, opts ...grpc.CallOption) (Milpa_AttachClient, error) {
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

func NewMockMilpaClient() MockMilpaClient {
	encoder := api.IndentingJsonCodec{}
	// Currently only pods are supported.
	podstore := make(map[string]*api.Pod)
	cli := MockMilpaClient{}
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
