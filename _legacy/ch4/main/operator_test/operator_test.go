package operator_test

import "testing"

// 用 == 比较数组
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	// fasle
	t.Log(a == b)
	//a == c (mismatched types [4]int and [5]int)
	//t.Log(a == c)
	// true
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

// 按位清零
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
