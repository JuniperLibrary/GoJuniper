# Go 进阶学习计划

> 当前六阶段（零～五）覆盖了 Go 语言本身的核心特性。
> 以下是学完基础后，进入"真正写 Go 项目"前的进阶路线。

---

## 阶段六：工程化与测试

> 目标：写出可维护、可测试、高性能的 Go 代码

| 模块 | 核心内容 | 难度 | 预估时间 | 实践方式 |
|------|---------|------|----------|---------|
| 6.1 测试进阶 | subtests + table-driven tests + `-shuffle` + `-count=1` 缓存控制 | ★★ | 2h | 给现有 internal 包补充边界 case |
| 6.2 Benchmark | `go test -bench`、benchmark 函数写法、benchstat 结果分析 | ★★ | 2h | 对比 strings.Builder vs bytes.Buffer |
| 6.3 Fuzzing | `f.Fuzz`、种子语料库、crash 分析 | ★★★ | 2h | 给 ParsePositiveInt 加 fuzz test |
| 6.4 Race Detector | `-race`、data race 复现与修复 | ★★★ | 2h | 给 syncx 的 Counter 故意制造 race 并修复 |
| 6.5 Profiling | `pprof` CPU/memory、`trace` 调度追踪 | ★★★ | 3h | 对 concurrency worker pool 做性能剖析 |
| 6.6 代码质量 | `go vet`、`staticcheck` / `golangci-lint`、`revive` | ★ | 1h | 配置 lint、修复现存 warning |
| 6.7 Test Double | 接口 mock、httptest 进阶、testify/mock | ★★★ | 3h | 给 httpx handler 写完备的 mock 测试 |

**阶段产出：** `internal/testingx/` 练习包 + `docs/07-工程化与测试/` 笔记

---

## 阶段七：Web 框架与 API 设计

> 目标：用 Go 搭建生产可用的 HTTP 服务
> 你已经有 gin_learn/ 基础，这里做系统性梳理

| 模块 | 核心内容 | 难度 | 预估时间 | 实践方式 |
|------|---------|------|----------|---------|
| 7.1 Gin 路由 | 路由分组、参数绑定、中间件链 | ★★ | 2h | 整理 gin_learn/ 为结构化笔记 |
| 7.2 请求校验 | binding/tags、自定义 validator | ★★ | 2h | 已有 gin_learn/binding/，整理成文档 |
| 7.3 RESTful 设计 | 资源路由、状态码选择、版本控制 | ★★ | 2h | 设计一个 Todo API |
| 7.4 中间件模式 | 日志/鉴权/限流/恢复中间件 | ★★★ | 3h | 实现 common middleware 链 |
| 7.5 OpenAPI / Swagger | swaggo 自动生成 API 文档 | ★★ | 2h | 给 Todo API 加 swagger 注解 |
| 7.6 JWT 鉴权 | access + refresh token、中间件校验 | ★★★ | 3h | 实现完整登录流程 |
| 7.7 Graceful Shutdown | `os.Signal` + `http.Server.Shutdown` | ★★ | 1h | 已有一个基础版在 cmd/server/ |
| 7.8 纯标准库路线 | 不用 Gin，用 `net/http` + `http.ServeMux`（Go 1.22+ 增强路由） | ★★ | 2h | 用标准库实现同样的 Todo API |

**阶段产出：** `cmd/todo-api/` 完整 REST 服务 + `docs/08-Web框架与API/` 笔记

---

## 阶段八：存储与数据

> 目标：掌握 Go 与数据库的交互方式

| 模块 | 核心内容 | 难度 | 预估时间 | 实践方式 |
|------|---------|------|----------|---------|
| 8.1 `database/sql` | 标准接口、连接池配置、事务 | ★★ | 3h | 用 sqlite 实现 Todo 持久化 |
| 8.2 SQL Migrations | `golang-migrate` / `goose`、up/down | ★★ | 2h | 给 Todo API 加 migration |
| 8.3 sqlx | 命名参数、结构体扫描、`In()` 子句 | ★★ | 2h | 把 database/sql 版本改为 sqlx |
| 8.4 GORM | AutoMigrate、关联、hook | ★★★ | 3h | 用 GORM 重写 Todo API |
| 8.5 Redis | `go-redis`、缓存模式、分布式锁 | ★★★ | 3h | 给 Todo API 加缓存层 |
| 8.6 连接池调优 | `SetMaxOpenConns`/`SetMaxIdleConns`/`SetConnMaxLifetime` | ★★ | 1h | 压测不同参数表现 |

