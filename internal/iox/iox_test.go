package iox_test

import (
	"bytes"
	"strings"
	"testing"

	"gojuniper/internal/iox"
)

// 这些测试覆盖 iox 的三个基础能力：
// - 从 io.Reader 读完整内容
// - 用 Scanner 按行读取
// - 用 Writer 按行写出并保证换行格式
func TestReadAllString(t *testing.T) {
	got, err := iox.ReadAllString(strings.NewReader("hello"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello" {
		t.Fatalf("got=%q, want %q", got, "hello")
	}
}

func TestReadLines(t *testing.T) {
	got, err := iox.ReadLines(strings.NewReader("a\nb\nc\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want %q", i, got[i], want[i])
		}
	}
}

func TestWriteLines(t *testing.T) {
	var buf bytes.Buffer
	if err := iox.WriteLines(&buf, []string{"a", "b", "c"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := strings.TrimRight(buf.String(), "\n"), "a\nb\nc"; got != want {
		t.Fatalf("got=%q, want %q", got, want)
	}
}
