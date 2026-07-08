// Package basics 提供一些 Go 入门阶段常见的"可测函数"，用来练习：
// - 变量/循环/switch
// - 切片与 rune（Unicode）
// - 错误处理与边界条件
// - 溢出检查（uint64）
//
// ⚠️ 注意：Go 里注释掉的代码不会被执行，但注释是给"读代码的人"看的。
// 本文件刻意把"为什么这样写"写进注释，方便复习时理解设计取舍。
package basics

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	// ErrNegativeN 表示参数 n 为负数（但该函数要求 n >= 0）。
	//
	// ⚠️ 注意：这是"哨兵错误（sentinel error）"模式——用包级导出的变量做错误比较。
	// 调用方用 errors.Is(err, basics.ErrNegativeN) 判断，而不是直接比较字符串。
	// 对应 Rust 里 Result<T, E> 的 E 枚举分支。
	ErrNegativeN = errors.New("n must be >= 0")

	// ErrOverflow 表示计算结果超过 uint64 的可表示范围。
	//
	// ⚠️ 注意：无符号整数溢出在 Go 里不会 panic，也不会报错，而是"回绕"（wrap around）。
	// 所以做阶乘/斐波那契这类指数增长的计算时，必须自己加溢出检查，否则会得到错误结果而非报错。
	ErrOverflow = errors.New("result overflows uint64")
)

// Sum 返回两个整数之和。
//
// ⚠️ 注意：Go 函数可以有多个返回值，但也可以像这样只有一个。
// int 在 64 位平台上是 64 位，和 Rust 的 i32/i64 不同——Go 的 int 宽度随平台变化。
func Sum(a, b int) int {
	return a + b
}

