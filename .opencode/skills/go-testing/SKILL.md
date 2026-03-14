---
name: go-testing
description: Go语言测试学习 - 单元测试、基准测试、覆盖率、mock
license: MIT
---

# Go 测试

## 1. 单元测试基础

### 1.1 测试文件命名

```
example.go          -> example_test.go
calculator.go       -> calculator_test.go
```

### 1.2 测试函数

```go
package mypackage_test

import (
    "testing"
    "mypackage"
)

func TestAdd(t *testing.T) {
    got := mypackage.Add(1, 2)
    want := 3
    if got != want {
        t.Errorf("Add(1, 2) = %d, want %d", got, want)
    }
}
```

### 1.3 子测试

```go
func TestMax(t *testing.T) {
    t.Run("positive numbers", func(t *testing.T) {
        got := Max(1, 2)
        want := 2
        if got != want {
            t.Errorf("Max(1, 2) = %d, want %d", got, want)
        }
    })
    
    t.Run("negative numbers", func(t *testing.T) {
        got := Max(-1, -2)
        want := -1
        if got != want {
            t.Errorf("Max(-1, -2) = %d, want %d", got, want)
        }
    })
}
```

### 1.4 表驱动测试

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name  string
        a, b  int
        want  int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", 1, -1, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### 1.5 测试辅助函数

```go
// Helper标记
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // 标记为辅助函数，错误行号指向调用处
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestAdd(t *testing.T) {
    assertEqual(t, Add(1, 2), 3)
}
```

## 2. 测试工具

### 2.1 testing.T 方法

```go
t.Error(args...)       // 记录错误，继续执行
t.Errorf(format, ...)  // 格式化错误，继续执行
t.Fatal(args...)       // 记录错误，停止当前测试
t.Fatalf(format, ...)  // 格式化错误，停止当前测试
t.Skip(args...)        // 跳过当前测试
t.Skipf(format, ...)   // 格式化跳过
t.Parallel()           // 标记测试可并行
t.Log(args...)         // 日志（verbose模式显示）
t.Logf(format, ...)    // 格式化日志
t.Run(name, func)      // 子测试
t.TempDir()            // 创建临时目录（自动清理）
```

### 2.2 运行测试

```bash
# 运行所有测试
go test ./...

# 运行指定包
go test ./mypackage

# 运行指定测试
go test -run TestAdd ./...
go test -run TestMax ./mypackage

# 详细输出
go test -v ./...

# 随机顺序
go test -shuffle on ./...

# 并行
go test -parallel 4 ./...

# 超时
go test -timeout 30s ./...
```

## 3. 基准测试

### 3.1 基本基准测试

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

### 3.2 运行基准测试

```bash
go test -bench=. ./...
go test -bench=Add ./...
go test -bench=. -benchmem ./...  # 显示内存分配
```

### 3.3 基准测试示例

```go
func BenchmarkStringConcat(b *testing.B) {
    b.ResetTimer()  // 重置计时器（排除初始化）
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "x"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        for j := 0; j < 100; j++ {
            builder.WriteString("x")
        }
        _ = builder.String()
    }
}

// 子基准测试
func BenchmarkAdd(b *testing.B) {
    b.Run("small", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Add(1, 2)
        }
    })
    
    b.Run("large", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Add(1000000, 2000000)
        }
    })
}
```

## 4. 测试覆盖率

### 4.1 生成覆盖率报告

```bash
# 基本覆盖率
go test -cover ./...

# 详细覆盖率
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# HTML报告
go tool cover -html=coverage.out

# 只覆盖特定函数
go test -coverprofile=coverage.out -run TestAdd ./...
```

### 4.2 覆盖率模式

```bash
# set: 是否覆盖
go test -covermode=set -coverprofile=coverage.out ./...

# count: 覆盖次数
go test -covermode=count -coverprofile=coverage.out ./...

# atomic: 原子计数（并发安全）
go test -covermode=atomic -coverprofile=coverage.out ./...
```

## 5. Example 测试

### 5.1 示例函数

```go
func ExampleAdd() {
    fmt.Println(Add(1, 2))
    // Output: 3
}

func ExampleAdd_negative() {
    fmt.Println(Add(-1, -2))
    // Output: -3
}
```

### 5.2 运行示例测试

```bash
go test -run Example ./...
```

### 5.3 无输出验证

