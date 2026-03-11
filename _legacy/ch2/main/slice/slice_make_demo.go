package main

import "fmt"

func main() {
	// 定义一个底层数组
	arr := [5]int{1, 2, 3, 4, 5}

	// 创建一个切片，引用 arr 的一部分
	slice := arr[1:3]

	fmt.Println("底层数组:", arr)
	fmt.Println("切片:", slice)
	fmt.Println("切片的长度:", len(slice))
	fmt.Println("切片的容量:", cap(slice))

	// 修改切片中的元素，会影响底层数组
	slice[0] = 99
	fmt.Println("修改切片后")
	fmt.Println("底层数组:", arr)
	fmt.Println("切片:", slice)

	// 使用内置函数 make 创建切片
	newSlice := make([]int, 3, 5)
	fmt.Println("使用 make 创建的切片:", newSlice)
	fmt.Println("新切片的长度:", len(newSlice))
	fmt.Println("新切片的容量:", cap(newSlice))

	// 向切片中添加元素
	newSlice = append(newSlice, 1, 2)
	fmt.Println("添加元素后切片:", newSlice)
	fmt.Println("添加元素后切片的长度:", len(newSlice))
	fmt.Println("添加元素后切片的容量:", cap(newSlice))
}
