// Package funcsx 提供 Go 函数进阶主题练习：
// - 闭包（closure）与捕获
// - defer 延时执行
// - 变参函数
// - panic / recover
package funcsx

import "fmt"

// Counter 返回一个闭包，每次调用计数 +1。
// 对应 Java 中用类字段保存状态（或 AtomicInteger）模拟闭包计数。
//
// ⚠️ 注意（闭包捕获的核心机制）：
// Counter 内部声明了局部变量 count，返回的匿名函数"捕获"了这个 count 的引用。
// 只要闭包还活着，count 就不会被回收，且【每次调用都操作同一个 count】。
// 关键点：调用两次 Counter() 会得到两个【独立的】闭包，各有各的 count，互不干扰。
// 对应 Java：Java 的 lambda 也能捕获外部 final/等效 final 局部变量，但只能读不能改（要改需包装成对象）；Go 闭包可直接读写捕获的变量。
func Counter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}

// RecoverFromPanic 执行 f，如果 f 触发了 panic，则 recover 并返回错误信息。
// 对应 Java 的 try-catch(Throwable) 兜底捕获（但 Go 的 recover 只能捕获同 goroutine 的 panic）。
//
// ⚠️ 注意（defer + recover 的写法要点）：
//  1. recover() 只能在【defer 函数】里调用才有效；放在别处返回 nil。
//  2. recover() 捕获的是 panic 传入的任意值（这里是字符串 "kaboom"），用 r 接收。
//  3. 这里用【命名返回值 err】，defer 里对它赋值，能在函数真正返回前"改写"返回结果——
//     这是 Go 里 recover 模式的标准套路。
//  4. panic/recover 不是 Go 的常规错误处理方式，只用于真正不可恢复/需要兜底防崩溃的场景（类似 Java 的未捕获异常兜底，不替代 try-catch 错误处理）。
func RecoverFromPanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panicked: %v", r)
		}
	}()
	f()
	return nil
}

// Sum 计算变参列表中所有整数的和。变参（variadic）对应 Java 的 `int... nums` 语法。
//
// ⚠️ 注意（变参规则）：
//  1. 变参 ...int 必须放在【参数列表最后一位】，不能写成 (nums ...int, sep string)。
//  2. 函数内部 nums 就是普通的 []int 切片。
//  3. 调用时可传多个值 Sum(1,2,3)，也可用切片展开 Sum(slice...)。
//  4. 一个函数最多只能有一个变参。
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
//
// ⚠️ 注意（defer 三个最容易错的点）：
//  1. 【执行顺序 LIFO】：循环里先注册 defer(1)、defer(2)、defer(3)，
//     执行时是反过来的——defer(3) → defer(2) → defer(1)，所以日志末尾是 "defer(3)defer(2)defer(1)"。
//  2. 【参数在注册时求值】：defer func() { log += ...(n) }() 里的 n 是【注册那一刻】的值（1,2,3），
//     不是循环结束时的值。如果写成 defer func(n int){...}(n) 或闭包捕获，结果取决于捕获时机——这里直接在 defer 调用里传 n，最安全。
//  3. 【能改命名返回值】：函数签名用了 (log string) 命名返回值，defer 里 append 到 log 会反映到最终返回。
//     若用普通返回值 `return log`，defer 的修改不会生效（因为 return 时已赋值给返回变量，但命名返回本身就是那个变量）。
func DeferInspect(nums []int) (log string) {
	for _, n := range nums {
		defer func() {
			log += fmt.Sprintf("defer(%d)", n)
		}()
	}
	log += "body"
	return
}

// ApplyFunc 把函数当作一等公民传参（高阶函数）。
//
// ⚠️ 注意：Go 的函数是"一等公民"，可以像 int 一样作为参数、返回值、存变量。
// 对应 Java 的 `Function<Integer, Integer>` 或方法引用 `BinaryOperator<Integer>` 作为参数。函数类型要写全签名：func(int, int) int。
func ApplyFunc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

// Factorial 用递归计算阶乘（n<=0 约定返回 1）。
//
// ⚠️ 注意：Go 没有"尾递归优化"，递归深度受栈大小限制（默认栈可增长到 GB 级，但深递归仍可能 stack overflow）。
// 大数阶乘应改用循环（见 basics.FactorialUint64）并做溢出检查。这里仅演示递归写法。
func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}
