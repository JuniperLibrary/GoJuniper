package basics_test

import (
	"errors"
	"fmt"
	"testing"

	"gojuniper/internal/basics"
)

// 这个测试演示“跨包全局变量”的推荐用法之一：导出的哨兵错误（sentinel error）。
// 重点不是错误文本，而是错误值本身；被包装后用 errors.Is 依然可识别。
func TestPackageLevelErrorsAreSentinels(t *testing.T) {
	if basics.ErrNegativeN == nil {
		t.Fatalf("ErrNegativeN must not be nil")
	}
	if basics.ErrOverflow == nil {
		t.Fatalf("ErrOverflow must not be nil")
	}

	wrapped := fmt.Errorf("wrap: %w", basics.ErrNegativeN)
	if !errors.Is(wrapped, basics.ErrNegativeN) {
		t.Fatalf("expected errors.Is(wrapped, ErrNegativeN)=true")
	}

	sameText := errors.New("n must be >= 0")
	if errors.Is(wrapped, sameText) {
		t.Fatalf("expected errors.Is(wrapped, errors.New(sameText))=false")
	}

	if basics.ErrNegativeN == sameText {
		t.Fatalf("expected ErrNegativeN != errors.New(sameText)")
	}
}
