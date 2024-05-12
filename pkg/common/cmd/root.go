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
	"fmt"
	"path/filepath"

	"github.com/Meikwei/aetim/pkg/common/config"
	"github.com/Meikwei/go-tools/errs"
	"github.com/Meikwei/go-tools/log"
	"github.com/spf13/cobra"
)

// RootCmd 结构体定义了根命令的基本配置，包括命令本身、进程名、端口等配置。
type RootCmd struct {
	Command        cobra.Command
	processName    string      // 进程名称
	port           int         // 服务监听端口
	prometheusPort int         // Prometheus 监听端口
	log            config.Log  // 日志配置
	index          int         // 进程启动序列号
}

// Index 返回进程启动序列号。
func (r *RootCmd) Index() int {
	return r.index
}

// Port 返回服务监听端口。
func (r *RootCmd) Port() int {
	return r.port
}

// CmdOpts 定义了命令选项的结构，包括日志前缀名和配置映射。
type CmdOpts struct {
	loggerPrefixName string // 日志文件前缀名
	configMap        map[string]any // 配置映射
}

// WithCronTaskLogName 为命令选项添加一个预设的日志前缀名，用于定时任务日志。
func WithCronTaskLogName() func(*CmdOpts) {
	return func(opts *CmdOpts) {
		opts.loggerPrefixName = "openim-crontask"
	}
}

// WithLogName 为命令选项指定一个自定义的日志前缀名。
func WithLogName(logName string) func(*CmdOpts) {
	return func(opts *CmdOpts) {
		opts.loggerPrefixName = logName
	}
}

// WithConfigMap 为命令选项添加一个配置映射。
func WithConfigMap(configMap map[string]any) func(*CmdOpts) {
	return func(opts *CmdOpts) {
		opts.configMap = configMap
	}
}

// NewRootCmd 创建一个新的根命令实例，支持通过选项函数进行配置。
func NewRootCmd(processName string, opts ...func(*CmdOpts)) *RootCmd {
	rootCmd := &RootCmd{processName: processName}
	cmd := cobra.Command{
		Use:  "Start openIM application", // 命令的名称
		Long: fmt.Sprintf(`Start %s `, processName),//命令的详细描述
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.persistentPreRun(cmd, opts...)},//执行前函数
		SilenceUsage:  true,//表示是否在执行时静默用法信息
		SilenceErrors: false,//表示是否在执行时静默错误
	}
	cmd.Flags().StringP(FlagConf, "c", "", "path of config directory")
	cmd.Flags().IntP(FlagTransferIndex, "i", 0, "process startup sequence number")

	rootCmd.Command = cmd
	return rootCmd
}

// persistentPreRun 在命令执行前的持久化预处理函数，负责初始化配置和日志。
func (r *RootCmd) persistentPreRun(cmd *cobra.Command, opts ...func(*CmdOpts)) error {
	cmdOpts := r.applyOptions(opts...)
	if err := r.initializeConfiguration(cmd, cmdOpts); err != nil {
		return err
	}

	if err := r.initializeLogger(cmdOpts); err != nil {
		return errs.WrapMsg(err, "failed to initialize logger")
	}

	return nil
}

// initializeConfiguration 加载配置文件。
func (r *RootCmd) initializeConfiguration(cmd *cobra.Command, opts *CmdOpts) error {
	// 配置文件名称
	configDirectory, _, err := r.getFlag(cmd)
	if err != nil {
		return err
	}
	// 加载配置项
	for configFileName, configStruct := range opts.configMap {
		err := config.LoadConfig(filepath.Join(configDirectory, configFileName),
			ConfigEnvPrefixMap[configFileName], configStruct)
		if err != nil {
			return err
		}
	}
	// 加载日志配置
	return config.LoadConfig(filepath.Join(configDirectory, LogConfigFileName),
		ConfigEnvPrefixMap[LogConfigFileName], &r.log)
}

// applyOptions 应用命令选项。
func (r *RootCmd) applyOptions(opts ...func(*CmdOpts)) *CmdOpts {
	// 初始日志前缀
	cmdOpts := defaultCmdOpts()
	for _, opt := range opts {
		opt(cmdOpts)
	}

	return cmdOpts
}

// initializeLogger 初始化日志系统。
func (r *RootCmd) initializeLogger(cmdOpts *CmdOpts) error {
	err := log.InitFromConfig(
		cmdOpts.loggerPrefixName,
		r.processName,
		r.log.RemainLogLevel,
		r.log.IsStdout,
		r.log.IsJson,
		r.log.StorageLocation,
		r.log.RemainRotationCount,
		r.log.RotationTime,
		config.Version,
	)
	if err != nil {
		return errs.Wrap(err)
	}
	return errs.Wrap(log.InitConsoleLogger(r.processName, r.log.RemainLogLevel, r.log.IsJson, config.Version))
}

// defaultCmdOpts 提供默认的命令选项设置。
func defaultCmdOpts() *CmdOpts {
	return &CmdOpts{
		loggerPrefixName: "openim-service-log",
	}
}

// getFlag 从命令行标志中获取配置目录路径和进程启动序列号。
func (r *RootCmd) getFlag(cmd *cobra.Command) (string, int, error) {
	configDirectory, err := cmd.Flags().GetString(FlagConf)
	if err != nil {
		return "", 0, errs.Wrap(err)
	}
	index, err := cmd.Flags().GetInt(FlagTransferIndex)
	if err != nil {
		return "", 0, errs.Wrap(err)
	}
	r.index = index
	return configDirectory, index, nil
}

// Execute 执行根命令。
func (r *RootCmd) Execute() error {
	return r.Command.Execute()
}
