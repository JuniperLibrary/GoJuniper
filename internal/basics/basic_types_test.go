package basics_test

import (
	"testing"
)

// 这组测试用于把“Go 基本数据类型”的关键行为固化成可运行用例：
// - 零值
// - string/byte/rune 与 UTF-8
// - 数值类型转换与未类型常量
// - complex、nil map 行为、array vs slice 的语义差异
func TestBasicTypes_ZeroValues(t *testing.T) {
	// 基本类型的零值（声明即有值）
	var b bool
	var i int
	var f float64
	var s string

	// 引用/指针类类型的零值通常是 nil
	var p *int
	var xs []int
	var m map[string]int
	var ch chan int
	var fn func() int
	var itf any

	if b != false {
		t.Fatalf("bool zero value must be false")
	}
	if i != 0 {
		t.Fatalf("int zero value must be 0")
	}
	if f != 0 {
		t.Fatalf("float64 zero value must be 0")
	}
	if s != "" {
		t.Fatalf("string zero value must be empty string")
	}
	if p != nil || xs != nil || m != nil || ch != nil || fn != nil || itf != nil {
		t.Fatalf("reference-like zero values must be nil")
	}
}

func TestBasicTypes_StringByteRune(t *testing.T) {
	// string 是字节序列；UTF-8 下 len(s) 返回字节数，不是“字符数”
	s := "你好Go"
	if got, want := len(s), 8; got != want {
		t.Fatalf("len(s)=%d, want %d", got, want)
	}
	if got, want := len([]rune(s)), 4; got != want {
		t.Fatalf("len([]rune(s))=%d, want %d", got, want)
	}

	// s[i] 的结果是 byte；把 byte 当 rune 用会得到错误的“字符”
	if got := rune(s[0]); got == '你' {
		t.Fatalf("s[0] is a byte, not a rune")
	}

	// 想按“字符”（码点）处理字符串：用 for range 或先转 []rune
	var first rune
	for _, r := range s {
		first = r
		break
	}
	if first != '你' {
		t.Fatalf("first rune=%q, want %q", first, '你')
	}
}

func TestBasicTypes_NumericConversions(t *testing.T) {
	// Go 不做隐式数值类型转换，不同整数类型之间赋值要显式转换
	var a int32 = 123
	var b int64 = int64(a)
	if b != 123 {
		t.Fatalf("b=%d, want 123", b)
	}
}

func TestBasicTypes_UntypedConstantFlexibility(t *testing.T) {
	// 未类型常量在赋值时可“适配”目标类型（只要值能表示）
	const x = 10
	var a int8 = x
	var b int64 = x
	if a != 10 || b != 10 {
		t.Fatalf("expected both values to be 10")
	}
}

func TestBasicTypes_ComplexNumbers(t *testing.T) {
	// complex 的实部/虚部分别用 real/imag 取出
	c := complex(1, 2)
	if real(c) != 1 || imag(c) != 2 {
		t.Fatalf("real/imag=%v/%v, want 1/2", real(c), imag(c))
	}
}

func TestBasicTypes_NilMapReadOkWritePanics(t *testing.T) {
	// nil map：读取返回 value 的零值；写入会 panic
	var m map[string]int
	if got := m["a"]; got != 0 {
		t.Fatalf("nil map read got=%d, want 0", got)
	}

	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic when writing to nil map")
		}
	}()
	m["a"] = 1
}

func TestBasicTypes_ArrayVsSliceSemantics(t *testing.T) {
	// array 赋值会拷贝全部元素；slice 赋值会共享底层数组
	a1 := [3]int{1, 2, 3}
	a2 := a1
	a2[0] = 9
	if a1[0] != 1 {
		t.Fatalf("array assignment should copy values")
	}

	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 9
	if s1[0] != 9 {
		t.Fatalf("slice assignment shares underlying array")
	}
}
