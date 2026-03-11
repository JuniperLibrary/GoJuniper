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