```go
func ExampleLogger() {
    logger := NewLogger()
    logger.Log("hello")
    // Output:
}
```

## 6. Fuzz 测试（Go 1.18+）

### 6.1 Fuzz 函数

```go
func FuzzAdd(f *testing.F) {
    // 添加种子语料
    f.Add(1, 2)
    f.Add(-1, 1)
    
    f.Fuzz(func(t *testing.T, a, b int) {
        result := Add(a, b)
        // 验证属性
        if result-a != b {
            t.Errorf("Add(%d, %d) = %d, property check failed", a, b, result)
        }
    })
}
```

### 6.2 运行 Fuzz 测试

```bash
# 运行fuzz测试
go test -fuzz=FuzzAdd

# 指定时间
go test -fuzz=FuzzAdd -fuzztime=30s

# 最小化失败用例
go test -fuzz=FuzzAdd -fuzzminimizetime=10s
```

## 7. Mock 和 Stub

### 7.1 接口 Mock

```go
// 接口
type UserRepository interface {
    Find(id int) (*User, error)
    Save(user *User) error
}

// Mock实现
type MockUserRepository struct {
    users map[int]*User
    err   error
}

func (m *MockUserRepository) Find(id int) (*User, error) {
    if m.err != nil {
        return nil, m.err
    }
    return m.users[id], nil
}

func (m *MockUserRepository) Save(user *User) error {
    if m.err != nil {
        return m.err
    }
    m.users[user.ID] = user
    return nil
}

// 测试
func TestUserService(t *testing.T) {
    mock := &MockUserRepository{
        users: map[int]*User{1: {ID: 1, Name: "Alice"}},
    }
    
    service := NewUserService(mock)
    user, err := service.GetUser(1)
    
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("got %s, want Alice", user.Name)
    }
}
```

### 7.2 httptest

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users/1", nil)
    w := httptest.NewRecorder()
    
    handler := UserHandler{}
    handler.ServeHTTP(w, req)
    
    resp := w.Result()
    body, _ := io.ReadAll(resp.Body)
    
    if resp.StatusCode != http.StatusOK {
        t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
    }
}
```

### 7.3 测试数据库

```go
// 使用SQLite内存数据库
func setupTestDB(t *testing.T) *sql.DB {
    t.Helper()
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatal(err)
    }
    
    // 创建表
    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT
        )
    `)
    if err != nil {
        t.Fatal(err)
    }
    
    t.Cleanup(func() {
        db.Close()
    })
    
    return db
}

func TestUserRepository(t *testing.T) {
    db := setupTestDB(t)
    repo := NewUserRepository(db)
    
    // 测试...
}
```

## 8. 测试最佳实践

### 8.1 测试结构

```
mypackage/
├── mypackage.go
├── mypackage_test.go       # 外部测试
├── internal_test.go        # 内部测试（package mypackage）
├── testdata/               # 测试数据
│   ├── input.json
│   └── golden.txt
└── benchmark_test.go       # 基准测试
```

### 8.2 Golden 文件测试

```go
func TestGolden(t *testing.T) {
    input := "test input"
    output := process(input)
    
    golden := filepath.Join("testdata", "golden.txt")
    if *update {
        os.WriteFile(golden, []byte(output), 0644)
    }
    
    expected, _ := os.ReadFile(golden)
    if output != string(expected) {
        t.Errorf("output mismatch")
    }
}
```

### 8.3 测试清理

```go
func TestWithCleanup(t *testing.T) {
    // 创建临时文件
    f, err := os.CreateTemp("", "test")
    if err != nil {
        t.Fatal(err)
    }
    
    // 注册清理函数
    t.Cleanup(func() {
        os.Remove(f.Name())
    })
    
    // 测试...
}
```

### 8.4 并行测试

```go
func TestParallel(t *testing.T) {
    t.Parallel()  // 标记可并行
    
    // 测试代码...
}

func TestParallelSubtests(t *testing.T) {
    tests := []struct{ name string }{
        {"test1"},
        {"test2"},
    }
    
    for _, tt := range tests {
        tt := tt  // 捕获变量
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            // 测试代码...
        })
    }
}
```

## 练习清单

1. 为一个计算器模块写单元测试
2. 使用表驱动测试重构测试代码
3. 写一个基准测试比较 `string` 拼接和 `strings.Builder`
4. 生成测试覆盖率报告并查看HTML
5. 使用 `httptest` 测试HTTP处理器