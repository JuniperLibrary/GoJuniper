package httpx_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gojuniper/internal/httpx"
)

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
