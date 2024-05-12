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
	"encoding/json"
	"time"

	"github.com/Meikwei/aetim/pkg/common/config"
	"github.com/Meikwei/go-tools/discovery"
	"github.com/Meikwei/go-tools/log"
	"github.com/Meikwei/go-tools/mcontext"
	"github.com/Meikwei/go-tools/mq/memamq"
	"github.com/Meikwei/go-tools/system/program"
	"github.com/Meikwei/go-tools/utils/idutil"
	"github.com/Meikwei/go-tools/utils/jsonutil"
	"github.com/Meikwei/go-tools/utils/timeutil"
	"github.com/Meikwei/protocol/constant"
	"github.com/Meikwei/protocol/msg"
	"github.com/Meikwei/protocol/sdkws"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// newContentTypeConf 生成一个映射，其将特定的通知类型映射到相应的配置。
//
// 参数:
// conf: 一个指向config.Notification的指针，包含了各种通知的配置信息。
//
// 返回值:
// 返回一个映射，键是通知类型的整数常量，值是对应的config.NotificationConfig结构体。
func newContentTypeConf(conf *config.Notification) map[int32]config.NotificationConfig {
	// 初始化一个通知类型到配置的映射
	return map[int32]config.NotificationConfig{
		// 配置与群组相关的通知类型
		constant.GroupCreatedNotification:                 conf.GroupCreated,
		constant.GroupInfoSetNotification:                 conf.GroupInfoSet,
		constant.JoinGroupApplicationNotification:         conf.JoinGroupApplication,
		constant.MemberQuitNotification:                   conf.MemberQuit,
		constant.GroupApplicationAcceptedNotification:     conf.GroupApplicationAccepted,
		constant.GroupApplicationRejectedNotification:     conf.GroupApplicationRejected,
		constant.GroupOwnerTransferredNotification:        conf.GroupOwnerTransferred,
		constant.MemberKickedNotification:                 conf.MemberKicked,
		constant.MemberInvitedNotification:                conf.MemberInvited,
		constant.MemberEnterNotification:                  conf.MemberEnter,
		constant.GroupDismissedNotification:               conf.GroupDismissed,
		constant.GroupMutedNotification:                   conf.GroupMuted,
		constant.GroupCancelMutedNotification:             conf.GroupCancelMuted,
		constant.GroupMemberMutedNotification:             conf.GroupMemberMuted,
		constant.GroupMemberCancelMutedNotification:       conf.GroupMemberCancelMuted,
		constant.GroupMemberInfoSetNotification:           conf.GroupMemberInfoSet,
		constant.GroupMemberSetToAdminNotification:        conf.GroupMemberSetToAdmin,
		constant.GroupMemberSetToOrdinaryUserNotification: conf.GroupMemberSetToOrdinary,
		constant.GroupInfoSetAnnouncementNotification:     conf.GroupInfoSetAnnouncement,
		constant.GroupInfoSetNameNotification:             conf.GroupInfoSetName,
		// 配置与用户相关的通知类型
		constant.UserInfoUpdatedNotification:  conf.UserInfoUpdated,
		constant.UserStatusChangeNotification: conf.UserStatusChanged,
		// 配置与好友相关的通知类型
		constant.FriendApplicationNotification:         conf.FriendApplicationAdded,
		constant.FriendApplicationApprovedNotification: conf.FriendApplicationApproved,
		constant.FriendApplicationRejectedNotification: conf.FriendApplicationRejected,
		constant.FriendAddedNotification:               conf.FriendAdded,
		constant.FriendDeletedNotification:             conf.FriendDeleted,
		constant.FriendRemarkSetNotification:           conf.FriendRemarkSet,
		constant.BlackAddedNotification:                conf.BlackAdded,
		constant.BlackDeletedNotification:              conf.BlackDeleted,
		constant.FriendInfoUpdatedNotification:         conf.FriendInfoUpdated,
		constant.FriendsInfoUpdateNotification:         conf.FriendInfoUpdated, // 使用相同的朋友信息更新配置
		// 配置与会话相关的通知类型
		constant.ConversationChangeNotification:      conf.ConversationChanged,
		constant.ConversationUnreadNotification:      conf.ConversationChanged,
		constant.ConversationPrivateChatNotification: conf.ConversationSetPrivate,
		// 配置与消息相关的通知类型，特定消息操作的通知配置
		constant.MsgRevokeNotification:  {IsSendMsg: false, ReliabilityLevel: constant.ReliableNotificationNoMsg},
		constant.HasReadReceipt:         {IsSendMsg: false, ReliabilityLevel: constant.ReliableNotificationNoMsg},
		constant.DeleteMsgsNotification: {IsSendMsg: false, ReliabilityLevel: constant.ReliableNotificationNoMsg},
	}
}

