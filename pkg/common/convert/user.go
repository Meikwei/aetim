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

package convert

import (
	"time"

	relationtb "github.com/Meikwei/aetim/pkg/common/db/table/relation"
	"github.com/Meikwei/protocol/sdkws"
)

// UsersDB2Pb 将用户数据库模型转换为protobuf消息格式
// 参数 users: 由relationtb.UserModel类型的用户模型组成的切片
// 返回值: 由sdkws.UserInfo类型的用户信息组成的切片
func UsersDB2Pb(users []*relationtb.UserModel) []*sdkws.UserInfo {
    result := make([]*sdkws.UserInfo, 0, len(users))
    for _, user := range users {
        userPb := &sdkws.UserInfo{
            UserID:           user.UserID,
            Nickname:         user.Nickname,
            FaceURL:          user.FaceURL,
            Ex:               user.Ex,
            CreateTime:       user.CreateTime.UnixMilli(),
            AppMangerLevel:   user.AppMangerLevel,
            GlobalRecvMsgOpt: user.GlobalRecvMsgOpt,
        }
        result = append(result, userPb)
    }
    return result
}

// UserPb2DB 将protobuf用户信息转换为数据库模型
// 参数 user: sdkws.UserInfo类型的用户信息
// 返回值: relationtb.UserModel类型的用户模型
func UserPb2DB(user *sdkws.UserInfo) *relationtb.UserModel {
    return &relationtb.UserModel{
        UserID:           user.UserID,
        Nickname:         user.Nickname,
        FaceURL:          user.FaceURL,
        Ex:               user.Ex,
        CreateTime:       time.UnixMilli(user.CreateTime),
        AppMangerLevel:   user.AppMangerLevel,
        GlobalRecvMsgOpt: user.GlobalRecvMsgOpt,
    }
}

// UserPb2DBMap 将protobuf用户信息转换为包含部分字段的map
// 参数 user: sdkws.UserInfo类型的用户信息
// 返回值: 包含用户部分字段的map[string]any类型
func UserPb2DBMap(user *sdkws.UserInfo) map[string]any {
    if user == nil {
        return nil
    }
    val := make(map[string]any)
    fields := map[string]any{
        "nickname":            user.Nickname,
        "face_url":            user.FaceURL,
        "ex":                  user.Ex,
        "app_manager_level":   user.AppMangerLevel,
        "global_recv_msg_opt": user.GlobalRecvMsgOpt,
    }
    // 筛选出非空字段添加到val中
    for key, value := range fields {
        if v, ok := value.(string); ok && v != "" {
            val[key] = v
        } else if v, ok := value.(int32); ok && v != 0 {
            val[key] = v
        }
    }
    return val
}

// UserPb2DBMapEx 将包含扩展信息的protobuf用户信息转换为包含部分字段的map
// 参数 user: sdkws.UserInfoWithEx类型的用户信息
// 返回值: 包含用户部分字段的map[string]any类型
func UserPb2DBMapEx(user *sdkws.UserInfoWithEx) map[string]any {
    if user == nil {
        return nil
    }
    val := make(map[string]any)

    // 从UserInfoWithEx中映射字段到val
    if user.Nickname != nil {
        val["nickname"] = user.Nickname.Value
    }
    if user.FaceURL != nil {
        val["face_url"] = user.FaceURL.Value
    }
    if user.Ex != nil {
        val["ex"] = user.Ex.Value
    }
    if user.GlobalRecvMsgOpt != nil {
        val["global_recv_msg_opt"] = user.GlobalRecvMsgOpt.Value
    }

    return val
}
