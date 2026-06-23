# Go 学习实验室

> 从零开始，动手学 Go。每个知识点都配有**可运行代码 + 单元测试 + 详细笔记**。
> 本仓库的 Go 代码包与 Rust 学习项目结构对应，方便做跨语言对照。

## 学习进度

| 阶段 | 核心内容 | 状态 | 代码 | 笔记 |
|------|---------|------|------|------|
| 阶段零 | 变量/函数/集合 | ⬜ 待学 | `internal/{basics,funcsx,collections}` | [📖](01-基础/) |
| 阶段一 | 类型系统与错误处理 | ⬜ 待学 | `internal/{typesx,errorsx}` | [📖](02-类型与错误/) |
| 阶段二 | I/O 与数据编解码 | ⬜ 待学 | `internal/{iox,jsonx}` | [📖](03-IO与数据/) |
| 阶段三 | 时间与上下文 | ⬜ 待学 | `internal/{timex,contextx}` | [📖](04-时间与上下文/) |
| 阶段四 | 并发编程 | ⬜ 待学 | `internal/{concurrency,channelsx,syncx}` | [📖](05-并发/) |
| 阶段五 | 网络与泛型 | ⬜ 待学 | `internal/{httpx,genericsx}` | [📖](06-网络与泛型/) |

## 目录结构

```
docs/
├── 01-基础/          # 阶段零：基本类型、函数进阶、集合操作
├── 02-类型与错误/     # 阶段一：struct/接口/组合、错误处理
├── 03-IO与数据/      # 阶段二：Reader/Writer、JSON 编解码
├── 04-时间与上下文/   # 阶段三：time 处理、Context 取消/超时
├── 05-并发/          # 阶段四：goroutine、channel 模式、同步原语
├── 06-网络与泛型/    # 阶段五：HTTP handler、泛型 Map/Filter/Reduce
├── TODO.md           # 学习路线图（详细任务分解）
└── index.md          # 本文件
```

## Rust → Go 对照

| Go 包 | 对应 Rust 概念 |
|-------|---------------|
| `internal/basics` | 变量/控制流/基本类型 |
| `internal/funcsx` | 函数/闭包/defer → `Fn` trait |
| `internal/collections` | Vec/HashMap |
| `internal/typesx` | struct/trait/enum |
| `internal/errorsx` | Result/Error |
| `internal/genericsx` | 泛型 + trait bound |
| `internal/iox` | std::io |
| `internal/jsonx` | serde / JSON |
| `internal/timex` | chrono / time |
| `internal/contextx` | 无直接对应（Go 特色） |
| `internal/concurrency` | std::thread |
| `internal/channelsx` | mpsc / crossbeam |
| `internal/syncx` | Arc / Mutex |
| `internal/httpx` | reqwest / axum |

## 进阶路线

完成以上六阶段后，进入 [FUTURE.md](FUTURE.md) 继续：

- **阶段六**：工程化与测试（benchmark/profiling/fuzzing/lint）
- **阶段七**：Web 框架与 API 设计（Gin/REST/JWT/OpenAPI）
- **阶段八**：存储与数据（database/sql/sqlx/GORM/Redis）
- **阶段九**：进阶架构（DI/observability/gRPC/消息队列）
- **实战项目**：URL Shortener / Blog API / Chat Room

---

## 学习原则

- **测试驱动**：每个模块都有 `*_test.go`，先看测试理解函数行为，再读实现
- **动手 > 看书**：运行 `go test`，修改代码，观察输出变化
- **渐进式**：阶段零 → 阶段五，不跳阶段
- **跨语言对照**：带着 Rust 经验学 Go，关注概念映射而非从头学
