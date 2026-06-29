package tests

import (
	"testing"

	"gojuniper/internal/02-funcsx"
)

func TestFuncsx_Counter(t *testing.T) {
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

func TestFuncsx_Sum(t *testing.T) {
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

func TestApplyFunc(t *testing.T) {
	add := func(a, b int) int { return a + b }
	got := funcsx.ApplyFunc(3, 4, add)
	if got != 7 {
		t.Fatalf("ApplyFunc(3,4,add)=%d, want 7", got)
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: -1, want: 1},
		{n: 0, want: 1},
		{n: 1, want: 1},
		{n: 5, want: 120},
	}
	for _, tt := range tests {
		if got := funcsx.Factorial(tt.n); got != tt.want {
			t.Fatalf("Factorial(%d)=%d, want %d", tt.n, got, tt.want)
		}
	}
}
