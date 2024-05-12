/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-04 15:23:07
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 15:07:48
 * @FilePath: \user\pkg\common\config\config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"fmt"
	"net"
	"time"

	"github.com/Meikwei/go-tools/db/mongoutil"
	"github.com/Meikwei/go-tools/db/redisutil"
	"github.com/Meikwei/go-tools/mq/kafka"
	"github.com/Meikwei/go-tools/s3/cos"
	"github.com/Meikwei/go-tools/s3/minio"
	"github.com/Meikwei/go-tools/s3/oss"
)

// CacheConfig 定义了缓存的配置结构，包括Topic, SlotNum, SlotSize, SuccessExpire和FailedExpire等字段。
// 这些字段用于配置缓存的主题、槽数量、槽大小、成功和失败的过期时间等。
type CacheConfig struct {
	Topic         string `mapstructure:"topic"` // 缓存主题
	SlotNum       int    `mapstructure:"slotNum"` // 槽数量
	SlotSize      int    `mapstructure:"slotSize"` // 槽大小
	SuccessExpire int    `mapstructure:"successExpire"` // 成功缓存项的过期时间
	FailedExpire  int    `mapstructure:"failedExpire"` // 失败缓存项的过期时间
}

// LocalCache 定义了本地缓存的配置结构，包括User, Group, Friend和Conversation等缓存配置。
// 这些配置分别用于不同类型的用户数据缓存。
type LocalCache struct {
	User         CacheConfig `mapstructure:"user"` // 用户缓存配置
	Group        CacheConfig `mapstructure:"group"` // 群组缓存配置
	Friend       CacheConfig `mapstructure:"friend"` // 好友缓存配置
	Conversation CacheConfig `mapstructure:"conversation"` // 聊天会话缓存配置
}

// Log 定义了日志的配置结构，包括StorageLocation, RotationTime, RemainRotationCount, RemainLogLevel, IsStdout, IsJson和WithStack等字段。
// 这些字段用于配置日志的存储位置、轮转时间、保留轮转次数、日志级别、是否输出到标准输出、是否以JSON格式记录以及是否记录堆栈信息。
type Log struct {
	StorageLocation     string `mapstructure:"storageLocation"` // 日志存储位置
	RotationTime        uint   `mapstructure:"rotationTime"` // 日志轮转时间
	RemainRotationCount uint   `mapstructure:"remainRotationCount"` // 保留的日志轮转次数
	RemainLogLevel      int    `mapstructure:"remainLogLevel"` // 保留的日志级别
	IsStdout            bool   `mapstructure:"isStdout"` // 是否输出到标准输出
	IsJson              bool   `mapstructure:"isJson"` // 是否以JSON格式记录日志
	WithStack           bool   `mapstructure:"withStack"` // 是否记录堆栈信息
}

// Minio 定义了MinIO服务的配置结构，包括Bucket, AccessKeyID, SecretAccessKey, SessionToken, InternalAddress和ExternalAddress等字段。
// 这些字段用于配置MinIO的存储桶、访问密钥、秘钥、会话令牌、内部地址和外部地址等。
type Minio struct {
	Bucket          string `mapstructure:"bucket"` // 存储桶名称
	AccessKeyID     string `mapstructure:"accessKeyID"` // 访问密钥ID
	SecretAccessKey string `mapstructure:"secretAccessKey"` // 访问密钥Secret
	SessionToken    string `mapstructure:"sessionToken"` // 会话令牌
	InternalAddress string `mapstructure:"internalAddress"` // 内部地址
	ExternalAddress string `mapstructure:"externalAddress"` // 外部地址
	PublicRead      bool   `mapstructure:"publicRead"` // 是否允许公共读
}

// Mongo 定义了MongoDB的配置结构，包括URI, Address, Database, Username, Password, MaxPoolSize和MaxRetry等字段。
// 这些字段用于配置MongoDB的连接URI、地址、数据库名称、用户名、密码、最大连接池大小和最大重试次数。
type Mongo struct {
	URI         string   `mapstructure:"uri"` // 连接URI
	Address     []string `mapstructure:"address"` // 服务地址列表
	Database    string   `mapstructure:"database"` // 数据库名称
	Username    string   `mapstructure:"username"` // 用户名
	Password    string   `mapstructure:"password"` // 密码
	MaxPoolSize int      `mapstructure:"maxPoolSize"` // 最大连接池大小
	MaxRetry    int      `mapstructure:"maxRetry"` // 最大重试次数
}

