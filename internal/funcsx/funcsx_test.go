package funcsx_test

import (
	"testing"

	"gojuniper/internal/funcsx"
)

func TestCounter(t *testing.T) {
	t.Run("闭包独立计数", func(t *testing.T) {
		c := funcsx.Counter()
		if got := c(); got != 1 {
			t.Fatalf("c()=%d, want 1", got)
		}
		if got := c(); got != 2 {
			t.Fatalf("c()=%d, want 2", got)
		}
		if got := c(); got != 3 {
			t.Fatalf("c()=%d, want 3", got)
		}
	})

	t.Run("多个闭包独立状态", func(t *testing.T) {
		a := funcsx.Counter()
		b := funcsx.Counter()
		if got := a(); got != 1 {
			t.Fatalf("a()=%d, want 1", got)
		}
		if got := b(); got != 1 {
			t.Fatalf("b()=%d, want 1", got)
		}
		if got := a(); got != 2 {
			t.Fatalf("a()=%d, want 2", got)
		}
	})
}

func TestRecoverFromPanic(t *testing.T) {
	t.Run("无 panic 返回 nil", func(t *testing.T) {
		err := funcsx.RecoverFromPanic(func() {})
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

	t.Run("捕获 panic 返回错误", func(t *testing.T) {
		err := funcsx.RecoverFromPanic(func() {
			panic("something went wrong")
		})
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("变参：空列表", func(t *testing.T) {
		if got := funcsx.Sum(); got != 0 {
			t.Fatalf("Sum()=%d, want 0", got)
		}
	})

	t.Run("变参：整数列表", func(t *testing.T) {
		got := funcsx.Sum(34, 50, 25, 100, 65)
		if got != 274 {
			t.Fatalf("Sum(34,50,25,100,65)=%d, want 274", got)
		}
	})

	t.Run("变参：展开切片", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		if got := funcsx.Sum(nums...); got != 15 {
			t.Fatalf("Sum(nums...)=%d, want 15", got)
		}
	})
}

func TestDeferInspect(t *testing.T) {
	t.Run("defer 后进先出", func(t *testing.T) {
		got := funcsx.DeferInspect([]int{1, 2, 3})
		// defer 在函数返回前执行，且顺序为 LIFO
		// 期望顺序：body + defer(3) + defer(2) + defer(1)
		want := "bodydefer(3)defer(2)defer(1)"
		if got != want {
			t.Fatalf("got=%q, want %q", got, want)
		}
	})
}
