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

func main() {
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
