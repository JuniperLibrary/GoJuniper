# Go 基础：并发（goroutine / worker pool / 取消）

这个文档对应 [internal/10-concurrency](./)。

配合阅读：
- 实现代码：[pool.go](./pool.go)
- 测试代码：本目录下暂无 `*_test.go`，建议先读实现，后续补测试（见练习清单）

---

## 1. goroutine 是什么

goroutine 可以理解成“轻量级线程”：

```go
go func() {
	// 并发执行
}()
```

特点：
- 创建成本低
- 数量可以很多
- 需要你自己管理退出（否则容易泄漏）

> ⚠️ **注意**：goroutine 是轻量的，但它**没有 RAII / 自动回收**。如果 goroutine 里没有退出路径（比如一直阻塞在某个 channel 接收上），它会**永久泄漏**，直到进程结束才释放。这点和 Java 的 `Thread`（或虚拟线程）类似，但 Go 不会帮你“join”或回收，必须自己用 channel / context 给出退出信号。

---

## 2. 并发最难的不是启动，而是“收尾”

初学者最常见的问题：
- goroutine 启动了，但没人等它结束（程序提前退出/测试结束）
- goroutine 一直阻塞不退出（泄漏）
- 多个 goroutine 之间缺少统一的停止信号

解决思路：
- 用 `context` 发取消信号
- 用 channel 传任务/结果
- 用 `sync.WaitGroup` 等待收尾（本仓库的 worker pool 把等待封装起来了）

> ⚠️ **注意**：`sync.WaitGroup` 的 `Add` / `Done` / `Wait` 必须**配对**。常见错误是 `Add` 的数量和 `Done` 的调用次数不一致：少 `Done` 会 `Wait` 永远阻塞（死锁），多 `Done` 会 panic。本仓库在 `pool.go:43` 用 `wg.Add(concurrency)` 一次性把 N 个 worker 都登记好，每个 worker 退出时 `defer wg.Done()`，保证严格 1:1。

**为什么**：并发程序的难点从来不是“怎么开 goroutine”，而是“怎么保证所有 goroutine 都干净地结束”。没有统一的等待与取消机制，程序行为会变得不确定。

---

## 3. worker pool 模式

worker pool 就是：
- 一个任务队列（通常用 channel 表示）
- N 个 worker goroutine 从队列取任务并执行
- 执行结果汇总（结果 channel 或回调）

优点：
- 控制并发度（N 就是最大并发）
- 更容易做取消与收尾

> ⚠️ **注意**：本仓库用 **无缓冲 channel `jobCh := make(chan Job)`** 作为任务队列，由生产者 goroutine 逐个投递（`pool.go:67`）。worker 用 `for job := range jobCh` 消费，当生产者 `close(jobCh)` 后循环自动退出。不要误以为 worker pool 一定要“预分配 N 个 job 到 channel”——真正控制并发度的是 **worker 的数量 N**，而不是 channel 的缓冲大小。

**常见坑**：把并发度理解成“channel 缓冲区大小”。缓冲区只是解耦生产者/消费者，真正限制“同时在跑的任务数”的是 worker goroutine 的个数。

---

## 4. “首错返回 + 取消其他任务”

在工程里很常见的需求：
- 并发做很多件事
- 一旦有一个失败，就取消其余任务，尽快返回错误

本模块就是围绕这个目标设计的。建议学习顺序：
1. 先跑测试，理解期望行为（本目录暂无测试，可直接读实现）
2. 再读实现，看它如何监听 `ctx.Done()`：[pool.go](./pool.go)

实现要点（`pool.go`）：
- 用 `ctx, cancel := context.WithCancel(ctx)` 派生一个“整组任务”的取消信号（`pool.go:33`）
- 每个 worker 在任务报错时，通过**非阻塞写**把错误塞进 `errCh`（容量 1），并调用 `cancel()` 通知其它 worker：

```go
select {
case errCh <- err:
    cancel()
default:
}
```

> ⚠️ **注意**：`errCh` 容量为 1 且写入用 `select { case ...: default: }` 非阻塞。原因是：多个 worker 可能**同时**失败，如果阻塞写，失败的 worker 会卡在“等别人来读 errCh”上，永远走不到 `return`，造成泄漏。非阻塞写保证“只记录第一个错误”，其余 worker 直接放弃。这是首错取消的经典写法。

**为什么**：如果用普通阻塞 channel 传错误，goroutine 之间会互相“等对方读”，一旦读者已退出，写者就永久挂起。非阻塞写 + 容量 1 的组合，让“谁先到谁赢”，其余一律丢弃。

---

## 5. 初学者常见坑

1. 忘记关闭 channel（或错误地关闭不该关闭的 channel）
2. worker 不监听取消信号，导致测试结束还在跑
3. 用共享变量但不加锁（数据竞争）

> ⚠️ **注意**：**关闭 channel 的黄金法则——只有发送方才应该关闭 channel**。本仓库里 `jobCh` 由生产者 goroutine `defer close(jobCh)` 关闭（`pool.go:68`），worker 只负责读，绝不关。重复关闭、或向已关闭的 channel 发送，都会 **panic**；从接收方关闭更容易导致“发送方还往里写就 panic”。

---

## 6. 练习清单

1. 写一个 `RunN(ctx, n, jobs)`：最多并发 n 个任务，返回第一个错误
2. 增加一个“结果收集”功能：让每个任务返回一个值，最终汇总成 slice
3. 给任务增加超时：每个任务最多执行 200ms

---

## 7. 自测命令

```bash
go test ./internal/10-concurrency -shuffle on
```

> 当前目录暂未包含 `*_test.go`，运行上述命令时若提示 “no test files” 属正常。建议完成练习清单后补上测试，再执行该命令验证。
