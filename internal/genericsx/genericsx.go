// Package genericsx 提供泛型（Go 1.18+）相关练习：
// - Map/Filter/Reduce 这类常见“高阶函数”
// - 通过类型参数让函数对多种类型复用
package genericsx

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
