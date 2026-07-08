# Go 基础：基本数据类型（含 byte/rune、零值、转换）

> ⚠️ **注意**：这个文档里的示例代码是"讲解片段"，大部分不是可独立运行的完整程序。
> 真正能跑、带断言验证的代码在 `internal/01-basics/basics_test.go`，建议边读边对照。

这个文档对应 [internal/01-basics](./)。

配合阅读：

- 实现代码：[basics.go](./basics.go)
- 测试代码：[basics_test.go](./basics_test.go)（本文档早先版本引用了不存在的 `basic_types_test.go`，已更正）

---

## 1. “基本类型”一览

Go 常见的基本类型可以先按这几类记：

- 布尔：`bool`
- 数字：
  - 整数：`int`、`int8/16/32/64`、`uint`、`uint8/16/32/64`、`uintptr`
  - 浮点：`float32`、`float64`
  - 复数：`complex64`、`complex128`
- 字符串：`string`
- 两个常用别名：
  - `byte` 是 `uint8` 的别名（通常用于“字节”）
  - `rune` 是 `int32` 的别名（通常用于“Unicode 码点”）

补充：`error` 是一个接口类型（不是基础类型），但它在工程里使用频率非常高。

> ⚠️ **注意**：`int` / `uint` 的**宽度跟随平台**（64 位机器上是 64 位，32 位上是 32 位），不像 Rust 的 `i32`/`i64` 固定。
> 所以涉及网络协议、文件格式、跨平台数据交换时，要显式写 `int32`/`int64`，不要用裸 `int`，否则不同平台上二进制布局会不一致。

> ⚠️ **注意**：`byte` 和 `rune` 只是别名，不是新类型。这意味着 `byte` 和 `uint8` 可以直接互用、`rune` 和 `int32` 可以直接互用，编译器视为同一类型。别把它们当成"strong typedef"。

---

## 2. 零值（zero value）

Go 的变量只要声明，就一定有值（零值），这对写出“少初始化代码 + 更可预测”的程序很重要：

- `bool` -> `false`
- 数字 -> `0` / `0.0`
- `string` -> `""`
- 引用/指针类（指针、slice、map、chan、func、interface）-> `nil`
- `struct` -> 每个字段都是零值

测试例子见：[basics_test.go](./basics_test.go)（`TestGetZeroValues` 验证了 `ZeroValues{}` 全零值）。

> ⚠️ **注意（最经典的坑）**：`var m map[string]int` 得到的是 **nil map**，声明即"零值"，但**不能写入**——`m["k"]=1` 会直接 panic。
> 要写入必须先 `m = make(map[string]int)` 或 `m = map[string]int{}`。指针、slice、chan 同理：nil slice 可以 `append`（append 内部会初始化），但 nil map/channel 写入会 panic。这是 Go 初学者踩得最多的雷。

> ⚠️ **注意**：结构体整体可以用 `T{}` 得到全零值，但**含 slice/map 字段的 struct 不能用 `==` 比较**（编译错误），要用 `reflect.DeepEqual`。这是零值机制带来的一个边界限制。

---

## 3. byte vs rune：字符串到底是什么

关键结论先记住两条：

1. `string` 是一段 **只读字节序列**（UTF-8 文本只是最常见的约定）
2. `len(s)` 返回的是 **字节数**，不是“字符数”

示例（注意 `len` 与 “rune 个数” 的差别）：

```go
s := "你好Go"
_ = len(s)         // 8（字节数：你=3 好=3 Go=2）
_ = len([]rune(s)) // 4（码点数）
```

再记一个常见坑：`s[i]` 的结果是 `byte`，不是 `rune`。

拿“第一个字符”，要用 `for range` 或先转 `[]rune`：

```go
var first rune
for _, r := range s {
	first = r
	break
}
```

对应测试见：[basics_test.go](./basics_test.go)（`TestReverseString` 验证了多字节字符反转不乱码）。

> ⚠️ **注意**：直接按下标 `s[i]` 取出来的是**单个字节（byte）**，对中文这种多字节字符来说，单个字节不是合法字符。要按"字符"处理，必须转 `[]rune` 或用 `for range`（range 会按 rune 迭代，自动处理多字节）。`ReverseString` 函数如果不转 rune 直接反转 `[]byte`，中文会变成乱码——这是字符串处理的第一坑。

> ⚠️ **注意**：`string` 是**只读**的，不能 `s[i] = 'x'`。要修改必须先转 `[]rune` 或 `[]byte`，改完再转回 `string`（每次转换都会分配新内存）。

---

## 4. 数字类型与转换（conversion）

Go 不做“隐式数值类型转换”，不同整数类型之间赋值通常要显式转换：

```go
var a int32 = 1
var b int64 = int64(a)
```

这样做的好处是：代码读起来更明确，不会悄悄丢精度或溢出。

复数的两个分量可通过 `real/imag` 取出：

```go
c := complex(1, 2)
_ = real(c) // 1
_ = imag(c) // 2
```

> ⚠️ **注意**：Go **没有** Rust/C++ 那种隐式数值提升。`int32` 不能直接赋给 `int64`，`int` 和 `int64` 也不是同一类型，必须显式 `int64(x)`。好处是不会"悄悄"丢精度，坏处是写起来啰嗦——但这是刻意的取舍（显式优于隐式）。

> ⚠️ **注意**：转换是"截断/重解释"，不是"安全转换"。`int64` 转 `int32` 若溢出会**静默丢高位**；浮点转整数向零截断（`int(3.9)==3`）。需要安全转换时要自己写范围检查。

---

## 4.5 数字字面量细节（下划线、`0b`/`0o`/`0x`、复数）

