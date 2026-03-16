package main

import "fmt"

func main() {
	// 创建一个长度和容量均为 5 的整数切片
	slice := make([]int, 5)
	fmt.Println(slice) // 输出: [0 0 0 0 0]

	// 创建一个长度为 3，容量为 5 的整数切片
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2) // 输出: [0 0 0]

	/*
		Go语言数组声明需要制定元素的类型和个数 语法格式
			var arrayName [size]dataType

		以下定义了数组 balance 长度为 10 类型为 float32：var balance [10]float32
	*/

	var numbers [5]int
	fmt.Println(numbers)

	var numbers1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	numbers2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)

	/*
		在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int 和 [10]int 是不同的类型。
	*/

	// 如果数组的长度不确定 可以使用 ... 来代替数组的长度 编译器会根据元素的个数自行推断
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)

	/*
		如果设置了数组的长度，我们还可以通过指定下标来初始化元素
	*/
	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance3)

	/*
		访问数组元素
	*/
	var n [10]int
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	/*
		示例
	*/
	
	balance4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出数组元素 */
	for i := 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance4[i])
	}

	balance5 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j := 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance5[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance6 := [5]float32{1: 2.0, 3: 7.0}
	for k := 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance6[k])
	}
}
