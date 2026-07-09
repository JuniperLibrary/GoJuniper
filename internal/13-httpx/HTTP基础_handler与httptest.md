# Go 基础：HTTP（http.Handler / ServeMux / httptest）

这个文档对应 [internal/13-httpx](./)。

配合阅读：
- 实现代码：[httpx.go](./httpx.go)
- 测试代码：本目录下暂无 `*_test.go`，建议先读实现，后续补测试（见练习清单）
- 可运行服务：[cmd/server/main.go](../../cmd/server/main.go)

---

## 1. HTTP Handler 是什么

在 Go 标准库里，一个 HTTP handler 就是实现了这个接口的东西：

```go
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
```

你通常用两种方式写 handler：
- `http.HandlerFunc(func(w, r) { ... })`
- 自己写一个类型，实现 `ServeHTTP`

> ⚠️ **注意**：`http.Handler` 接口**只需要一个方法 `ServeHTTP(w, r)`**。任何普通函数 `func(http.ResponseWriter, *http.Request)` 都可以通过 `http.HandlerFunc(...)` **适配**成 `Handler`——`http.HandlerFunc` 本身是个类型，它为这个函数类型实现了 `ServeHTTP`。所以“函数”和“Handler”之间不需要手写样板，一个适配器就搞定。这约等于 Java Spring 里用 `@GetMapping` 注解把普通方法注册成 handler 的思路，但 Go 是同步签名 + 接口适配（Java 用反射/注解驱动）。

**为什么**：接口极小（单方法），意味着任何东西只要能“处理一个请求并写响应”就能挂到路由上，扩展性极强。

---

## 2. ServeMux：路由分发器（最基础版本）

`http.ServeMux` 可以把不同路径分发到不同 handler：

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", handler)
```

本仓库用它构造一个简单的服务路由：[httpx.NewMux](./httpx.go)

```go
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("POST /echo", EchoHandler)
	return mux
}
```

> ⚠️ **注意**：本仓库用的是 **`"GET /health"` 这种 `方法 + 空格 + 路径` 的写法**，这是 **Go 1.22+** 才支持的增强路由——既匹配 HTTP 方法，也支持路径模式（如 `/users/{id}` 提取路径参数）。旧版本 `http.ServeMux` 不区分方法、也不支持路径参数。如果你的 Go 版本低于 1.22，这种写法会不生效或行为不同。

**常见坑**：路由注册的方法（GET/POST）必须和客户端实际请求的方法一致，否则会命中 405。不要以为 `HandleFunc("/echo", ...)` 能自动接收所有方法。

---

## 3. JSON 请求与响应（最常见业务形态）

关键点：
- 读请求体：`json.NewDecoder(r.Body).Decode(&req)`
- 写响应：设置 `Content-Type`，再 `json.NewEncoder(w).Encode(res)`
- 错误处理：JSON 解析失败通常返回 `400 Bad Request`

本仓库 `EchoHandler` 的实现：

```go
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	res := EchoResponse{Message: strings.ToUpper(req.Message)}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "encode failed", http.StatusInternalServerError)
		return
	}
}
```

> ⚠️ **注意**：**必须在 `Encode` 之前设置 `Content-Type`**。如果先 `Encode` 再 `Set`，Go 已经根据内容推断并写入了 header，你的 `Set` 不会生效（header 一旦发送就改不了）。另外优先用 `json.NewEncoder(w).Encode(res)` 而非 `json.Marshal` + `w.Write`——前者直接流式写入 `ResponseWriter`，更顺手。这对应 Java 里 Jackson 的 `ObjectMapper.writeValue(writer, res)` 或 Spring 的 `@ResponseBody` 自动序列化分工。

> ⚠️ **注意**：`Decode` 失败要**尽早 `return`**。否则后面的逻辑会基于一个未初始化的 `req` 继续跑，可能 panic 或返回错误结果。`http.Error` 会自动设置状态码和 `Content-Type: text/plain`，适合快速返回错误响应。

**常见坑**：忘记检查 `Decode` 返回的 `err`，直接信任请求体，导致非法输入污染业务逻辑；或重复写 header（多次 `WriteHeader`）引发 “superfluous WriteHeader call” 日志。

---

## 4. httptest：不启动真实端口也能测 HTTP

`httptest` 的好处：
- 测试更快
- 不依赖端口
- 可重复、稳定

常见做法：
- 用 `httptest.NewRecorder()` 当 ResponseWriter
- 用 `httptest.NewRequest()` 构造请求
- 直接调用 mux/handler 的 `ServeHTTP`

```go
func TestEcho(t *testing.T) {
	mux := httpx.NewMux()
	body := strings.NewReader(`{"message":"hello"}`)
	req := httptest.NewRequest(http.MethodPost, "/echo", body)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	// 断言 resp.StatusCode / w.Body.String() ...
}
```

> ⚠️ **注意**：`httptest.NewRecorder()` 返回的 `*ResponseRecorder` **本身实现了 `http.ResponseWriter` 接口**，所以可以直接喂给 `mux.ServeHTTP(w, req)` 做“单元级”测试，完全不占用网络端口。对于整服务级测试，可用 `httptest.NewServer(mux)` 起一个真实监听的临时服务（自动分配端口），测完 `defer ts.Close()`。两个工具按需选用：单 handler 用 Recorder，要测完整中间件链路用 NewServer。

**为什么**：HTTP handler 本质就是 `func(ResponseWriter, *Request)`，所以测试只需“拼一个请求 + 一个能记录响应的 writer”即可，无需真的起 socket。这让测试快、稳定、易断言。

---

## 5. 练习清单

1. 增加一个 `GET /time` 返回当前时间（RFC3339），并写测试
2. 增加一个 `POST /sum`：请求体 `{ "a": 1, "b": 2 }` 返回 `{ "sum": 3 }`
3. 把错误响应统一成 JSON 格式（例如 `{ "error": "..." }`）

---

## 6. 自测命令

```bash
go test ./internal/13-httpx -shuffle on
```

> 当前目录暂未包含 `*_test.go`，运行上述命令时若提示 “no test files” 属正常。建议完成练习清单后补上测试（参考第 4 节的 `httptest` 写法），再执行该命令验证。