// Max 返回切片的最大值；如果切片为空则返回 (0, false)。
//
// ⚠️ 注意：Go 习惯用 "值, ok" 这种双返回值来表示"可能不存在"。
// 对应 Rust 的 Option<T>：有值返回 (v, true)，无值返回 (零值, false)。
// 这里空切片返回的 0 是 int 的零值，调用方必须看第二个 bool 才知道是不是真的最大值。
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
//
// ⚠️ 注意（两个关键 Go 特性）：
//  1. switch 后面不写表达式（裸 switch），等价于 if/else if 链；
//     且 Go 的 case 默认【不穿透】——匹配到就执行完跳出，不需要写 break（和 C/Rust 相反）。
//  2. 判断 15 的倍数必须放在 3 和 5 之前，否则会被 i%3==0 / i%5==0 先截胡。
func FizzBuzz(n int) []string {
	if n <= 0 {
		return nil
	}
	// ⚠️ 注意：make([]string, 0, n) 预分配容量为 n，避免 append 时反复扩容（性能优化点）。
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
//
// ⚠️ 注意：循环只需到 sqrt(n)，且跳过所有偶数（从 3 开始每次 +2）。
// 这是性能常识，但在 Go 里写 math.Sqrt 要注意参数是 float64，有浮点精度问题——
// 对小整数足够，大整数判断素数是另一个话题。
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
//
// ⚠️ 注意（溢出检查写法）：
// 正确写法是用 res > math.MaxUint64/m 判断"再乘会爆"，而不是 res*m < res（后者已溢出，测不出来）。
// 这是无符号整数避免静默回绕的标准技巧。
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
//
// ⚠️ 注意：同样用 a > math.MaxUint64-b 检查加法溢出，而不是 a+b < a。
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

// ReverseString 以"字符"（rune）的粒度反转字符串，避免把 UTF-8 多字节字符拆坏。
//
// ⚠️ 注意（Go 字符串最易踩的坑）：
// Go 的 string 底层是 UTF-8 字节序列。中文/emoji 一个"字符"可能占 2~4 字节。
// 如果直接按 byte 反转（[]byte(s)），会把多字节字符拆成乱码。
// 必须先转成 []rune（每个 rune = 一个 Unicode 码点），反转后再转回 string。
func ReverseString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// CountWords 返回按空白分隔的单词数量。
//
// ⚠️ 注意：strings.Fields 会按任意空白（空格/制表符/连续多个空格）切分，
// 比手写 strings.Split 更贴合"单词数"语义。对应 Rust 的 s.split_whitespace().count()。
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// itoa 是一个简化版 int -> string，用于避免在入门阶段引入 strconv 的概念。
//
// ⚠️ 注意：真实项目请用 strconv.Itoa，不要自己写。这里仅作教学演示（递归 + 取模拼字节）。
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

// IotaDemo 演示 iota 的基础用法：在 const 块里，iota 从 0 开始每行 +1。
//
// ⚠️ 注意：iota 在每个 const (...) 块里【独立从 0 重新开始计数】，离开 const 块后归零。
// 它本质是"行号"，不是全局计数器。
func IotaDemo() (int, int, int) {
	const (
		a = iota // a = 0
		b        // b = 1（iota 自动 +1）
		c        // c = 2
	)
	return a, b, c
}

// SwapByPointer 通过指针交换两个 int 的值。
//
// ⚠️ 注意：Go 函数参数是【值传递】。想修改调用方的变量，必须传 *int 指针。
// 对应 Rust 的 &mut i32。这里 *a = *b 解引用后赋值，才真正改到原变量。
func SwapByPointer(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// TypeConvertDemo 演示 float64 -> int（截断）和 float64 -> string（格式化）。
//
// ⚠️ 注意：int(f) 是【向零截断】而非四舍五入，3.9 -> 3。这一点和很多语言一致，但容易忘记。
// strconv.FormatFloat(f, 'f', -1, 64) 中 -1 表示"用最少必要的位数"，不是保留 1 位小数。
func TypeConvertDemo(f float64) (int, string) {
	return int(f), strconv.FormatFloat(f, 'f', -1, 64)
}

// MakeVsNew 对比 make 和 new 的区别：
//   - make 只用于 slice/map/chan，返回【已初始化】的值（非 nil）
//   - new(T) 为任意类型分配零值内存，返回【指向它的指针】 *T
//
// ⚠️ 注意：对 slice/map 来说，new([]int) 得到的是 *[]int（指针），且指向 nil 切片；
// 而 make([]int, 3) 直接得到可用的切片。所以日常几乎都用 make，new 较少见。
func MakeVsNew() (string, string) {
	s := make([]int, 3)
	p := new([]int)
	return fmt.Sprintf("make=%v len=%d", s, len(s)), fmt.Sprintf("new=%v nil=%v", *p, *p == nil)
}

// ========== 补录：数字字面量细节（下划线、进制前缀、原始字符串） ==========

// NumericLiteralsDemo 演示数字字面量的写法细节。
//
// ⚠️ **注意（数字字面量易被忽视的细节，教材《Learning Go》第 2 章强调）**：
//  1. 下划线 `_` 可作数字分隔符：`1_000_000`、`0xFF_FF`、`1.23_45e6` —— 仅提升可读性，不影响值。
//  2. 进制前缀：`0b`/`0B` 二进制、`0o`/`0O` 八进制、`0x`/`0X` 十六进制。
//     老式八进制只写前导 `0`（如 `0777`）极易误读，**避免使用**，统一写 `0o777`。
//  3. 复数字面量：`1+2i`、`3.14i`，内部用 float32/float64。
//  4. 字面量是**无类型的**，赋值时才按上下文推断类型（整数→int，浮点→float64，复数→complex128）。
func NumericLiteralsDemo() (int, int, int, complex128) {
	_ = 1_000_000   // 十进制下划线
	_ = 0b1010_1010 // 二进制下划线
	_ = 0o777       // 八进制（推荐 0o 前缀）
	_ = 0xFF_FF     // 十六进制下划线
	return 1_000_000, 0b1010, 0o777, 1 + 2i
}

// RawStringDemo 演示原始字符串字面量（反引号）。
//
// ⚠️ **注意**：原始字符串用反引号 “ ` “ 定界，**没有转义**，换行、制表符、引号都按原样保留。
// 适合写多行文本、正则表达式、JSON 模板、Windows 路径（无需双写反斜杠）。
// 解释字符串（双引号）里要写换行需 `\n`，反斜杠需 `\\`；原始字符串直接写即可。
func RawStringDemo() string {
	return `这是原始字符串：
包含换行
制表符:	
Windows路径: C:\Users\Name
JSON: {"key": "value"}
`
}

// ========== 补录：浮点数陷阱（货币别用 float、比较用 epsilon） ==========

// FloatPitfallsDemo 演示浮点数常见陷阱。
//
// ⚠️ **注意（浮点数两大坑，教材《Learning Go》第 2 章反复强调）**：
//  1. **货币/精确小数绝对不要用 float32/float64** —— 二进制浮点无法精确表示 0.1 等十进制小数。
//     0.1 + 0.2 = 0.30000000000000004，会导致金额错账。必须用整数（分/厘）或第三方 decimal 库。
//  2. **浮点数比较别用 ==**，必须用 epsilon（容差）：`math.Abs(a-b) < 1e-9`。
//     即使是看似相等的计算（如 0.1+0.2），`== 0.3` 也会是 false。
func FloatPitfallsDemo() (bool, bool) {
	// 陷阱 1：货币计算 —— 必须用变量才能触发浮点误差
	// ⚠️ 注意：Go 无类型常量在编译期用任意精度运算，`0.1+0.2==0.3` 在编译期为 true！
	// 所以这里必须先用变量 a、b，再相加，才能复现运行时二进制浮点的误差。
	a := 0.1
	b := 0.2
	sum := a + b
	moneyWrong := sum == 0.3 // false！这是浮点误差（用变量才能触发）
	// 正确做法：用整数（分）或 decimal 库。这里演示 epsilon 比较。
	epsilon := 1e-9
	_ = math.Abs(sum-0.3) < epsilon // true

	// 陷阱 2：直接 == 比较
	directEq := a+b == 0.3       // false
	_ = math.Abs(a+b-0.3) < 1e-9 // true

	return moneyWrong, directEq // 返回 false, false（演示陷阱），正确用法见注释
}

// FloatEpsilonEqual 浮点数 epsilon 比较工具函数。
//
// ⚠️ **注意**：浮点比较永远用 `math.Abs(a-b) < epsilon`，epsilon 视精度需求设定（如 1e-9、1e-12）。
// 这是通用工具，推荐放入工具包复用。
func FloatEpsilonEqual(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
