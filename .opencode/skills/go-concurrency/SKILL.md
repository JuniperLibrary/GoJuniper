---
name: go-concurrency
description: Go语言并发学习 - goroutine、channel、select、sync包
license: MIT
---

# Go 并发编程

## 1. Goroutine

### 1.1 启动goroutine

```go
// 普通函数
go funcName()

// 匿名函数
go func() {
    fmt.Println("hello")
}()

// 带参数
go func(msg string) {
    fmt.Println(msg)
}("hello")
```

### 1.2 等待goroutine

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Printf("worker %d\n", id)
    }(i)
}

wg.Wait()
```

### 1.3 goroutine泄漏

```go
// 危险：无缓冲channel可能永久阻塞
func leak() {
    ch := make(chan int)
    go func() {
        ch <- 1  // 如果没人接收，永久阻塞
    }()
    // 函数返回，goroutine泄漏
}

// 安全：使用带缓冲channel或context
func safe() {
    ch := make(chan int, 1)  // 缓冲大小1
    go func() {
        ch <- 1
    }()
}
```

## 2. Channel

### 2.1 创建与基本操作

```go
// 无缓冲channel
ch := make(chan int)

// 有缓冲channel
ch := make(chan int, 10)

// 发送
ch <- 42

// 接收
v := <-ch
v, ok := <-ch  // ok=false表示channel已关闭

// 关闭
close(ch)
```

### 2.2 无缓冲 vs 有缓冲

```go
// 无缓冲：同步，发送方等待接收方
ch1 := make(chan int)
go func() { ch1 <- 1 }()  // 阻塞直到有人接收
v := <-ch1

// 有缓冲：异步，缓冲满才阻塞
ch2 := make(chan int, 2)
ch2 <- 1  // 不阻塞
ch2 <- 2  // 不阻塞
// ch2 <- 3  // 阻塞，缓冲已满
```

### 2.3 遍历channel

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)  // 必须关闭才能range遍历

for v := range ch {
    fmt.Println(v)
}
```

### 2.4 单向channel

```go
// 只发送
func send(ch chan<- int) {
    ch <- 1
}

// 只接收
func receive(ch <-chan int) {
    v := <-ch
    fmt.Println(v)
}

// 使用
ch := make(chan int)
go send(ch)
receive(ch)
```

### 2.5 channel常见模式

```go
// 生产者-消费者
func producer(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for v := range ch {
        fmt.Println(v)
    }
}

// 工作池
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}

// 使用
jobs := make(chan int, 100)
results := make(chan int, 100)

for i := 0; i < 3; i++ {
    go worker(i, jobs, results)
}

for j := 0; j < 10; j++ {
    jobs <- j
}
close(jobs)

for r := 0; r < 10; r++ {
    fmt.Println(<-results)
}
```

## 3. Select

### 3.1 基本用法

```go
select {
case v := <-ch1:
    fmt.Println("ch1:", v)
case ch2 <- 42:
    fmt.Println("sent to ch2")
case <-time.After(time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no activity")
}
```

### 3.2 超时模式

```go
select {
case v := <-ch:
    fmt.Println(v)
case <-time.After(5 * time.Second):
    fmt.Println("timeout")
}
```

### 3.3 非阻塞操作

```go
select {
case v := <-ch:
    fmt.Println(v)
default:
    fmt.Println("no value")
}
```

## 4. sync 包

### 4.1 Mutex（互斥锁）

```go
var (
    mu    sync.Mutex
    count int
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}
```

### 4.2 RWMutex（读写锁）

```go
var (
    rwmu sync.RWMutex
    data map[string]string
)

func read(key string) string {
    rwmu.RLock()         // 读锁
    defer rwmu.RUnlock()
    return data[key]
}

func write(key, value string) {
    rwmu.Lock()          // 写锁
    defer rwmu.Unlock()
    data[key] = value
}
```

### 4.3 WaitGroup

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // 工作...
    }(i)
}

wg.Wait()
```

### 4.4 Once

```go
var once sync.Once

func initSingleton() {
    once.Do(func() {
        // 只执行一次
        singleton = &Singleton{}
    })
}
```

### 4.5 Cond（条件变量）

```go
var (
    mu    sync.Mutex
    cond  = sync.NewCond(&mu)
    ready bool
)

// 等待
mu.Lock()
for !ready {
    cond.Wait()
}
mu.Unlock()

// 通知
mu.Lock()
ready = true
cond.Broadcast()  // 或 cond.Signal()
mu.Unlock()
```

### 4.6 Pool（对象池）

```go
var pool = sync.Pool{
    New: func() any {
        return &Object{}
    },
}

// 获取
obj := pool.Get().(*Object)

// 放回
pool.Put(obj)
```

### 4.7 Map（并发安全map）

```go
var m sync.Map

m.Store("key", "value")
v, ok := m.Load("key")
m.Delete("key")
m.Range(func(key, value any) bool {
    fmt.Println(key, value)
    return true
})
```

### 4.8 Atomic（原子操作）

```go
import "sync/atomic"

var counter int64

atomic.AddInt64(&counter, 1)
atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 0)
atomic.CompareAndSwapInt64(&counter, 0, 1)
```

## 5. Context

### 5.1 创建context

```go
// 背景
ctx := context.Background()

// TODO（不确定用哪个时）
ctx := context.TODO()

// 超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 截止时间
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
defer cancel()

// 取消
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// 携带值
ctx := context.WithValue(context.Background(), "key", "value")
```

### 5.2 使用context

```go
func worker(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()  // context.Canceled 或 context.DeadlineExceeded
        default:
            // 工作...
        }
    }
}

// HTTP请求取消
req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
client.Do(req)
```

## 6. 常见并发模式

### 6.1 Pipeline

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// 使用
nums := generate(1, 2, 3)
results := square(nums)
for r := range results {
    fmt.Println(r)
}
```

### 6.2 Fan-out / Fan-in

```go
// Fan-out: 多个goroutine读取同一channel
func fanOut(in <-chan int, n int) []<-chan int {
    var channels []<-chan int
    for i := 0; i < n; i++ {
        channels = append(channels, worker(in))
    }
    return channels
}

// Fan-in: 合并多个channel
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## 练习清单

1. 用goroutine+channel实现生产者-消费者
2. 实现一个工作池（worker pool）
3. 用select实现超时控制
4. 用sync.Map实现并发安全的计数器
5. 用context实现可取消的HTTP请求