package fib

import (
	"testing"
)

// 声明变量 在程序内部复制
var c int
var d int

/**
fib数列
*/
func TestFibFirst(t *testing.T) {
	// 声明常量
	//1. 第一种方式
	//var a int = 1
	//var b int = 1

	//2.第二种方式
	var (
		a int = 1
		b int = 1
	)
	// 3.第三种方式
	x := 1
	y := 1

	//fmt.Println(x, y)
	t.Log(x, y)

	//fmt.Print(a)
	t.Log(a)

	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(" ", b)

		temp := a
		a = b
		b = temp + a
	}
	t.Log()
}

//交换两个值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//temp := a
	//a = b
	//b = temp
	a, b = b, a
	t.Log(a, "->", b)
}
