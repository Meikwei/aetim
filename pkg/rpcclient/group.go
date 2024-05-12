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

	"github.com/Meikwei/aetim/pkg/common/servererrs"
	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/system/program"
	"github.com/Meikwei/go-tools/utils/datautil"
	"github.com/Meikwei/protocol/constant"
	"github.com/Meikwei/protocol/group"
	"github.com/Meikwei/protocol/sdkws"
)

// Group结构体定义了与群组服务交互所需的基本组件
type Group struct {
	Client group.GroupClient
	discov discovery.SvcDiscoveryRegistry
}

// NewGroup创建一个新的Group实例
// discov: 服务发现注册
// rpcRegisterName: RPC服务注册名称
// 返回值: 初始化的Group实例指针
func NewGroup(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) *Group {
	// 通过服务发现获取连接
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err)
	}
	// 使用连接创建群组客户端
	client := group.NewGroupClient(conn)
	return &Group{discov: discov, Client: client}
}

// GroupRpcClient是对Group结构的别名，用于定义不同的方法集合
type GroupRpcClient Group

// NewGroupRpcClient创建一个新的GroupRpcClient实例
// discov: 服务发现注册
// rpcRegisterName: RPC服务注册名称
// 返回值: 初始化的GroupRpcClient实例
func NewGroupRpcClient(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) GroupRpcClient {
	return GroupRpcClient(*NewGroup(discov, rpcRegisterName))
}

// GetGroupInfos获取群组信息列表
// ctx: 上下文
// groupIDs: 群组ID列表
// complete: 是否要求完整信息
// 返回值: 群组信息列表和错误信息
func (g *GroupRpcClient) GetGroupInfos(ctx context.Context, groupIDs []string, complete bool) ([]*sdkws.GroupInfo, error) {
	// 调用群组服务获取信息
	resp, err := g.Client.GetGroupsInfo(ctx, &group.GetGroupsInfoReq{
		GroupIDs: groupIDs,
	})
	if err != nil {
		return nil, err
	}
	// 如果需要完整信息，进行完整性检查
	if complete {
		if ids := datautil.Single(groupIDs, datautil.Slice(resp.GroupInfos, func(e *sdkws.GroupInfo) string {
			return e.GroupID
		})); len(ids) > 0 {
			return nil, servererrs.ErrGroupIDNotFound.WrapMsg(strings.Join(ids, ","))
		}
	}
	return resp.GroupInfos, nil
}

// GetGroupInfo获取单个群组信息
// ctx: 上下文
// groupID: 群组ID
// 返回值: 群组信息和错误信息
func (g *GroupRpcClient) GetGroupInfo(ctx context.Context, groupID string) (*sdkws.GroupInfo, error) {
	groups, err := g.GetGroupInfos(ctx, []string{groupID}, true)
	if err != nil {
		return nil, err
	}
	return groups[0], nil
}

// GetGroupInfoMap获取群组信息的映射表
// ctx: 上下文
// groupIDs: 群组ID列表
// complete: 是否要求完整信息
// 返回值: 群组ID到群组信息的映射表和错误信息
func (g *GroupRpcClient) GetGroupInfoMap(
	ctx context.Context,
	groupIDs []string,
	complete bool,
) (map[string]*sdkws.GroupInfo, error) {
	groups, err := g.GetGroupInfos(ctx, groupIDs, complete)
	if err != nil {
		return nil, err
	}
	return datautil.SliceToMap(groups, func(e *sdkws.GroupInfo) string {
		return e.GroupID
	}), nil
}

// GetGroupMemberInfos获取群组成员信息列表
// ctx: 上下文
// groupID: 群组ID
// userIDs: 用户ID列表
// complete: 是否要求完整信息
// 返回值: 群组成员信息列表和错误信息
func (g *GroupRpcClient) GetGroupMemberInfos(
	ctx context.Context,
	groupID string,
	userIDs []string,
	complete bool,
) ([]*sdkws.GroupMemberFullInfo, error) {
	resp, err := g.Client.GetGroupMembersInfo(ctx, &group.GetGroupMembersInfoReq{
		GroupID: groupID,
		UserIDs: userIDs,
	})
	if err != nil {
		return nil, err
	}
	// 如果需要完整信息，进行完整性检查
	if complete {
		if ids := datautil.Single(userIDs, datautil.Slice(resp.Members, func(e *sdkws.GroupMemberFullInfo) string {
			return e.UserID
		})); len(ids) > 0 {
			return nil, servererrs.ErrNotInGroupYet.WrapMsg(strings.Join(ids, ","))
		}
	}
	return resp.Members, nil
}

