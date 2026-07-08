package collections

import (
	"reflect"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{nil, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 3, 1, 2, 1}, []int{3, 1, 2}},
	}
	for _, tt := range tests {
		if got := UniqueInts(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("UniqueInts(%v)=%v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestFrequency(t *testing.T) {
	got := Frequency([]string{"a", "b", "a", "c", "b", "a"})
	want := map[string]int{"a": 3, "b": 2, "c": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Frequency=%v, want %v", got, want)
	}
}

func TestMapKeysSorted(t *testing.T) {
	m := map[string]int{"banana": 1, "apple": 2, "cherry": 3}
	if got := MapKeysSorted(m); !reflect.DeepEqual(got, []string{"apple", "banana", "cherry"}) {
		t.Errorf("MapKeysSorted=%v, want sorted", got)
	}
	// generic value type should not matter
	m2 := map[string]bool{"z": true, "a": false}
	if got := MapKeysSorted(m2); !reflect.DeepEqual(got, []string{"a", "z"}) {
		t.Errorf("MapKeysSorted[bool]=%v, want [a z]", got)
	}
}

func TestMakeVsNew(t *testing.T) {
	makeInit, newPtr := MakeVsNew()
	if !makeInit {
		t.Error("make(map) should be non-nil (initialized)")
	}
	if !newPtr {
		t.Error("new(map) should return non-nil pointer")
	}
}

func TestSliceAppendDemo(t *testing.T) {
	got := SliceAppendDemo()
	want := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SliceAppendDemo()=%v, want %v", got, want)
	}
}

func TestSliceCopyDemo(t *testing.T) {
	src := []int{1, 2, 3}
	dst := SliceCopyDemo(src)
	if !reflect.DeepEqual(dst, src) {
		t.Errorf("SliceCopyDemo=%v, want %v", dst, src)
	}
	// copy is a deep copy; mutating dst shouldn't affect src
	dst[0] = 99
	if src[0] != 1 {
		t.Error("SliceCopyDemo should be an independent copy")
	}
}
