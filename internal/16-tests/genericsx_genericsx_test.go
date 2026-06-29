package tests

import (
	"testing"

	"gojuniper/internal/14-genericsx"
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

func TestStack(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		s := genericsx.Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		v, ok := s.Pop()
		if !ok || v != 3 {
			t.Fatalf("Pop got %d, %v; want 3, true", v, ok)
		}
	})

	t.Run("Pop on empty stack", func(t *testing.T) {
		s := genericsx.Stack[string]{}
		_, ok := s.Pop()
		if ok {
			t.Fatal("expected ok=false on empty stack")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		s := genericsx.Stack[int]{}
		s.Push(42)
		v, ok := s.Peek()
		if !ok || v != 42 {
			t.Fatalf("Peek got %d, %v; want 42, true", v, ok)
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		var s genericsx.Stack[int]
		if !s.IsEmpty() {
			t.Fatal("expected empty")
		}
		s.Push(1)
		if s.IsEmpty() {
			t.Fatal("expected non-empty")
		}
	})
}

func TestSafeMap(t *testing.T) {
	m := genericsx.NewSafeMap[string, int]()
	m.Set("a", 1)
	m.Set("b", 2)

	v, ok := m.Get("a")
	if !ok || v != 1 {
		t.Fatalf("Get a got %d, %v; want 1, true", v, ok)
	}

	_, ok = m.Get("c")
	if ok {
		t.Fatal("expected ok=false for missing key")
	}

	keys := m.Keys()
	if len(keys) != 2 {
		t.Fatalf("Keys len=%d; want 2", len(keys))
	}

	m.Del("a")
	_, ok = m.Get("a")
	if ok {
		t.Fatal("expected ok=false after Del")
	}
}

func TestFindIndex(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	if got := genericsx.FindIndex(numbers, 3); got != 2 {
		t.Fatalf("got %d; want 2", got)
	}
	if got := genericsx.FindIndex(numbers, 99); got != -1 {
		t.Fatalf("got %d; want -1", got)
	}

	names := []string{"Alice", "Bob", "Charlie"}
	if got := genericsx.FindIndex(names, "Bob"); got != 1 {
		t.Fatalf("got %d; want 1", got)
	}
}

func TestAdd(t *testing.T) {
	if got := genericsx.Add(10, 20); got != 30 {
		t.Fatalf("got %d; want 30", got)
	}
	if got := genericsx.Add(3.14, 2.71); got != 5.85 {
		t.Fatalf("got %f; want 5.85", got)
	}
}
