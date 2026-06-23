package genericsx_test

import (
	"testing"

	"gojuniper/internal/genericsx"
)

// 这几个测试分别验证 Map/Filter/Reduce 三种最常见的泛型“高阶函数”模式。
func TestMap(t *testing.T) {
	got := genericsx.Map([]int{1, 2, 3}, func(v int) string {
		switch v {
		case 1:
			return "one"
		case 2:
			return "two"
		default:
			return "other"
		}
	})
	want := []string{"one", "two", "other"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want %q", i, got[i], want[i])
		}
	}
}

func TestFilter(t *testing.T) {
	got := genericsx.Filter([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
		return v%2 == 0
	})
	want := []int{2, 4, 6}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestReduce(t *testing.T) {
	got := genericsx.Reduce([]int{1, 2, 3, 4}, 0, func(acc int, v int) int {
		return acc + v
	})
	if got != 10 {
		t.Fatalf("got=%d, want 10", got)
	}
}

// TestGetLargest 对应 Rust 示例中 get_largest 的测试模式：
// - 用整数列表测试
// - 用 rune（字符）列表测试
func TestGetLargest(t *testing.T) {
	t.Run("empty slice 返回 false", func(t *testing.T) {
		_, ok := genericsx.GetLargest([]int{})
		if ok {
			t.Fatal("expected ok=false")
		}
	})

	t.Run("整数列表", func(t *testing.T) {
		numbers := []int{34, 50, 25, 100, 65}
		got, ok := genericsx.GetLargest(numbers)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if got != 100 {
			t.Fatalf("got=%d, want 100", got)
		}
	})

	t.Run("rune（字符）列表", func(t *testing.T) {
		chars := []rune{'y', 'm', 'a', 'q'}
		got, ok := genericsx.GetLargest(chars)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if got != 'y' {
			t.Fatalf("got=%c (%d), want 'y' (121)", got, got)
		}
	})

	t.Run("浮点数列表", func(t *testing.T) {
		floats := []float64{3.14, 2.71, 1.618, 0.577}
		got, ok := genericsx.GetLargest(floats)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if got != 3.14 {
			t.Fatalf("got=%f, want 3.14", got)
		}
	})
}
