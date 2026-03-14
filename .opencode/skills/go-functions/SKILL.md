---
name: go-functions
description: Go语言函数学习 - 函数、方法、闭包、defer、panic/recover
license: MIT
---

# Go 函数

## 1. 函数基础

### 1.1 函数声明

```go
func add(a, b int) int {
    return a + b
}

// 多返回值
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 命名返回值
func rectangle(width, height int) (area, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // 裸返回
}
```

### 1.2 可变参数

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// 调用
sum(1, 2, 3)           // 6
nums := []int{1, 2, 3}
sum(nums...)           // 展开slice
```

### 1.3 函数作为值

```go
// 函数类型
type Calculator func(int, int) int

func compute(a, b int, op Calculator) int {
    return op(a, b)
}

// 使用
add := func(a, b int) int { return a + b }
result := compute(1, 2, add)  // 3
```

## 2. 闭包（Closure）

### 2.1 基本闭包

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
c()  // 1
c()  // 2
c()  // 3
```

### 2.2 闭包陷阱

```go
// 错误示例：循环变量捕获问题
funcs := []func(){}
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i)  // 都会打印3
    })
}

// 正确做法：传参捕获
for i := 0; i < 3; i++ {
    funcs = append(funcs, func(i int) func() {
        return func() {
            fmt.Println(i)
        }
    }(i))
}
```

## 3. defer

### 3.1 基本用法

```go
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()  // 函数返回前执行
    
    // 使用文件...
    return nil
}
```

### 3.2 defer 执行顺序

```go
func example() {
    defer fmt.Println("1")  // 最后执行
    defer fmt.Println("2")  // 倒数第二
    defer fmt.Println("3")  // 最先执行
}
// 输出: 3, 2, 1 (LIFO栈顺序)
```

### 3.3 defer 参数预计算

```go
func example() {
    i := 0
    defer fmt.Println(i)  // 参数在defer时就确定了
    i++
}
// 输出: 0 (不是1)
```

### 3.4 defer 性能

```go
// 高性能场景避免defer
func slow() {
    mu.Lock()
    defer mu.Unlock()  // 有开销
}

func fast() {
    mu.Lock()
    // ... 代码
    mu.Unlock()  // 直接调用更快
}
```

## 4. panic 与 recover

### 4.1 panic

```go
func mustPositive(n int) {
    if n < 0 {
        panic("n must be positive")
    }
}
```

### 4.2 recover

```go
func safeOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()
    
    // 可能panic的代码
    riskyOperation()
    return nil
}
```

### 4.3 使用场景

- `panic`: 不可恢复的错误（如配置缺失、初始化失败）
- `recover`: HTTP服务器、goroutine等需要防止崩溃的场景
- **不要**用panic做普通错误处理

```go
// HTTP服务器recover示例
func handler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            log.Printf("panic: %v", err)
            http.Error(w, "Internal Server Error", 500)
        }
    }()
    // 处理请求...
}
```

## 5. 方法

### 5.1 方法定义

```go
type Person struct {
    Name string
    Age  int
}

// 值接收者
func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// 指针接收者（可修改）
func (p *Person) Birthday() {
    p.Age++
}
```

### 5.2 值接收者 vs 指针接收者

**使用指针接收者：**
- 需要修改接收者
- 结构体较大（避免拷贝）
- 一致性（如果一个方法用指针，其他也用指针）

**使用值接收者：**
- 不需要修改
- 小结构体
- 值类型语义（如time.Time）

## 6. init 函数

```go
var config Config

func init() {
    // 包初始化时自动执行
    config = loadConfig()
}
```

- 每个包可以有多个 `init`
- 执行顺序：依赖包 → 包级变量 → init → main
- 避免过度使用，推荐显式初始化

## 练习清单

1. 实现一个 `max` 函数，支持多返回值
2. 用闭包实现斐波那契生成器
3. 写一个带 `defer` 的文件处理函数
4. 实现 `panic/recover` 包装器
5. 为自定义类型添加方法，理解值/指针接收者