// Kafka 定义了Kafka的配置结构，包括Username, Password, ProducerAck, CompressType, Address, ToRedisTopic, ToMongoTopic, ToPushTopic等字段，以及用于TLS连接的TLSConfig子结构。
// 这些字段用于配置Kafka的用户名、密码、生产者确认、压缩类型、地址、不同目的的主题名称，以及TLS连接的相关配置。
type Kafka struct {
	Username       string    `mapstructure:"username"` // 用户名
	Password       string    `mapstructure:"password"` // 密码
	ProducerAck    string    `mapstructure:"producerAck"` // 生产者确认
	CompressType   string    `mapstructure:"compressType"` // 压缩类型
	Address        []string  `mapstructure:"address"` // Kafka地址列表
	ToRedisTopic   string    `mapstructure:"toRedisTopic"` // 发送到Redis的主题
	ToMongoTopic   string    `mapstructure:"toMongoTopic"` // 发送到MongoDB的主题
	ToPushTopic    string    `mapstructure:"toPushTopic"` // 发送推送的消息主题
	ToRedisGroupID string    `mapstructure:"toRedisGroupID"` // Redis消费者组ID
	ToMongoGroupID string    `mapstructure:"toMongoGroupID"` // MongoDB消费者组ID
	ToPushGroupID  string    `mapstructure:"toPushGroupID"` // 推送消费者组ID
	Tls            TLSConfig `mapstructure:"tls"` // TLS配置
}

// TLSConfig 定义了TLS连接的配置结构，包括EnableTLS, CACrt, ClientCrt, ClientKey, ClientKeyPwd和InsecureSkipVerify等字段。
// 这些字段用于配置是否启用TLS、CA证书路径、客户端证书路径、客户端密钥路径、客户端密钥密码和是否跳过TLS验证。
type TLSConfig struct {
	EnableTLS          bool   `mapstructure:"enableTLS"` // 是否启用TLS
	CACrt              string `mapstructure:"caCrt"` // CA证书路径
	ClientCrt          string `mapstructure:"clientCrt"` // 客户端证书路径
	ClientKey          string `mapstructure:"clientKey"` // 客户端密钥路径
	ClientKeyPwd       string `mapstructure:"clientKeyPwd"` // 客户端密钥密码
	InsecureSkipVerify bool   `mapstructure:"insecureSkipVerify"` // 是否跳过TLS验证
}

// API 结构定义了API配置的相关参数
type API struct {
	Api struct {
		ListenIP string `mapstructure:"listenIP"` // API监听的IP地址
		Ports    []int  `mapstructure:"ports"`    // API监听的端口列表
	} `mapstructure:"api"` // API配置的顶层标签
	Prometheus struct {
		Enable     bool   `mapstructure:"enable"`      // 是否启用Prometheus监控
		Ports      []int  `mapstructure:"ports"`       // Prometheus监听的端口列表
		GrafanaURL string `mapstructure:"grafanaURL"`  // Grafana的URL地址
	} `mapstructure:"prometheus"` // Prometheus配置的顶层标签
}

// CronTask 定义了定时任务的相关配置
type CronTask struct {
	ChatRecordsClearTime string `mapstructure:"chatRecordsClearTime"` // 聊天记录清除时间配置
	MsgDestructTime      string `mapstructure:"msgDestructTime"`      // 消息自毁时间配置
	RetainChatRecords    int    `mapstructure:"retainChatRecords"`    // 保留聊天记录的时间（天数）
	EnableCronLocker     bool   `yaml:"enableCronLocker"`             // 是否启用定时任务锁
}

// OfflinePushConfig 定义了离线推送的配置
type OfflinePushConfig struct {
	Enable bool   `mapstructure:"enable"` // 是否启用离线推送
	Title  string `mapstructure:"title"`  // 推送标题
	Desc   string `mapstructure:"desc"`   // 推送描述
	Ext    string `mapstructure:"ext"`    // 推送扩展信息
}

// NotificationConfig 定义了各种通知的配置
type NotificationConfig struct {
	IsSendMsg        bool              `mapstructure:"isSendMsg"`        // 是否发送消息通知
	ReliabilityLevel int               `mapstructure:"reliabilityLevel"` // 通知的可靠性级别
	UnreadCount      bool              `mapstructure:"unreadCount"`      // 是否显示未读计数
	OfflinePush      OfflinePushConfig `mapstructure:"offlinePush"`      // 离线推送配置
}

