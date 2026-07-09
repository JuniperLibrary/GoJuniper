# Go 基础：channel 模式（generator / pipeline / fan-in）

这个文档对应 [internal/11-channelsx](./)。

配合阅读：
- 实现代码：[channelsx.go](./channelsx.go)
- 测试代码：本目录下暂无 `*_test.go`，建议先读实现，后续补测试（见练习清单）

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

> ⚠️ **注意**：`make(chan T)` 创建的是**无缓冲** channel，发送和接收必须**两边同时就绪**才能通信（同步点）。需要缓冲时显式指定容量：`make(chan T, n)`。缓冲大小影响的是“解耦程度”，**不改变**关闭规则。另外，channel 可以指定方向：`<-chan int`（只读）、`chan<- int`（只写），本仓库的 `Generate`/`Square`/`Merge` 返回的就是 `<-chan int`，强制调用方只能读。

---

## 2. range 读 channel

当一个 channel 被关闭后，`for v := range ch` 会自然退出：

```go
for v := range ch {
	_ = v
}
```

这也是 pipeline 写法里最常用的结束方式。

> ⚠️ **注意**：**关闭 channel 后仍能读取**——读到的是该类型的零值，且第二返回值 `ok == false`（`v, ok := <-ch`）。`for-range` 之所以能退出，正是因为底层检测到 channel 已关闭且无剩余数据。反过来：对**未关闭**的 channel 做 `range`，接收方会**永远阻塞**，这是最常见的“goroutine 泄漏”来源之一。

**常见坑**：以为 `close(ch)` 会把已发送的数据“销毁”。不会——已发送的数据还在，关闭只是通知“没有新数据了”，`range` 会先把缓冲区/已发数据读完再退出。

---

## 3. generator：从无到有生产数据

generator 的典型形式：
- 创建 out channel
- 启动 goroutine 往 out 里写
- 写完 close(out)

对应函数：[Generate](./channelsx.go)

```go
func Generate(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()
	return out
}
```

> ⚠️ **注意**：generator **永远用 `defer close(out)`** 保证 channel 一定会被关闭（即使中途被 ctx 取消 `return` 了也会关）。如果忘了关闭，下游 `for-range` 会一直阻塞。同时发送侧在 `select` 里监听 `ctx.Done()`，这样外部取消时 goroutine 能立刻退出，不会泄漏。

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

对应函数：[Square](./channelsx.go)

```go
func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v * v:
			}
		}
	}()
	return out
}
```

**为什么**：每个 stage 都是“输入 range + 输出 defer close”的对称结构，前一个 stage 关闭输出就驱动后一个 stage 的 range 退出，整条管道像水流一样自然收尾。

---

## 5. fan-in：把多个输入合并成一个输出（Merge）

fan-in 常用于：
- 多个 worker 并发输出结果
- 最终统一汇总到一个 channel 让消费者处理

`Merge` 的关键点：
- 每个输入 channel 都由一个 goroutine 负责搬运到 out
- 用 `sync.WaitGroup` 等待所有搬运 goroutine 结束
- 等全部结束后再 close(out)

对应函数：[Merge](./channelsx.go)

```go
func Merge(ctx context.Context, ins ...<-chan int) <-chan int {
	out := make(chan int)
	if len(ins) == 0 {
		close(out)
		return out
	}

	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, ch := range ins {
		ch := ch // 拷贝循环变量，避免所有 goroutine 捕获同一个变量
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
		wg.Wait()
		close(out)
	}()

	return out
}
```

> ⚠️ **注意**：`Merge` 的 **close(out) 必须等所有输入搬运 goroutine 都结束后**再做（`wg.Wait()` 后再关）。如果在某个搬运者还在 `out <- v` 时就关掉 out，会 **panic：send on closed channel**。这也是为什么要用 `WaitGroup` 而非“随便关”。

> ⚠️ **注意**：循环变量 `ch := ch` 的拷贝看似多余，但**不能省**。Go 1.22 之前 `for` 循环变量的地址/引用在每次迭代中复用，多个 goroutine 可能全都拿到最后一个 channel。显式拷贝一份，确保每个 goroutine 捕获到正确的 channel（Go 1.22+ 虽已修复语义，但保留拷贝仍是稳妥写法）。

**常见坑**：fan-in 时**不要**由某个输入方去关闭 out，也不要在循环里反复 close。out 只能由“等你的人都结束”的协调 goroutine 关一次。

---

## 6. 一定要处理取消（ctx.Done）

初学者最容易写出“永远阻塞的 goroutine”。解决方式是：
- 在发送/接收时用 `select` 监听 `ctx.Done()`

本模块所有函数都带 `ctx`，就是为了练这个习惯。

> ⚠️ **注意**：凡是 `go func()` 里做阻塞发送（`out <- v`），**务必**用 `select` 包一层 `case <-ctx.Done(): return`。否则外部取消后，发送方会因“没人接收/没人关 out”而永久阻塞——经典的 goroutine 泄漏。channel 约等于 Java 的 `BlockingQueue`，但 Go 的 channel 原生支持**多读多写**，且关闭语义由你手动管理，不像 Java 靠 GC 回收队列引用。

---

## 7. 练习清单

1. 写 `FilterEven(ctx, in <-chan int) <-chan int`（只保留偶数）
2. 写 `FanOut(ctx, in <-chan int, n int) []<-chan int`（把输入分发给 n 个 worker）
3. 写一个三段 pipeline：`Generate -> Square -> FilterEven`

---

## 8. 自测命令

```bash
go test ./internal/11-channelsx -shuffle on
```

> 当前目录暂未包含 `*_test.go`，运行上述命令时若提示 “no test files” 属正常。建议完成练习清单后补上测试，再执行该命令验证。
