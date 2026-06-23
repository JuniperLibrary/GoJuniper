// Package funcsx 提供 Go 函数进阶主题练习：
// - 闭包（closure）与捕获
// - defer 延时执行
// - 变参函数
// - panic / recover
package funcsx

import "fmt"

// Counter 返回一个闭包，每次调用计数 +1。
// 对应 Rust 闭包 || { count += 1; count } 的模式。
func Counter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}

// RecoverFromPanic 执行 f，如果 f 触发了 panic，则 recover 并返回错误信息。
// 对应 Rust `catch_unwind` 的模式。
func RecoverFromPanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panicked: %v", r)
		}
	}()
	f()
	return nil
}

// Sum 计算变参列表中所有整数的和。变参（variadic）对应 Rust 中的 ... 语法。
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// DeferInspect 演示 defer 的"后进先出"执行顺序。
// 接收一个 []int，返回处理的日志字符串，用于验证执行顺序。
// 使用命名返回值，使 defer 能修改最终返回结果。
func DeferInspect(nums []int) (log string) {
	for _, n := range nums {
		defer func() {
			log += fmt.Sprintf("defer(%d)", n)
		}()
	}
	log += "body"
	return
}

func ApplyFunc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}
