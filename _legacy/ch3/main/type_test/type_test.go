package type_test

import "testing"

type MYINT int64

// 隐式类型转换和约束
func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	//cannot use a (type int) as type int64 in assignment
	//b = a
	b = int64(a)

	//cannot use a (type int) as type int64 in assignment
	//var c MYINT
	//c = b

	t.Log(a, b)
}

// 指针类型
func TestPoint(t *testing.T) {
	a := 1
	// 取址符
	aPtr := &a
	// 不支持指针运算
	// aPtr = aPtr + 1
	t.Log(a, " ", aPtr)
	t.Logf("%T %T", a, aPtr)
}

// string
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log(nil)
	}
}
