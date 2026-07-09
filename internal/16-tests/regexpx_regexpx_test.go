package tests

import (
	"testing"

	"gojuniper/internal/15-regexpx"
)

func TestMatchDigits(t *testing.T) {
	got := regexpx.MatchDigits("我有 3 个苹果和 5 个香蕉")
	want := []string{"3", "5"}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want %q", i, got[i], want[i])
		}
	}
}

func TestMatchDigits_NoMatch(t *testing.T) {
	got := regexpx.MatchDigits("hello")
	if len(got) != 0 {
		t.Fatalf("got %d matches, want 0", len(got))
	}
}

func TestReplaceWhitespace(t *testing.T) {
	got := regexpx.ReplaceWhitespace("Hello    World")
	if got != "Hello World" {
		t.Fatalf("got=%q, want %q", got, "Hello World")
	}
}

func TestReplaceWhitespace_Multiple(t *testing.T) {
	got := regexpx.ReplaceWhitespace("a   b\t\tc\n\nd")
	if got != "a b c d" {
		t.Fatalf("got=%q, want %q", got, "a b c d")
	}
}

func TestSplitByComma(t *testing.T) {
	got := regexpx.SplitByComma("apple,banana,orange")
	want := []string{"apple", "banana", "orange"}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want %q", i, got[i], want[i])
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	if !regexpx.IsAlphanumeric("Hello123") {
		t.Fatal("expected true")
	}
	if regexpx.IsAlphanumeric("Hello 123") {
		t.Fatal("expected false (contains space)")
	}
	if regexpx.IsAlphanumeric("") {
		t.Fatal("expected false (empty string)")
	}
	if regexpx.IsAlphanumeric("hello!") {
		t.Fatal("expected false (contains punctuation)")
	}
}
