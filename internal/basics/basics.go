// Package basics 提供一些 Go 入门阶段常见的“可测函数”，用来练习：
// - 变量/循环/switch
// - 切片与 rune（Unicode）
// - 错误处理与边界条件
// - 溢出检查（uint64）
package basics

import (
	"errors"
	"math"
	"strings"
)

var (
	// ErrNegativeN 表示参数 n 为负数（但该函数要求 n >= 0）。
	ErrNegativeN = errors.New("n must be >= 0")
	// ErrOverflow 表示计算结果超过 uint64 的可表示范围。
	ErrOverflow = errors.New("result overflows uint64")
)

// Sum 返回两个整数之和。
func Sum(a, b int) int {
	return a + b
}

// Max 返回切片的最大值；如果切片为空则返回 (0, false)。
func Max(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	m := xs[0]
	for _, v := range xs[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}

// FizzBuzz 返回从 1..n 的 FizzBuzz 序列。
// - 3 的倍数输出 "Fizz"
// - 5 的倍数输出 "Buzz"
// - 15 的倍数输出 "FizzBuzz"
// n <= 0 时返回 nil（这让调用方更容易用 len() 判断是否有结果）。
func FizzBuzz(n int) []string {
	if n <= 0 {
		return nil
	}
	out := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, itoa(i))
		}
	}
	return out
}

// IsPrime 判断 n 是否为质数。
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// FactorialUint64 计算 n 的阶乘。
// - n < 0 返回 ErrNegativeN
// - 结果溢出返回 ErrOverflow
func FactorialUint64(n int) (uint64, error) {
	if n < 0 {
		return 0, ErrNegativeN
	}
	var res uint64 = 1
	for i := 2; i <= n; i++ {
		m := uint64(i)
		if res > math.MaxUint64/m {
			return 0, ErrOverflow
		}
		res *= m
	}
	return res, nil
}

// FibonacciUint64 返回长度为 n 的斐波那契序列（从 0 开始）。
// - n < 0 返回 ErrNegativeN
// - 计算过程发生溢出返回 ErrOverflow
func FibonacciUint64(n int) ([]uint64, error) {
	if n < 0 {
		return nil, ErrNegativeN
	}
	if n == 0 {
		return []uint64{}, nil
	}
	out := make([]uint64, n)
	out[0] = 0
	if n == 1 {
		return out, nil
	}
	out[1] = 1
	for i := 2; i < n; i++ {
		a := out[i-1]
		b := out[i-2]
		if a > math.MaxUint64-b {
			return nil, ErrOverflow
		}
		out[i] = a + b
	}
	return out, nil
}

// ReverseString 以“字符”（rune）的粒度反转字符串，避免把 UTF-8 多字节字符拆坏。
func ReverseString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// CountWords 返回按空白分隔的单词数量。
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// itoa 是一个简化版 int -> string，用于避免在入门阶段引入 strconv 的概念。
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		return "-" + itoa(-n)
	}
	var b [32]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
