---
name: go-basics
description: Go语言基础语法学习 - 变量、常量、类型、控制流
license: MIT
---

# Go 基础语法

## 1. 变量（Variables）

### 1.1 变量声明方式

```go
// 方式1：完整声明
var name string = "Alice"
var age int = 25

// 方式2：类型推断
var name = "Alice"
var age = 25

// 方式3：短变量声明（函数内）
name := "Alice"
age := 25

// 方式4：变量组
var (
    host = "localhost"
    port = 8080
)
```

### 1.2 零值（Zero Values）

Go变量声明后自动初始化为零值：
- `int`, `float`: `0`
- `bool`: `false`
- `string`: `""`
- `pointer`, `slice`, `map`, `channel`, `function`, `interface`: `nil`

```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
var p *int      // nil
```

### 1.3 多重赋值

```go
a, b := 1, 2
a, b = b, a  // 交换值
```

## 2. 常量（Constants）

### 2.1 基本常量

```go
const Pi = 3.14159
const AppName = "MyApp"

// 常量组
const (
    StatusOK    = 200
    StatusError = 500
)
```

### 2.2 iota 枚举

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)

// 跳过值
const (
    _ = iota       // 0 丢弃
    KB = 1 << (10 * iota)  // 1024
    MB                      // 1048576
    GB                      // 1073741824
)
```

## 3. 基本数据类型

### 3.1 数值类型

```go
// 整数
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr

// 浮点
float32, float64

// 复数
complex64, complex128

// 别名
byte  // uint8
rune  // int32 (Unicode码点)
```

### 3.2 布尔与字符串

```go
var b bool = true
var s string = "Hello, Go!"

// 字符串是不可变的
// 多行字符串使用反引号
var multi = `第一行
第二行`
```

## 4. 控制流

### 4.1 if 语句

```go
if x > 0 {
    // ...
}

// 带初始化语句
if err := doSomething(); err != nil {
    return err
}
```

### 4.2 for 循环

```go
// 标准for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while风格
for x < 10 {
    x++
}

// 无限循环
for {
    break  // 或 continue
}

// 遍历slice/map
for index, value := range slice {
    // ...
}
for key, value := range map {
    // ...
}
```

### 4.3 switch 语句

```go
switch x {
case 1:
    fmt.Println("one")
case 2, 3:
    fmt.Println("two or three")
default:
    fmt.Println("other")
}

// 无条件switch
switch {
case x > 0:
    fmt.Println("positive")
case x < 0:
    fmt.Println("negative")
default:
    fmt.Println("zero")
}
```

### 4.4 select 语句（用于channel）

```go
select {
case msg := <-ch1:
    fmt.Println(msg)
case ch2 <- value:
    fmt.Println("sent")
default:
    fmt.Println("no activity")
}
```

## 5. 指针

```go
x := 10
p := &x       // p指向x
fmt.Println(*p)  // 输出10
*p = 20       // 修改x的值
```

## 6. 类型转换

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// 字符串转换
import "strconv"
s := strconv.Itoa(42)    // int -> string
i, _ := strconv.Atoi("42") // string -> int
```

## 练习清单

1. 用不同方式声明变量，体会 `var` vs `:=` 的区别
2. 写一个 `const` 枚举，用 `iota` 定义状态码
3. 用 `for range` 遍历字符串，观察 `rune` 的作用
4. 写一个 `switch` 判断成绩等级
5. 练习指针操作，理解值传递与引用传递