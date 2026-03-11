# Go 基础：context（取消与超时）

这个文档对应 [internal/contextx](file:///e:/dingchuan/GoJuniper/internal/contextx)。

配合阅读：
- 实现代码：[contextx.go](file:///e:/dingchuan/GoJuniper/internal/contextx/contextx.go)
- 测试代码：[contextx_test.go](file:///e:/dingchuan/GoJuniper/internal/contextx/contextx_test.go)

---

## 1. context 是用来解决什么问题的

在并发/网络/IO 场景里，经常需要：
- 任务超时自动取消
- 用户取消请求，后台任务也要尽快停止
- 多个 goroutine 协同工作，需要统一停止信号

`context.Context` 就是标准库统一的“取消/超时/请求范围值”的载体。

---

## 2. 你需要记住的三条规则（初学者版）

1. **ctx 作为第一个参数传下去**：`func Do(ctx context.Context, ...)`
2. **不要把 ctx 存到 struct 里长期持有**（它应该随请求/任务生命周期结束）
3. **ctx 不是用来传业务参数的**（业务参数用函数参数；ctx 只放跨边界的、请求范围的值）

---

## 3. 取消：WithCancel

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

当你调用 `cancel()`：
- `ctx.Done()` 会被关闭
- 监听 `Done()` 的 goroutine 应该退出

---

## 4. 超时：WithTimeout / WithDeadline

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
```

超时后 `ctx.Err()` 会变成 `context.DeadlineExceeded`。

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

---

## 6. 测试里用 t.Context()

在 Go 1.24+ 的测试里可以用 `t.Context()` 获取随测试生命周期自动取消的 ctx（本仓库用的就是这种写法）。

---

## 7. 练习清单

1. 写 `Work(ctx context.Context) error`：循环做 100 次小任务，但被取消就立刻退出
2. 写 `TimeoutCall(ctx context.Context, d time.Duration) error`：超时返回 `context.DeadlineExceeded`
3. 把你的一个“无限循环 goroutine”改造为监听 `ctx.Done()` 能退出

---

## 8. 自测命令

```bash
go test ./internal/contextx -shuffle on
```

