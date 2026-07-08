// Package genericsx 提供泛型（Go 1.18+）相关练习：
// - Map/Filter/Reduce 这类常见"高阶函数"
// - GetLargest 泛型约束示例（对应 Java getLargest<T extends Comparable<T>>）
// - 通过类型参数让函数对多种类型复用
package genericsx

import (
	"cmp"
	"sync"
)

// Map 把输入切片 xs 映射成另一个类型的切片。
func Map[A any, B any](xs []A, f func(A) B) []B {
	out := make([]B, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

// Filter 返回满足谓词 pred 的所有元素。
func Filter[T any](xs []T, pred func(T) bool) []T {
	out := make([]T, 0, len(xs))
	for _, x := range xs {
		if pred(x) {
			out = append(out, x)
		}
	}
	return out
}

// Reduce 将 xs 归并为一个值：
// - acc 是初始值
// - f 将 acc 与当前元素合并，得到新的 acc
func Reduce[T any, Acc any](xs []T, acc Acc, f func(Acc, T) Acc) Acc {
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// GetLargest 返回切片中的最大值（对应 Java getLargest<T extends Comparable<T>>）。
// - 空切片返回（零值, false）
// - 要求 T 实现 cmp.Ordered（int、string、float64、rune 等可比较类型）
func GetLargest[T cmp.Ordered](xs []T) (T, bool) {
	if len(xs) == 0 {
		var zero T
		return zero, false
	}
	largest := xs[0]
	for _, v := range xs[1:] {
		if v > largest {
			largest = v
		}
	}
	return largest, true
}

// Stack 泛型 Stack 实现
type Stack[T any] struct {
	elements []T
}

// Push 入栈
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop 出栈
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, true
}

// Peek 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// SafeMap 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// NewSafeMap 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set 设置键值对
func (s *SafeMap[K, V]) Set(k K, v V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[k] = v
}

// Get 获取
func (s *SafeMap[K, V]) Get(k K) (V, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.data[k]
	return v, ok
}

// Del 删除键
func (s *SafeMap[K, V]) Del(k K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, k)
}

// Keys 返回所有键
func (s *SafeMap[K, V]) Keys() []K {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	keys := make([]K, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

// FindIndex 在切片中查找目标元素的索引，未找到返回 -1
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// Number 数字类型联合约束
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Add 泛型数字加法
func Add[T Number](a, b T) T {
	return a + b
}
