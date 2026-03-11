package main

import "fmt"

func main() {
	// 创建一个长度和容量均为 5 的整数切片
	slice := make([]int, 5)
	fmt.Println(slice) // 输出: [0 0 0 0 0]

	// 创建一个长度为 3，容量为 5 的整数切片
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2) // 输出: [0 0 0]
}
