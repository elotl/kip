package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	grpcDialTimeout = 5 * time.Second
)

// Note: this can get called concurrently and cobra.Cmd.InheritedFlags
// is not safe for concurrent access.
func getMilpaClient(flags *pflag.FlagSet, needsLeader bool) (clientapi.MilpaClient, *grpc.ClientConn, error) {
	endpoints, err := flags.GetStringSlice("endpoints")
	if err != nil {
		return nil, nil, util.WrapError(err, "Error getting endpoints argument")
	}
	// We shuffle endpoints to do some weak loadbalancing
	rand.Seed(time.Now().UTC().UnixNano())
	order := rand.Perm(len(endpoints))
	for i, _ := range order {
		address := endpoints[i]
		var (
			client clientapi.MilpaClient
			conn   *grpc.ClientConn
		)
		client, conn, err = connectToServer(context.Background(), address, flags)
		if err != nil {
			// If we got an error with that server, just continue
			// trying other servers
			continue
		}

		if !needsLeader {
			return client, conn, nil
		} else {
			var foundLeader bool
			foundLeader, err = isLeader(context.Background(), client)
			if foundLeader {
				return client, conn, nil
			} else if err == nil {
				err = fmt.Errorf("leader required")
			}
		}
		_ = conn.Close()
	}
	msg := "Could not connect to the Milpa API server: " + err.Error()
	return nil, nil, fmt.Errorf(msg)
}

func isLeader(ctx context.Context, client clientapi.MilpaClient) (bool, error) {
	req := clientapi.IsLeaderRequest{}
	reply, err := client.IsLeader(ctx, &req)
	if err != nil {
		return false, fmt.Errorf("Error querying milpa server: %v", err.Error())
	}
	return reply.IsLeader, nil
}

func connectToServer(ctx context.Context, serverAddress string, flags *pflag.FlagSet) (clientapi.MilpaClient, *grpc.ClientConn, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, grpcDialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(
		timeoutCtx, serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}
	return clientapi.NewMilpaClient(conn), conn, nil
}
