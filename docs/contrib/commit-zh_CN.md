<!--
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-29 19:37:50
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-04-29 19:38:36
 * @FilePath: \open-im-server\docs\contrib\commit-zh_CN.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

## 提交标准

我们的项目 OpenIM 遵循 [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0) 标准。

> 中文翻译：[Conventional Commits：一种使提交日志更易于人类和机器理解的规范](https://tool.lu/en_US/article/2ac/preview)

除了遵守这些标准外，我们还鼓励所有 OpenIM 项目的参与者确保他们的提交信息清晰且描述性强。这有助于维护一个清晰且有意义的项目历史。每个提交信息都应该简洁地描述所做的更改，必要时还应说明这些更改背后的原因。

为了促进流程的简化，我们还建议根据 Conventional Commits 指南使用适当的提交类型，例如使用 `fix:` 表示修复漏洞，使用 `feat:` 表示新功能，等等。理解和使用这些约定有助于自动生成发布说明，使版本控制更加容易，并提高提交历史的可读性。

### 提交信息的结构

Conventional Commits 标准的提交信息通常包括三个部分：

1. **类型（Type）**：描述提交的类别，如 `fix`、`feat`、`docs`、`style`、`refactor`、`perf`、`test`、`build`、`ci` 或 `chore`。

2. **范围（Scope）**：（可选）指定提交影响的代码区域，如 `api`、`auth`、`deps` 等。

3. **主题（Subject）**：简短描述所做的更改。

### 示例

```markdown
fix(login): correct login endpoint authentication mechanism

- The authentication mechanism for the login endpoint was not functioning correctly. It has been corrected to use the proper token validation process.
- Updated the unit tests to reflect the new authentication flow.
```

在这个示例中，提交类型是 `fix`，范围是 `login`，主题是 `correct login endpoint authentication mechanism`。紧随主题的是一个详细的更改描述，可能包括多个更改点或原因。

### 提交信息的重要性

清晰的提交信息对于项目维护者和贡献者来说非常重要，因为它们提供了项目历史和变更的快速概览。遵循 Conventional Commits 标准可以帮助团队：

- **理解变更**：快速识别每个提交的目的和影响。
- **自动化文档**：生成变更日志和发布说明。
- **版本控制**：根据提交类型自动决定版本号的增加。
- **代码审查**：提供上下文信息，便于代码审查。

通过遵循这些提交标准，OpenIM 项目能够保持一个有组织、易于理解的提交历史，这对于所有项目参与者都是有益的。