Go 字面量支持多种进制和可读性分隔符，考试/面试常考，但日常代码容易忽略：

```go
// 来自 basics.go 的 NumericLiteralsDemo
_ = 1_000_000     // 千分位下划线，纯可读性，编译期去掉
_ = 0b1010        // 二进制
_ = 0o777         // 八进制（Go 1.13+ 推荐 0o 前缀，老的 0777 也行但易混）
_ = 0xFF          // 十六进制
_ = 1 + 2i        // 复数字面量（虚部级用 i 后缀）
```

对应测试见：[basics_test.go](./basics_test.go)（`TestNumericLiteralsDemo`）。

> ⚠️ **注意**：下划线 `_` 在字面量里**只能放在数字之间**（如 `1_000_000`），不能开头/结尾，也不能连续两个。它纯粹提升可读性，编译器直接忽略。复数 `1 + 2i` 里的 `i` 是虚部单位，不是变量——写成 `1+2i` 也合法，但建议加空格避免和标识符混淆。

> ⚠️ **注意（八进制坑）**：`0777` 这种老写法在 Go 里是**八进制**不是十进制！Go 1.13 后推荐显式 `0o777`，避免误读。同理 `0x` 是十六进制、`0b` 是二进制。这是《Learning Go》第 2 章特别强调的"字面量易错点"。

---

## 4.6 原始字符串字面量（反引号）

Go 有两种字符串字面量：解释型（双引号 `"`）和原始型（反引号 `` ` ``）。

```go
// 来自 basics.go 的 RawStringDemo
s := `这是原始字符串：
包含换行
制表符:	
Windows路径: C:\Users\Name
JSON: {"key": "value"}
`
```

对应测试见：[basics_test.go](./basics_test.go)（`TestRawStringDemo`）。

> ⚠️ **注意**：原始字符串**不做任何转义**——`\n`、`\t`、`\"`、甚至反斜杠本身都按字面字符保留。它天生适合写多行文本、正则表达式、JSON/SQL 模板、Windows 路径（不用双写反斜杠）。但**原始字符串里不能出现反引号**，也不能跨行做字符串拼接。解释字符串里写换行必须 `\n`，写反斜杠必须 `\\`。

---

## 4.7 浮点数陷阱（货币别用 float、比较用 epsilon）

这是工程里最容易造成"金额错账"的坑，教材《Learning Go》第 2 章反复强调：

```go
// 来自 basics.go 的 FloatPitfallsDemo 与 FloatEpsilonEqual
func FloatPitfallsDemo() (bool, bool) {
	a := 0.1
	b := 0.2
	sum := a + b
	moneyWrong := sum == 0.3          // false！浮点误差
	epsilon := 1e-9
	_ = math.Abs(sum-0.3) < epsilon    // true，正确比较法
	directEq := a+b == 0.3             // false
	return moneyWrong, directEq
}
func FloatEpsilonEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
```

对应测试见：[basics_test.go](./basics_test.go)（`TestFloatPitfallsDemo`）。

> ⚠️ **注意（两大坑）**：
> 1. **货币/精确小数绝对别用 `float32`/`float64`**。二进制浮点无法精确表示 `0.1` 这类十进制小数，`0.1+0.2 = 0.30000000000000004`，直接拿去算钱会错账。正确做法是用**整数（分/厘）**或第三方 `decimal` 库。
> 2. **浮点比较别用 `==`**，必须用容差：`math.Abs(a-b) < 1e-9`。即使 `0.1+0.2`，`== 0.3` 也是 false。

> ⚠️ **注意（Go 特有坑：常量折叠）**：Go 的**无类型常量**在编译期用**任意精度**运算，所以 `0.1+0.2 == 0.3` 在编译期其实是 **true**！只有用**变量**（`a := 0.1; b := 0.2; a+b`）才会触发运行时的二进制浮点误差。上面 `FloatPitfallsDemo` 特意用变量，才能复现陷阱。这是教材没明说、但极易踩的隐蔽点。

---

## 5. “值类型”与“引用语义”的快速对比（入门够用）

这不是严格的类型学分类，但对写代码很有帮助：

- 更偏“值语义”：数组 `array`、结构体 `struct`（赋值通常会拷贝内容）
- 更偏“引用语义”：切片 `slice`、映射 `map`、通道 `chan`、函数 `func`（赋值通常复制的是“描述符/引用”，底层数据共享）

一个入门级的验证方式：

- 数组拷贝后互不影响
- 切片赋值后可能共享同一底层数组（改动一处，另一处能看到）

对应测试见：[basics_test.go](./basics_test.go)（`TestSliceCopyDemo` 验证了 `copy` 是深拷贝）。

> ⚠️ **注意**：说 slice/map 是"引用语义"是**简化说法**。准确地说，slice 本身是个"描述符"（指针+长度+容量）的小结构体，赋值只是复制这个描述符，所以两个 slice 变量**可能共享底层数组**。这导致一个隐蔽 bug：对一个 slice 做 `append` 可能改写另一个 slice 看到的底层数据（当容量足够、未触发扩容时）。要真正独立，用 `copy()`。详见 [切片与映射_slice与map.md](./切片与映射_slice与map.md)。

> ⚠️ **注意（Rust 对照）**：Rust 里 `Vec<T>` 是拥有所有权的，赋值会 move（除非 clone）。Go 的 slice 没有所有权概念，赋值永远只是复制描述符、共享底层数组——这是两种语言内存模型的根本差异，习惯 Rust 的人最容易在这里误判"是不是独立副本"。

---

## 6. 自测命令

```bash
go test ./internal/01-basics -shuffle on
go test ./... -shuffle on
```
