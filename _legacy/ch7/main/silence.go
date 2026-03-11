package main

import "fmt"

/**
Slice 切片
切片和数组类似，可以把它理解为动态数组。切片是基于数组实现的，它的底层就是一个数组。对数组任意分隔，就可以得到一个切片

*/

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	// 左开 右闭
	//基于数组生成切片，包含索引start，但是不包含索引end
	//slice:=array[start:end]

	//array[:4] 等价于 array[0:4]
	//array[1:] 等价于 array[1:5]
	//array[:] 等价于 array[0:5]
	slice := array[2:5]
	fmt.Println(slice)

	// 切片修改
	slice2 := array[2:5]
	slice2[1] = "f"
	fmt.Println(array)

	// 切片申明
	// 声明了一个元素类型为 string 的切片，长度是 4，make 函数还可以传入一个容量参数
	slice1 := make([]string, 4)
	// 指定了新创建的切片 []string 容量为 8
	slice3 := make([]string, 4, 8)

	fmt.Println(slice1)
	fmt.Println(slice3)

}
