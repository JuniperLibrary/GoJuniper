// Package collections 提供一些集合相关的入门练习：
// - map 的使用（去重、计数）
// - 泛型（MapKeysSorted）
// - slice 的构造与容量预分配
package collections

import "sort"

// UniqueInts 返回去重后的切片，保持第一次出现时的顺序。
func UniqueInts(xs []int) []int {
	seen := make(map[int]struct{}, len(xs))
	out := make([]int, 0, len(xs))
	for _, v := range xs {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

// Frequency 统计每个字符串出现的次数。
func Frequency(xs []string) map[string]int {
	out := make(map[string]int, len(xs))
	for _, s := range xs {
		out[s]++
	}
	return out
}

// MapKeysSorted 返回 map 的 key（string）并按字典序排序。
// 这里用泛型让 value 类型不重要：只要 key 是 string 就行。
func MapKeysSorted[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// MakeVsNew 演示 make 与 new 的区别：
// make 返回初始化值（非 nil），new 返回指针。
func MakeVsNew() (makeIsInit bool, newIsPtr bool) {
	m := make(map[int]string)
	makeIsInit = m != nil

	p := new(map[int]string)
	newIsPtr = p != nil
	return
}

// SliceAppendDemo 演示 append 动态增长，返回最终切片。
func SliceAppendDemo() []int {
	var s []int
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	return s
}

// SliceCopyDemo 演示 copy 函数，将 src 复制到新切片并返回。
func SliceCopyDemo(src []int) []int {
	dst := make([]int, len(src), cap(src)*2)
	copy(dst, src)
	return dst
}
