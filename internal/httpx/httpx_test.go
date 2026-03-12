package httpx_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gojuniper/internal/httpx"
)

// 这些测试演示 net/http 的单元测试写法：
// - httptest.NewRequest 构造请求
// - httptest.NewRecorder 捕获响应
// - 直接调用 handler 验证状态码、响应体与 JSON 编解码
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

func TestEchoHandler(t *testing.T) {
	in := httpx.EchoRequest{Message: "hello"}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewReader(b))
	rr := httptest.NewRecorder()

	httpx.EchoHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d, want %d", rr.Code, http.StatusOK)
	}

	var out httpx.EchoResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Message != "HELLO" {
		t.Fatalf("message=%q, want %q", out.Message, "HELLO")
	}
}
