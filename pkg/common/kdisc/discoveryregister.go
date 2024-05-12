/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-04 17:52:12
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-04 17:56:16
 * @FilePath: \user\pkg\common\kdisc\discoveryregister.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package kdisc

import (
	"time"

	"github.com/Meikwei/aetim/pkg/common/config"
	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/discovery/zookeeper"
	"github.com/Meikwei/go-tools/errs"
)

const (
	zookeeperConst = "zookeeper"
	kubenetesConst = "k8s"
	directConst    = "direct"
)

// NewDiscoveryRegister 创建一个基于提供的环境类型的服务发现和注册客户端。
//
// 参数:
// - zookeeperConfig: 包含ZooKeeper配置信息的结构体指针。
// - share: 包含共享配置信息的结构体指针。
//
// 返回值:
// - discovery.SvcDiscoveryRegistry: 实现了服务发现和注册接口的客户端实例。
// - error: 在创建客户端时遇到错误时返回的错误信息。
func NewDiscoveryRegister(zookeeperConfig *config.ZooKeeper, share *config.Share) (discovery.SvcDiscoveryRegistry, error) {
	switch share.Env {
	case zookeeperConst:
		// 创建一个ZooKeeper客户端实例
		return zookeeper.NewZkClient(
			zookeeperConfig.Address,
			zookeeperConfig.Schema,
			zookeeper.WithFreq(time.Hour),
			zookeeper.WithUserNameAndPassword(zookeeperConfig.Username, zookeeperConfig.Password),
			zookeeper.WithRoundRobin(),
			zookeeper.WithTimeout(10),
		)
	//case directConst:
	// 返回一个直接连接客户端实例（当前未实现）
	default:
		// 如果不支持的服务发现类型，则返回错误
		return nil, errs.New("unsupported discovery type", "type", share.Env).Wrap()
	}
}
