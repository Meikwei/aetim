/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-27 17:02:46
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-01 12:07:04
 * @FilePath: \open-im-server\pkg\common\discoveryregister\zookeeper\zookeeper.go
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

package zookeeper

import (
	"os"
	"strings"
)

// getEnv 从环境变量中获取指定键的值。
// 如果键存在，则返回其值；如果键不存在，则返回 fallback 提供的默认值。
//
// 参数:
// key - 指定的环境变量键名。
// fallback - 当指定的环境变量不存在时，返回的默认值。
//
// 返回值:
// 获取到的环境变量值或默认值（如果环境变量不存在）。
func getEnv(key, fallback string) string {
	// 尝试从环境变量中获取键对应的值
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	// 如果键不存在于环境变量中，返回默认值
	return fallback
}

// getZkAddrFromEnv returns the Zookeeper addresses combined from the ZOOKEEPER_ADDRESS and ZOOKEEPER_PORT environment variables.
// getZkAddrFromEnv 从环境变量中获取 Zookeeper 地址列表。
// 如果环境变量 ZOOKEEPER_ADDRESS 和 ZOOKEEPER_PORT 都设置好了，就使用它们的值来构造地址列表；
// 如果有一个或两个环境变量没有设置，那么就返回 fallback 参数作为地址列表。
// 参数:
//   fallback []string - 当环境变量未设置时，返回的默认地址列表。
// 返回值:
//   []string - Zookeeper 地址列表。
func getZkAddrFromEnv(fallback []string) []string {
	// 从环境变量中获取 Zookeeper 的地址和端口
	address, addrExists := os.LookupEnv("ZOOKEEPER_ADDRESS")
	port, portExists := os.LookupEnv("ZOOKEEPER_PORT")

	// 当地址和端口环境变量都存在时，处理并返回它们
	if addrExists && portExists {
		addresses := strings.Split(address, ",")
		for i, addr := range addresses {
			// 组合地址和端口
			addresses[i] = addr + ":" + port
		}
		return addresses
	}
	// 如果环境变量不完整，返回默认地址列表
	return fallback
}