// Notification 定义了各种通知的配置集合
type Notification struct {
	GroupCreated              NotificationConfig `mapstructure:"groupCreated"`              // 群组创建通知配置
	GroupInfoSet              NotificationConfig `mapstructure:"groupInfoSet"`              // 群组信息设置通知配置
	JoinGroupApplication      NotificationConfig `mapstructure:"joinGroupApplication"`      // 加入群组申请通知配置
	MemberQuit                NotificationConfig `mapstructure:"memberQuit"`                // 群组成员退出通知配置
	GroupApplicationAccepted  NotificationConfig `mapstructure:"groupApplicationAccepted"`  // 群组申请通过通知配置
	GroupApplicationRejected  NotificationConfig `mapstructure:"groupApplicationRejected"`  // 群组申请拒绝通知配置
	GroupOwnerTransferred     NotificationConfig `mapstructure:"groupOwnerTransferred"`     // 群组所有权转移通知配置
	MemberKicked              NotificationConfig `mapstructure:"memberKicked"`              // 成员被移除通知配置
	MemberInvited             NotificationConfig `mapstructure:"memberInvited"`             // 成员被邀请通知配置
	MemberEnter               NotificationConfig `mapstructure:"memberEnter"`               // 成员进入群组通知配置
	GroupDismissed            NotificationConfig `mapstructure:"groupDismissed"`            // 群组解散通知配置
	GroupMuted                NotificationConfig `mapstructure:"groupMuted"`                // 群组被禁言通知配置
	GroupCancelMuted          NotificationConfig `mapstructure:"groupCancelMuted"`          // 群组取消禁言通知配置
	GroupMemberMuted          NotificationConfig `mapstructure:"groupMemberMuted"`          // 群组成员被禁言通知配置
	GroupMemberCancelMuted    NotificationConfig `mapstructure:"groupMemberCancelMuted"`    // 群组成员取消禁言通知配置
	GroupMemberInfoSet        NotificationConfig `mapstructure:"groupMemberInfoSet"`        // 群组成员信息设置通知配置
	GroupMemberSetToAdmin     NotificationConfig `yaml:"groupMemberSetToAdmin"`             // 群组成员被设置为管理员通知配置
	GroupMemberSetToOrdinary  NotificationConfig `yaml:"groupMemberSetToOrdinaryUser"`      // 群组成员被设置为普通成员通知配置
	GroupInfoSetAnnouncement  NotificationConfig `mapstructure:"groupInfoSetAnnouncement"`  // 群组公告设置通知配置
	GroupInfoSetName          NotificationConfig `mapstructure:"groupInfoSetName"`          // 群组名称设置通知配置
	FriendApplicationAdded    NotificationConfig `mapstructure:"friendApplicationAdded"`    // 好友申请添加通知配置
	FriendApplicationApproved NotificationConfig `mapstructure:"friendApplicationApproved"` // 好友申请通过通知配置
	FriendApplicationRejected NotificationConfig `mapstructure:"friendApplicationRejected"` // 好友申请拒绝通知配置
	FriendAdded               NotificationConfig `mapstructure:"friendAdded"`               // 好友添加通知配置
	FriendDeleted             NotificationConfig `mapstructure:"friendDeleted"`             // 好友删除通知配置
	FriendRemarkSet           NotificationConfig `mapstructure:"friendRemarkSet"`           // 好友备注设置通知配置
	BlackAdded                NotificationConfig `mapstructure:"blackAdded"`                // 黑名单添加通知配置
	BlackDeleted              NotificationConfig `mapstructure:"blackDeleted"`              // 黑名单删除通知配置
	FriendInfoUpdated         NotificationConfig `mapstructure:"friendInfoUpdated"`         // 好友信息更新通知配置
	UserInfoUpdated           NotificationConfig `mapstructure:"userInfoUpdated"`           // 用户信息更新通知配置
	UserStatusChanged         NotificationConfig `mapstructure:"userStatusChanged"`         // 用户状态改变通知配置
	ConversationChanged       NotificationConfig `mapstructure:"conversationChanged"`       // 聊天会话改变通知配置
	ConversationSetPrivate    NotificationConfig `mapstructure:"conversationSetPrivate"`    // 聊天会话设置为私有通知配置
}

