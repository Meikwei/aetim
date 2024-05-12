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

package prommetrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	SingleChatMsgProcessSuccessCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "single_chat_msg_process_success_total",
		Help: "The number of single chat msg successful processed",
	})
	SingleChatMsgProcessFailedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "single_chat_msg_process_failed_total",
		Help: "The number of single chat msg failed processed",
	})
	GroupChatMsgProcessSuccessCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "group_chat_msg_process_success_total",
		Help: "The number of group chat msg successful processed",
	})
	GroupChatMsgProcessFailedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "group_chat_msg_process_failed_total",
		Help: "The number of group chat msg failed processed",
	})
)
