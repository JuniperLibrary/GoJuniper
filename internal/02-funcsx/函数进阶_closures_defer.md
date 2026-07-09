# Go 基础：函数进阶（闭包、defer、变参、panic/recover）

这个文档对应 [internal/02-funcsx](./)。

配合阅读：
- 实现代码：[funcsx.go](./funcsx.go)
- 测试代码：[funcsx_test.go](./funcsx_test.go)

***

## 1. 闭包（Closure）

闭包是一个函数值，它"捕获"了其外部作用域中的变量。每次调用闭包时，它都能访问并修改这些被捕获的变量。

```go
counter := funcsx.Counter()
counter() // → 1
counter() // → 2
counter() // → 3
```

```java
// Java 对应：用类字段保存计数状态（lambda 只能捕获 final/等效 final 变量，不能自增）
class Counter {
    private int count = 0;
    public int next() { return ++count; }
}
```

关键特性：
- 每个闭包实例有独立的捕获状态（两个 `Counter()` 返回的闭包互不干扰）
- Go 的闭包直接读写捕获变量，类似 Java 用对象字段持有可变状态（Java lambda 自身不可变捕获，需借助对象）

> ⚠️ **注意**：Go 的闭包**自动按引用捕获**外部变量（Java 的 lambda 只能捕获 `final`/等效 final 的局部变量，且不能修改；要修改需包成对象）。循环里常见坑：如果多个闭包共享同一个循环变量（如 `for i := 0; ...` 里的 `i`），它们捕获的是**同一个变量**，等闭包真正执行时 `i` 已经是循环结束时的值。需要"每次迭代一份独立副本"时，要在循环体内再声明一个局部变量，或用函数参数传值。

> ⚠️ **注意**：调用两次 `Counter()` 会得到**两个独立的闭包**，各有各的 `count`，互不干扰。判断"是不是同一个状态"的关键，在于闭包捕获的是哪个作用域里的变量——捕获的变量声明在 `Counter` 内部，每次调用 `Counter` 都会新建一份。

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

> ⚠️ **注意**：**多个 `defer` 是 LIFO（后进先出）**。循环里先注册 `defer(1)`、`defer(2)`、`defer(3)`，执行时反着来——`defer(3)` → `defer(2)` → `defer(1)`。这是 Go 里最容易算错执行顺序的地方。

> ⚠️ **注意**：**`defer` 的参数在"注册时"求值，而不是执行时**。本例 `defer func(){ log += ...(n) }()` 里闭包捕获的 `n` 是 `range` 循环每次迭代的副本，所以拿到的是 1、2、3 各一份。如果闭包捕获的是一个在循环结束后才变的变量，就会全部拿到同一个终值——典型的反直觉 bug。

> ⚠️ **注意**：**`defer` 只能修改"命名返回值"**。`DeferInspect` 的签名是 `(log string)`，defer 里往 `log` 追加会反映到最终返回。若写成普通 `return log`（非命名返回），defer 对局部变量的修改不会成为返回值。为什么这样设计？因为命名返回值本身就是函数栈帧里的一个变量，defer 在 `return` 之前还能动它。

测试见：`TestDeferInspect`

***

## 3. 变参函数（Variadic）

变参函数允许传入可变数量的参数，使用 `...T` 语法声明。对应 Java 中的 `T...` 可变参数语法。

```go
funcsx.Sum(1, 2, 3)       // → 6
funcsx.Sum(10, 20)        // → 30
funcsx.Sum()              // → 0

// 也可以用切片展开调用
nums := []int{1, 2, 3, 4, 5}
funcsx.Sum(nums...)       // → 15
```

> ⚠️ **注意**：**变参 `...T` 必须放在参数列表的最后一位**，不能写成 `(nums ...int, sep string)`；而且**一个函数最多只能有一个变参**。函数体内 `nums` 就是普通的 `[]T` 切片，可以 `for range`。调用时既可用字面量 `Sum(1,2,3)`，也可用切片展开 `Sum(slice...)`。

> ⚠️ **注意**：Go 里**函数是一等公民**——可以像 `int` 一样作为参数、返回值、存进变量（见 `funcsx.ApplyFunc` 把 `func(int,int) int` 当参数）。对应 Java 的 `Function`/`BinaryOperator` 等函数式接口（Java 8+ 引入），只是 Go 的闭包签名直接写全即可，无需定义函数式接口类型。

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

对应 Java 中 `try-catch(Throwable)` 兜底捕获的模式。

最佳实践：
- 不要用 panic/recover 做普通的错误处理（那是 `error` 类型的工作）
- recover 只在 `defer` 中有意义，其他地方调用返回 `nil`
- 适合在"顶层 goroutine"或"中间件"中兜底

> ⚠️ **注意**：**`recover()` 只在 `defer` 函数里调用才有效**，放在别处调用永远返回 `nil`——因为它必须"拦截正在展开的 panic"。典型套路是 `defer func(){ if r := recover(); r != nil { err = ... } }()`，且依赖**命名返回值**把错误带出去（见 `RecoverFromPanic` 的 `(err error)`）。

> ⚠️ **注意**：**panic/recover 不是 Go 的常规错误处理方式**，只用于真正不可恢复、或需要兜底防整个程序崩溃的场景。普通错误请用 `error` 返回（类似 Java 里用 `try-catch`/`Exception` 而不是吞掉未捕获异常）。滥用 recover 会掩盖 bug。

测试见：`TestRecoverFromPanic`

***

## 5. 自测命令

```bash
go test ./internal/02-funcsx -shuffle on
go test ./... -shuffle on
```
