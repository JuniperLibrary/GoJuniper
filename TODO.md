# GoJuniper 学习路线图

## 阶段零：夯实基础（基本类型与语法）

- [ ] `internal/basics` — 变量/常量/基本类型/switch/循环
- [ ] `internal/funcsx` — 函数进阶：闭包/defer/变参/panic-recover
- [ ] `internal/collections` — slice + map 操作

## 阶段一：类型系统与错误处理

- [ ] `internal/typesx` — struct/方法/接口/组合/embedding
- [ ] `internal/errorsx` — 错误处理：`%w`/`errors.Is`/`errors.Join`

## 阶段二：I/O 与数据编解码

- [ ] `internal/iox` — io.Reader/io.Writer/bufio
- [ ] `internal/jsonx` — JSON tag/marshal/unmarshal

## 阶段三：时间与上下文

- [ ] `internal/timex` — time 解析/Duration/时区
- [ ] `internal/contextx` — Context 取消/超时/传值

## 阶段四：并发编程

- [ ] `internal/concurrency` — goroutine/worker pool/错误传播
- [ ] `internal/channelsx` — channel 模式：generator/pipeline/fan-in
- [ ] `internal/syncx` — Mutex/RWMutex/sync.Once

## 阶段五：网络与泛型

- [ ] `internal/httpx` — HTTP handler/ServeMux/httptest
- [ ] `internal/genericsx` — 泛型：Map/Filter/Reduce/GetLargest

---

> 阶段顺序：零 → 一 → 二 → 三 → 四 → 五（不允许跳阶段）
> 每个阶段学完输出 `docs/` 下的复习笔记 + git commit
> 进阶学习见 [docs/FUTURE.md](docs/FUTURE.md) — 阶段六～九 + 实战项目
