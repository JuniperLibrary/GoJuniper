# Go 基础：context（取消与超时）

这个文档对应 [internal/09-contextx](file:///e:/dingchuan/GoJuniper/internal/09-contextx)。

配合阅读：
- 实现代码：[contextx.go](file:///e:/dingchuan/GoJuniper/internal/09-contextx/contextx.go)

> ⚠️ **注意**：本仓库 `09-contextx` 目录下只有实现代码 `contextx.go`，没有 `_test.go`。学习时直接看 `contextx.go` 里的 `SleepOrDone` 即可，练习清单里的函数也写在同包内。

---

## 1. context 是用来解决什么问题的

在并发/网络/IO 场景里，经常需要：
- 任务超时自动取消
- 用户取消请求，后台任务也要尽快停止
- 多个 goroutine 协同工作，需要统一停止信号

`context.Context` 就是标准库统一的“取消/超时/请求范围值”的载体。

> ⚠️ **注意（Java 无直接对应）**：Context 是 Go 的特有机制，Java 里没有等价的标准机制。Java 通常用 `Thread.interrupt()` + 检查中断标志（或 `Future.cancel(true)`）来表达“取消信号”，而 Go 把这件事标准化成了一个接口 `context.Context`，并贯穿标准库（HTTP server、数据库驱动、gRPC 等都认它）。理解 Context 是写好 Go 并发/网络代码的前提。

---

## 2. 你需要记住的三条规则（初学者版）

1. **ctx 作为第一个参数传下去**：`func Do(ctx context.Context, ...)`
2. **不要把 ctx 存到 struct 里长期持有**（它应该随请求/任务生命周期结束）
3. **ctx 不是用来传业务参数的**（业务参数用函数参数；ctx 只放跨边界的、请求范围的值）

> ⚠️ **注意（第 2 条要理解 why）**：context 是“一次请求/任务”的生命周期载体，请求结束 ctx 就被取消。把 ctx 存进 struct 长期持有，会让下游goroutine 永远收不到取消信号，或收到一个早已过期的信号——这是典型的内存/协程泄漏来源。正确做法是每次调用都把 ctx **作为参数传进去**。

---

## 3. 取消：WithCancel

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

当你调用 `cancel()`：
- `ctx.Done()` 会被关闭
- 监听 `Done()` 的 goroutine 应该退出

> ⚠️ **注意（一定要调用 cancel）**：`WithCancel`/`WithTimeout`/`WithDeadline` 返回的 `cancel` 函数**必须调用**（通常用 `defer cancel()`），否则 context 关联的资源不会释放，可能造成泄漏。`cancel` 是幂等的，调多次也安全。

---

## 4. 超时：WithTimeout / WithDeadline

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
```

超时后 `ctx.Err()` 会变成 `context.DeadlineExceeded`。

> ⚠️ **注意**：`WithTimeout(parent, d)` 等价于 `WithDeadline(parent, time.Now().Add(d))`——它是“相对时长”；`WithDeadline` 是“绝对时刻”。两者超时后 `ctx.Err()` 返回 `context.DeadlineExceeded`（不是 `context.Canceled`）。区分两者：主动调 `cancel()` 得到 `Canceled`，到点自动触发得到 `DeadlineExceeded`。

---

## 5. 经典写法：“sleep or done”

很多时候你想“等待一段时间，但如果被取消就立刻退出”，写法是：

```go
select {
case <-ctx.Done():
	return ctx.Err()
case <-time.After(d):
	return nil
}
```

本仓库的练习函数就围绕这种模式展开，建议你先跑测试再读实现。

> ⚠️ **注意（ctx.Done() 的类型）**：`ctx.Done()` 返回的是 `<-chan struct{}`——一个只读 channel，被取消时这个 channel 被**关闭**（不是发送一个值）。`select` 监听它就能在取消时立刻收到信号。注意是“channel 关闭”触发，所以哪怕多次 select 也能正确收到。本仓库 `SleepOrDone` 用的就是 `case <-ctx.Done(): return ctx.Err()` 配对 `case <-timer.C:` 的写法。

> ⚠️ **注意（time.After 的泄漏陷阱）**：示例里 `time.After(d)` 在某些场景会泄漏——只要 `ctx.Done()` 先触发，`time.After` 创建的定时器要等 d 到期才被 GC。长生命周期循环里应改用 `time.NewTimer(d)` + `defer timer.Stop()`（本仓库 `SleepOrDone` 正是这么写的，值得对照学习）。

---

## 6. 测试里用 t.Context()

在 Go 1.24+ 的测试里可以用 `t.Context()` 获取随测试生命周期自动取消的 ctx（本仓库用的就是这种写法）。

> ⚠️ **注意**：`t.Context()` 是 Go 1.24 新增的测试辅助方法，它返回的 ctx 会在测试结束时自动取消，免去你手动 `defer cancel()`。但**业务代码里**不要依赖它——业务代码还是老老实实 `context.WithCancel`/`WithTimeout` + `defer cancel()`。

---

## 7. 练习清单

1. 写 `Work(ctx context.Context) error`：循环做 100 次小任务，但被取消就立刻退出
2. 写 `TimeoutCall(ctx context.Context, d time.Duration) error`：超时返回 `context.DeadlineExceeded`
3. 把你的一个“无限循环 goroutine”改造为监听 `ctx.Done()` 能退出

> ⚠️ **注意（WithValue 谨慎用）**：context 还能用 `context.WithValue` 携带“请求范围的值”，但它是 `any` 类型、key 也是 `any`，**完全无编译期类型安全**——取错 key 或类型断言失败运行时才暴露。业界共识：优先用函数参数传业务数据，只在真正“跨中间件/框架边界的请求元数据”（如 traceID、用户身份）才用 `WithValue`，且 key 用自定义私有类型避免冲突。

---

## 8. 自测命令

```bash
go test ./internal/09-contextx -shuffle on
```

> ⚠️ **注意**：当前目录没有测试文件，该命令会提示 “no test files”。先补 `contextx_test.go` 或把练习函数写好再运行。
