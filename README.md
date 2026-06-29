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
  internal/           # 仅仓库内部可导入的学习模块（每个 topic 配知识文档）
    01-basics/        # 变量/控制流/基本类型
    02-funcsx/        # 函数进阶：闭包/defer/变参
    03-collections/   # slice + map 操作
    04-typesx/        # struct/方法/接口/组合
    05-errorsx/       # 错误处理
    06-iox/           # I/O 操作
    07-jsonx/         # JSON 编解码
    08-timex/         # 时间处理
    09-contextx/      # Context 取消/超时
    10-concurrency/   # 并发编程
    11-channelsx/     # Channel 模式
    12-syncx/         # 同步原语
    13-httpx/         # HTTP 服务
    14-genericsx/     # 泛型
    15-regexpx/       # 正则表达式
    16-tests/         # 统一测试目录（代码与测试分离）
  _legacy/            # 历史学习代码封存（Go 工具链默认忽略 _ 开头目录）
  docs/               # 学习笔记 + TODO 路线图
  go.mod
  README.md
```

### 目录编号与文档对应

| 目录 | 内容 | 对应 docs/ |
|------|------|-----------|
| 01-basics | 变量/控制流/基本类型 | 01-基础/ |
| 02-funcsx | 函数进阶：闭包/defer/变参 | 01-基础/ |
| 03-collections | slice + map 操作 | 01-基础/ |
| 04-typesx | struct/方法/接口/组合 | 02-类型与错误/ |
| 05-errorsx | 错误处理 | 02-类型与错误/ |
| 06-iox | I/O 操作 | 03-IO与数据/ |
| 07-jsonx | JSON 编解码 | 03-IO与数据/ |
| 08-timex | 时间处理 | 04-时间与上下文/ |
| 09-contextx | Context 取消/超时 | 04-时间与上下文/ |
| 10-concurrency | 并发编程 | 05-并发/ |
| 11-channelsx | Channel 模式 | 05-并发/ |
| 12-syncx | 同步原语 | 05-并发/ |
| 13-httpx | HTTP 服务 | 06-网络与泛型/ |
| 14-genericsx | 泛型 | 06-网络与泛型/ |
| 15-regexpx | 正则表达式 | （额外） |
| 16-tests | 统一测试目录 | （测试） |

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

### 概念分类（Rust → Go 对照）

> 本仓库参照 Rust 学习项目的概念分类方法组织 Go 代码，每个包对应一个语言特性，
> 实现 + 测试 + 知识文档三位一体。右侧 Rust 对照列帮助你建立跨语言映射。

| Go 包 | 对应 Rust 概念 | Rust 对照示例 |
|---|---|---|
| `internal/01-basics` | 变量/控制流/基础类型 | `let x = 42`; `if`/`loop`/`match` |
| `internal/02-funcsx` | 函数/闭包/defer | `Fn` trait；`move \|x\| x + 1` |
| `internal/04-typesx` | struct/trait/enum | `struct User` / `trait Display` / `enum Option` |
| `internal/03-collections` | Vec/HashMap | `vec![]` / `HashMap::new()` |
| `internal/05-errorsx` | Result/Error | `Result<T,E>` / `?` / `anyhow` |
| `internal/14-genericsx` | 泛型 + trait bound | `fn largest<T: PartialOrd>` |
| `internal/06-iox` | std::io | `Read` / `Write` / `BufReader` |
| `internal/07-jsonx` | serde / JSON | `serde_json::from_str` |
| `internal/08-timex` | chrono / time | `chrono::DateTime` / `Duration` |
| `internal/09-contextx` | 无直接对应（Go 特色） | — |
| `internal/10-concurrency` | std::thread | `thread::spawn` / `Arc<Mutex<T>>` |
| `internal/11-channelsx` | mpsc / crossbeam | `sync::mpsc` / `select!` |
| `internal/12-syncx` | Arc / Mutex | `Arc::new(Mutex::new(v))` |
| `internal/13-httpx` | reqwest / axum | `reqwest::Client` / axum handlers |

### 包详情

- `internal/01-basics`：变量/循环/switch、rune 与字符串、边界条件、uint64 溢出检查（[basics.go](file:///e:/dingchuan/GoJuniper/internal/01-basics/basics.go)）
  - 知识文档：[常量与变量与作用域与初始化.md](file:///e:/dingchuan/GoJuniper/internal/01-basics/常量与变量与作用域与初始化.md)
- `internal/02-funcsx`：函数进阶：闭包、defer 后进先出、变参、panic/recover（[funcsx.go](file:///e:/dingchuan/GoJuniper/internal/02-funcsx/funcsx.go)）
  - 知识文档：[函数进阶_closures_defer.md](file:///e:/dingchuan/GoJuniper/internal/02-funcsx/函数进阶_closures_defer.md)
- `internal/03-collections`：slice + map（去重、计数）、key 排序、泛型 map key 提取（[collections.go](file:///e:/dingchuan/GoJuniper/internal/03-collections/collections.go)）
  - 知识文档：[切片与映射_slice与map.md](file:///e:/dingchuan/GoJuniper/internal/03-collections/切片与映射_slice与map.md)
- `internal/05-errorsx`：错误处理：`fmt.Errorf("%w")`、`errors.Is`、`errors.Join`（[errorsx.go](file:///e:/dingchuan/GoJuniper/internal/05-errorsx/errorsx.go)）
  - 知识文档：[错误处理_errors_wrap_is_join.md](file:///e:/dingchuan/GoJuniper/internal/05-errorsx/错误处理_errors_wrap_is_join.md)
- `internal/04-typesx`：类型系统：struct、方法（值/指针接收者）、embedding、实现 `fmt.Stringer`（[typesx.go](file:///e:/dingchuan/GoJuniper/internal/04-typesx/typesx.go)）
  - 知识文档：[结构体_方法_接口与组合.md](file:///e:/dingchuan/GoJuniper/internal/04-typesx/结构体_方法_接口与组合.md)
- `internal/06-iox`：I/O：`io.Reader/io.Writer`、`bufio.Scanner` 按行读取、缓冲写入（[iox.go](file:///e:/dingchuan/GoJuniper/internal/06-iox/iox.go)）
  - 知识文档：[IO基础_reader_writer_bufio.md](file:///e:/dingchuan/GoJuniper/internal/06-iox/IO基础_reader_writer_bufio.md)
- `internal/07-jsonx`：JSON：struct tag、`omitempty`、marshal/unmarshal（[jsonx.go](file:///e:/dingchuan/GoJuniper/internal/07-jsonx/jsonx.go)）
  - 知识文档：[JSON基础_tag与编解码.md](file:///e:/dingchuan/GoJuniper/internal/07-jsonx/JSON基础_tag与编解码.md)
- `internal/08-timex`：time：RFC3339 解析、Duration 计算、StartOfDay（[timex.go](file:///e:/dingchuan/GoJuniper/internal/08-timex/timex.go)）
  - 知识文档：[时间_time_parse_duration.md](file:///e:/dingchuan/GoJuniper/internal/08-timex/时间_time_parse_duration.md)
- `internal/09-contextx`：context：取消/超时、"sleep or done"常见模式（[contextx.go](file:///e:/dingchuan/GoJuniper/internal/09-contextx/contextx.go)）
  - 知识文档：[Context基础_取消与超时.md](file:///e:/dingchuan/GoJuniper/internal/09-contextx/Context基础_取消与超时.md)
- `internal/10-concurrency`：并发：worker pool，首错返回 + 取消其它任务（[pool.go](file:///e:/dingchuan/GoJuniper/internal/10-concurrency/pool.go)）
  - 知识文档：[并发基础_worker_pool与取消.md](file:///e:/dingchuan/GoJuniper/internal/10-concurrency/并发基础_worker_pool与取消.md)
- `internal/11-channelsx`：channel：generator、pipeline、fan-in（Merge）（[channelsx.go](file:///e:/dingchuan/GoJuniper/internal/11-channelsx/channelsx.go)）
  - 知识文档：[Channel模式_generator_pipeline_fanin.md](file:///e:/dingchuan/GoJuniper/internal/11-channelsx/Channel模式_generator_pipeline_fanin.md)
- `internal/12-syncx`：sync：`sync.Mutex` 保护共享变量、`sync.Once` lazy init（[syncx.go](file:///e:/dingchuan/GoJuniper/internal/12-syncx/syncx.go)）
  - 知识文档：[并发安全_mutex与once.md](file:///e:/dingchuan/GoJuniper/internal/12-syncx/并发安全_mutex与once.md)
- `internal/13-httpx`：HTTP：`http.Handler/ServeMux`、JSON 请求/响应、`httptest` 测试（[httpx.go](file:///e:/dingchuan/GoJuniper/internal/13-httpx/httpx.go)）
  - 知识文档：[HTTP基础_handler与httptest.md](file:///e:/dingchuan/GoJuniper/internal/13-httpx/HTTP基础_handler与httptest.md)
- `internal/14-genericsx`：泛型：Map/Filter/Reduce、GetLargest（对应 Rust `get_largest<T: PartialOrd>`）（[genericsx.go](file:///e:/dingchuan/GoJuniper/internal/14-genericsx/genericsx.go)）
  - 知识文档：[泛型基础_type_parameters.md](file:///e:/dingchuan/GoJuniper/internal/14-genericsx/泛型基础_type_parameters.md)

`GetLargest` 示例直接对应 Rust 经典泛型模式：

```go
numbers := []int{34, 50, 25, 100, 65}
got, _ := genericsx.GetLargest(numbers)   // → 100（整数列表）

chars := []rune{'y', 'm', 'a', 'q'}
got, _ = genericsx.GetLargest(chars)       // → 'y'（rune 列表）
```

```rust
// Rust 对应实现
let numbers = vec![34, 50, 25, 100, 65];
let result = get_largest(&numbers);         // → 100

let chars = vec!['y', 'm', 'a', 'q'];
let result = get_largest(&chars);           // → 'y'
```

## _legacy 目录说明

`_legacy/` 用来存放你早期“一个目录多份 demo”的历史代码。

因为 Go 在同一个包内不允许重复定义 `main()`，这些历史文件会导致 `go test ./...` 构建失败。
将其放在 `_legacy/` 后，Go 工具链会自动忽略该目录，从而保证新的学习模块保持可编译、可测试、可持续迭代。
