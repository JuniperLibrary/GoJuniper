# Go 基础：I/O（io.Reader / io.Writer / bufio）

这个文档对应 [internal/iox](file:///e:/dingchuan/GoJuniper/internal/iox)。

配合阅读：
- 实现代码：[iox.go](file:///e:/dingchuan/GoJuniper/internal/iox/iox.go)
- 测试代码：[iox_test.go](file:///e:/dingchuan/GoJuniper/internal/iox/iox_test.go)

---

## 1. 为什么 Go 的 I/O 这么“抽象”

Go 标准库把各种输入输出统一成两个核心接口：

- `io.Reader`：读数据
- `io.Writer`：写数据

它们是“最小抽象”，因此非常通用：
- 文件、网络连接、内存 buffer、字符串，都可以是 Reader/Writer

---

## 2. io.Reader

定义（你不用背，但要理解语义）：

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

核心理解：
- `Read` 把数据放进 `p`
- 返回实际读到的字节数 `n`
- `err == io.EOF` 表示读完了（通常 `n` 可能 > 0，也可能为 0）

---

## 3. io.Writer

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

核心理解：
- `Write` 把 `p` 的数据写出去
- 返回实际写入字节数 `n`

---

## 4. bufio：为什么需要缓冲

缓冲的目的通常是：减少系统调用次数，提高性能。

常见用法：
- `bufio.Scanner`：按行读取（适合文本）
- `bufio.Reader`：更灵活的读取
- `bufio.Writer`：缓冲写，最后记得 `Flush()`

`Scanner` 有一个初学者常见点：
- 默认 token（每行）有长度限制，超长行要调整 Buffer（工程里会遇到）

---

## 5. 本模块在练什么

请先看测试，再看实现：
- 测试：[iox_test.go](file:///e:/dingchuan/GoJuniper/internal/iox/iox_test.go)
- 实现：[iox.go](file:///e:/dingchuan/GoJuniper/internal/iox/iox.go)

关注点：
- 读完整个 Reader 并处理 EOF
- 用 Scanner 按行读取
- 用 Writer 写出结果

---

## 6. 练习清单

1. 写 `CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)`（参考 `io.CopyN` 的行为）
2. 写 `CountLines(r io.Reader) (int, error)`（用 Scanner）
3. 写 `UppercaseLines(r io.Reader, w io.Writer) error`（按行读、转大写、按行写）

---

## 7. 自测命令

```bash
go test ./internal/iox -shuffle on
```

