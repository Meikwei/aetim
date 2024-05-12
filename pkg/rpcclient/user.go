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
	"strings"

	"github.com/Meikwei/aetim/pkg/authverify"
	"github.com/Meikwei/aetim/pkg/common/servererrs"
	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/system/program"
	"github.com/Meikwei/go-tools/utils/datautil"
	"github.com/Meikwei/protocol/sdkws"
	"github.com/Meikwei/protocol/user"
	"google.golang.org/grpc"
)

// User 结构体用于存储 User RPC 客户端的连接详情。
type User struct {
	conn                  grpc.ClientConnInterface // gRPC连接接口，用于管理连接。
	Client                user.UserClient          // User RPC服务客户端，用于调用服务方法。
	Discov                discovery.SvcDiscoveryRegistry // 服务发现注册表，用于发现服务实例。
	MessageGateWayRpcName string                   // 消息网关RPC服务名称。
	imAdminUserID         []string                 // IM管理员用户ID列表。
}

// NewUser 根据提供的服务发现注册表初始化并返回一个User实例。
// discov: 服务发现注册表。
// rpcRegisterName: RPC服务注册名称。
// messageGateWayRpcName: 消息网关的RPC服务名称。
// imAdminUserID: 管理员用户ID列表。
func NewUser(discov discovery.SvcDiscoveryRegistry, rpcRegisterName, messageGateWayRpcName string,
	imAdminUserID []string) *User {
	// 使用服务发现注册表获取gRPC连接。
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err)
	}
	// 创建User服务客户端。
	client := user.NewUserClient(conn)
	return &User{Discov: discov, Client: client,
		conn:                  conn,
		MessageGateWayRpcName: messageGateWayRpcName,
		imAdminUserID:         imAdminUserID}
}

// UserRpcClient 表示User RPC客户端的结构体。
type UserRpcClient User

// NewUserRpcClientByUser 根据提供的User实例初始化一个UserRpcClient。
// user: 要基于的User实例。
func NewUserRpcClientByUser(user *User) *UserRpcClient {
	rpc := UserRpcClient(*user)
	return &rpc
}

// NewUserRpcClient 根据提供的服务发现注册表初始化一个UserRpcClient。
// client: 服务发现注册表。
// rpcRegisterName: RPC服务注册名称。
// imAdminUserID: 管理员用户ID列表。
func NewUserRpcClient(client discovery.SvcDiscoveryRegistry, rpcRegisterName string,
	imAdminUserID []string) UserRpcClient {
	// 利用NewUser初始化User实例，并转换为UserRpcClient。
	return UserRpcClient(*NewUser(client, rpcRegisterName, "", imAdminUserID))
}

// GetUsersInfo 根据用户ID列表获取多个用户的信息。
// ctx: 请求的上下文。
// userIDs: 需要获取信息的用户ID列表。
// 返回用户信息切片及错误。
func (u *UserRpcClient) GetUsersInfo(ctx context.Context, userIDs []string) ([]*sdkws.UserInfo, error) {
	// 如果没有提供用户ID，直接返回空切片。
	if len(userIDs) == 0 {
		return []*sdkws.UserInfo{}, nil
	}
	// 调用RPC服务获取指定用户信息。
	resp, err := u.Client.GetDesignateUsers(ctx, &user.GetDesignateUsersReq{
		UserIDs: userIDs,
	})
	if err != nil {
		return nil, err
	}
	// 检查响应中是否缺少请求的用户ID，并处理错误。
	if ids := datautil.Single(userIDs, datautil.Slice(resp.UsersInfo, func(e *sdkws.UserInfo) string {
		return e.UserID
	})); len(ids) > 0 {
		return nil, servererrs.ErrUserIDNotFound.WrapMsg(strings.Join(ids, ","))
	}
	return resp.UsersInfo, nil
}

