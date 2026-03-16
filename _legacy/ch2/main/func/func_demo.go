package main

import (
	"errors"
	"fmt"
)

/*
	Go 语言函数的定义格式
			func function_name ( [parameter list]) [retutn_types]{
				函数体
			}
*/

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else if num2 > num1 {
		result = num2
	}
	return result
}

/*函数返回多个值*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
	函数参数
		函数如果使用参数，改变量可称为函数的形参
			形参就像定义在函数体内的局部变量,调用函数，可以通过两种方式来传递参数
				值传递
				引用传递
*/

// main 是程序的入口点。
func main() {
	a, b := 10, 2
	// 调用 div 函数尝试进行除法运算。
	c, err := div(a, b)
	fmt.Println(c, err)

	x := 100
	// 调用 test 函数并获取一个闭包，该闭包后续会打印 x 的值。
	printX := test(x)
	printX()

	// 初始化一个容量为5的切片，用于演示切片的动态增长。
	splits := make([]int, 0, 5)
	for i := 0; i < 9; i++ {
		splits = append(splits, i)
	}
	fmt.Println(splits)

	// 初始化一个映射，用于演示映射的添加和查找操作。
	maps := make(map[string]int)
	maps["a"] = 1
	maps["b"] = 2
	maps["c"] = 3
	maps["d"] = 4
	maps["e"] = 5
	item, ok := maps["a"]
	fmt.Println(item, ok)

	// 从映射中删除一个元素，演示映射的删除操作。
	delete(maps, "a")

	/* 定义局部变量 */
	d := 100
	e := 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(d, e)

	fmt.Printf("最大值是 : %d\n", ret)

	f, g := swap("hello", "world")
	fmt.Println(f, g)
}

// div 尝试执行两个整数的除法操作。
// 如果除数 b 为 0，函数返回一个错误。
// 参数:
//
//	a - 被除数
//	b - 除数
//
// 返回值:
//
//	除法结果
//	错误对象，如果除数为 0 则返回一个非 nil 的错误对象。
func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// test 返回一个闭包函数，该闭包在调用时会打印传入 test 函数的 x 的值。
// 参数 x 被捕获并存储在闭包中，以便后续访问。
// 返回值:
//
//	一个函数，当调用该函数时，它会打印 x 的值。
func test(x int) func() {
	return func() {
		fmt.Println(x)
	}
}
