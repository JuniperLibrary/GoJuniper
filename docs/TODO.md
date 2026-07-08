# Go 学习路线图

> 从零基础到能写生产级 Go 项目的完整路线。
> 每个阶段都是下一阶段的前置，严格按照顺序推进。

---

## 第一阶段：语言基础

> 目标：掌握 Go 的核心语法和基本类型系统
> 对应代码：`internal/{basics, funcsx, collections}`

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 1.1 变量与类型 | 变量声明、零值、基础类型、rune/string、类型大小 | ★ | 2h | 读 `internal/01-basics` 源码 + 跑测试 |
| 1.2 控制流 | for 循环、if/switch、defer、break/continue | ★ | 1.5h | 跑 `TestFizzBuzz`，自己改逻辑 |
| 1.3 函数 | 多返回值、命名返回值、变参函数 | ★ | 1h | `internal/02-funcsx` 的 Sum 和 DeferInspect |
| 1.4 闭包与 defer | 闭包捕获、defer LIFO 顺序、panic/recover | ★★ | 2h | Counter 闭包、RecoverFromPanic |
| 1.5 切片与映射 | slice 预分配、map 去重/计数、key 排序 | ★★ | 2h | `internal/03-collections` 的三个函数 |

**阶段产出：** 完成 `docs/01-基础/` 三篇笔记阅读 + `go test ./internal/01-basics/... ./internal/02-funcsx/... ./internal/03-collections/...` 全部通过

---

## 第二阶段：类型系统与错误处理

> 目标：理解 Go 的类型设计哲学（组合取代继承）和错误处理惯用法
> 对应代码：`internal/{typesx, errorsx}`

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 2.1 struct 与方法 | 值接收者 vs 指针接收者、构造函数模式 | ★★ | 2h | `internal/04-typesx` 的 User 类型 |
| 2.2 接口 | `fmt.Stringer`、接口即约定、鸭子类型 | ★★ | 2h | 给 User 实现 String() |
| 2.3 组合（embedding） | struct 嵌入、字段提升、与继承的区别 | ★★ | 1.5h | Admin 嵌入 User |
| 2.4 哨兵错误 | `errors.New`、自定义错误变量 | ★ | 1h | `internal/05-errorsx` 的 ErrNotPositive |
| 2.5 错误包裹 | `fmt.Errorf("%w")`、`errors.Is` | ★★ | 1.5h | 解析错误时带上下文 |
| 2.6 错误合并 | `errors.Join`、多错误聚合 | ★★ | 1h | Join 函数 |

```go
// 学完本阶段能读懂并写出这样的代码：
type User struct { ID int; Name string }
func (u User) Greeting() string { return "hello, " + u.Name }
func NewUser(id int, name string) (User, error) { /* 校验后返回 */ }
```

**阶段产出：** `docs/02-类型与错误/` 笔记 + `go test ./internal/04-typesx/... ./internal/05-errorsx/...` 通过

---

## 第三阶段：I/O 与数据编解码

> 目标：掌握 Go 的 I/O 抽象和常见数据格式处理
> 对应代码：`internal/{iox, jsonx}`

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 3.1 io.Reader / io.Writer | 接口定义、组合、`io.ReadAll` | ★★ | 2h | `internal/06-iox` 的 ReadAllString |
| 3.2 bufio | Scanner 按行读取、buffer 调优 | ★★ | 1.5h | ReadLines 的 buffer 设置 |
| 3.3 缓冲写入 | bufio.NewWriter、Flush 语义 | ★★ | 1h | WriteLines 必须 Flush |
| 3.4 JSON 序列化 | `json.Marshal` / `Unmarshal`、struct tag | ★★ | 2h | `internal/07-jsonx` 基础用法 |
| 3.5 omitempty | 零值与omit、指针字段 | ★★ | 1h | JSON tag 的 omitempty 行为 |

**常见坑：** bufio.Writer 忘记 Flush 导致数据丢失、JSON 未导出字段不会被编解码、omitempty 对零值 struct 无效

**阶段产出：** `docs/03-IO与数据/` 笔记 + `go test ./internal/06-iox/... ./internal/07-jsonx/...` 通过

---

## 第四阶段：时间与上下文

