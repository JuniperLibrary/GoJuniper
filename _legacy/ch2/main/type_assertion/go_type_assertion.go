package main

import "fmt"

func main() {
	var i interface{} = "Hello, Go!"

	// 尝试将 i 断言为 string 类型
	s, ok := i.(string)
	if ok {
		fmt.Println("断言成功:", s)
	} else {
		fmt.Println("断言失败")
	}

	// 尝试将 i 断言为 int 类型
	n, ok := i.(int)
	if ok {
		fmt.Println("断言成功:", n)
	} else {
		fmt.Println("断言失败")
	}
	switch v := i.(type) {
	case int:
		fmt.Println("这是一个整数:", v)
	case string:
		fmt.Println("这是一个字符串:", v)
	default:
		fmt.Println("未知类型")
	}
}

func processInterface(i interface{}) {
	if s, ok := i.(string); ok {
		fmt.Println("处理字符串:", s)
	} else if n, ok := i.(int); ok {
		fmt.Println("处理整数:", n)
	} else {
		fmt.Println("无法处理的类型")
	}
}
