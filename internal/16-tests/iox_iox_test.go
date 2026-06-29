package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gojuniper/internal/06-iox"
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

func TestReadFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(path, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}
	got, err := iox.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello" {
		t.Fatalf("got=%q, want %q", got, "hello")
	}
}

func TestReadFile_NotFound(t *testing.T) {
	_, err := iox.ReadFile("/nonexistent/file")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestWriteFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")
	if err := iox.WriteFile(path, "hello"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hello" {
		t.Fatalf("got=%q, want %q", string(b), "hello")
	}
}

func TestAppendToFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "log.txt")
	if err := iox.AppendToFile(path, "line1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := iox.AppendToFile(path, "line2"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(b), "line1line2"; got != want {
		t.Fatalf("got=%q, want %q", got, want)
	}
}

func TestCopyFile(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "dst.txt")
	if err := os.WriteFile(src, []byte("hello world"), 0644); err != nil {
		t.Fatal(err)
	}
	n, err := iox.CopyFile(src, dst)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 11 {
		t.Fatalf("copied %d bytes, want 11", n)
	}
	b, err := os.ReadFile(dst)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hello world" {
		t.Fatalf("got=%q, want %q", string(b), "hello world")
	}
}

func TestCopyFile_SrcNotFound(t *testing.T) {
	dir := t.TempDir()
	_, err := iox.CopyFile("/nonexistent/src", filepath.Join(dir, "dst.txt"))
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestFileExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "exists.txt")
	if iox.FileExists(path) {
		t.Fatal("should not exist yet")
	}
	if err := os.WriteFile(path, []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	if !iox.FileExists(path) {
		t.Fatal("should exist now")
	}
}

func TestListDir(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.txt"), nil, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "b.txt"), nil, 0644); err != nil {
		t.Fatal(err)
	}
	got, err := iox.ListDir(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d entries, want 2", len(got))
	}
}

func TestListDir_NotFound(t *testing.T) {
	_, err := iox.ListDir("/nonexistent/dir")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
