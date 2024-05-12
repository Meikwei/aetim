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
	_ "net/http/pprof"

	"github.com/Meikwei/aetim/pkg/common/cmd"
	"github.com/Meikwei/go-tools/system/program"
)

// main 是程序的入口点。
// 该函数不接受参数，也没有返回值。
// 主要逻辑是创建一个新的ApiCmd实例并执行，如果执行过程中出现错误，则通过program.ExitWithError退出程序。
func main() {
    // 创建并执行ApiCmd，处理可能发生的错误
	if err := cmd.NewApiCmd().Exec(); err != nil {
		// 如果执行过程中出现错误，以错误方式退出程序
		program.ExitWithError(err)
	}

    // 如果没有错误发生，程序将继续运行
}
