package collections_test

import (
	"testing"

	"gojuniper/internal/collections"
)

// collections 包的测试聚焦于 slice/map 的常见题型：
// - 去重（保持首次出现顺序）
// - 计数（利用 map 读零值 + 自增）
// - map 的 key 排序（让遍历结果稳定可预期）
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

func TestMakeVsNew(t *testing.T) {
	makeIsInit, newIsPtr := collections.MakeVsNew()
	if !makeIsInit {
		t.Fatal("make 返回的 map 应该是已初始化的（非 nil）")
	}
	if !newIsPtr {
		t.Fatal("new 应该返回指针")
	}
}

func TestSliceAppendDemo(t *testing.T) {
	got := collections.SliceAppendDemo()
	want := []int{0, 1, 2, 3, 4}
	if len(got) != len(want) {
		t.Fatalf("len=%d; want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d; want %d", i, got[i], want[i])
		}
	}
}

func TestSliceCopyDemo(t *testing.T) {
	src := []int{1, 2, 3}
	dst := collections.SliceCopyDemo(src)
	if len(dst) < len(src) {
		t.Fatalf("dst len=%d; want >= %d", len(dst), len(src))
	}
	for i := range src {
		if dst[i] != src[i] {
			t.Fatalf("dst[%d]=%d; want %d", i, dst[i], src[i])
		}
	}
}
