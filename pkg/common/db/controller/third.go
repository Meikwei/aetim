/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-12 20:40:23
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 20:46:05
 * @FilePath: \aetim\pkg\common\db\controller\third.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

package controller

import (
	"context"
	"time"

	"github.com/Meikwei/aetim/pkg/common/db/cache"
	"github.com/Meikwei/aetim/pkg/common/db/table/relation"
	"github.com/Meikwei/go-tools/db/pagination"
)

type ThirdDatabase interface {
	FcmUpdateToken(ctx context.Context, account string, platformID int, fcmToken string, expireTime int64) error
	SetAppBadge(ctx context.Context, userID string, value int) error
	// about log for debug
	UploadLogs(ctx context.Context, logs []*relation.LogModel) error
	DeleteLogs(ctx context.Context, logID []string, userID string) error
	SearchLogs(ctx context.Context, keyword string, start time.Time, end time.Time, pagination pagination.Pagination) (int64, []*relation.LogModel, error)
	GetLogs(ctx context.Context, LogIDs []string, userID string) ([]*relation.LogModel, error)
}

type thirdDatabase struct {
	cache cache.ThirdCache
	logdb relation.LogInterface
}

// DeleteLogs implements ThirdDatabase.
func (t *thirdDatabase) DeleteLogs(ctx context.Context, logID []string, userID string) error {
	return t.logdb.Delete(ctx, logID, userID)
}

// GetLogs implements ThirdDatabase.
func (t *thirdDatabase) GetLogs(ctx context.Context, LogIDs []string, userID string) ([]*relation.LogModel, error) {
	return t.logdb.Get(ctx, LogIDs, userID)
}

// SearchLogs implements ThirdDatabase.
func (t *thirdDatabase) SearchLogs(ctx context.Context, keyword string, start time.Time, end time.Time, pagination pagination.Pagination) (int64, []*relation.LogModel, error) {
	return t.logdb.Search(ctx, keyword, start, end, pagination)
}

// UploadLogs implements ThirdDatabase.
func (t *thirdDatabase) UploadLogs(ctx context.Context, logs []*relation.LogModel) error {
	return t.logdb.Create(ctx, logs)
}

func NewThirdDatabase(cache cache.ThirdCache, logdb relation.LogInterface) ThirdDatabase {
	return &thirdDatabase{cache: cache, logdb: logdb}
}

func (t *thirdDatabase) FcmUpdateToken(ctx context.Context, account string, platformID int, fcmToken string, expireTime int64) error {
	return t.cache.SetFcmToken(ctx, account, platformID, fcmToken, expireTime)
}

func (t *thirdDatabase) SetAppBadge(ctx context.Context, userID string, value int) error {
	return t.cache.SetUserBadgeUnreadCountSum(ctx, userID, value)
}
