---
name: go-stdlib
description: Go语言标准库学习 - fmt、strings、io、net/http、encoding/json等
license: MIT
---

# Go 标准库

## 1. fmt 格式化输入输出

### 1.1 Print系列

```go
fmt.Print("hello")           // 无换行
fmt.Println("hello")         // 换行
fmt.Printf("%s %d\n", "age", 25)  // 格式化

// 格式化到字符串
s := fmt.Sprintf("%s %d", "age", 25)

// 格式化到io.Writer
fmt.Fprintf(os.Stdout, "hello %s\n", "world")
```

### 1.2 常用格式动词

| 动词 | 说明 |
|------|------|
| `%v` | 默认格式 |
| `%+v` | 结构体显示字段名 |
| `%#v` | Go语法表示 |
| `%T` | 类型 |
| `%d` | 整数 |
| `%f` | 浮点 |
| `%s` | 字符串 |
| `%q` | 带引号字符串 |
| `%x` | 十六进制 |
| `%p` | 指针 |
| `%t` | 布尔 |

### 1.3 Scan系列

```go
var name string
var age int
fmt.Scan(&name, &age)         // 空格分隔
fmt.Scanf("%s %d", &name, &age)  // 格式化输入

// 从字符串扫描
fmt.Sscan("Alice 25", &name, &age)
```

## 2. strings 字符串处理

### 2.1 查找与判断

```go
strings.Contains("hello", "ell")     // true
strings.HasPrefix("hello", "he")     // true
strings.HasSuffix("hello", "lo")     // true
strings.Index("hello", "l")          // 2
strings.LastIndex("hello", "l")      // 3
strings.Count("hello", "l")          // 2
```

### 2.2 转换

```go
strings.ToUpper("hello")             // "HELLO"
strings.ToLower("HELLO")             // "hello"
strings.Title("hello world")         // "Hello World"
strings.TrimSpace("  hello  ")       // "hello"
strings.Trim("xxxhelloxxx", "x")     // "hello"
strings.TrimLeft("xxxhello", "x")    // "hello"
strings.TrimRight("helloxxx", "x")   // "hello"
```

### 2.3 分割与连接

```go
strings.Split("a,b,c", ",")          // []string{"a","b","c"}
strings.SplitN("a,b,c", ",", 2)      // []string{"a","b,c"}
strings.Join([]string{"a","b"}, "-") // "a-b"
strings.Replace("hello", "l", "L", 1)// "heLlo"
strings.ReplaceAll("hello", "l", "L")// "heLLo"
strings.Repeat("ab", 3)              // "ababab"
```

### 2.4 Builder（高效构建）

```go
var b strings.Builder
b.WriteString("hello")
b.WriteString(" ")
b.WriteString("world")
fmt.Println(b.String())  // "hello world"
```

## 3. strconv 类型转换

```go
// 字符串 <-> 整数
i, err := strconv.Atoi("42")         // string -> int
s := strconv.Itoa(42)                // int -> string

// 字符串 <-> 其他类型
i, err := strconv.ParseInt("42", 10, 64)   // string -> int64
f, err := strconv.ParseFloat("3.14", 64)  // string -> float64
b, err := strconv.ParseBool("true")       // string -> bool

// 类型 -> 字符串
s := strconv.FormatInt(42, 10)       // int64 -> string
s := strconv.FormatFloat(3.14, 'f', 2, 64)  // float64 -> string
s := strconv.FormatBool(true)        // bool -> string

// 引号处理
s := strconv.Quote(`hello "world"`)  // `"hello \"world\""`
```

## 4. io 输入输出

### 4.1 io.Reader / io.Writer

```go
// 读取
var buf [128]byte
n, err := reader.Read(buf[:])

// 写入
n, err := writer.Write([]byte("hello"))

// 从Reader读到Writer
io.Copy(dstWriter, srcReader)

// 读取全部
data, err := io.ReadAll(reader)

// 读取字符串（以\0结尾）
s, err := io.ReadString(reader, '\n')
```

### 4.2 io.Copy 系列

```go
io.Copy(dst, src)              // Reader -> Writer
io.CopyN(dst, src, n)          // 最多复制n字节
io.CopyBuffer(dst, src, buf)   // 使用指定缓冲区
```

### 4.3 Pipe

```go
r, w := io.Pipe()
go func() {
    w.Write([]byte("hello"))
    w.Close()
}()
io.ReadAll(r)  // "hello"
```

## 5. os 文件系统

### 5.1 文件操作

```go
// 打开/创建
f, err := os.Open("file.txt")           // 只读
f, err := os.Create("file.txt")         // 创建/截断
f, err := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND, 0644)

// 读写
data := make([]byte, 100)
n, err := f.Read(data)
n, err := f.Write([]byte("hello"))
n, err := f.WriteString("hello")

// 关闭
defer f.Close()

// 读取全部
data, err := os.ReadFile("file.txt")
// 写入文件
err := os.WriteFile("file.txt", []byte("hello"), 0644)
```

### 5.2 文件信息

