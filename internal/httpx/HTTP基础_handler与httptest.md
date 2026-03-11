# Go 基础：HTTP（http.Handler / ServeMux / httptest）

这个文档对应 [internal/httpx](file:///e:/dingchuan/GoJuniper/internal/httpx)。

配合阅读：
- 实现代码：[httpx.go](file:///e:/dingchuan/GoJuniper/internal/httpx/httpx.go)
- 测试代码：[httpx_test.go](file:///e:/dingchuan/GoJuniper/internal/httpx/httpx_test.go)
- 可运行服务：[cmd/server/main.go](file:///e:/dingchuan/GoJuniper/cmd/server/main.go)

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

---

## 2. ServeMux：路由分发器（最基础版本）

`http.ServeMux` 可以把不同路径分发到不同 handler：

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", handler)
```

本仓库用它构造一个简单的服务路由：[httpx.NewMux](file:///e:/dingchuan/GoJuniper/internal/httpx/httpx.go)

---

## 3. JSON 请求与响应（最常见业务形态）

关键点：
- 读请求体：`json.NewDecoder(r.Body).Decode(&req)`
- 写响应：设置 `Content-Type`，再 `json.NewEncoder(w).Encode(res)`
- 错误处理：JSON 解析失败通常返回 `400 Bad Request`

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

本仓库的测试就是这样写的：[httpx_test.go](file:///e:/dingchuan/GoJuniper/internal/httpx/httpx_test.go)

---

## 5. 练习清单

1. 增加一个 `GET /time` 返回当前时间（RFC3339），并写测试
2. 增加一个 `POST /sum`：请求体 `{ "a": 1, "b": 2 }` 返回 `{ "sum": 3 }`
3. 把错误响应统一成 JSON 格式（例如 `{ "error": "..." }`）

---

## 6. 自测命令

```bash
go test ./internal/httpx -shuffle on
```

