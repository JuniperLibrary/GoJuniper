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
