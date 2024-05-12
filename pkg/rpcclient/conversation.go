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
	"fmt"

	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/errs"
	"github.com/Meikwei/go-tools/system/program"
	pbconversation "github.com/Meikwei/protocol/conversation"
	"google.golang.org/grpc"
)

// Conversation 是一个用于与对话服务进行交互的结构体。
type Conversation struct {
	Client pbconversation.ConversationClient
	conn   grpc.ClientConnInterface
	discov discovery.SvcDiscoveryRegistry
}

// NewConversation 创建并返回一个新的Conversation实例。
// discov: 服务发现注册表，用于获取与对话服务的连接。
// rpcRegisterName: 注册服务的名称。
// 返回: 初始化后的Conversation实例指针。
func NewConversation(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) *Conversation {
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		program.ExitWithError(err)
	}
	client := pbconversation.NewConversationClient(conn)
	return &Conversation{discov: discov, conn: conn, Client: client}
}

// ConversationRpcClient 是Conversation的一个别名，用于RPC调用。
type ConversationRpcClient Conversation

// NewConversationRpcClient 创建并返回一个ConversationRpcClient实例。
// discov: 服务发现注册表，用于获取与对话服务的连接。
// rpcRegisterName: 注册服务的名称。
// 返回: 初始化后的ConversationRpcClient实例。
func NewConversationRpcClient(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) ConversationRpcClient {
	return ConversationRpcClient(*NewConversation(discov, rpcRegisterName))
}

// GetSingleConversationRecvMsgOpt 获取指定对话的接收消息选项。
// ctx: 上下文。
// userID: 用户ID。
// conversationID: 对话ID。
// 返回: 接收消息选项和错误信息。
func (c *ConversationRpcClient) GetSingleConversationRecvMsgOpt(ctx context.Context, userID, conversationID string) (int32, error) {
	var req pbconversation.GetConversationReq
	req.OwnerUserID = userID
	req.ConversationID = conversationID
	conversation, err := c.Client.GetConversation(ctx, &req)
	if err != nil {
		return 0, err
	}
	return conversation.GetConversation().RecvMsgOpt, err
}

// SingleChatFirstCreateConversation 创建单聊对话。
// ctx: 上下文。
// recvID: 接收者ID。
// sendID: 发送者ID。
// conversationID: 对话ID。
// conversationType: 对话类型。
// 返回: 错误信息。
func (c *ConversationRpcClient) SingleChatFirstCreateConversation(ctx context.Context, recvID, sendID,
	conversationID string, conversationType int32) error {
	_, err := c.Client.CreateSingleChatConversations(ctx,
		&pbconversation.CreateSingleChatConversationsReq{
			RecvID: recvID, SendID: sendID, ConversationID: conversationID,
			ConversationType: conversationType,
		})
	return err
}

// GroupChatFirstCreateConversation 创建群聊对话。
// ctx: 上下文。
// groupID: 群组ID。
// userIDs: 用户ID列表。
// 返回: 错误信息。
func (c *ConversationRpcClient) GroupChatFirstCreateConversation(ctx context.Context, groupID string, userIDs []string) error {
	_, err := c.Client.CreateGroupChatConversations(ctx, &pbconversation.CreateGroupChatConversationsReq{UserIDs: userIDs, GroupID: groupID})
	return err
}

// SetConversationMaxSeq 设置对话的最大序列号。
// ctx: 上下文。
// ownerUserIDs: 拥有者用户ID列表。
// conversationID: 对话ID。
// maxSeq: 最大序列号。
// 返回: 错误信息。
func (c *ConversationRpcClient) SetConversationMaxSeq(ctx context.Context, ownerUserIDs []string, conversationID string, maxSeq int64) error {
	_, err := c.Client.SetConversationMaxSeq(ctx, &pbconversation.SetConversationMaxSeqReq{OwnerUserID: ownerUserIDs, ConversationID: conversationID, MaxSeq: maxSeq})
	return err
}

// SetConversations 设置对话信息。
// ctx: 上下文。
// userIDs: 用户ID列表。
// conversation: 对话信息请求。
// 返回: 错误信息。
func (c *ConversationRpcClient) SetConversations(ctx context.Context, userIDs []string, conversation *pbconversation.ConversationReq) error {
	_, err := c.Client.SetConversations(ctx, &pbconversation.SetConversationsReq{UserIDs: userIDs, Conversation: conversation})
	return err
}

