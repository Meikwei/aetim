/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-04 16:56:56
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-11 19:52:26
 * @FilePath: \user\pkg\common\kdisc\direct\direct_resolver.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package direct

import (
	"context"
	"math/rand"
	"strings"

	"github.com/Meikwei/go-tools/log"
	"google.golang.org/grpc/resolver"
)

// 定义常量
const (
	slashSeparator = "/" // 文件路径分隔符
	// EndpointSepChar 是端点中的分隔符字符。
	EndpointSepChar = ','

	subsetSize = 32 // 子集大小
	scheme     = "direct" // 使用的方案
)

type ResolverDirect struct {
}

// NewResolverDirect 创建一个新的ResolverDirect实例。
//
// 参数: 无
//
// 返回值: 返回一个初始化的*ResolverDirect指针。
func NewResolverDirect() *ResolverDirect {
	return &ResolverDirect{}
}

// Build 创建一个新的ResolverDirect实例，并基于给定的目标信息初始化解析器的地址。
//
// target: 表示服务的目标信息，用于获取服务的端点。
// cc: 客户端连接，用于后续更新解析器的状态。
// _: BuildOptions参数，此处未使用。
//
// 返回解析器实例和可能发生的错误。如果成功，错误为nil。
func (rd *ResolverDirect) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (
	resolver.Resolver, error) {
	// 记录Build方法的调用信息
	log.ZDebug(context.Background(), "Build", "target", target)

	// 使用特定字符分割目标字符串获取端点列表
	endpoints := strings.FieldsFunc(GetEndpoints(target), func(r rune) bool {
		return r == EndpointSepChar
	})

	// 对端点列表进行子集化处理，只保留指定数量的端点
	endpoints = subset(endpoints, subsetSize)

	// 初始化地址列表
	addrs := make([]resolver.Address, 0, len(endpoints))

	// 将端点字符串转换为resolver.Address类型并添加到地址列表中
	for _, val := range endpoints {
		addrs = append(addrs, resolver.Address{
			Addr: val,
		})
	}

	// 更新客户端连接的状态为新的地址列表
	if err := cc.UpdateState(resolver.State{
		Addresses: addrs,
	}); err != nil {
		return nil, err // 如果更新状态失败，返回错误
	}

	// 返回一个不执行任何操作的解析器实例和nil错误
	return &nopResolver{cc: cc}, nil
}
// init 函数是在程序启动时调用的初始化函数。
// 该函数无参数，也不返回任何值。
// 在该函数内，通过 resolver.Register 注册了一个 ResolverDirect 实例。
func init() {
	resolver.Register(&ResolverDirect{})
}
func (rd *ResolverDirect) Scheme() string {
	return scheme // return your custom scheme name
}

// GetEndpoints returns the endpoints from the given target.
func GetEndpoints(target resolver.Target) string {
	return strings.Trim(target.URL.Path, slashSeparator)
}
func subset(set []string, sub int) []string {
	rand.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})
	if len(set) <= sub {
		return set
	}

	return set[:sub]
}

type nopResolver struct {
	cc resolver.ClientConn
}

func (n nopResolver) ResolveNow(options resolver.ResolveNowOptions) {

}

func (n nopResolver) Close() {

}
