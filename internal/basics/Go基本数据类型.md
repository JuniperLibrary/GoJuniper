# Go 基础：基本数据类型（含 byte/rune、零值、转换）

这个文档对应 [internal/basics](file:///e:/dingchuan/GoJuniper/internal/basics)。

配合阅读：

- 实现代码：[basics.go](file:///e:/dingchuan/GoJuniper/internal/basics/basics.go)
- 测试代码：[basic_types_test.go](file:///e:/dingchuan/GoJuniper/internal/basics/basic_types_test.go)

***

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

***

## 2. 零值（zero value）

Go 的变量只要声明，就一定有值（零值），这对写出“少初始化代码 + 更可预测”的程序很重要：

- `bool` -> `false`
- 数字 -> `0` / `0.0`
- `string` -> `""`
- 引用/指针类（指针、slice、map、chan、func、interface）-> `nil`
- `struct` -> 每个字段都是零值

测试例子见：[basic_types_test.go](file:///e:/dingchuan/GoJuniper/internal/basics/basic_types_test.go)

***

## 3. byte vs rune：字符串到底是什么

关键结论先记住两条：

1. `string` 是一段 **只读字节序列**（UTF-8 文本只是最常见的约定）
2. `len(s)` 返回的是 **字节数**，不是“字符数”

示例（注意 `len` 与 “rune 个数” 的差别）：

```go
s := "你好Go"
_ = len(s)         // 8（字节数）
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

对应测试见：[basic_types_test.go](file:///e:/dingchuan/GoJuniper/internal/basics/basic_types_test.go)

***

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

***

## 5. “值类型”与“引用语义”的快速对比（入门够用）

这不是严格的类型学分类，但对写代码很有帮助：

- 更偏“值语义”：数组 `array`、结构体 `struct`（赋值通常会拷贝内容）
- 更偏“引用语义”：切片 `slice`、映射 `map`、通道 `chan`、函数 `func`（赋值通常复制的是“描述符/引用”，底层数据共享）

一个入门级的验证方式：

- 数组拷贝后互不影响
- 切片赋值后可能共享同一底层数组（改动一处，另一处能看到）

对应测试见：[basic_types_test.go](file:///e:/dingchuan/GoJuniper/internal/basics/basic_types_test.go)

***

## 6. 自测命令

```bash
go test ./internal/basics -shuffle on
go test ./... -shuffle on
```
