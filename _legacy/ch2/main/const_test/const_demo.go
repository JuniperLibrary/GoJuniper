package main

import (
	"fmt"
)

// 变量
var userName string = "牛逼"

// 常量
const password string = "123123123"
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

//const (
//	a = "abv"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func add(a int, b int) int {
	return a + b
}

// 实现  只要实现了接口里的方法，就自动实现接口
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("wang")
}

func main() {
	fmt.Println("Hello, 世界")
	//var x int
	//x = 788
	// 声明变量+自动推导类型+赋值
	y := 10
	const Pi float64 = 3.1415
	fmt.Println(y)
	fmt.Println(Pi)

	// %d 表示整型数字 %s 表示字符串
	var stockcode = 123
	var enddate = "2002-12-31"
	var url = "Code=%d,enddate=%s"
	var target_url_foramt = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println("url= " + target_url_foramt)

	// Go 数据类型
	// 布尔型
	var bo bool = true
	fmt.Println(bo)
	// 数字类型 整数型：int、浮点型：float32、float64、复数：

	// 字符串类型
	// 派生类型
	//    a、指针类型 Pointer 是用来保存变量的内存地址，主要是为了：修改原变量和减少内存复制
	z := 20
	p := &z
	fmt.Println(p)  // 地址
	fmt.Println(*p) // 10
	//    b、数组类型 Array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0])
	//    c、结构化类型 struct 相当于 Java的class
	type UserDemo struct {
		Name string
		Age  int
	}
	u := UserDemo{Name: "Jack", Age: 18}
	fmt.Println(u)
	//    d、Channel 类型 通道 是 Go并发编程的核心
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	fmt.Println(<-ch)

	//    e、函数类型

	f := add
	fmt.Println(f(1, 2))
	//    f、切片类型 Slice 动态数组 长度可变
	s := []int{1, 2, 3}
	s = append(s, 4)
	//    g、接口类型 interface 接口定义 行为规范
	type Animal interface {
		Speak()
	}

	//    h、Map 类型
	m := map[string]int{
		"Tom":  16,
		"Lucy": 34,
	}
	fmt.Println(m["Tom"])

	//var name string
	//name = "Bobby"
	//
	//fmt.Println(name)
	//
	//var name1 string = "golang"
	//fmt.Println(name1)
	//
	//var name2 = "golang"
	//fmt.Println(name2)
	//
	//name3 := "golang"
	//fmt.Println(name3)
	//
	//const LENGTH int = 10
	//const WIDTH int = 5
	//var area int
	//const a1, b1, c1 = 1, false, "str"
	//area = LENGTH * WIDTH
	//fmt.Printf("面积为：%d\n", area)
	//println()
	//println(a1, b1, c1)
	//
	//println(a, b, c)
	//
	//fmt.Println("i=", i)
	//fmt.Println("j=", j)
	//fmt.Println("k=", k)
	//fmt.Println("l=", l)
	//
	//// Go语言运算符
	//a2 := 21
	//b2 := 10
	//var c int
	//
	//c = a2 + b2
	//fmt.Printf("第一行 - c 的值为 %d\n", c)
	//c = a2 - b2
	//fmt.Printf("第二行 - c 的值为 %d\n", c)
	//c = a2 * b2
	//fmt.Printf("第三行 - c 的值为 %d\n", c)
	//c = a2 / b2
	//fmt.Printf("第四行 - c 的值为 %d\n", c)
	//c = a2 % b2
	//fmt.Printf("第五行 - c 的值为 %d\n", c)
	//a2++
	//fmt.Printf("第六行 - a 的值为 %d\n", a2)
	//a2 = 21 // 为了方便测试，a 这里重新赋值为 21
	//a2--
	//fmt.Printf("第七行 - a 的值为 %d\n", a2)

	//var a int = 21
	//var b int = 10
	//
	//if a == b {
	//	fmt.Printf("第一行 - a 等于 b\n")
	//} else {
	//	fmt.Printf("第一行 - a 不等于 b\n")
	//}
	//if a < b {
	//	fmt.Printf("第二行 - a 小于 b\n")
	//} else {
	//	fmt.Printf("第二行 - a 不小于 b\n")
	//}
	//
	//if a > b {
	//	fmt.Printf("第三行 - a 大于 b\n")
	//} else {
	//	fmt.Printf("第三行 - a 不大于 b\n")
	//}
	///* Lets change value of a and b */
	//a = 5
	//b = 20
	//if a <= b {
	//	fmt.Printf("第四行 - a 小于等于 b\n")
	//}
	//if b >= a {
	//	fmt.Printf("第五行 - b 大于等于 a\n")
	//}
	//
	//var a bool = true
	//var b bool = false
	//if a && b {
	//	fmt.Printf("第一行 - 条件为 true\n")
	//}
	//if a || b {
	//	fmt.Printf("第二行 - 条件为 true\n")
	//}
	///* 修改 a 和 b 的值 */
	//a = false
	//b = true
	//if a && b {
	//	fmt.Printf("第三行 - 条件为 true\n")
	//} else {
	//	fmt.Printf("第三行 - 条件为 false\n")
	//}
	//if !(a && b) {
	//	fmt.Printf("第四行 - 条件为 true\n")
	//}

	//var a uint = 60 /* 60 = 0011 1100 */
	//var b uint = 13 /* 13 = 0000 1101 */
	//var c uint = 0
	//
	//println(a, b, c)
	//
	//c = a & b /* 12 = 0000 1100 */
	//fmt.Printf("第一行 - c 的值为 %d\n", c)
	//
	//c = a | b /* 61 = 0011 1101 */
	//fmt.Printf("第二行 - c 的值为 %d\n", c)
	//
	//c = a ^ b /* 49 = 0011 0001 */
	//fmt.Printf("第三行 - c 的值为 %d\n", c)
	//
	//c = a << 2 /* 240 = 1111 0000 */
	//fmt.Printf("第四行 - c 的值为 %d\n", c)
	//
	//c = a >> 2 /* 15 = 0000 1111 */
	//fmt.Printf("第五行 - c 的值为 %d\n", c)

	//var a int = 21
	//var c int
	//
	//c = a
	//fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)
	//
	//c += a
	//fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)
	//
	//c -= a
	//fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)
	//
	//c *= a
	//fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)
	//
	//c /= a
	//fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)
	//
	//c = 200
	//
	//c <<= 2
	//fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)
	//
	//c >>= 2
	//fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)
	//
	//c &= 2
	//fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)
	//
	//c ^= 2
	//fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)
	//
	//c |= 2
	//fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)

	/*
		&  返回变量存储地址  &a ；将给出变量的实际地址
		*  指针变量，  *a ；是一个指针变量
	*/

	a := 4
	var b int32
	var c float32
	var ptr *int
	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/* & 和 * 运算符实例*/
	ptr = &a
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)

}
