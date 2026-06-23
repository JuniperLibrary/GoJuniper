# Go 基础：函数进阶（闭包、defer、变参、panic/recover）

这个文档对应 [internal/funcsx](file:///e:/dingchuan/GoJuniper/internal/funcsx)。

配合阅读：
- 实现代码：[funcsx.go](file:///e:/dingchuan/GoJuniper/internal/funcsx/funcsx.go)
- 测试代码：[funcsx_test.go](file:///e:/dingchuan/GoJuniper/internal/funcsx/funcsx_test.go)

***

## 1. 闭包（Closure）

闭包是一个函数值，它"捕获"了其外部作用域中的变量。每次调用闭包时，它都能访问并修改这些被捕获的变量。

```go
counter := funcsx.Counter()
counter() // → 1
counter() // → 2
counter() // → 3
```

```rust
// Rust 对应：move 闭包
let mut counter = || {
    count += 1;
    count
};
```

关键特性：
- 每个闭包实例有独立的捕获状态（两个 `Counter()` 返回的闭包互不干扰）
- 对应 Rust `Fn` / `FnMut` / `FnOnce` 三组 trait —— Go 只区分"闭包"与"普通函数"

测试见：`TestCounter`

***

## 2. defer：延迟执行

`defer` 将函数调用推迟到外层函数返回之前执行。常用于资源释放（关闭文件、解锁）。

```go
func DeferInspect(nums []int) (log string) {
    for _, n := range nums {
        defer func() {
            log += fmt.Sprintf("defer(%d)", n)
        }()
    }
    log += "body"
    return
}
```

两个核心规则：

1. **LIFO（后进先出）**：多个 `defer` 按声明顺序的逆序执行
2. **defer 可读写命名返回值**：如上面 `log` 是命名返回值，defer 修改它会影响最终返回

`DeferInspect([]int{1, 2, 3})` 的结果就是 `"bodydefer(3)defer(2)defer(1)"`：

```
执行顺序：           log 的值
log += "body"       → "body"
return              → 捕捉 log
defer(3) 执行       → "bodydefer(3)"
defer(2) 执行       → "bodydefer(3)defer(2)"
defer(1) 执行       → "bodydefer(3)defer(2)defer(1)"
返回                → "bodydefer(3)defer(2)defer(1)"
```

测试见：`TestDeferInspect`

***

## 3. 变参函数（Variadic）

变参函数允许传入可变数量的参数，使用 `...T` 语法声明。对应 Rust 中的 `...` 语法（macro 中的重复模式）。

```go
funcsx.Sum(1, 2, 3)       // → 6
funcsx.Sum(10, 20)        // → 30
funcsx.Sum()              // → 0

// 也可以用切片展开调用
nums := []int{1, 2, 3, 4, 5}
funcsx.Sum(nums...)       // → 15
```

测试见：`TestSum`

***

## 4. panic / recover

- `panic`：触发运行时错误，展开调用栈，执行所有已注册的 `defer`
- `recover`：在 `defer` 中调用，阻止 panic 继续传播，返回 panic 的值

```go
// 安全执行一个可能 panic 的函数
err := funcsx.RecoverFromPanic(func() {
    panic("something went wrong")
})
// err → "panicked: something went wrong"
```

对应 Rust 中 `std::panic::catch_unwind` 的模式。

最佳实践：
- 不要用 panic/recover 做普通的错误处理（那是 `error` 类型的工作）
- recover 只在 `defer` 中有意义，其他地方调用返回 `nil`
- 适合在"顶层 goroutine"或"中间件"中兜底

测试见：`TestRecoverFromPanic`

***

## 5. 自测命令

```bash
go test ./internal/funcsx -shuffle on
go test ./... -shuffle on
```
