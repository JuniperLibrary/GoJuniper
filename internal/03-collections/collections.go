// Package collections 提供一些集合相关的入门练习：
// - map 的使用（去重、计数）
// - 泛型（MapKeysSorted）
// - slice 的构造与容量预分配
package collections

import "sort"

// UniqueInts 返回去重后的切片，保持第一次出现时的顺序。
//
// ⚠️ 注意（map 做集合的惯用法）：
//  1. 用 map[int]struct{} 当 set——value 用空结构体 struct{}，它【不占内存】（0 字节），
//     比 map[int]bool 省空间。这是 Go 里"集合"的标准写法。
//  2. 判断是否存在用 `if _, ok := seen[v]; ok`，下划线丢弃 value，只看第二个 bool。
//  3. 用 make(map[int]struct{}, len(xs)) 预分配容量，减少 map 扩容 rehash。
//  4. 遍历顺序：Go 的 map 遍历是【随机顺序】，所以要用 out 切片按原顺序记录，才能得到稳定输出。
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
//
// ⚠️ 注意：map 的 value 是 int，循环里 `out[s]++` 对不存在的 key 会从零值 0 自增到 1，
// 不需要先判断 key 是否存在再初始化——这是 map 的零值特性带来的便利。
func Frequency(xs []string) map[string]int {
	out := make(map[string]int, len(xs))
	for _, s := range xs {
		out[s]++
	}
	return out
}

// MapKeysSorted 返回 map 的 key（string）并按字典序排序。
// 这里用泛型让 value 类型不重要：只要 key 是 string 就行。
//
// ⚠️ 注意（泛型要点）：
//  1. 函数签名 [V any] 声明了一个【类型参数 V】，调用时由编译器根据实参推断。
//  2. 因为只用到 key（string），value 类型 V 被约束为 any（任意类型），调用方可传 map[string]int / map[string]bool 等。
//  3. map 遍历顺序随机，所以必须 sort.Strings 才能得到确定有序的结果。
//  4. sort 包只内置了对少数类型的排序（Strings/Ints/Float64s），自定义类型需传 Less 函数。
func MapKeysSorted[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// MakeVsNew 演示 make 与 new 的区别：
//   - make 返回初始化值（非 nil）
//   - new 返回指针
//
// ⚠️ 注意：对 map 来说 make(map[int]string) 返回【已初始化、可直接写入】的 map；
// 而 new(map[int]string) 返回 *map，且指向的是 nil map（写入会 panic）。
// 所以 map 永远用 make，绝不用 new。返回值用命名返回 (makeIsInit, newIsPtr) 只为测试断言方便。
func MakeVsNew() (makeIsInit bool, newIsPtr bool) {
	m := make(map[int]string)
	makeIsInit = m != nil

	p := new(map[int]string)
	newIsPtr = p != nil
	return
}

// SliceAppendDemo 演示 append 动态增长，返回最终切片。
//
// ⚠️ 注意：append 可能在底层数组容量不够时【分配新数组并复制】，此时原切片和新切片底层数组分离。
// 所以 append 的结果【必须接收返回值】——s = append(s, x)，不能只写 append(s, x) 而丢弃返回值，否则可能"没生效"。
// 变参展开：append(s, 2, 3, 4) 一次追加多个元素。
func SliceAppendDemo() []int {
	var s []int
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	return s
}

// SliceCopyDemo 演示 copy 函数，将 src 复制到新切片并返回。
//
// ⚠️ 注意（slice 是引用语义，copy 才是真拷贝）：
//  1. slice 本身是个"描述符"（指针+长度+容量），直接赋值 dst := src 只是复制描述符，
//     两个变量【共享底层数组】——改一个会影响另一个。要用 copy() 做深拷贝。
//  2. copy(dst, src) 复制的元素个数 = min(len(dst), len(src))。这里 dst 用 len(src) 长度创建，能完整复制。
//  3. cap(dst)*2 只是把容量留大一点，和 copy 的元素数无关（copy 只看 len）。
func SliceCopyDemo(src []int) []int {
	dst := make([]int, len(src), cap(src)*2)
	copy(dst, src)
	return dst
}
