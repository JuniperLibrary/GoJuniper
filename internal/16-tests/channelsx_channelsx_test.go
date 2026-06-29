package tests

import (
	"context"
	"reflect"
	"sort"
	"testing"
	"time"

	"gojuniper/internal/11-channelsx"
)

// 这个测试把 channels 的几个经典模式串起来练：
// - generator 生成数据
// - pipeline 做变换（平方）
// - fan-in 合并多个输入
func TestPipeline_GenerateSquareMerge(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	a := channelsx.Square(ctx, channelsx.Generate(ctx, 5)) // 0..4 -> square
	b := channelsx.Square(ctx, channelsx.Generate(ctx, 5)) // 0..4 -> square
	out := channelsx.Merge(ctx, a, b)

	var got []int
	for v := range out {
		got = append(got, v)
	}

	if len(got) != 10 {
		t.Fatalf("len=%d, want 10", len(got))
	}

	sort.Ints(got)
	want := []int{0, 0, 1, 1, 4, 4, 9, 9, 16, 16}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestFibonacciSelect(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	got, err := channelsx.FibonacciSelect(ctx, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v, want=%v", got, want)
	}
}

func TestFibonacciChannel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	got := make([]int, 0, 10)
	for v := range channelsx.FibonacciChannel(ctx, 10) {
		got = append(got, v)
	}
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v, want=%v", got, want)
	}
}
