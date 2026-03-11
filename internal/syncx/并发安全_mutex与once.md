# Go 基础：并发安全（sync.Mutex / sync.Once）

这个文档对应 [internal/syncx](file:///e:/dingchuan/GoJuniper/internal/syncx)。

配合阅读：
- 实现代码：[syncx.go](file:///e:/dingchuan/GoJuniper/internal/syncx/syncx.go)
- 测试代码：[syncx_test.go](file:///e:/dingchuan/GoJuniper/internal/syncx/syncx_test.go)

---

## 1. 为什么需要锁（Mutex）

当多个 goroutine 同时读写同一份数据，就可能出现：
- 数据竞争（race）
- 结果不一致（丢失更新）

例如“计数器 +1”看似简单，但它不是原子操作：
1. 读旧值
2. 旧值 + 1
3. 写回新值

并发下如果不保护，会互相覆盖。

---

## 2. sync.Mutex 的基本用法

最常见模式：

```go
mu.Lock()
// 访问共享数据
mu.Unlock()
```

在本仓库里：
- [Counter](file:///e:/dingchuan/GoJuniper/internal/syncx/syncx.go#L8-L34) 用 `mu` 保护 `n`
- `Inc/Add/Value` 都先 Lock 再操作

初学者先记住：同一份共享数据，只要被多个 goroutine 写，就必须加锁或换一种并发设计（例如 channel）。

---

## 3. sync.Once：只初始化一次

有时你希望：
- 某个值只计算一次（懒加载）
- 即使多个 goroutine 同时调用，也只执行一次初始化逻辑

就可以用 `sync.Once`：

```go
once.Do(func() {
	// 只执行一次
})
```

本仓库的 [OnceValue](file:///e:/dingchuan/GoJuniper/internal/syncx/syncx.go#L36-L48) 就是一个典型练习：第一次计算后缓存，后续直接返回。

---

## 4. 初学者常见坑

1. Lock 后忘记 Unlock（会导致死锁）
2. 把 Lock 范围写太大（影响并发性能）
3. 只读操作也需要锁吗？
   - 如果并发写存在：读也要锁，否则仍然 race
   - 如果保证没有写：读可以不锁

---

## 5. 练习清单

1. 给 `Counter` 增加 `Reset()` 方法，并补测试
2. 写一个线程安全的 `Set`（内部用 `map[T]struct{}` + Mutex）
3. 写一个 `OnceValue` 的变体：缓存 `func() (T, error)` 的错误（只执行一次）

---

## 6. 自测命令

```bash
go test ./internal/syncx -shuffle on
```

