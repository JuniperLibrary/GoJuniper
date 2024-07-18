package main

import "fmt"

// user 结构体代表一个用户，包含姓名和年龄信息。
// 结构体(struct)可匿名嵌入其他类型。
type user struct {
	name string
	age  int
}

// manager 结构体代表一个管理者，它嵌入了 user 结构体，增加了职称信息。
// 通过嵌入 user 结构体，manager 继承了 user 的所有字段和方法。
type manager struct {
	user
	title string
}

func main() {
	// 初始化一个 manager 实例
	//var m manager
	m := manager{} // 使用默认值初始化manager类型的变量m

	// 设置 manager 实例的姓名、年龄和职称
	m.name = "Tom"
	m.age = 29
	m.title = "Go"

	// 输出 manager 实例的信息
	fmt.Println(m)
}