// newSessionTypeConf 创建并返回一个映射，该映射将不同的通知类型关联到相应的会话消息类型。
// 该映射主要用于确定接收到特定类型的通知时，应该以哪种聊天类型（单聊或群聊）来处理该消息。
// 
// 返回值:
//   - map[int32]int32: 一个映射，其中key是通知类型（int32），value是相应的会话消息类型（int32）。
func newSessionTypeConf() map[int32]int32 {
	// 初始化并返回一个包含预定义通知类型到会话类型映射的字典
	return map[int32]int32{
		// 配置群组相关的通知类型到对应的会话类型
		constant.GroupCreatedNotification:                 constant.ReadGroupChatType,
		constant.GroupInfoSetNotification:                 constant.ReadGroupChatType,
		constant.JoinGroupApplicationNotification:         constant.SingleChatType,
		constant.MemberQuitNotification:                   constant.ReadGroupChatType,
		constant.GroupApplicationAcceptedNotification:     constant.SingleChatType,
		constant.GroupApplicationRejectedNotification:     constant.SingleChatType,
		constant.GroupOwnerTransferredNotification:        constant.ReadGroupChatType,
		constant.MemberKickedNotification:                 constant.ReadGroupChatType,
		constant.MemberInvitedNotification:                constant.ReadGroupChatType,
		constant.MemberEnterNotification:                  constant.ReadGroupChatType,
		constant.GroupDismissedNotification:               constant.ReadGroupChatType,
		constant.GroupMutedNotification:                   constant.ReadGroupChatType,
		constant.GroupCancelMutedNotification:             constant.ReadGroupChatType,
		constant.GroupMemberMutedNotification:             constant.ReadGroupChatType,
		constant.GroupMemberCancelMutedNotification:       constant.ReadGroupChatType,
		constant.GroupMemberInfoSetNotification:           constant.ReadGroupChatType,
		constant.GroupMemberSetToAdminNotification:        constant.ReadGroupChatType,
		constant.GroupMemberSetToOrdinaryUserNotification: constant.ReadGroupChatType,
		constant.GroupInfoSetAnnouncementNotification:     constant.ReadGroupChatType,
		constant.GroupInfoSetNameNotification:             constant.ReadGroupChatType,
		// 配置用户相关的通知类型到对应的会话类型
		constant.UserInfoUpdatedNotification:  constant.SingleChatType,
		constant.UserStatusChangeNotification: constant.SingleChatType,
		// 配置好友相关的通知类型到对应的会话类型
		constant.FriendApplicationNotification:         constant.SingleChatType,
		constant.FriendApplicationApprovedNotification: constant.SingleChatType,
		constant.FriendApplicationRejectedNotification: constant.SingleChatType,
		constant.FriendAddedNotification:               constant.SingleChatType,
		constant.FriendDeletedNotification:             constant.SingleChatType,
		constant.FriendRemarkSetNotification:           constant.SingleChatType,
		constant.BlackAddedNotification:                constant.SingleChatType,
		constant.BlackDeletedNotification:              constant.SingleChatType,
		constant.FriendInfoUpdatedNotification:         constant.SingleChatType,
		constant.FriendsInfoUpdateNotification:         constant.SingleChatType,
		// 配置会话相关的通知类型到对应的会话类型
		constant.ConversationChangeNotification:      constant.SingleChatType,
		constant.ConversationUnreadNotification:      constant.SingleChatType,
		constant.ConversationPrivateChatNotification: constant.SingleChatType,
		// 配置删除相关的通知类型到对应的会话类型
		constant.DeleteMsgsNotification: constant.SingleChatType,
	}
}

// Message 表示一个消息结构体，包含了与消息服务相关的gRPC连接和客户端
type Message struct {
	conn   grpc.ClientConnInterface // gRPC连接接口
	Client msg.MsgClient            // 消息服务的gRPC客户端
	discov discovery.SvcDiscoveryRegistry // 服务发现注册接口，用于获取gRPC连接
}

// NewMessage 创建一个新的Message实例。
// discov: 用于服务发现的注册接口。
// rpcRegisterName: 注册服务的名称，用于通过服务发现获取gRPC连接。
// 返回值: 初始化后的Message指针。
func NewMessage(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) *Message {
	// 通过服务发现获取gRPC连接
	conn, err := discov.GetConn(context.Background(), rpcRegisterName)
	if err != nil {
		// 如果获取连接失败，则程序退出
		program.ExitWithError(err)
	}
	// 根据获取的连接创建消息服务的gRPC客户端
	client := msg.NewMsgClient(conn)
	return &Message{discov: discov, conn: conn, Client: client}
}

