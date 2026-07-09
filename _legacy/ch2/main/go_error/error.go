package main

import (
	"errors"
	"fmt"
)

/*
	Go 错误处理
		Go 语言通过内置的错误接口非常简单的错误处理机制。
		Go 语言的错误处理采用显式返回错误的方式，而非传统的异常处理机制。这种设计使代码逻辑更清晰，便于开发者在编译时或运行时明确处理错误。
		Go 的错误处理主要围绕以下机制展开：
			1. error 接口： 标准的错误表示
			2. 显式返回值： 通过函数的返回值返回错误
			3. 自定义错误：可以通过标准库或自定义的方式创建错误
			4. panic 和 recover 处理不可恢复的严重错误
*/

/*
	1.error 接口
		Go 标准库定义了一个 error 接口，表示一个错误的抽象。
		error 类型式一个接口类型，这是它的定义：
			type error interface{
				Error() string
			}
		实现 error 接口：任何实现了 Error() 方法的类型都可以作为错误
		Error() 方法返回一个描述错误的字符串
*/

func main() {

	/*
		2.使用errors包创建错误
	*/
	err := errors.New("this is an error")
	fmt.Println(err.Error())

	/*
		3.自定义错误
	*/
	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println(err2) // 输出：cannot divide 10 by 0
	}
}

type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}
