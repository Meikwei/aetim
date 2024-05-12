<!--
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-29 19:42:04
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-04-29 20:36:59
 * @FilePath: \open-im-server\docs\contrib\directory-zh_CN.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

## 🤖 项目目录结构

标准化项目的文件和目录组织:

```bash
.go-project-layout
├── CONTRIBUTING.md          # 贡献指南，提供给想要为项目贡献代码的人的指南和步骤
├── LICENSE                  # 项目的许可证文件，说明了项目遵循的法律条款
├── Makefile                 # 一个用于自动化项目构建和测试等流程的配置文件
├── README.md                # 项目概述，通常用英语编写，介绍项目的基本信息
├── README_zh-CN.md          # 项目概述的中文版本
├── api                      # 与项目 API 相关的文件和文档
│   ├── OWNERS               # 标识 API 部分的责任人或团队
│   └── README.md            # API 的文档说明
├── assets                   # 存放静态资源，如图片和样式表
│   └── README.md            # 静态资源的文档说明
├── build                    # 包含与项目构建过程相关的文件
│   ├── OWNERS               # 构建过程的责任人或团队
│   └── README.md            # 构建过程的文档说明
├── cmd                      # 存放命令行工具和程序的入口点
│   ├── OWNERS               # 命令行工具的责任人或团队
│   └── README.md            # 命令行工具的文档说明
├── configs                  # 存放配置文件
│   ├── OWNERS               # 配置文件的责任人或团队
│   ├── README.md            # 配置文件的文档说明
│   └── config.yaml          # 主要的配置文件
├── deploy                   # 与项目部署相关的文件
│   ├── OWNERS               #  部署过程的责任人或团队
│   └── README.md            # 部署过程的文档说明
├── docs                     # 存放项目文档
│   ├── OWNERS               # 文档的责任人或团队
│   └── README.md            # 文档索引
├── examples                 # 包含示例代码和用法
│   ├── OWNERS               # 示例代码的责任人或团队
│   └── README.md            # 示例代码的文档说明
├── init                     # 包含初始化文件
│   ├── OWNERS               # 初始化文件的责任人或团队
│   └── README.md            # 初始化文件的文档说明
├── internal                 # 包含内部应用程序代码
│   ├── OWNERS               # 内部代码的责任人或团队
│   ├── README.md            # 内部代码的文档说明
│   ├── app                  # 应用程序逻辑相关的代码
│   ├── pkg                  # 内部使用的软件包
│   └── utils                # 实用工具函数和帮助程序
├── pkg                      # 包含公共的软件包和库
│   ├── OWNERS               # 软件包的责任人或团队
│   ├── README.md            # 软件包的文档说明
│   ├── common               # 通用的实用工具和帮助程序
│   ├── log                  # 日志相关的实用工具
│   ├── tools                # 开发工具和脚本
│   ├── utils                # 通用的实用函数
│   └── version              # 版本信息
├── scripts                  # 包含开发和自动化脚本
│   ├── LICENSE_TEMPLATES    # 许可证模板文件
│   ├── OWNERS               # 脚本的责任人或团队
│   ├── README.md            # 脚本的文档说明
│   ├── githooks             # 用于开发的 Git 钩子
│   └── make-rules           # Makefile 规则和脚本
├── test                     # 包含测试文件和与测试相关的实用工具
│   ├── OWNERS               # 测试文件的责任人或团队
│   └── README.md            # 测试文件的文档说明
├── third_party              # 包含第三方依赖和库
│   └── README.md            # 第三方组件的文档说明
├── tools                    # 包含开发工具和实用工具
│   └── README.md            # Tool documentation
└── web                      # 包含与 Web 相关的文件，如 HTML 和 CSS
    ├── OWNERS               # Web 相关文件的责任人或团队
    └── README.md            # Web 相关文件的文档说明
```
