# Go 学习实验室

> 从零开始，动手学 Go。每个知识点都配有**可运行代码 + 单元测试 + 详细笔记**。
> 本仓库的 Go 代码包与 Rust 学习项目结构对应，方便做跨语言对照。

## 学习进度

| 阶段 | 核心内容 | 状态 | 代码 | 笔记 |
|------|---------|------|------|------|
| 一 | 语言基础 | ⬜ 待学 | `internal/{basics,funcsx,collections}` | [📖](01-基础/) |
| 二 | 类型系统与错误处理 | ⬜ 待学 | `internal/{typesx,errorsx}` | [📖](02-类型与错误/) |
| 三 | I/O 与数据编解码 | ⬜ 待学 | `internal/{iox,jsonx}` | [📖](03-IO与数据/) |
| 四 | 时间与上下文 | ⬜ 待学 | `internal/{timex,contextx}` | [📖](04-时间与上下文/) |
| 五 | 并发编程 | ⬜ 待学 | `internal/{concurrency,channelsx,syncx}` | [📖](05-并发/) |
| 六 | 网络与泛型 | ⬜ 待学 | `internal/{httpx,genericsx}` | [📖](06-网络与泛型/) |
| 七 | 工程化与测试 | ⬜ 待学 | `internal/testingx/` | 待创建 |
| 八 | Web 框架与 API | ⬜ 待学 | `cmd/todo-api/` | 待创建 |
| 九 | 存储与数据 | ⬜ 待学 | 接入 PostgreSQL/SQLite | 待创建 |
| 十 | 进阶架构 | ⬜ 待学 | gRPC/Observability/DI | 待创建 |

完整学习路线见 [TODO.md](TODO.md)。

## 目录结构

```
docs/
├── 01-基础/          # 阶段一：基本类型、函数进阶、集合操作
├── 02-类型与错误/     # 阶段二：struct/接口/组合、错误处理
├── 03-IO与数据/      # 阶段三：Reader/Writer、JSON 编解码
├── 04-时间与上下文/   # 阶段四：time 处理、Context 取消/超时
├── 05-并发/          # 阶段五：goroutine、channel 模式、同步原语
├── 06-网络与泛型/    # 阶段六：HTTP handler、泛型 Map/Filter/Reduce
├── TODO.md           # 完整学习路线图（一～十阶段 + 实战项目）
└── index.md          # 本文件
```

## Rust → Go 对照

| Go 包 | 对应 Rust 概念 |
|-------|---------------|
| `internal/01-basics` | 变量/控制流/基本类型 |
| `internal/02-funcsx` | 函数/闭包/defer → `Fn` trait |
| `internal/03-collections` | Vec/HashMap |
| `internal/04-typesx` | struct/trait/enum |
| `internal/05-errorsx` | Result/Error |
| `internal/14-genericsx` | 泛型 + trait bound |
| `internal/06-iox` | std::io |
| `internal/07-jsonx` | serde / JSON |
| `internal/08-timex` | chrono / time |
| `internal/09-contextx` | 无直接对应（Go 特色） |
| `internal/10-concurrency` | std::thread |
| `internal/11-channelsx` | mpsc / crossbeam |
| `internal/12-syncx` | Arc / Mutex |
| `internal/13-httpx` | reqwest / axum |

---

## 学习原则

- **测试驱动**：每个模块都有 `*_test.go`，先看测试理解函数行为，再读实现
- **动手 > 看书**：运行 `go test`，修改代码，观察输出变化
- **渐进式**：阶段一 → 阶段十，不跳阶段
- **跨语言对照**：带着 Rust 经验学 Go，关注概念映射而非从头学
