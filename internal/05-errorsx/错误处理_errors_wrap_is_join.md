# Go 基础：错误处理（errors / %w / Is / Join）

这个文档对应 [internal/05-errorsx](../05-errorsx)。

配合阅读：
- 实现代码：[errorsx.go](../05-errorsx/errorsx.go)
- 测试代码：[errorsx_errorsx_test.go](../16-tests/errorsx_errorsx_test.go)

> ⚠️ **注意**：本目录（05-errorsx）只放实现与知识文档；对应的单元测试统一放在 `internal/16-tests/` 下，文件名是 `errorsx_errorsx_test.go`。本文档里的测试链接已指向真实位置。

---

## 1. Go 的 error 是什么

Go 的 `error` 是一个接口：

```go
type error interface {
	Error() string
}
```

当一个函数可能失败时，常见签名是：

```go
func Do() (T, error)
```

调用方需要显式处理：

```go
v, err := Do()
if err != nil {
	return err
}
_ = v
```

> ⚠️ **注意**：Go 没有 Java 的受检异常（`throws`/`try-catch`）或 `Optional`，错误是**普通的多返回值**，必须手写 `if err != nil` 处理。忘写 `err != nil` 检查，编译能通过，但程序可能在错误状态下继续运行——这是 Go 里最容易被忽视的雷。对应 Java：`(T, error)` ≈ `try-catch` 捕获 `Exception`，但 Go 把错误处理完全交给程序员（没有编译器强制检查）。

---

## 2. 创建错误：errors.New 与 fmt.Errorf

### 2.1 errors.New（哨兵错误）

```go
// 来自 errorsx.go
var ErrNotPositive = errors.New("value must be positive")
```

这种通常叫“哨兵错误”（sentinel error），用于让调用方做判断。

> ⚠️ **注意**：哨兵错误要用 `var` 定义成**包级变量**（首字母大写导出），并习惯以 `Err` 开头命名。调用方应当用 `errors.Is(err, ErrNotPositive)` 判断，而不是比较字符串。原因在于：每次 `errors.New(...)` 都会生成一个**新的、互不相等**的实例，直接 `==` 字符串或拿不同的 `errors.New` 比对都会失败。

### 2.2 fmt.Errorf（带上下文）

```go
// 来自 errorsx.go 的 ParsePositiveInt
return fmt.Errorf("parse int: %w", err)
```

它可以带上下文信息（更利于排查）。

---

## 3. 错误包装（wrap）：%w

Go 推荐你在往上层返回时保留“原始错误”，做法是 `%w`：

```go
// 来自 errorsx.go
n, err := strconv.Atoi(s)
if err != nil {
	return 0, fmt.Errorf("parse int: %w", err)
}
```

这比 `%v` 更强，因为它让你后续可以用 `errors.Is` / `errors.As` 判断“根因是否是某个错误”。

> ⚠️ **注意**：**wrap 用 `%w`，不要用 `%v`**。`%v` 只是把错误拼进字符串，原始的 error 链断了，上层再用 `errors.Is` / `errors.As` 就找不到根因。换句话说：`%w` 保留结构，`%v` 只保留文本。

---

## 4. errors.Is：判断错误是否为某类错误

不要用字符串判断（例如 `err.Error() == "xxx"`），因为：
- 文案可能变化
- 包装后字符串更复杂

用 `errors.Is`：

```go
if errors.Is(err, ErrNotPositive) {
	// 处理 NotPositive
}
```

> ⚠️ **注意**：`errors.Is` 会**沿着 `%w` 包装链一层层往下找**，只要链路中某个错误等于目标哨兵就返回 true。即使错误被包裹了三层，也能正确识别根因。字符串比较则完全做不到这点——一旦中间包了上下文，比较就失败。对应 Java：`errors.Is` ≈ `catch (MyException e)` 或 `if (err instanceof MyException)`，按类型匹配根因。

### 自定义错误类型用 errors.As

`errorsx.go` 里还有一个自定义错误：

```go
type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}
```

想提取它的具体字段，用 `errors.As`：

```go
var de *DivideError
if errors.As(err, &de) {
	fmt.Println("除零发生在:", de.Dividend, de.Divisor)
}
```

> ⚠️ **注意**：**提取具体错误类型用 `errors.As`，判断是否为某哨兵用 `errors.Is`**，二者不要混。而且自定义错误的方法集接收者要一致——`DivideError` 用的是 `*DivideError` 接收者，那么 `SafeDivide` 返回的是 `&DivideError{...}`（指针），`errors.As` 的第二个参数也得是 `**DivideError` 对应的指针变量。

---

## 5. errors.Join：合并多个错误

当你需要“同时返回多个错误”（例如：关闭多个资源时都失败了），可以用 `errors.Join`：

```go
// 来自 errorsx.go 的 Join
func Join(a, b error) error {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return errors.Join(a, b)
}
```

Join 出来的错误也能被 `errors.Is` 识别（只要其中包含某个哨兵错误）。

> ⚠️ **注意**：`errors.Join` 对 nil 有短路逻辑——只要有一个是 nil 就返回另一个；两个都 nil 才返回 nil。`errors.Is` 作用在 Join 结果上时会**逐个检查**被合并的每个错误，所以即便多个错误被拼成一个，仍能识别出里面的某个哨兵错误。

---

## 6. 初学者最容易犯的错误（请对照）

1. 忽略 err：`_ = err` 或直接不处理（容易埋雷，编译不报错但运行时状态错乱）
2. 用字符串比较判断错误（`err.Error() == "xxx"`），包装后立刻失效
3. wrap 用 `%v`，导致上层 `errors.Is` 无法判断根因
4. 把错误当日志打完就吞掉（上层不知道失败了）
5. 用 `errors.New` 在每次返回时现造错误去做 `==` 比较（每次都是新实例，比较永远不等）

---

## 7. 练习清单

1. 写一个 `ParsePort(s string) (int, error)`：非数字或范围不对就返回带上下文的错误（用 `%w` 包裹 `strconv` 的原始错误）
2. 定义 `var ErrEmptyInput = errors.New("empty input")`，让调用方用 `errors.Is` 判断
3. 写一个关闭函数 `CloseAll(cs ...io.Closer) error`，把多个关闭错误用 `errors.Join` 合并返回

---

## 8. 自测命令

```bash
go test ./internal/16-tests -run Errorsx -shuffle on
```
