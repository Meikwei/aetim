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
	"github.com/Meikwei/protocol/third"
	"google.golang.org/grpc"
)

// Third 结构体用于维护与第三方服务交互所需的信息
type Third struct {
	conn       grpc.ClientConnInterface // conn 代表与第三方服务的gRPC连接
	Client     third.ThirdClient        // Client 是通过conn创建的第三方服务的客户端
	discov     discovery.SvcDiscoveryRegistry // discov 用于服务发现，帮助获取与第三方服务的连接
	GrafanaUrl string                   // GrafanaUrl 存储Grafana的URL，用于监控等用途
}

// NewThird 创建并初始化Third结构体实例
// discov: 用于服务发现的实例，帮助动态获取与第三方服务的连接
// rpcRegisterName: 第三方服务在服务发现中的注册名称
// grafanaUrl: Grafana的访问URL
// 返回值: 初始化后的Third结构体指针
func NewThird(discov discovery.SvcDiscoveryRegistry, rpcRegisterName, grafanaUrl string) *Third {
	// 通过服务发现获取与第三方服务的gRPC连接
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err) // 如果获取连接失败，则退出程序
	}
	// 使用获取的连接创建第三方服务的客户端
	client := third.NewThirdClient(conn)
	if err != nil {
		program.ExitWithError(err) // 如果创建客户端失败，则退出程序
	}
	// 返回初始化后的Third结构体实例
	return &Third{discov: discov, Client: client, conn: conn, GrafanaUrl: grafanaUrl}
}
