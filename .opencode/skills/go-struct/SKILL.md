---
name: go-struct
description: Go语言结构体学习 - 结构体、方法、接口、组合
license: MIT
---

# Go 结构体与接口

## 1. 结构体（Struct）

### 1.1 定义与初始化

```go
type Person struct {
    Name string
    Age  int
    City string
}

// 初始化方式
p1 := Person{Name: "Alice", Age: 25}           // 指定字段
p2 := Person{"Bob", 30, "Beijing"}             // 按顺序
p3 := new(Person)                               // 指针，零值
p4 := &Person{Name: "Charlie"}                  // 指针字面量
```

### 1.2 匿名字段与嵌入

```go
type Address struct {
    City    string
    Country string
}

type User struct {
    Name    string
    Address          // 匿名嵌入（组合）
}

u := User{Name: "Alice", Address: Address{City: "Beijing"}}
fmt.Println(u.City)  // 直接访问嵌入字段: "Beijing"
```

### 1.3 结构体标签（Tag）

```go
type User struct {
    Name string `json:"name" validate:"required"`
    Age  int    `json:"age" validate:"min=0,max=150"`
}

// 读取标签
import "reflect"
t := reflect.TypeOf(User{})
field, _ := t.FieldByName("Name")
fmt.Println(field.Tag.Get("json"))  // "name"
```

## 2. 方法

### 2.1 方法定义

```go
type Rectangle struct {
    Width, Height float64
}

// 值接收者
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 指针接收者
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### 2.2 方法集规则

| 接收者类型 | 值变量 | 指针变量 |
|-----------|-------|---------|
| 值接收者 | ✓ | ✓ |
| 指针接收者 | ✗ | ✓ |

```go
r := Rectangle{Width: 10, Height: 5}
pr := &r

r.Area()    // ✓ 值接收者，值变量
pr.Area()   // ✓ 值接收者，指针变量（自动解引用）

r.Scale(2)  // ✓ 指针接收者，值变量（Go自动取址）
pr.Scale(2) // ✓ 指针接收者，指针变量
```

## 3. 接口（Interface）

### 3.1 接口定义

```go
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// 组合接口
type ReadWriter interface {
    Reader
    Writer
}
```

### 3.2 隐式实现

```go
type FileLogger struct{}

func (f *FileLogger) Write(p []byte) (int, error) {
    // 实现Write方法，自动实现Writer接口
    return len(p), nil
}

// 使用
func Log(w Writer, msg string) {
    w.Write([]byte(msg))
}

Log(&FileLogger{}, "hello")
```

### 3.3 空接口

```go
// interface{} 或 any（Go 1.18+）
func PrintAny(v any) {
    fmt.Println(v)
}

// 类型断言
func Process(v any) {
    if s, ok := v.(string); ok {
        fmt.Println("string:", s)
    }
    
    // type switch
    switch v := v.(type) {
    case string:
        fmt.Println("string:", v)
    case int:
        fmt.Println("int:", v)
    default:
        fmt.Println("unknown:", v)
    }
}
```

### 3.4 接口 nil 陷阱

```go
var w Writer          // nil接口
var p *FileLogger     // nil指针

w = p                 // w不是nil！（有类型信息）
fmt.Println(w == nil) // false!

// 正确检查
func isNil(w Writer) bool {
    if w == nil {
        return true
    }
    rv := reflect.ValueOf(w)
    return rv.Kind() == reflect.Ptr && rv.IsNil()
}
```

## 4. 组合与嵌入

### 4.1 组合模式

```go
type Engine interface {
    Start() error
    Stop() error
}

type Car struct {
    Engine          // 组合：Car has Engine
    Brand   string
}

func (c *Car) Start() error {
    fmt.Printf("%s starting...\n", c.Brand)
    return c.Engine.Start()  // 委托
}
```

### 4.2 嵌入 vs 继承

```go
type Animal struct {
    Name string
}

func (a *Animal) Speak() string {
    return "..."
}

type Dog struct {
    Animal  // 嵌入（组合，不是继承）
}

func (d *Dog) Speak() string {
    return "Woof!"
}

// 使用
dog := Dog{Animal{Name: "Max"}}
dog.Name       // "Max"（嵌入字段提升）
dog.Speak()    // "Woof!"（方法覆盖）
dog.Animal.Speak()  // "..."（访问嵌入类型的方法）
```

## 5. 类型断言与类型转换

### 5.1 类型断言

```go
var i interface{} = "hello"

s, ok := i.(string)  // 安全断言
if ok {
    fmt.Println(s)
}

s = i.(string)       // 不安全，失败会panic
```

### 5.2 类型转换

```go
type MyInt int

var a int = 10
var b MyInt = MyInt(a)  // 显式转换
```

## 6. 常用接口

### 6.1 Stringer

```go
type Stringer interface {
    String() string
}

type Point struct{ X, Y int }

func (p Point) String() string {
    return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
```

### 6.2 error

```go
type error interface {
    Error() string
}

type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}
```

### 6.3 io.Reader / io.Writer

```go
func Copy(dst io.Writer, src io.Reader) (int64, error) {
    return io.Copy(dst, src)
}
```

## 练习清单

1. 定义一个 `BankAccount` 结构体，实现存取款方法
2. 为结构体实现 `Stringer` 接口
3. 用组合模式实现 `Manager` 包含 `Employee`
4. 实现一个自定义 `error` 类型
5. 用 `type switch` 处理多种类型的函数参数