// Prometheus 结构定义了Prometheus监控的配置
type Prometheus struct {
	Enable bool  `mapstructure:"enable"` // 是否启用Prometheus监控
	Ports  []int `mapstructure:"ports"`  // Prometheus监听的端口列表
}

// MsgGateway 定义了消息网关的配置
type MsgGateway struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // RPC注册的IP地址
		Ports      []int  `mapstructure:"ports"`      // RPC监听的端口列表
	} `mapstructure:"rpc"` // RPC配置
	Prometheus  Prometheus `mapstructure:"prometheus"`  // Prometheus监控配置
	ListenIP    string     `mapstructure:"listenIP"`    // 消息网关监听的IP地址
	LongConnSvr struct {
		Ports               []int `mapstructure:"ports"`               // 长连接服务器监听的端口列表
		WebsocketMaxConnNum int   `mapstructure:"websocketMaxConnNum"` // WebSocket最大连接数
		WebsocketMaxMsgLen  int   `mapstructure:"websocketMaxMsgLen"`  // WebSocket最大消息长度
		WebsocketTimeout    int   `mapstructure:"websocketTimeout"`    // WebSocket超时时间
	} `mapstructure:"longConnSvr"` // 长连接服务器配置
	MultiLoginPolicy int `mapstructure:"multiLoginPolicy"` // 多设备登录策略
}

// MsgTransfer 定义了消息传输的配置
type MsgTransfer struct {
	Prometheus Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
}

// Push 结构体定义了推送服务的配置参数
type Push struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听IP地址
		Ports      []int  `mapstructure:"ports"`      // 使用的端口号列表
	} `mapstructure:"rpc"` // RPC服务配置
	Prometheus           Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
	MaxConcurrentWorkers int        `mapstructure:"maxConcurrentWorkers"` // 最大并发工作器数量
	Enable               string     `mapstructure:"enable"` // 启用标志
	GeTui                struct { // GeTui推送服务配置
		PushUrl      string `mapstructure:"pushUrl"`      // 推送URL
		MasterSecret string `mapstructure:"masterSecret"` // 主密钥
		AppKey       string `mapstructure:"appKey"`       // 应用密钥
		Intent       string `mapstructure:"intent"`       // 意图参数
		ChannelID    string `mapstructure:"channelID"`    // 频道ID
		ChannelName  string `mapstructure:"channelName"`  // 频道名称
	} `mapstructure:"geTui"`
	FCM struct { // FCM推送服务配置
		ServiceAccount string `mapstructure:"serviceAccount"` // 服务帐户文件路径
	} `mapstructure:"fcm"`
	JPNS struct { // JPNS推送服务配置
		AppKey       string `mapstructure:"appKey"`       // 应用密钥
		MasterSecret string `mapstructure:"masterSecret"` // 主密钥
		PushURL      string `mapstructure:"pushURL"`      // 推送URL
		PushIntent   string `mapstructure:"pushIntent"`   // 意图参数
	} `mapstructure:"jpns"`
	IOSPush struct { // iOS推送服务配置
		PushSound  string `mapstructure:"pushSound"`  // 推送声音
		BadgeCount bool   `mapstructure:"badgeCount"` // 是否显示角标计数
		Production bool   `mapstructure:"production"`  // 是否使用生产环境
	} `mapstructure:"iosPush"`
}

// Auth 结构体定义了认证服务的配置参数
type Auth struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听IP地址
		Ports      []int  `mapstructure:"ports"`      // 使用的端口号列表
	} `mapstructure:"rpc"` // RPC服务配置
	Prometheus  Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
	TokenPolicy struct {
		Expire int64 `mapstructure:"expire"` // Token过期时间
	} `mapstructure:"tokenPolicy"` // Token策略配置
}

// Conversation 结构体定义了会话服务的配置参数
type Conversation struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听IP地址
		Ports      []int  `mapstructure:"ports"`      // 使用的端口号列表
	} `mapstructure:"rpc"` // RPC服务配置
	Prometheus Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
}

// Friend 结构体定义了好友服务的配置参数
type Friend struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听IP地址
		Ports      []int  `mapstructure:"ports"`      // 使用的端口号列表
	} `mapstructure:"rpc"` // RPC服务配置
	Prometheus Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
}

// Group 结构体定义了群组服务的配置参数
type Group struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听IP地址
		Ports      []int  `mapstructure:"ports"`      // 使用的端口号列表
	} `mapstructure:"rpc"` // RPC服务配置
	Prometheus Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
}

