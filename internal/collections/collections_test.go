package collections_test

import (
	"testing"

	"gojuniper/internal/collections"
)

func TestUniqueInts(t *testing.T) {
	// 这个测试顺便验证“保持顺序”的约定：3/1/2/4 是首次出现的顺序。
	got := collections.UniqueInts([]int{3, 3, 1, 2, 1, 2, 4})
	want := []int{3, 1, 2, 4}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestFrequency(t *testing.T) {
	got := collections.Frequency([]string{"a", "b", "a", "c", "b", "a"})
	if got["a"] != 3 || got["b"] != 2 || got["c"] != 1 {
		t.Fatalf("unexpected frequency map: %#v", got)
	}
}

func TestMapKeysSorted(t *testing.T) {
	m := map[string]int{"b": 2, "a": 1, "c": 3}
	got := collections.MapKeysSorted(m)
	want := []string{"a", "b", "c"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want %q", i, got[i], want[i])
		}
	}
}
