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

package msggateway

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Meikwei/aetim/pkg/common/servererrs"

	"github.com/Meikwei/go-tools/utils/encrypt"
	"github.com/Meikwei/go-tools/utils/stringutil"
	"github.com/Meikwei/go-tools/utils/timeutil"
	"github.com/Meikwei/protocol/constant"
)

type UserConnContext struct {
	RespWriter http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	RemoteAddr string
	ConnID     string
}

func (c *UserConnContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *UserConnContext) Done() <-chan struct{} {
	return nil
}

func (c *UserConnContext) Err() error {
	return nil
}

func (c *UserConnContext) Value(key any) any {
	switch key {
	case constant.OpUserID:
		return c.GetUserID()
	case constant.OperationID:
		return c.GetOperationID()
	case constant.ConnID:
		return c.GetConnID()
	case constant.OpUserPlatform:
		return constant.PlatformIDToName(stringutil.StringToInt(c.GetPlatformID()))
	case constant.RemoteAddr:
		return c.RemoteAddr
	default:
		return ""
	}
}

func newContext(respWriter http.ResponseWriter, req *http.Request) *UserConnContext {
	return &UserConnContext{
		RespWriter: respWriter,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
		RemoteAddr: req.RemoteAddr,
		ConnID:     encrypt.Md5(req.RemoteAddr + "_" + strconv.Itoa(int(timeutil.GetCurrentTimestampByMill()))),
	}
}

func newTempContext() *UserConnContext {
	return &UserConnContext{
		Req: &http.Request{URL: &url.URL{}},
	}
}

func (c *UserConnContext) GetRemoteAddr() string {
	return c.RemoteAddr
}

func (c *UserConnContext) Query(key string) (string, bool) {
	var value string
	if value = c.Req.URL.Query().Get(key); value == "" {
		return value, false
	}
	return value, true
}

func (c *UserConnContext) GetHeader(key string) (string, bool) {
	var value string
	if value = c.Req.Header.Get(key); value == "" {
		return value, false
	}
	return value, true
}

func (c *UserConnContext) SetHeader(key, value string) {
	c.RespWriter.Header().Set(key, value)
}

func (c *UserConnContext) ErrReturn(error string, code int) {
	http.Error(c.RespWriter, error, code)
}

func (c *UserConnContext) GetConnID() string {
	return c.ConnID
}

func (c *UserConnContext) GetUserID() string {
	return c.Req.URL.Query().Get(WsUserID)
}

func (c *UserConnContext) GetPlatformID() string {
	return c.Req.URL.Query().Get(PlatformID)
}

func (c *UserConnContext) GetOperationID() string {
	return c.Req.URL.Query().Get(OperationID)
}

func (c *UserConnContext) SetOperationID(operationID string) {
	values := c.Req.URL.Query()
	values.Set(OperationID, operationID)
	c.Req.URL.RawQuery = values.Encode()
}

func (c *UserConnContext) GetToken() string {
	return c.Req.URL.Query().Get(Token)
}

func (c *UserConnContext) GetCompression() bool {
	compression, exists := c.Query(Compression)
	if exists && compression == GzipCompressionProtocol {
		return true
	} else {
		compression, exists := c.GetHeader(Compression)
		if exists && compression == GzipCompressionProtocol {
			return true
		}
	}
	return false
}

func (c *UserConnContext) ShouldSendResp() bool {
	errResp, exists := c.Query(SendResponse)
	if exists {
		b, err := strconv.ParseBool(errResp)
		if err != nil {
			return false
		} else {
			return b
		}
	}
	return false
}

func (c *UserConnContext) SetToken(token string) {
	c.Req.URL.RawQuery = Token + "=" + token
}

func (c *UserConnContext) GetBackground() bool {
	b, err := strconv.ParseBool(c.Req.URL.Query().Get(BackgroundStatus))
	if err != nil {
		return false
	}
	return b
}
func (c *UserConnContext) ParseEssentialArgs() error {
	_, exists := c.Query(Token)
	if !exists {
		return servererrs.ErrConnArgsErr.WrapMsg("token is empty")
	}
	_, exists = c.Query(WsUserID)
	if !exists {
		return servererrs.ErrConnArgsErr.WrapMsg("sendID is empty")
	}
	platformIDStr, exists := c.Query(PlatformID)
	if !exists {
		return servererrs.ErrConnArgsErr.WrapMsg("platformID is empty")
	}
	_, err := strconv.Atoi(platformIDStr)
	if err != nil {
		return servererrs.ErrConnArgsErr.WrapMsg("platformID is not int")

	}
	return nil
}
