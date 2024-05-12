// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpcclient

import (
	"context"

	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/system/program"
	"github.com/Meikwei/protocol/push"
	"google.golang.org/grpc"
)

type Push struct {
	conn   grpc.ClientConnInterface
	Client push.PushMsgServiceClient
	discov discovery.SvcDiscoveryRegistry
}

func NewPush(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) *Push {
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err)
	}
	return &Push{
		discov: discov,
		conn:   conn,
		Client: push.NewPushMsgServiceClient(conn),
	}
}

type PushRpcClient Push

func NewPushRpcClient(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) PushRpcClient {
	return PushRpcClient(*NewPush(discov, rpcRegisterName))
}

func (p *PushRpcClient) DelUserPushToken(ctx context.Context, req *push.DelUserPushTokenReq) (*push.DelUserPushTokenResp, error) {
	return p.Client.DelUserPushToken(ctx, req)
}