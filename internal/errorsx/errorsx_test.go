package errorsx_test

import (
	"errors"
	"testing"

	"gojuniper/internal/errorsx"
)

func TestParsePositiveInt(t *testing.T) {
	// 子测试（t.Run）用于让不同场景的失败信息更清晰。
	t.Run("non-number", func(t *testing.T) {
		_, err := errorsx.ParsePositiveInt("x")
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("zero", func(t *testing.T) {
		_, err := errorsx.ParsePositiveInt("0")
		if !errors.Is(err, errorsx.ErrNotPositive) {
			t.Fatalf("expected ErrNotPositive, got %v", err)
		}
	})

	t.Run("positive", func(t *testing.T) {
		got, err := errorsx.ParsePositiveInt("42")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 42 {
			t.Fatalf("got=%d, want 42", got)
		}
	})
}

func TestJoin(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")

	if got := errorsx.Join(nil, nil); got != nil {
		t.Fatalf("expected nil")
	}
	if got := errorsx.Join(a, nil); !errors.Is(got, a) {
		t.Fatalf("expected a")
	}
	if got := errorsx.Join(nil, b); !errors.Is(got, b) {
		t.Fatalf("expected b")
	}
	if got := errorsx.Join(a, b); !errors.Is(got, a) || !errors.Is(got, b) {
		t.Fatalf("expected joined error to contain a and b, got %v", got)
	}
}