// GetGroupMemberInfo获取单个群组成员信息
// ctx: 上下文
// groupID: 群组ID
// userID: 用户ID
// 返回值: 群组成员信息和错误信息
func (g *GroupRpcClient) GetGroupMemberInfo(
	ctx context.Context,
	groupID string,
	userID string,
) (*sdkws.GroupMemberFullInfo, error) {
	members, err := g.GetGroupMemberInfos(ctx, groupID, []string{userID}, true)
	if err != nil {
		return nil, err
	}
	return members[0], nil
}

// GetGroupMemberInfoMap获取群组成员信息的映射表
// ctx: 上下文
// groupID: 群组ID
// userIDs: 用户ID列表
// complete: 是否要求完整信息
// 返回值: 用户ID到群组成员信息的映射表和错误信息
func (g *GroupRpcClient) GetGroupMemberInfoMap(
	ctx context.Context,
	groupID string,
	userIDs []string,
	complete bool,
) (map[string]*sdkws.GroupMemberFullInfo, error) {
	members, err := g.GetGroupMemberInfos(ctx, groupID, userIDs, true)
	if err != nil {
		return nil, err
	}
	return datautil.SliceToMap(members, func(e *sdkws.GroupMemberFullInfo) string {
		return e.UserID
	}), nil
}

// GetOwnerAndAdminInfos获取群组的所有者和管理员信息列表
// ctx: 上下文
// groupID: 群组ID
// 返回值: 群组所有者和管理员信息列表和错误信息
func (g *GroupRpcClient) GetOwnerAndAdminInfos(
	ctx context.Context,
	groupID string,
) ([]*sdkws.GroupMemberFullInfo, error) {
	resp, err := g.Client.GetGroupMemberRoleLevel(ctx, &group.GetGroupMemberRoleLevelReq{
		GroupID:    groupID,
		RoleLevels: []int32{constant.GroupOwner, constant.GroupAdmin},
	})
	if err != nil {
		return nil, err
	}
	return resp.Members, nil
}

// GetOwnerInfo获取群组所有者信息
// ctx: 上下文
// groupID: 群组ID
// 返回值: 群组所有者信息和错误信息
func (g *GroupRpcClient) GetOwnerInfo(ctx context.Context, groupID string) (*sdkws.GroupMemberFullInfo, error) {
	resp, err := g.Client.GetGroupMemberRoleLevel(ctx, &group.GetGroupMemberRoleLevelReq{
		GroupID:    groupID,
		RoleLevels: []int32{constant.GroupOwner},
	})
	return resp.Members[0], err
}

// GetGroupMemberIDs获取群组内的成员ID列表
// ctx: 上下文
// groupID: 群组ID
// 返回值: 群组成员ID列表和错误信息
func (g *GroupRpcClient) GetGroupMemberIDs(ctx context.Context, groupID string) ([]string, error) {
	resp, err := g.Client.GetGroupMemberUserIDs(ctx, &group.GetGroupMemberUserIDsReq{
		GroupID: groupID,
	})
	if err != nil {
		return nil, err
	}
	return resp.UserIDs, nil
}

// GetGroupInfoCache获取群组信息的缓存
// ctx: 上下文
// groupID: 群组ID
// 返回值: 群组信息和错误信息
func (g *GroupRpcClient) GetGroupInfoCache(ctx context.Context, groupID string) (*sdkws.GroupInfo, error) {
	resp, err := g.Client.GetGroupInfoCache(ctx, &group.GetGroupInfoCacheReq{
		GroupID: groupID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GroupInfo, nil
}

// GetGroupMemberCache获取群组成员信息的缓存
// ctx: 上下文
// groupID: 群组ID
// groupMemberID: 群组成员ID
// 返回值: 群组成员信息和错误信息
func (g *GroupRpcClient) GetGroupMemberCache(ctx context.Context, groupID string, groupMemberID string) (*sdkws.GroupMemberFullInfo, error) {
	resp, err := g.Client.GetGroupMemberCache(ctx, &group.GetGroupMemberCacheReq{
		GroupID:       groupID,
		GroupMemberID: groupMemberID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Member, nil
}

// DismissGroup解散群组
// ctx: 上下文
// groupID: 群组ID
// 返回值: 错误信息
func (g *GroupRpcClient) DismissGroup(ctx context.Context, groupID string) error {
	_, err := g.Client.DismissGroup(ctx, &group.DismissGroupReq{
		GroupID:      groupID,
		DeleteMember: true,
	})
	return err
}

// NotificationUserInfoUpdate更新用户信息通知
// ctx: 上下文
// userID: 用户ID
// 返回值: 错误信息
func (g *GroupRpcClient) NotificationUserInfoUpdate(ctx context.Context, userID string) error {
	_, err := g.Client.NotificationUserInfoUpdate(ctx, &group.NotificationUserInfoUpdateReq{
		UserID: userID,
	})
	return err
}
