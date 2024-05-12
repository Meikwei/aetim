/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-11 19:36:53
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 22:21:33
 * @FilePath: \user\pkg\rpcclient\auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package rpcclient

import (
	"context"

	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/system/program"
	"github.com/Meikwei/protocol/auth"
	pbAuth "github.com/Meikwei/protocol/auth"
	"google.golang.org/grpc"
)

// NewAuth 创建一个新的Auth实例。
//
// 参数:
// - discov: 服务发现注册接口，用于获取RPC服务的连接。
// - rpcRegisterName: RPC服务的注册名称。
//
// 返回值:
// - 返回一个初始化好的Auth指针。
func NewAuth(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) *Auth {
    // 通过服务发现获取与RPC服务的连接
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err) // 如果获取连接失败，则退出程序
	}
	client := auth.NewAuthClient(conn) // 使用连接创建Auth客户端
	return &Auth{discov: discov, conn: conn, Client: client}
}

// Auth 定义了Auth结构体，存储与认证服务相关的客户端连接和实例。
type Auth struct {
	conn   grpc.ClientConnInterface // RPC连接
	Client auth.AuthClient         // Auth客户端
	discov discovery.SvcDiscoveryRegistry // 服务发现注册接口
}

// ParseToken 解析令牌。
//
// 参数:
// - ctx: 上下文，用于控制请求的生命周期。
// - token: 需要解析的令牌字符串。
//
// 返回值:
// - 返回解析令牌后的响应信息和可能的错误。
func (a *Auth) ParseToken(ctx context.Context, token string) (*pbAuth.ParseTokenResp, error) {
	req := pbAuth.ParseTokenReq{
		Token: token, // 设置请求的令牌字段
	}
	resp, err := a.Client.ParseToken(ctx, &req) // 调用ParseToken方法
	if err != nil {
		return nil, err
	}
	return resp, err
}

// InvalidateToken 使令牌失效。
//
// 参数:
// - ctx: 上下文，用于控制请求的生命周期。
// - preservedToken: 需要保留的令牌。
// - userID: 用户ID。
// - platformID: 平台ID。
//
// 返回值:
// - 返回使令牌失效操作的响应信息和可能的错误。
func (a *Auth) InvalidateToken(ctx context.Context, preservedToken, userID string, platformID int) (*pbAuth.InvalidateTokenResp, error) {
	req := pbAuth.InvalidateTokenReq{
		PreservedToken: preservedToken, // 设置请求的保留令牌字段
		UserID:         userID,         // 设置请求的用户ID字段
		PlatformID:     int32(platformID), // 设置请求的平台ID字段
	}
	resp, err := a.Client.InvalidateToken(ctx, &req) // 调用InvalidateToken方法
	if err != nil {
		return nil, err
	}
	return resp, err
}


