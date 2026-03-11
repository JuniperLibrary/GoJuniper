// Package httpx 提供 net/http 主题的基础练习：
// - Handler / ServeMux
// - JSON 请求/响应
// - httptest 的配套测试
package httpx

import (
	"encoding/json"
	"net/http"
	"strings"
)

// EchoRequest 是 /echo 的请求体。
type EchoRequest struct {
	Message string `json:"message"`
}

// EchoResponse 是 /echo 的响应体。
type EchoResponse struct {
	Message string `json:"message"`
}

// NewMux 返回一个包含基础路由的 http.Handler。
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("POST /echo", EchoHandler)
	return mux
}

// HealthHandler 用于健康检查。
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// EchoHandler 读取 JSON 请求，做一个简单变换后返回 JSON。
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