// MessageRpcClient 是Message的一个别名，用于创建RPC客户端
type MessageRpcClient Message

// NewMessageRpcClient 创建一个MessageRpcClient实例。
// discov: 用于服务发现的注册接口。
// rpcRegisterName: 注册服务的名称，用于通过服务发现获取gRPC连接。
// 返回值: 初始化后的MessageRpcClient实例。
func NewMessageRpcClient(discov discovery.SvcDiscoveryRegistry, rpcRegisterName string) MessageRpcClient {
	// 通过调用NewMessage来创建MessageRpcClient实例
	return MessageRpcClient(*NewMessage(discov, rpcRegisterName))
}

// SendMsg sends a message through the gRPC client and returns the response.
// It wraps any encountered error for better error handling and context understanding.
func (m *MessageRpcClient) SendMsg(ctx context.Context, req *msg.SendMsgReq) (*msg.SendMsgResp, error) {
	resp, err := m.Client.SendMsg(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMaxSeq retrieves the maximum sequence number from the gRPC client.
// Errors during the gRPC call are wrapped to provide additional context.
func (m *MessageRpcClient) GetMaxSeq(ctx context.Context, req *sdkws.GetMaxSeqReq) (*sdkws.GetMaxSeqResp, error) {
	resp, err := m.Client.GetMaxSeq(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MessageRpcClient) GetMaxSeqs(ctx context.Context, conversationIDs []string) (map[string]int64, error) {
	log.ZDebug(ctx, "GetMaxSeqs", "conversationIDs", conversationIDs)
	resp, err := m.Client.GetMaxSeqs(ctx, &msg.GetMaxSeqsReq{
		ConversationIDs: conversationIDs,
	})
	return resp.MaxSeqs, err
}

func (m *MessageRpcClient) GetHasReadSeqs(ctx context.Context, userID string, conversationIDs []string) (map[string]int64, error) {
	resp, err := m.Client.GetHasReadSeqs(ctx, &msg.GetHasReadSeqsReq{
		UserID:          userID,
		ConversationIDs: conversationIDs,
	})
	return resp.MaxSeqs, err
}

func (m *MessageRpcClient) GetMsgByConversationIDs(ctx context.Context, docIDs []string, seqs map[string]int64) (map[string]*sdkws.MsgData, error) {
	resp, err := m.Client.GetMsgByConversationIDs(ctx, &msg.GetMsgByConversationIDsReq{
		ConversationIDs: docIDs,
		MaxSeqs:         seqs,
	})
	return resp.MsgDatas, err
}

// PullMessageBySeqList retrieves messages by their sequence numbers using the gRPC client.
// It directly forwards the request to the gRPC client and returns the response along with any error encountered.
func (m *MessageRpcClient) PullMessageBySeqList(ctx context.Context, req *sdkws.PullMessageBySeqsReq) (*sdkws.PullMessageBySeqsResp, error) {
	resp, err := m.Client.PullMessageBySeqs(ctx, req)
	if err != nil {
		// Wrap the error to provide more context if the gRPC call fails.
		return nil, err
	}
	return resp, nil
}

func (m *MessageRpcClient) GetConversationMaxSeq(ctx context.Context, conversationID string) (int64, error) {
	resp, err := m.Client.GetConversationMaxSeq(ctx, &msg.GetConversationMaxSeqReq{ConversationID: conversationID})
	if err != nil {
		return 0, err
	}
	return resp.MaxSeq, nil
}

type NotificationSender struct {
	contentTypeConf map[int32]config.NotificationConfig
	sessionTypeConf map[int32]int32
	sendMsg         func(ctx context.Context, req *msg.SendMsgReq) (*msg.SendMsgResp, error)
	getUserInfo     func(ctx context.Context, userID string) (*sdkws.UserInfo, error)
	queue           *memamq.MemoryQueue
}

func WithQueue(queue *memamq.MemoryQueue) NotificationSenderOptions {
	return func(s *NotificationSender) {
		s.queue = queue
	}
}

type NotificationSenderOptions func(*NotificationSender)

func WithLocalSendMsg(sendMsg func(ctx context.Context, req *msg.SendMsgReq) (*msg.SendMsgResp, error)) NotificationSenderOptions {
	return func(s *NotificationSender) {
		s.sendMsg = sendMsg
	}
}

func WithRpcClient(msgRpcClient *MessageRpcClient) NotificationSenderOptions {
	return func(s *NotificationSender) {
		s.sendMsg = msgRpcClient.SendMsg
	}
}

func WithUserRpcClient(userRpcClient *UserRpcClient) NotificationSenderOptions {
	return func(s *NotificationSender) {
		s.getUserInfo = userRpcClient.GetUserInfo
	}
}

const (
	notificationWorkerCount = 2
	notificationBufferSize  = 200
)

func NewNotificationSender(conf *config.Notification, opts ...NotificationSenderOptions) *NotificationSender {
	notificationSender := &NotificationSender{contentTypeConf: newContentTypeConf(conf), sessionTypeConf: newSessionTypeConf()}
	for _, opt := range opts {
		opt(notificationSender)
	}
	if notificationSender.queue == nil {
		notificationSender.queue = memamq.NewMemoryQueue(notificationWorkerCount, notificationBufferSize)
	}
	return notificationSender
}

type notificationOpt struct {
	WithRpcGetUsername bool
}

type NotificationOptions func(*notificationOpt)

func WithRpcGetUserName() NotificationOptions {
	return func(opt *notificationOpt) {
		opt.WithRpcGetUsername = true
	}
}

func (s *NotificationSender) send(ctx context.Context, sendID, recvID string, contentType, sessionType int32, m proto.Message, opts ...NotificationOptions) {
	ctx = mcontext.WithMustInfoCtx([]string{mcontext.GetOperationID(ctx), mcontext.GetOpUserID(ctx), mcontext.GetOpUserPlatform(ctx), mcontext.GetConnID(ctx)})
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(5))
	defer cancel()
	n := sdkws.NotificationElem{Detail: jsonutil.StructToJsonString(m)}
	content, err := json.Marshal(&n)
	if err != nil {
		log.ZWarn(ctx, "json.Marshal failed", err, "sendID", sendID, "recvID", recvID, "contentType", contentType, "msg", jsonutil.StructToJsonString(m))
		return
	}
	notificationOpt := &notificationOpt{}
	for _, opt := range opts {
		opt(notificationOpt)
	}
	var req msg.SendMsgReq
	var msg sdkws.MsgData
	var userInfo *sdkws.UserInfo
	if notificationOpt.WithRpcGetUsername && s.getUserInfo != nil {
		userInfo, err = s.getUserInfo(ctx, sendID)
		if err != nil {
			log.ZWarn(ctx, "getUserInfo failed", err, "sendID", sendID)
			return
		}
		msg.SenderNickname = userInfo.Nickname
		msg.SenderFaceURL = userInfo.FaceURL
	}
	var offlineInfo sdkws.OfflinePushInfo
	msg.SendID = sendID
	msg.RecvID = recvID
	msg.Content = content
	msg.MsgFrom = constant.SysMsgType
	msg.ContentType = contentType
	msg.SessionType = sessionType
	if msg.SessionType == constant.ReadGroupChatType {
		msg.GroupID = recvID
	}
	msg.CreateTime = timeutil.GetCurrentTimestampByMill()
	msg.ClientMsgID = idutil.GetMsgIDByMD5(sendID)
	optionsConfig := s.contentTypeConf[contentType]
	if sendID == recvID && contentType == constant.HasReadReceipt {
		optionsConfig.ReliabilityLevel = constant.UnreliableNotification
	}
	options := config.GetOptionsByNotification(optionsConfig)
	s.SetOptionsByContentType(ctx, options, contentType)
	msg.Options = options
	msg.OfflinePushInfo = &offlineInfo
	req.MsgData = &msg
	_, err = s.sendMsg(ctx, &req)
	if err != nil {
		log.ZWarn(ctx, "SendMsg failed", err, "req", req.String())
	}
}

func (s *NotificationSender) NotificationWithSessionType(ctx context.Context, sendID, recvID string, contentType, sessionType int32, m proto.Message, opts ...NotificationOptions) {
	s.queue.Push(func() { s.send(ctx, sendID, recvID, contentType, sessionType, m, opts...) })
}

func (s *NotificationSender) Notification(ctx context.Context, sendID, recvID string, contentType int32, m proto.Message, opts ...NotificationOptions) {
	s.NotificationWithSessionType(ctx, sendID, recvID, contentType, s.sessionTypeConf[contentType], m, opts...)
}

func (s *NotificationSender) SetOptionsByContentType(_ context.Context, options map[string]bool, contentType int32) {
	switch contentType {
	case constant.UserStatusChangeNotification:
		options[constant.IsSenderSync] = false
	default:
	}
}
