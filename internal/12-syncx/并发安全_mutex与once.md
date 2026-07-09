# Go 基础：并发安全（sync.Mutex / sync.Once）

这个文档对应 [internal/12-syncx](./)。

配合阅读：
- 实现代码：[syncx.go](./syncx.go)
- 测试代码：本目录下暂无 `*_test.go`，建议先读实现，后续补测试（见练习清单）

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

> ⚠️ **注意**：数据竞争不一定**立刻**出错，它表现为**偶发、不可复现**的 bug，是最难调试的一类问题。一定要用 `go test -race` 或 `go run -race` 去抓。光看代码“大部分时候对”没用——竞争窗口一旦被调度命中，结果就错。

**为什么**：`n++` 在机器层面是“读-改-写”三步，多个 goroutine 交错执行会丢失更新（A 读 0、B 读 0、A 写 1、B 写 1，最终是 1 而非 2）。锁把这段临界区变成“串行”，才能正确。

---

## 2. sync.Mutex 的基本用法

最常见模式：

```go
mu.Lock()
// 访问共享数据
mu.Unlock()
```

在本仓库里：
- [Counter](./syncx.go#L8-L34) 用 `mu` 保护 `n`
- `Inc/Add/Value` 都先 Lock 再操作

```go
func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}
```

初学者先记住：同一份共享数据，只要被多个 goroutine 写，就必须加锁或换一种并发设计（例如 channel）。

> ⚠️ **注意**：Go 的 `sync.Mutex` **没有 RAII 守卫**（不像 Java 的 `synchronized` 块结束自动释放锁，或 `ReentrantLock` 需配合 `try-finally`/`AutoCloseable`）。必须**显式 `Unlock()`**，而且最稳妥的写法是 `mu.Lock(); defer mu.Unlock()`。如果中间 `return` 或 panic 而没有 `Unlock`，其它 goroutine 会**永久阻塞**在这把锁上（死锁）。本仓库正因为用了 `defer`，即使 `Value()` 提前返回也能正常解锁。

> ⚠️ **注意**：**`sync.Mutex` 不能被复制**。本仓库 `Counter` 把 `mu sync.Mutex` 作为字段，对 `Counter` 做**值拷贝**（如 `c2 := c1` 或当作函数参数值传递）会把锁一起拷走，导致两份“假锁”、并且 `go vet` 会报警。正确做法：含 Mutex 的 struct 一律用**指针**传递（`*Counter`）。这点和 Java 的 `synchronized` 不同——Java 的对象锁绑定在对象实例上，复制引用仍指向同一把锁；Go 的 Mutex 是值类型，复制即出错。

**常见坑**：把 `Lock` 范围写太大（整个函数一把锁），会严重损害并发性能；但范围太小又可能漏保护。原则是“临界区越小越好，但必须完整包含对共享数据的全部读写”。

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

本仓库的 [OnceValue](./syncx.go#L36-L48) 就是一个典型练习：第一次计算后缓存，后续直接返回。

```go
type OnceValue[T any] struct {
	once sync.Once
	v    T
}

func (o *OnceValue[T]) Get(f func() T) T {
	o.once.Do(func() {
		o.v = f()
	})
	return o.v
}
```

> ⚠️ **注意**：`sync.Once.Do` 保证 `f` **绝对只执行一次**，即使有 100 个 goroutine 同时调用 `Get`。第一个进来的执行 `f`，其余全部阻塞等它完成、然后直接拿缓存值。`Once` 本身也是不能复制的（见上一条 Mutex 同理，它内部含锁）。

> ⚠️ **注意**：`OnceValue[T any]` 用的是 **Go 1.21+ 的泛型 Once 变体思路**（标准库有 `sync.OnceFunc` / `sync.OnceValue`）。本仓库自己实现了泛型版 `OnceValue[T]`，注意它和 `sync.Once` 一样——`Get` 一旦执行过 `f`，之后传不同的 `f` 也**不会再执行**，只返回第一次的结果。别误以为“每次传新函数都会算”。

**为什么**：懒加载（lazy init）常用于“昂贵的一次性初始化”（如建连接、读配置）。用 `Once` 比手写 `if initialized { ... }` + 锁更不易出错，因为并发安全由标准库兜底。

---

## 4. 初学者常见坑

1. Lock 后忘记 Unlock（会导致死锁）
2. 把 Lock 范围写太大（影响并发性能）
3. 只读操作也需要锁吗？
   - 如果并发写存在：读也要锁，否则仍然 race
   - 如果保证没有写：读可以不锁

> ⚠️ **注意**：第 3 点最容易被忽略——**只要存在并发写，读也必须加锁**。很多“只读为什么还要锁”的崩溃，就是因为读操作读到了写操作写了一半的数据（撕裂读）。本仓库 `Counter.Value()` 也 `Lock` 了，正是这个原因。对应 Java：`ReadWriteLock` 的读锁（`readLock()`）同样是必须的，且同样防止撕裂读。

---

## 5. 练习清单

1. 给 `Counter` 增加 `Reset()` 方法，并补测试
2. 写一个线程安全的 `Set`（内部用 `map[T]struct{}` + Mutex）
3. 写一个 `OnceValue` 的变体：缓存 `func() (T, error)` 的错误（只执行一次）

---

## 6. 自测命令

```bash
go test ./internal/12-syncx -shuffle on
```

> 当前目录暂未包含 `*_test.go`，运行上述命令时若提示 “no test files” 属正常。建议完成练习清单后补上测试，再执行该命令验证。