// Msg 结构体定义了消息服务的配置参数
type Msg struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听IP地址
		Ports      []int  `mapstructure:"ports"`      // 使用的端口号列表
	} `mapstructure:"rpc"` // RPC服务配置
	Prometheus   Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
	FriendVerify bool       `mapstructure:"friendVerify"` // 好友验证标志
}

// Third 定义了与第三方服务配置相关的结构体
type Third struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册时使用的IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听的IP地址
		Ports      []int  `mapstructure:"ports"`      // 监听的端口
	} `mapstructure:"rpc"` // RPC配置
	Prometheus Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
	Object     struct {
		Enable string `mapstructure:"enable"` // 是否启用对象存储
		Cos    Cos    `mapstructure:"cos"`    // COS配置
		Oss    Oss    `mapstructure:"oss"`    // OSS配置
		Kodo   struct {
			Endpoint        string `mapstructure:"endpoint"`        // 服务端点
			Bucket          string `mapstructure:"bucket"`          // 存储桶名称
			BucketURL       string `mapstructure:"bucketURL"`       // 存储桶URL
			AccessKeyID     string `mapstructure:"accessKeyID"`     // 访问密钥ID
			AccessKeySecret string `mapstructure:"accessKeySecret"` // 访问密钥Secret
			SessionToken    string `mapstructure:"sessionToken"`    // 会话令牌
			PublicRead      bool   `mapstructure:"publicRead"`      // 是否允许公共读取
		} `mapstructure:"kodo"` // Kodo配置
		Aws struct {
			Endpoint        string `mapstructure:"endpoint"`        // 服务端点
			Region          string `mapstructure:"region"`          // 区域
			Bucket          string `mapstructure:"bucket"`          // 存储桶名称
			AccessKeyID     string `mapstructure:"accessKeyID"`     // 访问密钥ID
			AccessKeySecret string `mapstructure:"accessKeySecret"` // 访问密钥Secret
			PublicRead      bool   `mapstructure:"publicRead"`      // 是否允许公共读取
		} `mapstructure:"aws"` // Aws配置
	} `mapstructure:"object"` // 对象存储配置
}

// Cos 定义了COS（Cloud Object Storage）服务的配置项
type Cos struct {
	BucketURL    string `mapstructure:"bucketURL"`    // 存储桶URL
	SecretID     string `mapstructure:"secretID"`     // 秘钥ID
	SecretKey    string `mapstructure:"secretKey"`    // 秘钥Key
	SessionToken string `mapstructure:"sessionToken"` // 会话令牌
	PublicRead   bool   `mapstructure:"publicRead"`   // 是否允许公共读取
}

// Oss 定义了OSS（Object Storage Service）服务的配置项
type Oss struct {
	Endpoint        string `mapstructure:"endpoint"`        // 服务端点
	Bucket          string `mapstructure:"bucket"`          // 存储桶名称
	BucketURL       string `mapstructure:"bucketURL"`       // 存储桶URL
	AccessKeyID     string `mapstructure:"accessKeyID"`     // 访问密钥ID
	AccessKeySecret string `mapstructure:"accessKeySecret"` // 访问密钥Secret
	SessionToken    string `mapstructure:"sessionToken"`    // 会话令牌
	PublicRead      bool   `mapstructure:"publicRead"`      // 是否允许公共读取
}

// User 定义了用户服务的配置结构体
type User struct {
	RPC struct {
		RegisterIP string `mapstructure:"registerIP"` // 注册时使用的IP地址
		ListenIP   string `mapstructure:"listenIP"`   // 监听的IP地址
		Ports      []int  `mapstructure:"ports"`      // 监听的端口
	} `mapstructure:"rpc"` // RPC配置
	Prometheus Prometheus `mapstructure:"prometheus"` // Prometheus监控配置
}

// Redis 定义了Redis服务的配置项
type Redis struct {
	Address        []string `mapstructure:"address"`        // Redis服务地址列表
	Username       string   `mapstructure:"username"`       // Redis用户名
	Password       string   `mapstructure:"password"`       // Redis密码
	EnablePipeline bool     `mapstructure:"enablePipeline"` // 是否启用Pipeline
	ClusterMode    bool     `mapstructure:"clusterMode"`    // 是否为集群模式
	DB             int      `mapstructure:"db"`             // 数据库索引
	MaxRetry       int      `mapstructure:"MaxRetry"`       // 最大重试次数
}

