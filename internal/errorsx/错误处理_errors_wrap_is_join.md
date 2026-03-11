# Go 基础：错误处理（errors / %w / Is / Join）

这个文档对应 [internal/errorsx](file:///e:/dingchuan/GoJuniper/internal/errorsx)。

配合阅读：
- 实现代码：[errorsx.go](file:///e:/dingchuan/GoJuniper/internal/errorsx/errorsx.go)
- 测试代码：[errorsx_test.go](file:///e:/dingchuan/GoJuniper/internal/errorsx/errorsx_test.go)

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

---

## 2. 创建错误：errors.New 与 fmt.Errorf

### 2.1 errors.New

```go
var ErrNotFound = errors.New("not found")
```

这种通常叫“哨兵错误”（sentinel error），用于让调用方做判断。

### 2.2 fmt.Errorf

```go
return fmt.Errorf("load user: %v", err)
```

它可以带上下文信息（更利于排查）。

---

## 3. 错误包装（wrap）：%w

Go 推荐你在往上层返回时保留“原始错误”，做法是 `%w`：

```go
return fmt.Errorf("load config: %w", err)
```

这比 `%v` 更强，因为它让你后续可以用 `errors.Is` 判断“根因是否是某个错误”。

---

## 4. errors.Is：判断错误是否为某类错误

不要用字符串判断（例如 `err.Error() == "xxx"`），因为：
- 文案可能变化
- 包装后字符串更复杂

用 `errors.Is`：

```go
if errors.Is(err, ErrNotFound) {
	// 处理 NotFound
}
```

---

## 5. errors.Join：合并多个错误

当你需要“同时返回多个错误”（例如：关闭多个资源时都失败了），可以用 `errors.Join`：

```go
return errors.Join(err1, err2, err3)
```

Join 出来的错误也能被 `errors.Is` 识别（只要其中包含某个哨兵错误）。

---

## 6. 初学者最容易犯的错误（请对照）

1. 忽略 err：`_ = err` 或直接不处理（容易埋雷）
2. 用字符串比较判断错误
3. wrap 用 `%v`，导致上层 `errors.Is` 无法判断
4. 把错误当日志打完就吞掉（上层不知道失败了）

---

## 7. 练习清单

1. 写一个 `ParsePort(s string) (int, error)`：非数字或范围不对就返回带上下文的错误
2. 定义 `var ErrEmptyInput = errors.New("empty input")`，让调用方用 `errors.Is` 判断
3. 写一个关闭函数 `CloseAll(cs ...io.Closer) error`，把多个关闭错误用 `errors.Join` 合并返回

---

## 8. 自测命令

```bash
go test ./internal/errorsx -shuffle on
```

