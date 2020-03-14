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

package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/util"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	grpcDialTimeout = 5 * time.Second
)

// Note: this can get called concurrently and cobra.Cmd.InheritedFlags
// is not safe for concurrent access.
func getKipClient(flags *pflag.FlagSet, needsLeader bool) (clientapi.KipClient, *grpc.ClientConn, error) {
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
			client clientapi.KipClient
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
	msg := "Could not connect to the Kip API server: " + err.Error()
	return nil, nil, fmt.Errorf(msg)
}

func isLeader(ctx context.Context, client clientapi.KipClient) (bool, error) {
	req := clientapi.IsLeaderRequest{}
	reply, err := client.IsLeader(ctx, &req)
	if err != nil {
		return false, fmt.Errorf("Error querying kip server: %v", err.Error())
	}
	return reply.IsLeader, nil
}

func connectToServer(ctx context.Context, serverAddress string, flags *pflag.FlagSet) (clientapi.KipClient, *grpc.ClientConn, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, grpcDialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(
		timeoutCtx, serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}
	return clientapi.NewKipClient(conn), conn, nil
}
