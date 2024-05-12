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

package cmd

import (
	"strings"
)

var (
	FileName                         string
	NotificationFileName             string
	ShareFileName                    string
	WebhooksConfigFileName           string
	LocalCacheConfigFileName         string
	KafkaConfigFileName              string
	RedisConfigFileName              string
	ZookeeperConfigFileName          string
	MongodbConfigFileName            string
	MinioConfigFileName              string
	LogConfigFileName                string
	OpenIMAPICfgFileName             string
	OpenIMCronTaskCfgFileName        string
	OpenIMMsgGatewayCfgFileName      string
	OpenIMMsgTransferCfgFileName     string
	OpenIMPushCfgFileName            string
	OpenIMRPCAuthCfgFileName         string
	OpenIMRPCConversationCfgFileName string
	OpenIMRPCFriendCfgFileName       string
	OpenIMRPCGroupCfgFileName        string
	OpenIMRPCMsgCfgFileName          string
	OpenIMRPCThirdCfgFileName        string
	OpenIMRPCUserCfgFileName         string
)

var ConfigEnvPrefixMap map[string]string

// init 函数初始化各种配置文件的名称，并构建环境变量映射。
// 该函数没有参数和返回值。
func init() {
    // 初始化配置文件名称
    FileName = "config.yaml"
    NotificationFileName = "notification.yml"
    ShareFileName = "share.yml"
    WebhooksConfigFileName = "webhooks.yml"
    LocalCacheConfigFileName = "local-cache.yml"
    KafkaConfigFileName = "kafka.yml"
    RedisConfigFileName = "redis.yml"
    ZookeeperConfigFileName = "zookeeper.yml"
    MongodbConfigFileName = "mongodb.yml"
    MinioConfigFileName = "minio.yml"
    LogConfigFileName = "log.yml"
    OpenIMAPICfgFileName = "openim-api.yml"
    OpenIMCronTaskCfgFileName = "openim-crontask.yml"
    OpenIMMsgGatewayCfgFileName = "openim-msggateway.yml"
    OpenIMMsgTransferCfgFileName = "openim-msgtransfer.yml"
    OpenIMPushCfgFileName = "openim-push.yml"
    OpenIMRPCAuthCfgFileName = "openim-rpc-auth.yml"
    OpenIMRPCConversationCfgFileName = "openim-rpc-conversation.yml"
    OpenIMRPCFriendCfgFileName = "openim-rpc-friend.yml"
    OpenIMRPCGroupCfgFileName = "openim-rpc-group.yml"
    OpenIMRPCMsgCfgFileName = "openim-rpc-msg.yml"
    OpenIMRPCThirdCfgFileName = "openim-rpc-third.yml"
    OpenIMRPCUserCfgFileName = "openim-rpc-user.yml"

    // 构建配置文件名到环境变量名的映射
    ConfigEnvPrefixMap = make(map[string]string)
    fileNames := []string{
        FileName, NotificationFileName, ShareFileName, WebhooksConfigFileName,
        KafkaConfigFileName, RedisConfigFileName, ZookeeperConfigFileName,
        MongodbConfigFileName, MinioConfigFileName, LogConfigFileName,
        OpenIMAPICfgFileName, OpenIMCronTaskCfgFileName, OpenIMMsgGatewayCfgFileName,
        OpenIMMsgTransferCfgFileName, OpenIMPushCfgFileName, OpenIMRPCAuthCfgFileName,
        OpenIMRPCConversationCfgFileName, OpenIMRPCFriendCfgFileName, OpenIMRPCGroupCfgFileName,
        OpenIMRPCMsgCfgFileName, OpenIMRPCThirdCfgFileName, OpenIMRPCUserCfgFileName,
    }

    // 遍历文件名列表，构建环境变量映射
    for _, fileName := range fileNames {
        envKey := strings.TrimSuffix(strings.TrimSuffix(fileName, ".yml"), ".yaml")
        envKey = "IMENV_" + envKey
        envKey = strings.ToUpper(strings.ReplaceAll(envKey, "-", "_"))
        ConfigEnvPrefixMap[fileName] = envKey
    }
}

const (
	FlagConf          = "config_folder_path"
	FlagTransferIndex = "index"
)
