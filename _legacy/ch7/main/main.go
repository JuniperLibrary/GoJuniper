package main

import "fmt"

/*
*
数组存放的是固定长度、相同类型的数据，而且这些存放的元素是连续的。所存放的数据类型没有限制，可以是整型、字符串甚至自定义
*/
func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])

	// 数组的长度可以忽略
	array1 := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array1[2])

	// 针对特定的索引指定元素
	array2 := [5]string{1: "b", 3: "d"}
	fmt.Println(array2[1])

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}

	// 数组循环
	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	//如果返回的值 用不到 可以使用 _ 代替
	for _, v := range array {
		fmt.Printf("对应值:%s\n", v)
	}
}
