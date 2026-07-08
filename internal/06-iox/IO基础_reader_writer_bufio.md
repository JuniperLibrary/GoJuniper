# Go 基础：I/O（io.Reader / io.Writer / bufio）

这个文档对应 [internal/06-iox](file:///e:/dingchuan/GoJuniper/internal/06-iox)。

配合阅读：
- 实现代码：[iox.go](file:///e:/dingchuan/GoJuniper/internal/06-iox/iox.go)

> ⚠️ **注意**：本仓库 `06-iox` 目录下只有实现代码 `iox.go`，没有 `_test.go`。学习时直接看 `iox.go` 里的函数实现（`ReadAllString`、`ReadLines`、`WriteLines` 等）即可，练习清单里的函数也写在同包内。

---

## 1. 为什么 Go 的 I/O 这么“抽象”

Go 标准库把各种输入输出统一成两个核心接口：

- `io.Reader`：读数据
- `io.Writer`：写数据

它们是“最小抽象”，因此非常通用：
- 文件、网络连接、内存 buffer、字符串，都可以是 Reader/Writer

> ⚠️ **注意**：`io.Reader` / `io.Writer` 是**接口**，只要一个类型实现了 `Read(p []byte) (n int, err error)` 就自动满足 `io.Reader`，无需显式 `implements` 声明。这意味着 `*os.File`、`bytes.Buffer`、`strings.Reader`、甚至你自定义的类型都能直接传给任何接收 `io.Reader` 的函数——这是 Go 鸭子类型的核心威力，也是和 Java 里 `implements` 不同的地方（Java 需要显式 `implements` 接口，且接口方法需在类中声明）。

**为什么**要把 I/O 抽象成接口？因为这样 `io.Copy(dst, src)` 能在文件↔网络↔内存之间零耦合地搬运数据，你的业务函数只要接收 `io.Reader` 而不是 `*os.File`，测试时就能传 `strings.NewReader("...")`，完全不碰真实文件系统。

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

> ⚠️ **注意**：`Read` **不保证一次填满 `p`**。正确调用循环是：反复读直到 `err == io.EOF`，且每次只处理前 `n` 个字节。新手常误以为一次 `Read` 就能拿到全部数据。想一次性读完全部内容，直接用 `io.ReadAll(r)`（本仓库 `ReadAllString` 就是这么做的）。

**常见坑**：`io.EOF` 经常和 `n > 0` 同时出现（最后一次读到部分数据并读完）。所以判断“是否读完”要看 `err`，不要看 `n == 0`。

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

> ⚠️ **注意**：`Write` 同样**不保证一次写完 `p` 全部字节**。标准实现要求 `n < len(p)` 时必须返回 `err`，调用方需要重试剩余部分。不过 `bufio.Writer` 和 `*os.File` 等常见实现会替你处理好这一点，日常使用基本可以当“整段写入”看待。

**为什么** `Write` 返回 `n`？因为底层（如网络 socket）可能因缓冲区满只接收了一部分，返回 `n` 让上层知道“还差多少没写”。

---

## 4. bufio：为什么需要缓冲

缓冲的目的通常是：减少系统调用次数，提高性能。

常见用法：
- `bufio.Scanner`：按行读取（适合文本）
- `bufio.Reader`：更灵活的读取
- `bufio.Writer`：缓冲写，最后记得 `Flush()`

`Scanner` 有一个初学者常见点：
- 默认 token（每行）有长度限制，超长行要调整 Buffer（工程里会遇到）

> ⚠️ **注意（bufio.Writer 必须 Flush）**：`bufio.NewWriter(w)` 会把数据攒在内存 buffer 里，**只有调用 `Flush()` 才会真正写出去**。本仓库 `WriteLines` 在结尾 `return bw.Flush()` 就是在刷盘——如果忘了 Flush，buffer 里的数据会**静默丢失**，程序还不报错，这是最隐蔽的 bug 之一。

> ⚠️ **注意（Scanner 行长度限制）**：`bufio.Scanner` 默认单行上限是 `bufio.MaxScanTokenSize`（64KB）。读超长行会触发 `bufio.ErrTooLong`，`Scan()` 返回 false 且 `Err()` 报错。本仓库 `ReadLines` 用 `sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)` 显式把上限提到 1MB——生产环境读取不可信输入（如大日志、CSV）时必须这么做。

**为什么**需要缓冲？每次 `Write` 都直接落到磁盘/网络意味着一次系统调用，高频小写入时开销巨大。缓冲攒一批再一次性写出，把 N 次系统调用压成 1 次。

---

## 5. 本模块在练什么

请先看实现，再自己动手写练习函数（同包内）：
- 实现：[iox.go](file:///e:/dingchuan/GoJuniper/internal/06-iox/iox.go)

关注点：
- 读完整个 Reader 并处理 EOF（用 `io.ReadAll` 最简单）
- 用 Scanner 按行读取（`ReadLines` 是范例）
- 用 Writer 写出结果（记得 `Flush`）

> ⚠️ **注意**：`os.ReadFile` / `os.WriteFile` 是一次性读写整个文件的便捷封装（本仓库 `ReadFile`/`WriteFile` 用的就是它），适合小文件；大文件请改用 `io.Reader`/`io.Writer` 流式处理，避免把整文件读进内存。

---

## 6. 练习清单

1. 写 `CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)`（参考 `io.CopyN` 的行为）
2. 写 `CountLines(r io.Reader) (int, error)`（用 Scanner）
3. 写 `UppercaseLines(r io.Reader, w io.Writer) error`（按行读、转大写、按行写）

---

## 7. 自测命令

```bash
go test ./internal/06-iox -shuffle on
```

> ⚠️ **注意**：当前目录没有测试文件，该命令会提示 “no test files”。先把练习清单里的函数补上、或自行补 `iox_test.go`，命令才会真正跑起来。
