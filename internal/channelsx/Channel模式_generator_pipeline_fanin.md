# Go 基础：channel 模式（generator / pipeline / fan-in）

这个文档对应 [internal/channelsx](file:///e:/dingchuan/GoJuniper/internal/channelsx)。

配合阅读：
- 实现代码：[channelsx.go](file:///e:/dingchuan/GoJuniper/internal/channelsx/channelsx.go)
- 测试代码：[channelsx_test.go](file:///e:/dingchuan/GoJuniper/internal/channelsx/channelsx_test.go)

---

## 1. channel 是什么

channel 是 goroutine 之间传递数据的一种类型安全管道：

```go
ch := make(chan int)
```

基本操作：
- 发送：`ch <- v`
- 接收：`v := <-ch`
- 关闭：`close(ch)`（表示“以后不会再发送数据了”）

重要规则：
- **发送方负责关闭**（通常是谁创建谁关闭）
- **接收方不要关闭**（否则容易 panic）

---

## 2. range 读 channel

当一个 channel 被关闭后，`for v := range ch` 会自然退出：

```go
for v := range ch {
	_ = v
}
```

这也是 pipeline 写法里最常用的结束方式。

---

## 3. generator：从无到有生产数据

generator 的典型形式：
- 创建 out channel
- 启动 goroutine 往 out 里写
- 写完 close(out)

对应函数：[Generate](file:///e:/dingchuan/GoJuniper/internal/channelsx/channelsx.go)

---

## 4. pipeline：一段段处理数据

pipeline 的思路是：

```
Generate -> Square -> (更多 stage) -> 最终消费
```

每一段 stage：
- 从输入 channel range 读数据
- 处理后写到输出 channel
- 结束时关闭输出 channel

对应函数：[Square](file:///e:/dingchuan/GoJuniper/internal/channelsx/channelsx.go)

---

## 5. fan-in：把多个输入合并成一个输出（Merge）

fan-in 常用于：
- 多个 worker 并发输出结果
- 最终统一汇总到一个 channel 让消费者处理

`Merge` 的关键点：
- 每个输入 channel 都由一个 goroutine 负责搬运到 out
- 用 `sync.WaitGroup` 等待所有搬运 goroutine 结束
- 等全部结束后再 close(out)

对应函数：[Merge](file:///e:/dingchuan/GoJuniper/internal/channelsx/channelsx.go)

---

## 6. 一定要处理取消（ctx.Done）

初学者最容易写出“永远阻塞的 goroutine”。解决方式是：
- 在发送/接收时用 `select` 监听 `ctx.Done()`

本模块所有函数都带 `ctx`，就是为了练这个习惯。

---

## 7. 练习清单

1. 写 `FilterEven(ctx, in <-chan int) <-chan int`（只保留偶数）
2. 写 `FanOut(ctx, in <-chan int, n int) []<-chan int`（把输入分发给 n 个 worker）
3. 写一个三段 pipeline：`Generate -> Square -> FilterEven`

---

## 8. 自测命令

```bash
go test ./internal/channelsx -shuffle on
```