// BeforeConfig 定义了前置任务的配置项
type BeforeConfig struct {
	Enable         bool `mapstructure:"enable"`         // 是否启用前置任务
	Timeout        int  `mapstructure:"timeout"`        // 超时时间
	FailedContinue bool `mapstructure:"failedContinue"` // 失败是否继续
}

// AfterConfig 定义了后置任务的配置项
type AfterConfig struct {
	Enable  bool `mapstructure:"enable"`  // 是否启用后置任务
	Timeout int  `mapstructure:"timeout"` // 超时时间
}

// Share 定义了共享配置的结构体
type Share struct {
	Secret          string          `mapstructure:"secret"`          // 共享密钥
	Env             string          `mapstructure:"env"`             // 环境
	RpcRegisterName RpcRegisterName `mapstructure:"rpcRegisterName"` // RPC注册名称
	IMAdminUserID   []string        `mapstructure:"imAdminUserID"`   // IM管理员用户ID列表
}

// RpcRegisterName 定义了RPC服务的注册名称
type RpcRegisterName struct {
	User           string `mapstructure:"user"`           // 用户服务
	Friend         string `mapstructure:"friend"`         // 好友服务
	Msg            string `mapstructure:"msg"`            // 消息服务
	Push           string `mapstructure:"push"`           // 推送服务
	MessageGateway string `mapstructure:"messageGateway"` // 消息网关服务
	Group          string `mapstructure:"group"`          // 群组服务
	Auth           string `mapstructure:"auth"`           // 认证服务
	Conversation   string `mapstructure:"conversation"`   // 聊天服务
	Third          string `mapstructure:"third"`          // 第三方服务
}

// GetServiceNames 返回RpcRegisterName结构体中所有服务的名称列表
func (r *RpcRegisterName) GetServiceNames() []string {
	return []string{
		r.User,
		r.Friend,
		r.Msg,
		r.Push,
		r.MessageGateway,
		r.Group,
		r.Auth,
		r.Conversation,
		r.Third,
	}
}

