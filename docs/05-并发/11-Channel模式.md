# 11 — Channel 模式

## 是什么

channel 是 Go 的 CSP 通信原语，类型为 `chan T`。goroutine 通过 channel 发送/接收值，实现同步和数据传递。channel 本身是并发安全的。

## 怎么用

`internal/11-channelsx` 实现了三种经典 channel 模式（`channelsx.go`）：

### 生成器（Generator）

```go
func Generate(ctx context.Context, n int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)          // 发送完毕关闭 channel
        for i := 0; i < n; i++ {
            select {
            case <-ctx.Done():
                return
            case out <- i:
            }
        }
    }()
    return out                    // 返回只读 channel
}
```

> ⚠️ **注意**：**channel 由"发送方"负责关闭**。关闭后不能再发送（会 panic），所以生成器在发送完毕后 `defer close(out)`。接收方不能关闭它不拥有的 channel，否则多发送方场景下会 panic。

### Pipeline（流水线）

```go
func Square(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for v := range in {       // 从上游读取
            select {
            case <-ctx.Done():
                return
            case out <- v * v:    // 处理后发送到下游
            }
        }
    }()
    return out
}
```

### Fan-in（多路合并）

```go
func Merge(ctx context.Context, ins ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, ch := range ins {
        ch := ch
        go func() {
            defer wg.Done()
            for v := range ch {
                select {
                case <-ctx.Done():
                    return
                case out <- v:
                }
            }
        }()
    }
    go func() {
        wg.Wait()        // 所有输入 channel 关闭后
        close(out)       // 关闭输出 channel
    }()
    return out
}
```

> ⚠️ **注意**：**fan-in 用 WaitGroup 等所有输入关闭后再关闭 out**。只有所有上游都 `close` 了，下游 `for range` 才会正常退出。如果提前或重复关闭 out，会 panic 或让接收方误判结束。`Merge` 里 `wg.Wait() → close(out)` 是关键时序。

> ⚠️ **注意**：fan-in 的循环变量要拷贝：`for _, ch := range ins { ch := ch; go func(){...}() }`。虽然 Go 1.22 已修复循环变量捕获，但显式拷贝让"每个 goroutine 拿到独立 channel"的意图更清晰、向后兼容老版本。

组合使用：

```go
a := Square(ctx, Generate(ctx, 5))   // 0..4 -> 平方
b := Square(ctx, Generate(ctx, 5))   // 另一路
out := Merge(ctx, a, b)              // 合并两路
for v := range out {
    fmt.Println(v)                   // 0, 0, 1, 1, 4, 4, 9, 9, 16, 16
}
```

## 为什么

对比 Java：

| 特性 | Go | Java |
|------|----|------|
| 多生产者 | channel 天然支持多 goroutine 发送 | `BlockingQueue`（多生产者多消费者） |
| 多消费者 | 多个 goroutine 从同一 channel 接收 | `BlockingQueue`（天然多消费者） |
| 有无缓冲 | `make(chan T)` vs `make(chan T, N)` | `ArrayBlockingQueue(N)` / `LinkedBlockingQueue` |
| 类型方向 | `<-chan T`（只读）/ `chan<- T`（只写） | 无类型级区分（靠接口设计） |
| 选择 | `select { case <-ch1: ... case <-ch2: ... }` | 无语言级 select；用 `ExecutorService` / `CompletableFuture` 组合 |

Go 的 channel 和 Java 的 `BlockingQueue` 都用于生产者-消费者模型，但 Go 的设计更通用：一个 channel 可以被任意多个 goroutine 发送/接收（无限制多生产者/消费者），且 `select` 能同时等待多个 channel。Java 的 `BlockingQueue` 天然支持多生产多消费，但没有语言级的 `select`，需借助 `CompletableFuture`/`ExecutorService` 实现多路等待。

Go 的 channel 方向（`<-chan` / `chan<-`）在编译期检查，但只是接口层面的约定，底层仍是同一个 channel。

## 常见坑

1. **向已关闭的 channel 发送会 panic**：只能由发送方关闭 channel。关闭后不能再发数据。

2. **从已关闭的 channel 接收不会阻塞**：会立即返回零值（类型默认值），需用 `v, ok := <-ch` 区分 channel 是否关闭（`ok == false` 表示已关闭且无数据）。`for range ch` 会自动在关闭后退出。

3. **死锁**：无缓冲 channel 必须同时有发送方和接收方，否则会永久阻塞。`fatal error: all goroutines are asleep - deadlock!`

4. **忘记关闭 channel**：导致 `for range` 循环永不退出。如果接收方数量为 0，还可能导致发送方死锁。

5. **无缓冲 vs 有缓冲**：无缓冲 channel 是同步的（发送等待接收）；有缓冲 channel 当缓冲区未满时发送不阻塞。选错场景可能导致性能问题或死锁。

---

对应的代码文件：`internal/11-channelsx/channelsx.go` / `internal/11-channelsx/channelsx_test.go`

```
cd /Users/dingchuan/Documents/Repos/GoJuniper && go test ./internal/11-channelsx/ -v
```