// GetUserInfo 根据提供的用户ID获取单个用户的信息。
// ctx: 请求的上下文。
// userID: 需要获取信息的用户ID。
// 返回用户信息对象及错误。
func (u *UserRpcClient) GetUserInfo(ctx context.Context, userID string) (*sdkws.UserInfo, error) {
	// 调用GetUsersInfo获取单个用户信息并返回。
	users, err := u.GetUsersInfo(ctx, []string{userID})
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

// GetUsersInfoMap 根据用户ID列表获取用户信息，并以map形式返回。
// ctx: 请求的上下文。
// userIDs: 需要获取信息的用户ID列表。
// 返回按用户ID索引的用户信息映射及错误。
func (u *UserRpcClient) GetUsersInfoMap(ctx context.Context, userIDs []string) (map[string]*sdkws.UserInfo, error) {
	// 转换用户信息切片为映射。
	users, err := u.GetUsersInfo(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return datautil.SliceToMap(users, func(e *sdkws.UserInfo) string {
		return e.UserID
	}), nil
}

// GetPublicUserInfos 根据用户ID列表获取多个用户的公开信息。
// ctx: 请求的上下文。
// userIDs: 需要获取公开信息的用户ID列表。
// complete: 是否获取完整的公开信息。
// 返回公开用户信息切片及错误。
func (u *UserRpcClient) GetPublicUserInfos(
	ctx context.Context,
	userIDs []string,
	complete bool,
) ([]*sdkws.PublicUserInfo, error) {
	// 获取用户信息并转换为公开用户信息。
	users, err := u.GetUsersInfo(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return datautil.Slice(users, func(e *sdkws.UserInfo) *sdkws.PublicUserInfo {
		return &sdkws.PublicUserInfo{
			UserID:   e.UserID,
			Nickname: e.Nickname,
			FaceURL:  e.FaceURL,
			Ex:       e.Ex,
		}
	}), nil
}

// GetPublicUserInfo 根据提供的用户ID获取单个用户的公开信息。
// ctx: 请求的上下文。
// userID: 需要获取公开信息的用户ID。
// 返回公开用户信息对象及错误。
func (u *UserRpcClient) GetPublicUserInfo(ctx context.Context, userID string) (*sdkws.PublicUserInfo, error) {
	// 调用GetPublicUserInfos获取单个用户的公开信息并返回。
	users, err := u.GetPublicUserInfos(ctx, []string{userID}, true)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

// GetPublicUserInfoMap 根据用户ID列表获取用户的公开信息，并以map形式返回。
// ctx: 请求的上下文。
// userIDs: 需要获取公开信息的用户ID列表。
// complete: 是否获取完整的公开信息。
// 返回按用户ID索引的公开用户信息映射及错误。
func (u *UserRpcClient) GetPublicUserInfoMap(
	ctx context.Context,
	userIDs []string,
	complete bool,
) (map[string]*sdkws.PublicUserInfo, error) {
	// 转换公开用户信息切片为映射。
	users, err := u.GetPublicUserInfos(ctx, userIDs, complete)
	if err != nil {
		return nil, err
	}
	return datautil.SliceToMap(users, func(e *sdkws.PublicUserInfo) string {
		return e.UserID
	}), nil
}


// GetUserGlobalMsgRecvOpt 根据提供的用户ID获取用户的全局消息接收选项。
// 参数:
// - ctx: 用于取消和设置截止时间的上下文。
// - userID: 用户的唯一标识符。
// 返回:
// - int32: 用户的全局消息接收选项。
// - 错误: 如果检索失败，返回错误。
func (u *UserRpcClient) GetUserGlobalMsgRecvOpt(ctx context.Context, userID string) (int32, error) {
	resp, err := u.Client.GetGlobalRecvMessageOpt(ctx, &user.GetGlobalRecvMessageOptReq{
		UserID: userID,
	})
	if err != nil {
		return 0, err
	}
	return resp.GlobalRecvMsgOpt, nil
}

// Access 验证给定用户ID的访问权限。
// 参数:
// - ctx: 用于取消和设置截止时间的上下文。
// - ownerUserID: 需要验证访问权限的用户ID。
// 返回:
// - 错误: 如果权限检查失败，返回错误。
func (u *UserRpcClient) Access(ctx context.Context, ownerUserID string) error {
	_, err := u.GetUserInfo(ctx, ownerUserID)
	if err != nil {
		return err
	}
	return authverify.CheckAccessV3(ctx, ownerUserID, u.imAdminUserID)
}

// GetAllUserIDs 分页获取所有用户ID。
// 参数:
// - ctx: 用于取消和设置截止时间的上下文。
// - pageNumber: 分页的页码。
// - showNumber: 每页显示的用户数。
// 返回:
// - []string: 用户ID列表。
// - 错误: 如果检索失败，返回错误。
func (u *UserRpcClient) GetAllUserIDs(ctx context.Context, pageNumber, showNumber int32) ([]string, error) {
	resp, err := u.Client.GetAllUserID(ctx, &user.GetAllUserIDReq{Pagination: &sdkws.RequestPagination{PageNumber: pageNumber, ShowNumber: showNumber}})
	if err != nil {
		return nil, err
	}
	return resp.UserIDs, nil
}

// SetUserStatus 根据提供的用户ID、状态和平台ID设置用户的状态。
// 参数:
// - ctx: 用于取消和设置截止时间的上下文。
// - userID: 用户的唯一标识符。
// - status: 用户的新状态。
// - platformID: 用户的平台标识符。
// 返回:
// - 错误: 如果设置用户状态失败，返回错误。
func (u *UserRpcClient) SetUserStatus(ctx context.Context, userID string, status int32, platformID int) error {
	_, err := u.Client.SetUserStatus(ctx, &user.SetUserStatusReq{
		UserID: userID,
		Status: status, PlatformID: int32(platformID),
	})
	return err
}

// GetNotificationByID 根据用户ID获取用户的提醒账户详情。
// 参数:
// - ctx: 用于取消和设置截止时间的上下文。
// - userID: 用户的唯一标识符。
// 返回:
// - 错误: 如果检索失败，返回错误。
func (u *UserRpcClient) GetNotificationByID(ctx context.Context, userID string) error {
	_, err := u.Client.GetNotificationAccount(ctx, &user.GetNotificationAccountReq{
		UserID: userID,
	})
	return err
}
