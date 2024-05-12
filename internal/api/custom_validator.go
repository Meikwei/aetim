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

package api

import (
	"github.com/Meikwei/protocol/constant"
	"github.com/go-playground/validator/v10"
)

// RequiredIf 根据会话类型验证指定字段是否必需。
//
// 参数:
// fl - validator.FieldLevel类型，提供访问当前验证字段及其相关上下文的能力。
//
// 返回值:
// 返回一个布尔值，如果字段根据会话类型判断为必需且未提供有效值，则返回false；否则返回true。
func RequiredIf(fl validator.FieldLevel) bool {
    // 获取“SessionType”字段的值
    sessionType := fl.Parent().FieldByName("SessionType").Int()

    // 根据SessionType的值决定哪些字段是必需的
    switch sessionType {
    case constant.SingleChatType, constant.NotificationChatType:
        // 在单聊和通知聊天类型中，如果字段“RecvID”存在但为空，则验证失败
        return fl.FieldName() != "RecvID" || fl.Field().String() != ""
    case constant.WriteGroupChatType, constant.ReadGroupChatType:
        // 在写入和读取群组聊天类型中，如果字段“GroupID”存在但为空，则验证失败
        return fl.FieldName() != "GroupID" || fl.Field().String() != ""
    default:
        // 对于其他会话类型，默认所有字段都不必填
        return true
    }
}