> 目标：掌握 Go 的时间处理和 Context 并发协调
> 对应代码：`internal/{timex, contextx}`

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 4.1 时间解析/格式化 | RFC3339、time.Layout、时区 | ★★ | 2h | `internal/08-timex` 的 ParseRFC3339 |
| 4.2 Duration 计算 | Sub、Add、Since/Until | ★ | 1h | Until 函数 |
| 4.3 时间截断 | Date() 分解、StartOfDay 模式 | ★★ | 1h | StartOfDay 实现 |
| 4.4 Context 接口 | Deadline/Done/Err/Value | ★★ | 1.5h | `internal/09-contextx` 的 SleepOrDone |
| 4.5 select 多路复用 | select + ctx.Done() 模式 | ★★ | 2h | 等待超时或被取消 |

```go
// Context 是 Go 并发编程的基石：
func SleepOrDone(ctx context.Context, d time.Duration) error {
    select {
    case <-time.After(d): return nil
    case <-ctx.Done():    return ctx.Err()
    }
}
```

**阶段产出：** `docs/04-时间与上下文/` 笔记 + `go test ./internal/08-timex/... ./internal/09-contextx/...` 通过

---

## 第五阶段：并发编程

> 目标：理解 goroutine 和 channel 的 CSP 模型，写出安全的并发代码
> 对应代码：`internal/{concurrency, channelsx, syncx}`

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 5.1 goroutine | go 关键字、M:N 调度、栈大小 | ★★ | 1h | 简单的 go func() 启动 |
| 5.2 worker pool | 固定并发度、job channel、错误传播 | ★★★ | 3h | `internal/10-concurrency` 的 Run 函数 |
| 5.3 首错取消 | context.WithCancel、errCh 非阻塞写入 | ★★★ | 2h | Run 中的取消逻辑 |
| 5.4 生成器模式 | channel 作为迭代器、close 语义 | ★★ | 1.5h | `internal/11-channelsx` 的 Generate |
| 5.5 pipeline | channel 链式处理、goroutine 生命周期 | ★★ | 2h | Generate → Square 串联 |
| 5.6 fan-in | 多路合并、sync.WaitGroup 协调 | ★★★ | 2h | Merge 函数 |
| 5.7 Mutex | 共享变量保护、锁粒度 | ★★ | 1.5h | `internal/12-syncx` 的 Counter |
| 5.8 sync.Once | 懒加载、线程安全的一次执行 | ★★ | 1h | OnceValue |

**对比 Java：** goroutine 类似 `new Thread(...)` / 虚拟线程（`Thread.ofVirtual()`）但更轻量（KB 级栈 vs MB 级）；channel 类似 `BlockingQueue` 但语言内建、支持多生产者多消费者；Mutex 没有 `synchronized` 的块级作用域，必须显式 Lock/Unlock（或直接使用 `synchronized`）。

**阶段产出：** `docs/05-并发/` 笔记 + `go test ./internal/10-concurrency/... ./internal/11-channelsx/... ./internal/12-syncx/...` 通过

---

## 第六阶段：网络与泛型

> 目标：掌握 HTTP 服务构建和泛型编程
> 对应代码：`internal/{httpx, genericsx}`

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 6.1 Handler 接口 | `http.Handler`、`HandlerFunc` 适配器 | ★★ | 1.5h | `internal/13-httpx` 的 Handler |
| 6.2 ServeMux 路由 | Go 1.22+ 增强路由、方法匹配 | ★★ | 1.5h | NewMux 的路由注册 |
| 6.3 JSON 请求/响应 | json.Decoder/Encoder、Content-Type | ★★ | 1.5h | EchoHandler 实现 |
| 6.4 httptest | 不启动真实端口测试 handler | ★★ | 1.5h | httpx 的测试文件 |
| 6.5 泛型函数 | 类型参数、Map/Filter/Reduce | ★★ | 2h | `internal/14-genericsx` 的三个高阶函数 |
| 6.6 类型约束 | `cmp.Ordered`、自定义约束、interface 作为约束 | ★★★ | 2h | GetLargest 的约束写法 |

```go
// 泛型让 Go 也能写高阶函数：
func Map[A, B any](xs []A, f func(A) B) []B
func Filter[T any](xs []T, pred func(T) bool) []T
func Reduce[T, Acc any](xs []T, acc Acc, f func(Acc, T) Acc) Acc
```

