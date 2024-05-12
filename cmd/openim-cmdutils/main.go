/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-12 23:04:01
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 23:10:18
 * @FilePath: \aetim\cmd\openim-cmdutils\main.go
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

package main

import (
	"github.com/Meikwei/aetim/pkg/common/cmd"
	"github.com/Meikwei/go-tools/system/program"
)

func main() {
	msgUtilsCmd := cmd.NewMsgUtilsCmd("openIMCmdUtils", "openIM cmd utils", nil)
	getCmd := cmd.NewGetCmd()
	fixCmd := cmd.NewFixCmd()
	clearCmd := cmd.NewClearCmd()
	seqCmd := cmd.NewSeqCmd()
	msgCmd := cmd.NewMsgCmd()
	getCmd.AddCommand(seqCmd.GetSeqCmd(), msgCmd.GetMsgCmd())
	getCmd.AddSuperGroupIDFlag()
	getCmd.AddUserIDFlag()
	getCmd.AddConfigDirFlag()
	getCmd.AddIndexFlag()
	getCmd.AddBeginSeqFlag()
	getCmd.AddLimitFlag()
	// openIM get seq --userID=xxx
	// openIM get seq --superGroupID=xxx
	// openIM get msg --userID=xxx --beginSeq=100 --limit=10
	// openIM get msg --superGroupID=xxx --beginSeq=100 --limit=10

	fixCmd.AddCommand(seqCmd.FixSeqCmd())
	fixCmd.AddSuperGroupIDFlag()
	fixCmd.AddUserIDFlag()
	fixCmd.AddConfigDirFlag()
	fixCmd.AddIndexFlag()
	fixCmd.AddFixAllFlag()
	// openIM fix seq --userID=xxx
	// openIM fix seq --superGroupID=xxx
	// openIM fix seq --fixAll

	clearCmd.AddCommand(msgCmd.ClearMsgCmd())
	clearCmd.AddSuperGroupIDFlag()
	clearCmd.AddUserIDFlag()
	clearCmd.AddConfigDirFlag()
	clearCmd.AddIndexFlag()
	clearCmd.AddClearAllFlag()
	clearCmd.AddBeginSeqFlag()
	clearCmd.AddLimitFlag()
	// openIM clear msg --userID=xxx --beginSeq=100 --limit=10
	// openIM clear msg --superGroupID=xxx --beginSeq=100 --limit=10
	// openIM clear msg --clearAll
	msgUtilsCmd.AddCommand(&getCmd.Command, &fixCmd.Command, &clearCmd.Command)
	if err := msgUtilsCmd.Execute(); err != nil {
		program.ExitWithError(err)
	}
}
