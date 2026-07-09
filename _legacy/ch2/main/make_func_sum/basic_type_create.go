package main

import "fmt"

func main() {

	/*
		基本类型（值类型）
	*/
	// 1. 布尔型
	var b bool = true
	fmt.Println("布尔型：", b) // 输出：true

	// 2. 整型（int/int8/int16/int32/int64/uint/uint8等）
	var i int = 10                // 默认int（32/64位系统对应不同长度）
	var i8 int8 = -128            // 范围：-128~127
	var u8 uint8 = 255            // 无符号，范围：0~255
	fmt.Println("整型：", i, i8, u8) // 输出：10 -128 255

	// 3. 浮点型（float32/float64）
	var f32 float32 = 3.14
	var f64 float64 = 3.1415926535
	fmt.Println("浮点型：", f32, f64) // 输出：3.14 3.1415926535

	// 4. 字符串（string）
	var s string = "Hello Go"
	fmt.Println("字符串：", s) // 输出：Hello Go

	// 5. 字符型（rune/byte）
	var c1 byte = 'a'                                   // 等价uint8，存ASCII字符
	var c2 rune = '中'                                   // 等价int32，存Unicode字符（包括中文）
	fmt.Println("字符型：", c1, string(c1), c2, string(c2)) // 输出：97 a 20013 中

	/*
		复合类型（引用 / 聚合类型）
	*/
	// 1. 数组（Array，固定长度，值类型）
	var arr1 [3]int = [3]int{1, 2, 3} // 定长数组
	arr2 := [5]string{"a", "b", "c"}  // 简写，未赋值的元素为默认值（""）
	fmt.Println("数组：", arr1, arr2)    // 输出：[1 2 3] [a b c  ]

	// 2. 切片（Slice，动态数组，引用类型，需make创建或基于数组切片）
	// 方式1：make创建
	s1 := make([]int, 2, 5) // len=2（初始长度），cap=5（容量）
	s1[0] = 10
	s1[1] = 20
	// 方式2：基于数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	s2 := arr[1:3]                               // 切片：[2,3]（左闭右开）
	fmt.Println("切片：", s1, len(s1), cap(s1), s2) // 输出：[10 20] 2 5 [2 3]

	// 3. 指针（Pointer，保存变量地址）
	var num int = 100
	var p *int = &num         // 取num的地址赋值给指针p
	fmt.Println("指针：", p, *p) // 输出：0xc000018078 100（地址不同正常）

	// 4. 结构体（Struct，自定义复合类型）
	type Person struct {
		name string
		age  int
	}
	// 创建结构体实例
	p1 := Person{name: "李四", age: 25}
	p2 := Person{"王五", 30}      // 简写（按字段顺序）
	fmt.Println("结构体：", p1, p2) // 输出：{李四 25} {王五 30}

	// 5. 通道（Channel，用于goroutine通信，引用类型）
	ch := make(chan int, 3) // 创建带缓冲的通道，容量3
	ch <- 1                 // 向通道发送数据
	ch <- 2
	fmt.Println("通道：", <-ch) // 从通道接收数据，输出：1

	// 6. 函数（Function，一等公民，可作为参数/返回值）
	add := func(a, b int) int { // 匿名函数赋值给变量
		return a + b
	}
	fmt.Println("函数：", add(3, 5)) // 输出：8

}