**阶段产出：** `docs/06-网络与泛型/` 笔记 + `go test ./internal/13-httpx/... ./internal/14-genericsx/...` 通过

---

## 第七阶段：工程化与测试

> 目标：写出可维护、可测试、高性能的 Go 代码
> 前置条件：已完成第一～六阶段

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 7.1 测试进阶 | subtests、table-driven tests、`-shuffle`、`-count=1` | ★★ | 2h | 给现有 internal 包补充边界 case |
| 7.2 Benchmark | `go test -bench`、benchmark 函数写法、benchstat | ★★ | 2h | 对比 strings.Builder vs bytes.Buffer |
| 7.3 Fuzzing | `f.Fuzz`、种子语料库、crash 分析 | ★★★ | 2h | 给 ParsePositiveInt 加 fuzz test |
| 7.4 Race Detector | `-race`、data race 复现与修复 | ★★★ | 2h | 给 syncx 的 Counter 制造 race 并修复 |
| 7.5 Profiling | `pprof` CPU/memory、`trace` 调度追踪 | ★★★ | 3h | 对 worker pool 做性能剖析 |
| 7.6 代码质量 | `go vet`、`staticcheck` / `golangci-lint`、`revive` | ★ | 1h | 配置 lint、修复 warning |
| 7.7 Test Double | 接口 mock、httptest 进阶、testify/mock | ★★★ | 3h | 给 httpx handler 写 mock 测试 |

**阶段产出：** `internal/testingx/` 练习包 + `docs/07-工程化与测试/` 笔记

---

## 第八阶段：Web 框架与 API 设计

> 目标：用 Go 搭建生产可用的 HTTP 服务
> 前置条件：已完成第六阶段（net/http 基础）
> 你已有 gin_learn/ 基础，这里做系统性梳理

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 8.1 Gin 路由 | 路由分组、参数绑定、中间件链 | ★★ | 2h | 整理 gin_learn/ 为结构化笔记 |
| 8.2 请求校验 | binding/tags、自定义 validator | ★★ | 2h | gin_learn/binding/ 的 14 个例子 |
| 8.3 RESTful 设计 | 资源路由、状态码选择、版本控制 | ★★ | 2h | 设计一个 Todo API |
| 8.4 中间件模式 | 日志/鉴权/限流/恢复 | ★★★ | 3h | 实现 common middleware 链 |
| 8.5 OpenAPI / Swagger | swaggo 自动生成 API 文档 | ★★ | 2h | 给 Todo API 加 swagger 注解 |
| 8.6 JWT 鉴权 | access + refresh token、中间件校验 | ★★★ | 3h | 实现完整登录流程 |
| 8.7 Graceful Shutdown | `os.Signal` + `http.Server.Shutdown` | ★★ | 1h | 已有基础版在 `cmd/server/` |
| 8.8 纯标准库路线 | Go 1.22+ 增强 `http.ServeMux` 实现同样 API | ★★ | 2h | 不依赖 Gin 实现 Todo API |

**阶段产出：** `cmd/todo-api/` 完整 REST 服务 + `docs/08-Web框架与API/` 笔记

---

## 第九阶段：存储与数据

> 目标：掌握 Go 与数据库的交互方式
> 前置条件：已完成第八阶段（HTTP 服务）

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 9.1 `database/sql` | 标准接口、连接池配置、事务 | ★★ | 3h | 用 sqlite 实现 Todo 持久化 |
| 9.2 SQL Migrations | `golang-migrate` / `goose`、up/down | ★★ | 2h | 给 Todo API 加 migration |
| 9.3 sqlx | 命名参数、结构体扫描、`In()` | ★★ | 2h | 把 database/sql 改为 sqlx |
| 9.4 GORM | AutoMigrate、关联、hook | ★★★ | 3h | 用 GORM 重写 Todo API |
| 9.5 Redis | `go-redis`、缓存模式、分布式锁 | ★★★ | 3h | 给 Todo API 加缓存层 |
| 9.6 连接池调优 | `SetMaxOpenConns` 等参数 | ★★ | 1h | 压测不同参数表现 |

**阶段产出：** Todo API 接入 PostgreSQL/SQLite + `docs/09-存储与数据/` 笔记