// Webhooks 定义了所有事件前后的配置信息
type Webhooks struct {
	URL                      string       `mapstructure:"url"` // Webhook的URL地址
	BeforeSendSingleMsg      BeforeConfig `mapstructure:"beforeSendSingleMsg"` // 发送单聊消息前的配置
	BeforeUpdateUserInfoEx   BeforeConfig `mapstructure:"beforeUpdateUserInfoEx"` // 更新用户信息扩展字段前的配置
	AfterUpdateUserInfoEx    AfterConfig  `mapstructure:"afterUpdateUserInfoEx"` // 更新用户信息扩展字段后的配置
	AfterSendSingleMsg       AfterConfig  `mapstructure:"afterSendSingleMsg"` // 发送单聊消息后的配置
	BeforeSendGroupMsg       BeforeConfig `mapstructure:"beforeSendGroupMsg"` // 发送群聊消息前的配置
	BeforeMsgModify          BeforeConfig `mapstructure:"beforeMsgModify"` // 消息修改前的配置
	AfterSendGroupMsg        AfterConfig  `mapstructure:"afterSendGroupMsg"` // 发送群聊消息后的配置
	AfterUserOnline          AfterConfig  `mapstructure:"afterUserOnline"` // 用户上线后的配置
	AfterUserOffline         AfterConfig  `mapstructure:"afterUserOffline"` // 用户下线后的配置
	AfterUserKickOff         AfterConfig  `mapstructure:"afterUserKickOff"` // 用户被踢出后的配置
	BeforeOfflinePush        BeforeConfig `mapstructure:"beforeOfflinePush"` // 离线推送前的配置
	BeforeOnlinePush         BeforeConfig `mapstructure:"beforeOnlinePush"` // 在线推送前的配置
	BeforeGroupOnlinePush    BeforeConfig `mapstructure:"beforeGroupOnlinePush"` // 群组在线推送前的配置
	BeforeAddFriend          BeforeConfig `mapstructure:"beforeAddFriend"` // 添加好友前的配置
	BeforeUpdateUserInfo     BeforeConfig `mapstructure:"beforeUpdateUserInfo"` // 更新用户信息前的配置
	AfterUpdateUserInfo      AfterConfig  `mapstructure:"afterUpdateUserInfo"` // 更新用户信息后的配置
	BeforeCreateGroup        BeforeConfig `mapstructure:"beforeCreateGroup"` // 创建群组前的配置
	AfterCreateGroup         AfterConfig  `mapstructure:"afterCreateGroup"` // 创建群组后的配置
	BeforeMemberJoinGroup    BeforeConfig `mapstructure:"beforeMemberJoinGroup"` // 成员加入群组前的配置
	BeforeSetGroupMemberInfo BeforeConfig `mapstructure:"beforeSetGroupMemberInfo"` // 设置群组成员信息前的配置
	AfterSetGroupMemberInfo  AfterConfig  `mapstructure:"afterSetGroupMemberInfo"` // 设置群组成员信息后的配置
	AfterQuitGroup           AfterConfig  `mapstructure:"afterQuitGroup"` // 群组成员退出后的配置
	AfterKickGroupMember     AfterConfig  `mapstructure:"afterKickGroupMember"` // 踢出群组成员后的配置
	AfterDismissGroup        AfterConfig  `mapstructure:"afterDismissGroup"` // 解散群组后的配置
	BeforeApplyJoinGroup     BeforeConfig `mapstructure:"beforeApplyJoinGroup"` // 申请加入群组前的配置
	AfterGroupMsgRead        AfterConfig  `mapstructure:"afterGroupMsgRead"` // 群消息被读后的配置
	AfterSingleMsgRead       AfterConfig  `mapstructure:"afterSingleMsgRead"` // 单消息被读后的配置
	BeforeUserRegister       BeforeConfig `mapstructure:"beforeUserRegister"` // 用户注册前的配置
	AfterUserRegister        AfterConfig  `mapstructure:"afterUserRegister"` // 用户注册后的配置
	AfterTransferGroupOwner  AfterConfig  `mapstructure:"afterTransferGroupOwner"` // 转移群组所有权后的配置
	BeforeSetFriendRemark    BeforeConfig `mapstructure:"beforeSetFriendRemark"` // 设置好友备注前的配置
	AfterSetFriendRemark     AfterConfig  `mapstructure:"afterSetFriendRemark"` // 设置好友备注后的配置
	AfterGroupMsgRevoke      AfterConfig  `mapstructure:"afterGroupMsgRevoke"` // 撤销群消息后的配置
	AfterJoinGroup           AfterConfig  `mapstructure:"afterJoinGroup"` // 加入群组后的配置
	BeforeInviteUserToGroup  BeforeConfig `mapstructure:"beforeInviteUserToGroup"` // 邀请用户加入群组前的配置
	AfterSetGroupInfo        AfterConfig  `mapstructure:"afterSetGroupInfo"` // 设置群组信息后的配置
	BeforeSetGroupInfo       BeforeConfig `mapstructure:"beforeSetGroupInfo"` // 设置群组信息前的配置
	AfterRevokeMsg           AfterConfig  `mapstructure:"afterRevokeMsg"` // 撤销消息后的配置
	BeforeAddBlack           BeforeConfig `mapstructure:"beforeAddBlack"` // 添加黑名单前的配置
	AfterAddFriend           AfterConfig  `mapstructure:"afterAddFriend"` // 添加好友后的配置
	BeforeAddFriendAgree     BeforeConfig `mapstructure:"beforeAddFriendAgree"` // 同意添加好友请求前的配置
	AfterDeleteFriend        AfterConfig  `mapstructure:"afterDeleteFriend"` // 删除好友后的配置
	BeforeImportFriends      BeforeConfig `mapstructure:"beforeImportFriends"` // 导入好友前的配置
	AfterImportFriends       AfterConfig  `mapstructure:"afterImportFriends"` // 导入好友后的配置
	AfterRemoveBlack         AfterConfig  `mapstructure:"afterRemoveBlack"` // 从黑名单中移除后的配置
}

// ZooKeeper 定义了与ZooKeeper交互的配置
type ZooKeeper struct {
	Schema   string   `mapstructure:"schema"` // 协议
	Address  []string `mapstructure:"address"` // 地址列表
	Username string   `mapstructure:"username"` // 用户名
	Password string   `mapstructure:"password"` // 密码
}

// Build 根据Mongo结构体构建mongoutil.Config对象
func (m *Mongo) Build() *mongoutil.Config {
	return &mongoutil.Config{
		Uri:         m.URI, // MongoDB的URI
		Address:     m.Address, // MongoDB的地址列表
		Database:    m.Database, // 数据库名称
		Username:    m.Username, // 用户名
		Password:    m.Password, // 密码
		MaxPoolSize: m.MaxPoolSize, // 连接池最大大小
		MaxRetry:    m.MaxRetry, // 最大重试次数
	}
}

