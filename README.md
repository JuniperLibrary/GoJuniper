# GoJuniper

一个用于学习与复习 Go（从基础到进阶/高级）的练习仓库。

特点：
- 以 `internal/<topic>` 的方式组织“主题包”，每个主题都有注释与单元测试
- 以 `cmd/<app>` 的方式提供可运行示例
- 每个 topic 目录下会有一个知识文档（*.md），用于记录该主题的知识点与练习清单
- 旧的历史学习代码封存在 `_legacy/`（Go 工具链会忽略以下划线开头目录，避免影响构建）

## 目录结构

本仓库采用 Go 社区常见的应用组织方式（更利于持续扩展与保持可测试性）：

```
GoJuniper/
  cmd/                # 可执行程序入口（一个子目录 = 一个可运行程序）
    hello/
    server/
  internal/           # 仅仓库内部可导入的学习模块（每个 topic 配测试）
    basics/
    collections/
    errorsx/
    typesx/
    iox/
    jsonx/
    timex/
    contextx/
    concurrency/
    channelsx/
    syncx/
    httpx/
    genericsx/
  _legacy/            # 历史学习代码封存（Go 工具链默认忽略 _ 开头目录）
  go.mod
  README.md
```

为什么这样分：
- `cmd/`：避免“一个目录多个 main()”的编译冲突；每个程序独立运行、独立演示
- `internal/`：强制这些包只能被本仓库使用，适合做学习/练习，不用考虑对外 API 稳定性
- `_legacy/`：保留旧代码备查，但不影响 `go test ./...` 的全量构建

## 环境与运行

需要 Go 1.24+（本仓库 `go.mod` 为 `go 1.24`）。

常用命令：

```bash
gofmt -w .
go test -shuffle on ./...
go vet ./...
```

运行示例：

```bash
# 最小 demo
go run ./cmd/hello

# 启动 http 服务
go run ./cmd/server -addr :8080
```

## 可执行示例（cmd）

