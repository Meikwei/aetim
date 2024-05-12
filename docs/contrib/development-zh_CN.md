<!--
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-29 19:31:11
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-04-29 19:33:05
 * @FilePath: \open-im-server\docs\contrib\development-zh_CN.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

# 开发指南

由于 OpenIM 是用 Go 语言编写的，可以合理假设 Go 工具是贡献该项目所需的全部。不幸的是，当需要测试或构建本地更改时，这一点不再成立。本文档详细说明了 OpenIM 开发所需的工具。

## 非 Linux 环境先决条件

此仓库中的所有测试和构建脚本都是在 GNU Linux 开发环境中创建的。因此，建议使用此仓库的 [Vagrantfile](https://developer.hashicorp.com/vagrant/docs/vagrantfile) 定义的虚拟机来使用它们。

无论如何，如果有人仍然希望在非 Linux 环境中构建和测试 OpenIM，需要遵循特定的设置。

### Windows 设置

只有在支持 Windows Subsystem for Linux (WSL) 的 Windows 版本上才可能构建 OpenIM。如果开发环境使用的是 Windows 10，版本 2004，构建 19041 或更高版本，请[按照这些说明安装 WSL2](https://docs.microsoft.com/en-us/windows/wsl/install-win10)；否则，使用 Linux 虚拟机。

### macOS 设置

负责构建和测试流程的 shell 脚本依赖于 GNU 工具（例如 `sed`），[在 macOS 上略有不同](https://unix.stackexchange.com/a/79357)，这意味着在使用它们之前必须进行一些调整。

首先，安装 GNU 工具：

```sh
brew install coreutils findutils gawk gnu-sed gnu-tar grep make
```

然后更新 shell 初始化脚本（例如 `.bashrc`）以在 `$PATH` 变量前添加 GNU Utils

```sh
GNUBINS="$(find /usr/local/opt -type d -follow -name gnubin -print)"

for bindir in ${GNUBINS[@]}; do
  PATH=$bindir:$PATH
done

export PATH
```

## 安装所需软件

### Go

众所周知，OpenIM 是用 [Go](http://golang.org) 编写的。请按照 [Go 开始使用指南](https://golang.org/doc/install) 安装和设置用于编译和运行测试套件的 Go 工具。

| OpenIM      | 需要 Go |
| ----------- | ------- |
| 2.24 - 3.00 | 1.15 +  |
| 3.30 +      | 1.18 +  |

### Docker

OpenIM 构建和测试流程的开发需要 Docker 来运行某些步骤。[按照 Docker 网站说明在开发环境中安装 Docker](https://docs.docker.com/get-docker/)。

### Vagrant

如[测试文档](https://github.com/openimsdk/open-im-server/tree/main/test/readme)所述，所有的冒烟测试都是在由 Vagrant 管理的虚拟机中运行的。要在开发环境中安装 Vagrant，[按照 Hashicorp 网站的说明进行](https://www.vagrantup.com/downloads)，以及以下任何虚拟化软件：

- [VirtualBox](https://www.virtualbox.org/)
- [libvirt](https://libvirt.org/) 以及 [vagrant-libvirt 插件](https://github.com/vagrant-libvirt/vagrant-libvirt#installation)

## 依赖管理

OpenIM 使用 [go modules](https://github.com/golang/go/wiki/Modules) 来管理依赖。

---

请注意，由于网络原因，一些链接可能无法直接解析。如果您需要访问这些资源，建议您检查网络连接或直接访问相关网站以获取更多信息。如果您在设置开发环境或安装所需软件时遇到问题，可以查阅官方文档或寻求社区支持。
