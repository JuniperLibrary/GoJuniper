# 09 — Context

## 是什么

`context.Context` 是 Go 在 goroutine 之间传递**取消信号、超时截止和请求级元数据**的标准接口。
它是一个不可变的树状结构：从根节点派生（derive）子 context，父取消则子也取消。

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}    // 返回一个只读 channel，取消时 close
    Err() error               // 取消后返回原因（Canceled 或 DeadlineExceeded）
    Value(key any) any        // 取值（不推荐滥用）
}
```

## 怎么用

### 基本模式：select + Done()

```go
func SleepOrDone(ctx context.Context, d time.Duration) error {
    timer := time.NewTimer(d)
    defer timer.Stop()

    select {
    case <-timer.C:        // 正常到期
        return nil
    case <-ctx.Done():     // 被取消
        return ctx.Err()
    }
}
```

核心模式是 `select` 同时监听两个 channel：业务操作 vs `ctx.Done()`。
谁先返回就执行谁的分支。这使调用方可以随时取消正在执行的函数。

### 创建 Context

```go
// 根节点：空 context，永不取消
ctx := context.Background()

// 可手动取消
ctx, cancel := context.WithCancel(ctx)
defer cancel()     // 务必 defer，防止资源泄漏

// 带超时（time.Duration）
ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
defer cancel()

// 带截止时间（具体时间点）
ctx, cancel := context.WithDeadline(ctx, time.Now().Add(100*time.Millisecond))
defer cancel()
```

### 传递 Context 的惯例

- 函数第一个参数传 `ctx context.Context`
- 通过显式参数传递，不存 struct 字段
- 只往下传，不上传（parent 创建 child，child 不修改 parent）

## 为什么

### Go 为什么需要 context

Go 的核心并发模型是 goroutine（轻量线程）。当一个 HTTP 请求进来，服务端会启动多个 goroutine
协作处理（读 DB、调下游 API、写响应）。如果客户端断开连接或超时，这些 goroutine 应该**被通知停止工作**，
否则会浪费资源甚至导致泄漏。

Context 解决的就是这个"如何通知 goroutine 停下来"的问题。它本质是一个**只读的取消信号广播机制**：

```
Request → context.WithTimeout → goroutine A → goroutine B
                                     ↓              ↓
                               select + Done()   select + Done()
                               ← 同时收到取消信号 →
```

### Rust 没有 context

Rust 的异步模型基于 `Future + poll`：一个 future 不做 poll 就不会推进，扔掉 future 就自然取消。
Go 的 goroutine 一旦启动就会运行（除非自己阻塞），所以需要外部机制通知它退出。

Rust 中 closest 的对应是 `tokio_util::sync::CancellationToken`，但这是第三方库而非语言标准。
Go 的 context 是标准库，所有框架和工具都遵守同一约定。

### context.Value 的争议

`context.WithValue` 可以往 context 中存键值对。但 Go 社区普遍**不推荐**用它传函数参数：
- 类型不安全（key/value 都是 `interface{}`）
- 隐式依赖难追踪
- 只在请求全链路元数据（trace ID、认证信息）等场景使用

## 常见坑

1. **`cancel` 必须 defer 调用**：
   `context.WithCancel/WithTimeout` 返回的 cancel 函数不调用会导致资源泄漏
   （内部会泄漏 timer goroutine）。

2. **忘记在 select 中监听 `ctx.Done()`**：
   传了 context 进去却不检查 Done，context 就毫无意义。

3. **用 `context.WithValue` 传可选参数**：
   这是反模式。可选参数应该用函数选项（functional options）或显式参数。

4. **把 context 存在 struct 里**：
   Context 应该显式传递，而不是作为 struct 字段。标准库的 `http.Request` 是个例外
   （`r.Context()`），因为请求对象本身就是 context 的自然载体。

5. **`ctx.Err()` 只在 `<-ctx.Done()` 之后有意义**：
   Done() 未关闭时调用 `ctx.Err()` 返回 `nil`。不要用它做轮询判断。

---

对应的代码文件：`internal/contextx/contextx.go` / `internal/contextx/contextx_test.go`

```
cd /Users/dingchuan/Documents/Repos/GoJuniper && go test ./internal/contextx/ -v
```
