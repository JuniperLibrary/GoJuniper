package main

import "fmt"

func main() {
	// 创建一个空的字符串到整数的映射
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m) // 输出: map[one:1 two:2]

	// 创建一个初始容量为 10 的映射
	m2 := make(map[string]int, 10)
	fmt.Println(m2) // 输出: map[]
}