// Build 根据Redis结构体构建redisutil.Config对象
func (r *Redis) Build() *redisutil.Config {
	return &redisutil.Config{
		ClusterMode: r.ClusterMode, // 是否为集群模式
		Address:     r.Address, // Redis地址列表
		Username:    r.Username, // 用户名
		Password:    r.Password, // 密码
		DB:          r.DB, // 数据库索引
		MaxRetry:    r.MaxRetry, // 最大重试次数
	}
}

// Build 根据Kafka结构体构建kafka.Config对象
func (k *Kafka) Build() *kafka.Config {
	return &kafka.Config{
		Username:     k.Username, // 用户名
		Password:     k.Password, // 密码
		ProducerAck:  k.ProducerAck, // 生产者确认模式
		CompressType: k.CompressType, // 压缩类型
		Addr:         k.Address, // Kafka地址
		TLS: kafka.TLSConfig{ // TLS配置
			EnableTLS:          k.Tls.EnableTLS, // 是否启用TLS
			CACrt:              k.Tls.CACrt, // CA证书
			ClientCrt:          k.Tls.ClientCrt, // 客户端证书
			ClientKey:          k.Tls.ClientKey, // 客户端密钥
			ClientKeyPwd:       k.Tls.ClientKeyPwd, // 客户端密钥密码
			InsecureSkipVerify: k.Tls.InsecureSkipVerify, // 是否跳过证书验证
		},
	}
}
// Build 构建Minio的配置信息。
// 返回值为配置好的minio.Config指针。
func (m *Minio) Build() *minio.Config {
    // 初始化Minio配置
    conf := minio.Config{
        Bucket:          m.Bucket,
        AccessKeyID:     m.AccessKeyID,
        SecretAccessKey: m.SecretAccessKey,
        SessionToken:    m.SessionToken,
        PublicRead:      m.PublicRead,
    }

    // 设置内部地址为Endpoint，如果内部地址是有效的主机和端口则添加http协议
    if _, _, err := net.SplitHostPort(m.InternalAddress); err == nil {
        conf.Endpoint = fmt.Sprintf("http://%s", m.InternalAddress)
    } else {
        conf.Endpoint = m.InternalAddress
    }

    // 设置外部签名地址，如果外部地址是有效的主机和端口则添加http协议
    if _, _, err := net.SplitHostPort(m.ExternalAddress); err == nil {
        conf.SignEndpoint = fmt.Sprintf("http://%s", m.ExternalAddress)
    } else {
        conf.SignEndpoint = m.ExternalAddress
    }

    return &conf
}

// Build 构建Cos的配置信息。
// 返回值为配置好的cos.Config指针。
func (c *Cos) Build() *cos.Config {
    // 初始化Cos配置
    return &cos.Config{
        BucketURL:    c.BucketURL,
        SecretID:     c.SecretID,
        SecretKey:    c.SecretKey,
        SessionToken: c.SessionToken,
        PublicRead:   c.PublicRead,
    }
}

// Build 构建Oss的配置信息。
// 返回值为配置好的oss.Config指针。
func (o *Oss) Build() *oss.Config {
    // 初始化Oss配置
    return &oss.Config{
        Endpoint:        o.Endpoint,
        Bucket:          o.Bucket,
        BucketURL:       o.BucketURL,
        AccessKeyID:     o.AccessKeyID,
        AccessKeySecret: o.AccessKeySecret,
        SessionToken:    o.SessionToken,
        PublicRead:      o.PublicRead,
    }
}

// Failed 返回失败缓存的过期时间。
// 返回值为失败缓存的过期时间Duration。
func (l *CacheConfig) Failed() time.Duration {
    // 计算失败缓存过期时间
    return time.Second * time.Duration(l.FailedExpire)
}

// Success 返回成功缓存的过期时间。
// 返回值为成功缓存的过期时间Duration。
func (l *CacheConfig) Success() time.Duration {
    // 计算成功缓存过期时间
    return time.Second * time.Duration(l.SuccessExpire)
}

// Enable 检查缓存配置是否启用。
// 返回值表示缓存是否启用的布尔值。
func (l *CacheConfig) Enable() bool {
    // 检查是否启用了缓存（非空的主题、槽数量和槽大小）
    return l.Topic != "" && l.SlotNum > 0 && l.SlotSize > 0
}
