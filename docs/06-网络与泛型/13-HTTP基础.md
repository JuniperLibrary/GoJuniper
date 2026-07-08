# 13 — HTTP 基础

## 是什么

Go 标准库 `net/http` 提供了完整的 HTTP 客户端和服务端实现，无需第三方依赖即可构建 Web 服务。核心概念是 `http.Handler` 接口（只有一个 `ServeHTTP` 方法）和 `http.ServeMux`（路由复用器）。

## 怎么用

### 定义 handler

handler 是实现了 `http.Handler` 接口的任何值，但最常用的是函数式 handler：

```go
// net/http 要求 handler 签名固定：func(w http.ResponseWriter, r *http.Request)
func HealthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
}
```

`http.ResponseWriter` 是一个接口，写入时依次：设置 Header → 调用 WriteHeader（状态码）→ Write（响应体）。WriteHeader 调用后不能再修改 Header。

> ⚠️ **注意**：**`http.Handler` 是一个接口**，只有一个方法 `ServeHTTP(w, r)`。普通函数本身不是 Handler，但 `http.HandlerFunc` 是适配器类型，能把 `func(w, r)` 转成 Handler，所以 `mux.HandleFunc(...)` 才能直接注册函数。理解"函数 → HandlerFunc → Handler"这层适配，是读懂 Go HTTP 代码的关键。

### 路由注册（Go 1.22+ 支持方法+路径模式）

```go
func NewMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /health", HealthHandler)
    mux.HandleFunc("POST /echo", EchoHandler)
    return mux
}
```

Go 1.22 之前只能用 `/health` 这种路径模式，不支持限定 HTTP 方法。1.22 之后 `"GET /health"` 写法自动匹配 GET 请求。

### JSON 请求/响应

```go
// 读请求体 JSON
var req EchoRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "invalid json", http.StatusBadRequest)
    return
}

// 写 JSON 响应
w.Header().Set("Content-Type", "application/json; charset=utf-8")
if err := json.NewEncoder(w).Encode(res); err != nil {
    http.Error(w, "encode failed", http.StatusInternalServerError)
    return
}
```

标准库 `encoding/json` 配合 `json.NewDecoder`/`json.NewEncoder` 直接对流进行读写，无需中间字节缓冲区。

### 测试：httptest

```go
func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rr := httptest.NewRecorder()

    httpx.HealthHandler(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("status=%d, want %d", rr.Code, http.StatusOK)
    }
    if body := rr.Body.String(); body != "ok" {
        t.Fatalf("body=%q, want %q", body, "ok")
    }
}
```

`httptest.NewRequest` 构造请求，`httptest.NewRecorder` 模拟 `ResponseWriter`。不需要启动真实端口即可完成 handler 单元测试。

> ⚠️ **注意**：**`ServeMux` 是路由分发中心**，所有路由都注册在它上面；handler 本身不持有路由。`http.ListenAndServe(":8080", mux)` 的第二个参数就是 mux。忘记把 mux 传给 Server 会导致所有请求落到默认 404 handler。

## 为什么

Go 把 HTTP 放在标准库里，不像 Java 需要选 Spring Boot / Javalin / Spark（标准库只有底层 `HttpURLConnection`/`HttpClient`，Web 框架需第三方）。这体现了 Go "电池内置"的理念：标准库足够完成 80% 的工作。

对比 Java 的 Spring Boot：
- Spring 的 handler 通过注解（`@GetMapping`/`@RequestParam`）和参数绑定实现类型安全的解析；Go 直接操作 `http.Request`，手动解析。
- Spring 的路由通过注解/配置在应用启动时注册；Go 的 `ServeMux` 是运行时注册。
- Spring（Web MVC）使用每请求一线程（或虚拟线程）模型；Go 的 `http.Server` 内部自带 goroutine-per-connection 模型。
- httptest 类似于 axum 的 `TestApp`（或 reqwest 的 mock server），但 Go 的 httptest 是标准库自带，开箱即用。

## 常见坑

> ⚠️ **注意**：**生产环境必须给 `http.Server` 设置超时**（`ReadTimeout` / `WriteTimeout` / `IdleTimeout`）。标准库默认**无超时**，慢客户端或恶意连接会一直占用连接直到资源耗尽。本仓库 `cmd/server` 才有完整超时配置，学习用的 `NewMux` 仅做路由演示。

1. **ResponseWriter 的操作顺序**：必须先设置 Header，再 WriteHeader，最后 Write。先 Write 后设置 Header 会不生效，因为 Write 会隐式调用 WriteHeader(200)。

2. **没有默认超时**：生产环境必须显式设置 `http.Server` 的 `ReadTimeout`、`WriteTimeout`、`IdleTimeout`，否则慢客户端可能耗尽连接。

3. **`r.Body` 必须关闭**：但 Go 的文档说标准库会自动关闭，手动 close 也不会错，双重 close 不会 panic。

4. **ServeMux 的模式匹配**：Go 1.22 之前 `/path` 会匹配所有 `/path/...` 子路径。Go 1.22 引入的 `"GET /path"` 模式更精确，但需注意版本兼容性。

5. **handler 签名固定**：必须是 `func(http.ResponseWriter, *http.Request)`，签名不符无法注册到 `HandleFunc`。

---

**参考代码**：`internal/13-httpx/httpx.go` · `internal/13-httpx/httpx_test.go`
**运行测试**：`go test ./internal/13-httpx/`
