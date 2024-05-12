/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-27 17:02:46
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 21:43:11
 * @FilePath: \open-im-server\pkg\common\cmd\api.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

package cmd

import (
	"context"

	"github.com/Meikwei/aetim/internal/api"
	"github.com/Meikwei/aetim/pkg/common/config"
	"github.com/Meikwei/go-tools/system/program"
	"github.com/spf13/cobra"
)

// ApiCmd 定义了 API 命令结构体，封装了 API 相关的配置和上下文信息。
type ApiCmd struct {
	*RootCmd                // RootCmd 是基础命令结构体，提供了一些通用的功能。
	ctx       context.Context // ctx 是 API 命令执行的上下文，用于传递请求生命周期中的信息。
	configMap map[string]any  // configMap 用于存储不同类型的配置信息，通过文件名映射到具体的配置结构体字段。
	apiConfig *api.Config     // apiConfig 存储了 API 配置，包括 RPC、Zookeeper 和 Share 配置等。
}

// NewApiCmd 创建并返回一个新的 ApiCmd 实例。
// 这个函数初始化了 ApiCmd 结构体，并设置了相关的配置和上下文。
func NewApiCmd() *ApiCmd {
	var apiConfig api.Config
	ret := &ApiCmd{apiConfig: &apiConfig} // 初始化 ApiCmd 结构体。

	// 配置映射，将不同的配置文件名映射到 apiConfig 的相应字段。
	ret.configMap = map[string]any{
		OpenIMAPICfgFileName:    &apiConfig.RpcConfig,
		ZookeeperConfigFileName: &apiConfig.ZookeeperConfig,
		ShareFileName:           &apiConfig.Share,
	}

	// 设置 RootCmd，包括程序名称和配置映射。
	ret.RootCmd = NewRootCmd(program.GetProcessName(), WithConfigMap(ret.configMap))

	// 创建并设置上下文，传递版本信息。
	ret.ctx = context.WithValue(context.Background(), "version", config.Version)

	// 设置命令的执行函数。
	ret.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return ret.runE()
	}
	return ret
}

// Exec 启动 API 命令的执行。
// 它实际上是调用了 Execute 方法，实现了命令执行的入口点。
func (a *ApiCmd) Exec() error {
	return a.Execute()
}

// runE 是命令执行的实际逻辑函数。
// 它调用了 api.Start 函数来启动 API 服务。
func (a *ApiCmd) runE() error {
	return api.Start(a.ctx, a.Index(), a.apiConfig)
}
