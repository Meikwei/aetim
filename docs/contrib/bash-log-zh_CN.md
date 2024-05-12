<!--
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-04-28 22:36:30
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-04-28 22:40:42
 * @FilePath: \open-im-server\docs\contrib\bash-log-zh_CN.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

## OpenIM 的日志系统设计与使用

路径：scripts/lib/logging.sh

**简介**
OpenIM 作为一个复杂的项目，需要一套强大的日志机制来诊断问题、维护系统健康并提供洞察力。嵌入于 OpenIM 中的自定义日志系统保证了日志的一致性和结构化。下面我们将深入了解这个日志系统的设计及其各种功能和使用场景。

**设计概览**

- **初始化**：系统首先通过`OPENIM_VERBOSE`变量确定日志详细程度。若未设置，默认值为 5，这决定了日志详情的深度。
- **日志文件设置**：日志存储在由`OPENIM_OUTPUT`变量指定的目录下，若未明确设置，则默认为脚本所在位置的\_output 目录。每个日志文件按日期命名，便于识别。
- **日志记录函数**：`echo_log()`函数负责将消息同时写入控制台（stdout）和日志文件，并在消息前添加时间戳。日志文件路径默认为\_output/logs/\*，默认开启日志记录。如需关闭，可设置`export ENABLE_LOGGING=false`。

**关键函数与使用场景**

- **错误处理**：

  - `openim::log::errexit()`：在命令执行出错时激活，打印调用栈，显示引发错误的函数序列，随后调用`openim::log::error_exit()`并附带相关信息。
  - `openim::log::install_errexit()`：设置错误捕获陷阱，确保错误处理器（errexit）传播到各种脚本结构，如函数、展开和子 shell 中。

- **日志级别**：

  - `openim::log::error()`：记录带有时间戳的错误消息，以'!!!'开头标示严重性。
  - `openim::log::info()`：提供信息性消息，显示与否取决于日志详细程度（OPENIM_VERBOSE）设定。
  - `openim::log::progress()`：设计用于记录进度信息或生成进度条。
  - `openim::log::status()`：记录带有时间戳的状态消息，每条前缀为'+++'，易于识别。
  - `openim::log::success()`：以明亮的绿色前缀突出成功操作，视觉上标示操作已完成。

- **退出与堆栈追踪**：

  - `openim::log::error_exit()`：记录错误消息，转储调用堆栈，并以指定退出码退出脚本。
  - `openim::log::stack()`：打印调用堆栈，展示调用层次结构。

- **使用信息**：

  - `openim::log::usage()` & `openim::log::usage_from_stdin()`：分别直接接受参数或从 stdin 读取参数来显示使用说明。

- **测试函数**：
  - `openim::log::test_log()`：测试套件，验证所有日志功能是否正常工作。

**使用场景**
设想 OpenIM 操作失败，需要确定原因的情境。有了日志系统，你可以：

- 检查特定日期的日志文件，寻找以'!!!'为前缀的错误消息。
- 查看调用树和堆栈追踪，追溯导致失败的操作序列。
- 利用详细程度级别过滤掉无关细节，专注于问题核心。

这种系统化和结构化的方法极大地简化了调试过程，提高了系统的维护效率。

**结论**
OpenIM 的日志系统彰显了在复杂项目中结构化和详细日志的重要性。通过使用这套日志机制，开发者和系统管理员可以简化故障排查流程，确保 OpenIM 项目的顺畅运行。
