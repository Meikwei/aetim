<!--
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-28 22:59:31
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-04-28 23:06:37
 * @FilePath: \open-im-server\docs\README-zh_CN.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

# OpenIM Server Docs

OpenIM Server 是一个开源的即时通讯（IM）服务器项目，它提供了一系列的文档和指南，旨在帮助开发者和用户充分利用 OpenIM 的功能。

### 文档中心

- **贡献指南（Contrib）**：为开发者提供如何贡献代码、设置环境和遵循相关流程的详细指南。

  - [代码约定（Code Conventions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/code-conventions.md)：OpenIM 编写代码的规则和约定。
  - [开发指南（Development Guide）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/development.md)：在 OpenIM 中进行开发的指南。
  - [Git Cherry Pick 指南（Git Cherry Pick）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/gitcherry-pick.md)：关于 Cherry Pick 操作的指南。
  - [Git 工作流（Git Workflow）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/git-workflow.md)：OpenIM 中的 Git 工作流。
  - [初始化配置（Initialization Configurations）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/init-config.md)：设置和初始化 OpenIM 的指南。
  - [Docker 安装（Docker Installation）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/install-docker.md)：如何在机器上安装 Docker。
  - [Linux 开发环境（Linux Development Environment）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/linux-development.md)：在 Linux 上设置开发环境的指南。
  - [本地操作（Local Actions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/local-actions.md)：执行某些常见本地操作的指南。
  - [离线部署（Offline Deployment）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/offline-deployment.md)：部署 OpenIM 的离线方法。
  - [Protoc 工具（Protoc Tools）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/protoc-tools.md)：使用 Protoc 工具的指南。
  - [Go 工具（Go Tools）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/util-go.md)：OpenIM 中的 Go 工具和库。
  - [Makefile 工具（Makefile Tools）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/util-makefile.md)：Makefile 的最佳实践和工具。
  - [脚本工具（Script Tools）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/util-scripts.md)：脚本的最佳实践和工具。

- **转换（Conversions）**：介绍了 OpenIM 内部的各种约定和政策，包括代码、日志、版本等。
  - [API 转换（API Conversions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/api.md)：API 转换的指南和方法。
  - [日志策略（Logging Policy）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/bash-log.md)：OpenIM 中的日志策略和约定。
  - [CI/CD 操作（CI/CD Actions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/cicd-actions.md)：CI/CD 的流程和约定。
  - [提交约定（Commit Conventions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/commit.md)：OpenIM 中代码提交的约定。
  - [目录约定（Directory Conventions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/directory.md)：OpenIM 内部的目录结构和约定。
  - [错误代码（Error Codes）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/error-code.md)：错误代码的列表和描述。
  - [Go 代码转换（Go Code Conversions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/go-code.md)：Go 代码的约定和转换。
  - [Docker 镜像策略（Docker Image Strategy）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/images.md)：OpenIM Docker 镜像的管理策略，涵盖多种架构和镜像仓库。
  - [日志约定（Logging Conventions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/logging.md)：更详细的日志约定。
  - [版本约定（Version Conventions）](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/version.md)：OpenIM 版本的命名和管理策略。

### 开发者、贡献者和社区维护者

- **开发者和贡献者**：如果你是开发者或有意贡献的人：

  - 熟悉我们的[代码约定](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/code-conventions.md)和[Git 工作流](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/git-workflow.md)，以确保顺利贡献。
  - 深入[开发指南](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/development.md)，了解 OpenIM 的开发实践。

- **社区维护者**：作为社区维护者：
  - 确保贡献符合我们文档中概述的标准。
  - 定期查看[日志策略](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/bash-log.md)和[错误代码](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/error-code.md)，保持更新。

### 用户

- **用户**：用户应特别注意：
  - [Docker 安装](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/install-docker.md)：如果你计划使用 OpenIM 的 Docker 镜像，这是必要的。
  - [Docker 镜像策略](https://github.com/openimsdk/open-im-server/blob/main/docs/contrib/images.md)：了解不同镜像的可用性以及如何选择适合你架构的正确镜像。

OpenIM Server 的文档中心提供了全面的指南和手册，旨在帮助用户和开发者充分利用 OpenIM 的功能。如果你需要更多帮助或想要贡献代码，可以访问 OpenIM Server 的 [GitHub 仓库](https://github.com/openimsdk/open-im-server)。
