/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-12 20:40:23
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 21:00:05
 * @FilePath: \aetim\pkg\common\db\controller\auth.go
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

	"github.com/Meikwei/aetim/pkg/authverify"
	"github.com/Meikwei/aetim/pkg/common/db/cache"
	"github.com/Meikwei/go-tools/errs"
	"github.com/Meikwei/go-tools/tokenverify"
	"github.com/Meikwei/protocol/constant"
	"github.com/golang-jwt/jwt/v4"
)

type AuthDatabase interface {
	// If the result is empty, no error is returned.
	GetTokensWithoutError(ctx context.Context, userID string, platformID int) (map[string]int, error)
	// Create token
	CreateToken(ctx context.Context, userID string, platformID int) (string, error)

	SetTokenMapByUidPid(ctx context.Context, userID string, platformID int, m map[string]int) error
}

type authDatabase struct {
	cache        cache.TokenModel
	accessSecret string
	accessExpire int64
}

func NewAuthDatabase(cache cache.TokenModel, accessSecret string, accessExpire int64) AuthDatabase {
	return &authDatabase{cache: cache, accessSecret: accessSecret, accessExpire: accessExpire}
}

// If the result is empty.
func (a *authDatabase) GetTokensWithoutError(ctx context.Context, userID string, platformID int) (map[string]int, error) {
	return a.cache.GetTokensWithoutError(ctx, userID, platformID)
}

func (a *authDatabase) SetTokenMapByUidPid(ctx context.Context, userID string, platformID int, m map[string]int) error {
	return a.cache.SetTokenMapByUidPid(ctx, userID, platformID, m)
}

// Create Token.
func (a *authDatabase) CreateToken(ctx context.Context, userID string, platformID int) (string, error) {
	tokens, err := a.cache.GetTokensWithoutError(ctx, userID, platformID)
	if err != nil {
		return "", err
	}
	var deleteTokenKey []string
	for k, v := range tokens {
		_, err = tokenverify.GetClaimFromToken(k, authverify.Secret(a.accessSecret))
		if err != nil || v != constant.NormalToken {
			deleteTokenKey = append(deleteTokenKey, k)
		}
	}
	if len(deleteTokenKey) != 0 {
		err = a.cache.DeleteTokenByUidPid(ctx, userID, platformID, deleteTokenKey)
		if err != nil {
			return "", err
		}
	}

	claims := tokenverify.BuildClaims(userID, platformID, a.accessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.accessSecret))
	if err != nil {
		return "", errs.WrapMsg(err, "token.SignedString")
	}
	return tokenString, a.cache.AddTokenFlag(ctx, userID, platformID, tokenString, constant.NormalToken)
}