**阶段产出：** Todo API 接入 PostgreSQL/SQLite + `docs/09-存储与数据/` 笔记

---

## 阶段九：进阶架构

> 目标：理解大型 Go 项目的组织方式和周边生态

| 模块 | 核心内容 | 难度 | 预估时间 | 实践方式 |
|------|---------|------|----------|---------|
| 9.1 依赖注入 | `wire` 编译期 DI / `uber-fx` 运行时 DI | ★★★ | 3h | 重构 Todo API 用 DI |
| 9.2 配置管理 | `viper` + env + yaml + 命令行覆盖 | ★★ | 2h | 给 Todo API 加多来源配置 |
| 9.3 结构化日志 | `slog`（标准库）/ `zap` 高性能 | ★★ | 2h | 替换 fmt.Println → 结构化日志 |
| 9.4 Metrics | `prometheus` 客户端、自定义指标 | ★★★ | 2h | 给 Todo API 加请求计数/延迟 |
| 9.5 分布式追踪 | `otel` trace、context 传播 | ★★★ | 3h | HTTP 请求链路 trace |
| 9.6 gRPC + protobuf | `.proto` 定义、server/client 代码生成 | ★★★ | 4h | Todo 服务加 gRPC 接口 |
| 9.7 WebSocket | `gorilla/websocket` / `nhooyr.io/websocket` | ★★★ | 3h | 加实时推送通知 |
| 9.8 消息队列 | `asynq`（Redis 任务队列） / Kafka 客户端 | ★★★ | 3h | 异步处理耗时任务 |
| 9.9 整洁架构 | handler → service → repository 分层 | ★★ | 2h | 按 clean arch 重构 |

**阶段产出：** 完整的 Todo 应用（REST + gRPC + WS + 可观测性）+ `docs/10-进阶架构/` 笔记

---

## 实战项目 Ideas

> 综合运用以上所有知识点

| 项目 | 涉及阶段 | 核心练习点 |
|------|---------|-----------|
| **URL Shortener** | 六～七 | Gin + 路由 + 重定向 + 缓存 |
| **CLI Todo** | 六 | cobra + viper + 文件持久化 |
| **RESTful Blog API** | 六～八 | 用户/文章/评论 CRUD + JWT + 数据库 |
| **Chat Room** | 七～九 | WebSocket + 房间管理 + 广播 |
| **Rate Limiter Proxy** | 七 | 中间件模式 + 滑动窗口 + Redis |
| **Mini Monitoring** | 九 | prometheus 指标采集 + 可视化 |
| **gRPC Gateway** | 九 | gRPC + REST 双协议 + 网关 |

---

## 推荐学习资源

| 资源 | 说明 |
|------|------|
| [**Go By Example**](https://gobyexample.com/) | 边看边练，适合快速查阅 |
| [**Effective Go**](https://go.dev/doc/effective_go) | 官方风格指南，必读 |
| [**Go 语言设计与实现**](https://draveness.me/golang/) | 底层原理，进阶必看 |
| [**Go 语言高级编程**](https://github.com/chai2010/advanced-go-programming-book) | CGO/汇编/架构 |
| [**Uber Go Style Guide**](https://github.com/uber-go/guide) | 工程化最佳实践 |
| [**Go 标准库文档**](https://pkg.go.dev/std) | 最权威的参考 |

---

## 学习建议

1. **不要一口气学完** — 每完成一个阶段，用对应的 mini project 巩固
2. **六阶段基础要牢** — 阶段零～五不牢靠，不要跳过
3. **标准库优先** — 学 Gin 前，先吃透 `net/http`；学 GORM 前，先吃透 `database/sql`
4. **带着 Rust 经验找差异** — 你已经有系统编程基础，重点理解 Go 的设计取舍（组合 vs 继承、goroutine vs thread、interface vs trait）
5. **每阶段写完要提交** — commit message 格式：`feat: 阶段X - <主题>`