```go
info, err := os.Stat("file.txt")
info.Name()       // 文件名
info.Size()       // 大小
info.ModTime()    // 修改时间
info.IsDir()      // 是否目录
info.Mode()       // 文件模式

// 判断文件存在
_, err := os.Stat("file.txt")
if os.IsNotExist(err) {
    // 文件不存在
}
```

### 5.3 目录操作

```go
// 创建目录
os.Mkdir("dir", 0755)
os.MkdirAll("dir/subdir", 0755)

// 删除
os.Remove("file.txt")
os.RemoveAll("dir")

// 列出目录
entries, _ := os.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name(), entry.IsDir())
}

// 临时目录/文件
os.TempDir()                        // 临时目录路径
f, _ := os.CreateTemp("", "prefix") // 临时文件
d, _ := os.MkdirTemp("", "prefix")  // 临时目录
```

### 5.4 环境变量与命令行

```go
// 环境变量
os.Getenv("PATH")
os.Setenv("KEY", "value")
os.Environ()  // 所有环境变量

// 命令行参数
os.Args  // []string{程序名, 参数...}

// 执行命令
cmd := exec.Command("ls", "-la")
output, _ := cmd.Output()
```

## 6. encoding/json

### 6.1 编码（序列化）

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "Alice", Age: 25}

// 到[]byte
data, err := json.Marshal(p)
data, err := json.MarshalIndent(p, "", "  ")  // 格式化

// 到io.Writer
json.NewEncoder(writer).Encode(p)

// 到字符串
s, _ := json.Marshal(p)
fmt.Println(string(s))
```

### 6.2 解码（反序列化）

```go
// 从[]byte
var p Person
err := json.Unmarshal(data, &p)

// 从io.Reader
var p Person
json.NewDecoder(reader).Decode(&p)

// 解析到map/interface
var result map[string]interface{}
json.Unmarshal(data, &result)
```

### 6.3 自定义编码

```go
func (p Person) MarshalJSON() ([]byte, error) {
    type Alias Person
    return json.Marshal(&struct {
        *Alias
        UpperName string `json:"upper_name"`
    }{
        Alias:     (*Alias)(&p),
        UpperName: strings.ToUpper(p.Name),
    })
}

func (p *Person) UnmarshalJSON(data []byte) error {
    type Alias Person
    aux := &struct {
        *Alias
    }{
        Alias: (*Alias)(p),
    }
    return json.Unmarshal(data, &aux)
}
```

## 7. net/http

### 7.1 HTTP客户端

```go
// GET
resp, err := http.Get("http://example.com")
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// POST
resp, err := http.Post("http://example.com", "application/json", bytes.NewReader(data))

// 自定义请求
req, _ := http.NewRequest("GET", "http://example.com", nil)
req.Header.Set("Authorization", "Bearer token")
resp, _ := http.DefaultClient.Do(req)

// 带超时
client := &http.Client{Timeout: 10 * time.Second}
resp, _ := client.Get("http://example.com")
```

### 7.2 HTTP服务器

```go
// 简单处理
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
})
http.ListenAndServe(":8080", nil)

// 使用路由器
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api", apiHandler)
http.ListenAndServe(":8080", mux)

// 优雅关闭
srv := &http.Server{Addr: ":8080", Handler: mux}
go srv.ListenAndServe()

// 优雅关闭
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

### 7.3 中间件

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// 使用
mux := http.NewServeMux()
mux.HandleFunc("/", handler)
http.ListenAndServe(":8080", loggingMiddleware(mux))
```

## 8. time 时间处理

### 8.1 时间创建

```go
now := time.Now()
specific := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
```

### 8.2 格式化与解析

```go
// 格式化（使用参考时间 2006-01-02 15:04:05）
s := now.Format("2006-01-02 15:04:05")
s := now.Format(time.RFC3339)

// 解析
t, _ := time.Parse("2006-01-02", "2024-01-01")
t, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")
```

### 8.3 时间运算

```go
now.Add(24 * time.Hour)           // 加1天
now.Add(-time.Hour)               // 减1小时
now.AddDate(1, 2, 3)              // 加1年2月3天

sub := t2.Sub(t1)                 // 时间差

now.Before(t)                     // 是否早于
now.After(t)                      // 是否晚于
now.Equal(t)                      // 是否相等
```

### 8.4 定时器

```go
// Timer
timer := time.NewTimer(5 * time.Second)
<-timer.C

// Ticker
ticker := time.NewTicker(1 * time.Second)
for t := range ticker.C {
    fmt.Println(t)
}

// After
select {
case <-time.After(5 * time.Second):
    fmt.Println("timeout")
}

// Sleep
time.Sleep(1 * time.Second)
```

## 练习清单

1. 用 `fmt.Printf` 格式化输出各种类型
2. 用 `strings` 包处理字符串（分割、替换、连接）
3. 用 `os` 包读写文件
4. 用 `encoding/json` 实现结构体序列化/反序列化
5. 用 `net/http` 写一个简单的HTTP服务器