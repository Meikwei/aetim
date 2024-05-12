# 持续集成与自动化

OpenIM 仓库的每次更改，无论是通过拉取请求还是直接推送，都会触发仓库内定义的持续集成流水线。不用说，所有 OpenIM 的贡献都必须在所有检查通过（即拥有绿色构建）后才能合并。

## CI 平台

目前，有两个不同的平台参与运行 CI 流程：

- GitHub Actions
- CNCF 基础设施上的 Drone 流水线

### GitHub Actions

所有现有的 GitHub Actions 都定义为 `.github/workflows` 目录下的 YAML 文件。这些可以归类为：

- **PR 检查**：这些操作在 PR 创建和更新时运行所有必需的验证。涵盖了 DCO 合规性检查、`x86_64` 测试套件（单元测试、集成测试、冒烟测试）和代码覆盖率。
- **仓库自动化**：目前，它只涵盖问题和史诗（epic）整理。

所有操作都在 GitHub 提供的运行器上运行；因此，测试限于在 `x86_64` 架构上运行。

## 在本地运行

为了加快拉取请求过程，贡献者应该在本地验证他们的更改。幸运的是，除了发布步骤外，所有 CI 步骤都可以通过以下任一方法在本地环境中运行：

**用户 Makefile：**

```bash
root@PS2023EVRHNCXG:~/workspaces/openim/Open-IM-Server# make help 😊

Usage: make <TARGETS> <OPTIONS> ...

Targets:

all                          Run tidy, gen, add-copyright, format, lint, cover, build 🚀
build                        Build binaries by default 🛠️
multiarch                    Build binaries for multiple platforms. See option PLATFORMS. 🌍
tidy                         tidy go.mod ✨
vendor                       vendor go.mod 📦
style                        code style -> fmt,vet,lint 💅
fmt                          Run go fmt against code. ✨
vet                          Run go vet against code. ✅
lint                         Check syntax and styling of go sources. ✔️
format                       Gofmt (reformat) package sources (exclude vendor dir if existed). 🔄
test                         Run unit test. 🧪
cover                        Run unit test and get test coverage. 📊
updates                      Check for updates to go.mod dependencies 🆕
imports                      task to automatically handle import packages in Go files using goimports tool 📥
clean                        Remove all files that are created by building. 🗑️
image                        Build docker images for host arch. 🐳
image.multiarch              Build docker images for multiple platforms. See option PLATFORMS. 🌍🐳
push                         Build docker images for host arch and push images to registry. 📤🐳
push.multiarch               Build docker images for multiple platforms and push images to registry. 🌍📤🐳
tools                        Install dependent tools. 🧰
gen                          Generate all necessary files. 🧩
swagger                      Generate swagger document. 📖
serve-swagger                Serve swagger spec and docs. 🚀📚
verify-copyright             Verify the license headers for all files. ✅
add-copyright                Add copyright ensure source code files have license headers. 📄
release                      release the project 🎉
help                         Show this help info. ℹ️
help-all                     Show all help details info. ℹ️📚

Options:

DEBUG            Whether or not to generate debug symbols. Default is 0. ❓

BINS             Binaries to build. Default is all binaries under cmd. 🛠️
This option is available when using: make {build}(.multiarch) 🧰
Example: make build BINS="openim-api openim_cms_api".

PLATFORMS        Platform to build for. Default is linux_arm64 and linux_amd64. 🌍
This option is available when using: make {build}.multiarch 🌍
Example: make multiarch PLATFORMS="linux_s390x linux_mips64
linux_mips64le darwin_amd64 windows_amd64 linux_amd64 linux_arm64".

V                Set to 1 enable verbose build. Default is 0. 📝
```

如何使用 Makefile 帮助贡献者快速构建项目 😊

`make help` 命令是一个实用工具，它提供了如何有效使用 Makefile 的有用信息。通过运行此命令，贡献者将了解各种目标和选项，以便快速构建项目。

以下是 Makefile 提供的目标和选项的分解：

**目标 😃**

1. `all`：此目标运行多个任务，如 `tidy`、`gen`、`add-copyright`、`format`、`lint`、`cover` 和 `build`。它确保全面构建项目。
2. `build`：主要目标，默认情况下编译二进制文件。它特别适用于创建必要的可执行文件。
3. `multiarch`：为目标平台构建二进制文件的目标。贡献者可以使用 `PLATFORMS` 选项指定所需的平台。
4. `tidy`：此目标清理 `go.mod` 文件，确保其一致性。
5. `vendor`：根据 `go.mod` 文件更新项目依赖的目标。
6. `style`：使用 `fmt`、`vet` 和 `lint` 等工具检查代码风格。它确保整个项目中代码风格的一致性。
7. `fmt`：使用 `go fmt` 命令格式化代码，确保适当的缩进和格式化。
8. `vet`：运行 `go vet` 命令识别代码中的常见错误。
9. `lint`：使用 linter 验证 Go 源文件的语法和风格。
10. `format`：使用 `gofmt` 重新格式化包源代码。如果存在，它将排除 vendor 目录。
11. `test`：执行单元测试以确保代码的功能性和稳定性。
12. `cover`：执行单元测试并计算代码的测试覆盖率。
13. `updates`：检查 `go.mod` 文件中指定的项目依赖项的更新。
14. `imports`：使用 `goimports` 工具自动处理 Go 文件中的导入包。
15. `clean`：删除构建过程中创建的所有文件，有效清理项目目录。
16. `image`：为主机架构构建 Docker 镜像。
17. `image.multiarch`：与 `image` 目标类似，但它为多个平台构建 Docker 镜像。贡献者可以使用 `PLATFORMS` 选项指定所需的平台。
18. `push`：为主机架构构建 Docker 镜像并将它们推送到注册表。
19. `push.multiarch`：为多个平台构建 Docker 镜像并将它们推送到注册表。贡献者可以使用 `PLATFORMS` 选项指定所需的平台。
20. `tools`：安装项目所需的工具或依赖。
21. `gen`：自动生成所有必需的文件。
22. `swagger`：为项目生成 swagger 文档。
23. `serve-swagger`：提供 swagger 规范和文档。
24. `verify-copyright`：验证所有项目文件的许可证头。
25. `add-copyright`：为源代码文件添加版权头。
26. `release`：发布项目，假定是为了分发。
27. `help`：显示有关可用目标和选项的信息。
28. `help-all`：显示所有可用目标和选项的详细信息。

**选项 😄**

1. `DEBUG`：一个布尔选项，用于确定是否生成调试符号。默认值为 0（false）。
2. `BINS`：指定要构建的二进制文件。默认情况下，它在 `cmd` 目录下构建所有二进制文件。贡献者可以使用此选项提供特定二进制文件的列表。
3. `PLATFORMS`：指定要为其构建的平台。默认平台是 `linux_arm64` 和 `linux_amd64`。贡献者可以通过提供以空格分隔的平台名称列表来指定多个平台。
4. `V`：一个布尔选项，当设置为 1（true）时启用详细构建输出。默认值为 0（false）。

有了这些目标和选项，贡献者可以有效地使用 Makefile 构建项目。编码愉快！🚀😊