// GetConversationIDs 获取用户的所有对话ID。
// ctx: 上下文。
// ownerUserID: 用户ID。
// 返回: 对话ID列表和错误信息。
func (c *ConversationRpcClient) GetConversationIDs(ctx context.Context, ownerUserID string) ([]string, error) {
	resp, err := c.Client.GetConversationIDs(ctx, &pbconversation.GetConversationIDsReq{UserID: ownerUserID})
	if err != nil {
		return nil, err
	}
	return resp.ConversationIDs, nil
}

// GetConversation 获取指定对话的信息。
// ctx: 上下文。
// ownerUserID: 用户ID。
// conversationID: 对话ID。
// 返回: 对话信息和错误信息。
func (c *ConversationRpcClient) GetConversation(ctx context.Context, ownerUserID, conversationID string) (*pbconversation.Conversation, error) {
	resp, err := c.Client.GetConversation(ctx, &pbconversation.GetConversationReq{OwnerUserID: ownerUserID, ConversationID: conversationID})
	if err != nil {
		return nil, err
	}
	return resp.Conversation, nil
}

// GetConversationsByConversationID 通过对话ID获取对话信息。
// ctx: 上下文。
// conversationIDs: 对话ID列表。
// 返回: 对话信息列表和错误信息。
func (c *ConversationRpcClient) GetConversationsByConversationID(ctx context.Context, conversationIDs []string) ([]*pbconversation.Conversation, error) {
	if len(conversationIDs) == 0 {
		return nil, nil
	}
	resp, err := c.Client.GetConversationsByConversationID(ctx, &pbconversation.GetConversationsByConversationIDReq{ConversationIDs: conversationIDs})
	if err != nil {
		return nil, err
	}
	if len(resp.Conversations) == 0 {
		return nil, errs.ErrRecordNotFound.WrapMsg(fmt.Sprintf("conversationIDs: %v not found", conversationIDs))
	}
	return resp.Conversations, nil
}

// GetConversationOfflinePushUserIDs 获取对话的离线推送用户ID列表。
// ctx: 上下文。
// conversationID: 对话ID。
// userIDs: 用户ID列表。
// 返回: 离线推送的用户ID列表和错误信息。
func (c *ConversationRpcClient) GetConversationOfflinePushUserIDs(ctx context.Context, conversationID string, userIDs []string) ([]string, error) {
	resp, err := c.Client.GetConversationOfflinePushUserIDs(ctx, &pbconversation.GetConversationOfflinePushUserIDsReq{ConversationID: conversationID, UserIDs: userIDs})
	if err != nil {
		return nil, err
	}
	return resp.UserIDs, nil
}

// GetConversations 获取用户指定对话ID列表的对话信息。
// ctx: 上下文。
// ownerUserID: 用户ID。
// conversationIDs: 对话ID列表。
// 返回: 对话信息列表和错误信息。
func (c *ConversationRpcClient) GetConversations(ctx context.Context, ownerUserID string, conversationIDs []string) ([]*pbconversation.Conversation, error) {
	if len(conversationIDs) == 0 {
		return nil, nil
	}
	resp, err := c.Client.GetConversations(
		ctx,
		&pbconversation.GetConversationsReq{OwnerUserID: ownerUserID, ConversationIDs: conversationIDs},
	)
	if err != nil {
		return nil, err
	}
	return resp.Conversations, nil
}

// GetConversationNotReceiveMessageUserIDs 获取对话中未接收消息的用户ID列表。
// ctx: 上下文。
// conversationID: 对话ID。
// 返回: 未接收消息的用户ID列表和错误信息。
func (c *ConversationRpcClient) GetConversationNotReceiveMessageUserIDs(ctx context.Context, conversationID string) ([]string, error) {
	resp, err := c.Client.GetConversationNotReceiveMessageUserIDs(ctx, &pbconversation.GetConversationNotReceiveMessageUserIDsReq{ConversationID: conversationID})
	if err != nil {
		return nil, err
	}
	return resp.UserIDs, nil
}
