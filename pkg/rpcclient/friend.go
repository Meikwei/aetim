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
	"github.com/Meikwei/protocol/friend"
	sdkws "github.com/Meikwei/protocol/sdkws"
	"google.golang.org/grpc"
)

// Friend 结构体代表一个朋友服务客户端，封装了与朋友服务通过gRPC交互所需的所有组件。
type Friend struct {
	conn   grpc.ClientConnInterface // gRPC通信的连接接口。
	Client friend.FriendClient      // FriendClient 是生成的用于朋友服务的gRPC客户端。
	discov discovery.SvcDiscoveryRegistry // 用于服务发现的功能注册表。
}

// NewFriend 使用发现的连接初始化一个新的Friend结构体实例，并创建必要的gRPC客户端。
//
// 参数:
// - discov: 服务发现功能的SvcDiscoveryRegistry实例。
// - rpcRegisterName: 要注册或发现的服务名称。
//
// 返回:
// - 初始化的Friend结构体指针。
func NewFriend(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) *Friend {
	// 获取与朋友服务的gRPC连接。
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err) // 如果获取连接时出错，退出程序。
	}
	// 使用建立的连接创建新的FriendClient。
	client := friend.NewFriendClient(conn)
	return &Friend{discov: discov, conn: conn, Client: client}
}

// FriendRpcClient 是Friend结构体的别名，用于RPC客户端操作。
type FriendRpcClient Friend

// NewFriendRpcClient 初始化一个新FriendRpcClient实例，封装了Friend结构体的朋友服务操作功能。
//
// 参数:
// - discov: 服务发现功能的SvcDiscoveryRegistry实例。
// - rpcRegisterName: 要注册或发现的服务名称。
//
// 返回:
// - 初始化的FriendRpcClient实例。
func NewFriendRpcClient(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) FriendRpcClient {
	return FriendRpcClient(*NewFriend(discov, rpcRegisterName))
}

// GetFriendsInfo 获取指定朋友用户ID的友人信息。
//
// 参数:
// - ctx: 取消和截止时间的上下文。
// - ownerUserID: 朋友所有者的用户ID。
// - friendUserID: 朋友的用户ID。
//
// 返回:
// - resp: 友人信息响应的指针。
// - err: 操作期间发生的任何错误。
func (f *FriendRpcClient) GetFriendsInfo(
	ctx context.Context,
	ownerUserID, friendUserID string,
) (*sdkws.FriendInfo, error) {
	// 从朋友服务请求朋友信息。
	r, err := f.Client.GetDesignatedFriends(
		ctx,
		&friend.GetDesignatedFriendsReq{OwnerUserID: ownerUserID, FriendUserIDs: []string{friendUserID}},
	)
	if err != nil {
		return nil, err
	}
	// 提取并返回请求的友人信息。
	resp := r.FriendsInfo[0]
	return resp, nil
}

// IsFriend 检查possibleFriendUserID是否是userID的好友。
//
// 参数:
// - ctx: 取消和截止时间的上下文。
// - possibleFriendUserID: 需要检查是否为好友的用户ID。
// - userID: 正在检查其好友列表的用户ID。
//
// 返回:
// - 如果possibleFriendUserID是userID的好友，返回true。
// - 操作期间发生的任何错误。
func (f *FriendRpcClient) IsFriend(ctx context.Context, possibleFriendUserID, userID string) (bool, error) {
	// 向朋友服务查询possibleFriendUserID是否在userID的好友列表中。
	resp, err := f.Client.IsFriend(ctx, &friend.IsFriendReq{UserID1: userID, UserID2: possibleFriendUserID})
	if err != nil {
		return false, err
	}
	// 返回结果，表示possibleFriendUserID是否是好友。
	return resp.InUser1Friends, nil
}

// GetFriendIDs 获取指定所有者用户ID的友人用户ID列表。
//
// 参数:
// - ctx: 取消和截止时间的上下文。
// - ownerUserID: 需要获取友人ID的用户ID。
//
// 返回:
// - friendIDs: 一个字符串切片，表示友人用户ID。
// - err: 操作期间发生的任何错误。
func (f *FriendRpcClient) GetFriendIDs(ctx context.Context, ownerUserID string) ([]string, error) {
	// 从朋友服务请求友人用户ID列表。
	req := friend.GetFriendIDsReq{UserID: ownerUserID}
	resp, err := f.Client.GetFriendIDs(ctx, &req)
	if err != nil {
		return nil, err
	}
	// 提取并返回友人用户ID列表。
	return resp.FriendIDs, nil
}

// IsBlack 检查possibleBlackUserID是否在userID的黑名单中。
//
// 参数:
// - ctx: 取消和截止时间的上下文。
// - possibleBlackUserID: 需要检查是否在黑名单中的用户ID。
// - userID: 正在检查其黑名单的用户ID。
//
// 返回:
// - 如果possibleBlackUserID在黑名单中，返回true。
// - 操作期间发生的任何错误。
func (b *FriendRpcClient) IsBlack(ctx context.Context, possibleBlackUserID, userID string) (bool, error) {
	// 向朋友服务查询possibleBlackUserID是否在userID的黑名单中。
	r, err := b.Client.IsBlack(ctx, &friend.IsBlackReq{UserID1: possibleBlackUserID, UserID2: userID})
	if err != nil {
		return false, err
	}
	// 返回结果，表示possibleBlackUserID是否在黑名单中。
	return r.InUser2Blacks, nil
}
