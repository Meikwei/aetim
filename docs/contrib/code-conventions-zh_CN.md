<!--
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-29 19:25:14
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-04-29 19:28:57
 * @FilePath: \open-im-server\docs\contrib\code-conventions-zh_CN.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

# 代码约定(Code conventions)

### POSIX shell

- 提到了一个风格指南，用于 POSIX shell 脚本的编写。

### Go

- 推荐阅读 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) 和 [Effective Go](https://golang.org/doc/effective_go.html)。
- 避免 [Go 语言的常见陷阱](https://gist.github.com/lavalamp/4bd23295a9f32706a48f)。
- 代码注释很重要，应该遵循 [Go 的注释约定](http://blog.golang.org/godoc-documenting-go-code)。
- 命令行标志应使用连字符（`-`），而不是下划线（`_`）。
- 命名约定，包括包名、接口名和目录名的选择。

### OpenIM 命名约定指南

- 介绍了 OpenIM 项目遵循的最佳实践和标准化命名约定，以保持清晰、一致，并与行业标准保持一致。

#### 1. 通用文件命名

- 文件名中可以使用连字符（`-`）和下划线（`_`），通常偏好使用下划线以提高可读性和兼容性。

#### 2. 特殊文件类型

- 脚本和 Markdown 文件应使用连字符（`-`）。
- 大写 Markdown 文档文件（如 `README`）可以使用下划线（`_`）分隔单词。

#### 3. 目录命名

- 目录名必须使用连字符（`-`）。

#### 4. 配置文件

- 配置文件应使用连字符（`-`）。

#### 最佳实践

- 名称应简洁但具有描述性，足以一目了然地传达文件的目的或内容。
- 避免在名称中使用空格；改用连字符或下划线以提高不同操作系统和环境中的兼容性。
- 尽可能使用小写命名，以保持一致性并避免与大小写敏感的系统发生问题。
- 如果文件需要更新，文件名中应包含版本号或日期。

### 目录和文件约定

- 避免使用通用的实用程序包，而应选择一个明确描述其目的的名称。
- 所有文件名、脚本文件、配置文件和目录应使用小写字母，并使用连字符（`-`）作为分隔符。
- Go 语言文件的文件名应使用小写字母并使用下划线（`_`）。
- 包名应与目录名匹配，以确保一致性。

### 测试约定

- 有关测试的约定，请参考 [TESTING.md](https://github.com/openimsdk/open-im-server/tree/main/test/readme) 文档。

这些约定有助于项目维护者、开发者和贡献者遵循一致的编码风格，从而使得代码库更加易于理解和维护。如果您需要查看具体的测试约定，可以访问提供的链接。