---

## 第十阶段：进阶架构

> 目标：理解大型 Go 项目的组织方式和周边生态
> 前置条件：已完成第九阶段

| 模块 | 核心内容 | 难度 | 预计时间 | 实践方式 |
|------|---------|------|---------|---------|
| 10.1 依赖注入 | `wire` 编译期 DI / `uber-fx` 运行时 DI | ★★★ | 3h | 重构 Todo API 用 DI |
| 10.2 配置管理 | `viper` + env + yaml + 命令行 | ★★ | 2h | 给 Todo API 加多来源配置 |
| 10.3 结构化日志 | `slog`（标准库）/ `zap` 高性能 | ★★ | 2h | fmt.Println → 结构化日志 |
| 10.4 Metrics | `prometheus` 客户端、自定义指标 | ★★★ | 2h | 加请求计数/延迟 |
| 10.5 分布式追踪 | `otel` trace、context 传播 | ★★★ | 3h | HTTP 请求链路 trace |
| 10.6 gRPC + protobuf | `.proto` 定义、代码生成 | ★★★ | 4h | Todo 服务加 gRPC 接口 |
| 10.7 WebSocket | `gorilla/websocket` | ★★★ | 3h | 加实时推送通知 |
| 10.8 消息队列 | `asynq` / Kafka 客户端 | ★★★ | 3h | 异步处理耗时任务 |
| 10.9 整洁架构 | handler → service → repository 分层 | ★★ | 2h | 按 clean arch 重构 |

**阶段产出：** 完整 Todo 应用（REST + gRPC + WS + 可观测性）+ `docs/10-进阶架构/` 笔记

---

## 实战项目集

> 综合运用以上所有知识点的独立项目

| 项目 | 需要完成阶段 | 核心练习点 |
|------|------------|-----------|
| **CLI 单词计数器** | 一～三 | 文件读取、strings 操作、`flag` 包 |
| **URL Shortener** | 六～八 | Gin + 重定向 + 缓存 + 存储 |
| **CLI Todo 管理器** | 六～七 | cobra + viper + 文件持久化 |
| **RESTful Blog API** | 六～九 | 用户/文章/评论 CRUD + JWT + DB |
| **WebSocket Chat Room** | 八～十 | WS + 房间管理 + 广播 |
| **Rate Limiter Proxy** | 七～八 | 中间件模式 + 滑动窗口 + Redis |
| **Mini Monitoring** | 十 | prometheus 采集 + 可视化 |
| **gRPC Gateway** | 十 | gRPC + REST 双协议 + 网关 |

---

## 学习建议

1. **先跑通再深究** — 遇到看不懂的概念，先运行 `go test`，看输出，再读源码
2. **标准库优先** — 学 Gin 前先吃透 `net/http`；学 GORM 前先吃透 `database/sql`
3. **带着 Java 经验看差异** — 重点理解 Go 的设计取舍：组合 vs 继承、goroutine vs Thread、interface vs interface（隐式实现 vs 显式 implements）、error vs Exception
4. **不要跳阶段** — 每个阶段都是下一阶段的前置条件
5. **每阶段完成后提交** — commit message 格式：`docs: 第X阶段 - <阶段名>`
6. **善用测试** — 本仓库的每个 `internal/` 包都有测试，修改代码后跑测试验证
7. **不要怕犯错** — data race、死锁、panic 都是学习的一部分，用 `-race` 和 vet 抓住它们

---

## 推荐资源

| 资源 | 适用阶段 | 说明 |
|------|---------|------|
| [Go By Example](https://gobyexample.com/) | 一～三 | 边看边练 |
| [Effective Go](https://go.dev/doc/effective_go) | 四～六 | 官方风格指南 |
| [Go 标准库文档](https://pkg.go.dev/std) | 全阶段 | 最权威参考 |
| [Go 语言设计与实现](https://draveness.me/golang/) | 五～七 | 底层原理 |
| [Uber Go Style Guide](https://github.com/uber-go/guide) | 七～十 | 工程化最佳实践 |
| [Go 语言高级编程](https://github.com/chai2010/advanced-go-programming-book) | 十 | CGO/汇编/架构 |

---

**进度记录：** 始于 2026-06-23 | 10 阶段 + 实战项目