- `cmd/hello`：最小可运行入口，演示导入 `internal/*` 包并调用函数（[cmd/hello/main.go](file:///e:/dingchuan/GoJuniper/cmd/hello/main.go)）
- `cmd/server`：`net/http` 服务 + `flag` 参数解析 + `os/signal` 优雅退出（[cmd/server/main.go](file:///e:/dingchuan/GoJuniper/cmd/server/main.go)）
  - `GET /health` -> `ok`
  - `POST /echo` -> 将 `{"message":"hello"}` 变成 `{"message":"HELLO"}`

## 学习模块（internal）

每个模块都配套 `*_test.go`，建议学习顺序：先看测试用例，再读实现，最后自己改动让测试继续通过。

- `internal/basics`：变量/循环/switch、rune 与字符串、边界条件、uint64 溢出检查（[basics.go](file:///e:/dingchuan/GoJuniper/internal/basics/basics.go)）
  - 知识文档：[常量与变量与作用域与初始化.md](file:///e:/dingchuan/GoJuniper/internal/basics/常量与变量与作用域与初始化.md)
- `internal/collections`：slice + map（去重、计数）、key 排序、泛型 map key 提取（[collections.go](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go)）
  - 知识文档：[切片与映射_slice与map.md](file:///e:/dingchuan/GoJuniper/internal/collections/切片与映射_slice与map.md)
- `internal/errorsx`：错误处理：`fmt.Errorf("%w")`、`errors.Is`、`errors.Join`（[errorsx.go](file:///e:/dingchuan/GoJuniper/internal/errorsx/errorsx.go)）
  - 知识文档：[错误处理_errors_wrap_is_join.md](file:///e:/dingchuan/GoJuniper/internal/errorsx/错误处理_errors_wrap_is_join.md)
- `internal/typesx`：类型系统：struct、方法（值/指针接收者）、embedding、实现 `fmt.Stringer`（[typesx.go](file:///e:/dingchuan/GoJuniper/internal/typesx/typesx.go)）
  - 知识文档：[结构体_方法_接口与组合.md](file:///e:/dingchuan/GoJuniper/internal/typesx/结构体_方法_接口与组合.md)
- `internal/iox`：I/O：`io.Reader/io.Writer`、`bufio.Scanner` 按行读取、缓冲写入（[iox.go](file:///e:/dingchuan/GoJuniper/internal/iox/iox.go)）
  - 知识文档：[IO基础_reader_writer_bufio.md](file:///e:/dingchuan/GoJuniper/internal/iox/IO基础_reader_writer_bufio.md)
- `internal/jsonx`：JSON：struct tag、`omitempty`、marshal/unmarshal（[jsonx.go](file:///e:/dingchuan/GoJuniper/internal/jsonx/jsonx.go)）
  - 知识文档：[JSON基础_tag与编解码.md](file:///e:/dingchuan/GoJuniper/internal/jsonx/JSON基础_tag与编解码.md)
- `internal/timex`：time：RFC3339 解析、Duration 计算、StartOfDay（[timex.go](file:///e:/dingchuan/GoJuniper/internal/timex/timex.go)）
  - 知识文档：[时间_time_parse_duration.md](file:///e:/dingchuan/GoJuniper/internal/timex/时间_time_parse_duration.md)
- `internal/contextx`：context：取消/超时、“sleep or done”常见模式（[contextx.go](file:///e:/dingchuan/GoJuniper/internal/contextx/contextx.go)）
  - 知识文档：[Context基础_取消与超时.md](file:///e:/dingchuan/GoJuniper/internal/contextx/Context基础_取消与超时.md)
- `internal/concurrency`：并发：worker pool，首错返回 + 取消其它任务（[pool.go](file:///e:/dingchuan/GoJuniper/internal/concurrency/pool.go)）
  - 知识文档：[并发基础_worker_pool与取消.md](file:///e:/dingchuan/GoJuniper/internal/concurrency/并发基础_worker_pool与取消.md)
- `internal/channelsx`：channel：generator、pipeline、fan-in（Merge）（[channelsx.go](file:///e:/dingchuan/GoJuniper/internal/channelsx/channelsx.go)）
  - 知识文档：[Channel模式_generator_pipeline_fanin.md](file:///e:/dingchuan/GoJuniper/internal/channelsx/Channel模式_generator_pipeline_fanin.md)
- `internal/syncx`：sync：`sync.Mutex` 保护共享变量、`sync.Once` lazy init（[syncx.go](file:///e:/dingchuan/GoJuniper/internal/syncx/syncx.go)）
  - 知识文档：[并发安全_mutex与once.md](file:///e:/dingchuan/GoJuniper/internal/syncx/并发安全_mutex与once.md)
- `internal/httpx`：HTTP：`http.Handler/ServeMux`、JSON 请求/响应、`httptest` 测试（[httpx.go](file:///e:/dingchuan/GoJuniper/internal/httpx/httpx.go)）
  - 知识文档：[HTTP基础_handler与httptest.md](file:///e:/dingchuan/GoJuniper/internal/httpx/HTTP基础_handler与httptest.md)
- `internal/genericsx`：泛型：Map/Filter/Reduce（[genericsx.go](file:///e:/dingchuan/GoJuniper/internal/genericsx/genericsx.go)）
  - 知识文档：[泛型基础_type_parameters.md](file:///e:/dingchuan/GoJuniper/internal/genericsx/泛型基础_type_parameters.md)

## _legacy 目录说明

`_legacy/` 用来存放你早期“一个目录多份 demo”的历史代码。

因为 Go 在同一个包内不允许重复定义 `main()`，这些历史文件会导致 `go test ./...` 构建失败。
将其放在 `_legacy/` 后，Go 工具链会自动忽略该目录，从而保证新的学习模块保持可编译、可测试、可持续迭代。
