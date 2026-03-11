# Go 基础：并发（goroutine / worker pool / 取消）

这个文档对应 [internal/concurrency](file:///e:/dingchuan/GoJuniper/internal/concurrency)。

配合阅读：
- 实现代码：[pool.go](file:///e:/dingchuan/GoJuniper/internal/concurrency/pool.go)
- 测试代码：[pool_test.go](file:///e:/dingchuan/GoJuniper/internal/concurrency/pool_test.go)

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

---

## 3. worker pool 模式

worker pool 就是：
- 一个任务队列（通常用 channel 表示）
- N 个 worker goroutine 从队列取任务并执行
- 执行结果汇总（结果 channel 或回调）

优点：
- 控制并发度（N 就是最大并发）
- 更容易做取消与收尾

---

## 4. “首错返回 + 取消其他任务”

在工程里很常见的需求：
- 并发做很多件事
- 一旦有一个失败，就取消其余任务，尽快返回错误

本模块就是围绕这个目标设计的。建议学习顺序：
1. 先跑测试，理解期望行为：[pool_test.go](file:///e:/dingchuan/GoJuniper/internal/concurrency/pool_test.go)
2. 再读实现，看它如何监听 `ctx.Done()`：[pool.go](file:///e:/dingchuan/GoJuniper/internal/concurrency/pool.go)

---

## 5. 初学者常见坑

1. 忘记关闭 channel（或错误地关闭不该关闭的 channel）
2. worker 不监听取消信号，导致测试结束还在跑
3. 用共享变量但不加锁（数据竞争）

---

## 6. 练习清单

1. 写一个 `RunN(ctx, n, jobs)`：最多并发 n 个任务，返回第一个错误
2. 增加一个“结果收集”功能：让每个任务返回一个值，最终汇总成 slice
3. 给任务增加超时：每个任务最多执行 200ms

---

## 7. 自测命令

```bash
go test ./internal/concurrency -shuffle on
```

