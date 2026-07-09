package main

import "fmt"

//避免大数据结构的拷贝： 当函数参数是一个大数据结构时，如果传递值会造成拷贝，增加开销，这时候可以传递指针，避免拷贝
//type LargeStruct struct {
//	Data [1000]int
//}
//
//func ProcessStruct(s *LargeStruct) {
//	// 处理数据
//}
//
//func main() {
//	var ls LargeStruct
//	ProcessStruct(&ls)
//}

// 修改函数参数的值： 如果需要在函数内部修改参数的值，需要传递指针。
//func Increment(p *int) {
//	*p++
//}
//
//func main() {
//	x := 10
//	Increment(&x)
//	fmt.Println(x) // 输出 11
//}

// 与C语言交互： Go语言中可以通过CGo与C语言交互，很多C语言库需要传递指针
///*
//#include <stdio.h>
//
//void PrintMessage(const char* msg) {
//    printf("%s\n", msg);
//}
//*/
//import "C"
//
//func main() {
//	message := "Hello, C!"
//	C.PrintMessage(C.CString(message))
//}

// 链表或树等数据结构： 这些数据结构通常使用指针来连接节点。
//type Node struct {
//	Value int
//	Next  *Node
//}
//
//func main() {
//	node1 := &Node{Value: 1}
//	node2 := &Node{Value: 2}
//	node1.Next = node2
//
//	fmt.Println(node1.Value)      // 输出 1
//	fmt.Println(node1.Next.Value) // 输出 2
//}

// 方法接收者： 当方法需要修改接收者的值时，接收者类型应为指针
//type Counter struct {
//	Count int
//}
//
//func (c *Counter) Increment() {
//	c.Count++
//}
//
//func main() {
//	counter := &Counter{}
//	counter.Increment()
//	fmt.Println(counter.Count) // 输出 1
//}

//const MAX int = 3
//
//func main() {
//
//	a := []int{10, 100, 200}
//	var i int
//
//	for i = 0; i < MAX; i++ {
//		fmt.Printf("a[%d] = %d\n", i, a[i])
//	}
//}

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	fmt.Println("=====================================================")

	/*
		Go语言指针
	*/

	var b int = 10
	fmt.Printf("变量的地址：%x\n", &b)

	fmt.Println("===================================================")

	/*
		什么是指针？
			一个指针变量就是指向了一个值的内存地址。类似与常量和变量，在使用指针前你需要声明指针。
		格式如下：
			var var_name *var-type
		var-type 为指针类型；var_name 为指针变量名；* 号用于指定变量是作为一个指针;
		以下是有效的指针声明：
		var ip *int 	// 指向整形
		var fp *float32 // 指向浮点型
	*/

	/*
		如何使用指针？

		指针使用流程：
			定义指针变量。
			为指针变量赋值。
			访问指针变量中指向地址的值。
			在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
	*/
	var a1 int = 100 // 声明实际变量
	var ip *int      // 声明指针变量

	ip = &a1 // 指针变量的存储地址
	fmt.Printf("a 变量的地址是: %x\n", &a1)

	/*指针变量的存储地址*/
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	fmt.Println("==================================================")
	/*
		Go 空指针
			当一个指针被定义之后没有分配到任何变量，它的值为nil。
			nil 指针也称为空指针。nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
			一个指针变量通常缩写为 ptr2。
	*/
	var ptr2 *int
	fmt.Printf("ptr2 的值为 : %x\n", ptr2)

	fmt.Println("==================================================")

	/*
		Go 语言指针数组
	*/

	ptrArray := [...]int{100, 200, 100, 200}
	var index int
	for index = 0; index < MAX; index++ {
		fmt.Printf("a[%d] = %d\n", index, ptrArray[index])
	}
	fmt.Println("==================================================")
	/*
			有一种情况，我们可能需要保存数组，这样我们就需要使用到指针.
			以下声明了整型指针数组：
		 		var ptr [MAX]*int;
				ptr 为整型指针数组;
	*/

	var ptr3 [MAX]*int
	for index = 0; index < MAX; index++ {
		ptr3[index] = &ptrArray[index] // 整数地址赋值给指针数组
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr3[i])
	}
	fmt.Println("===============Go语言指向指针的指针===================================")
	/*
		Go语言指向指针的指针
			如果一个指针变量存放又是另一个指针变量的地址，则称为这个指针变量为指向指针的指针变量。
			当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
			Pointer 			Pointer             Variable
			address    ---->    address    ----->   Value
		指向指针的指针变量声明格式如下：
			var ptr **int
	*/

	var a2 int
	var ptr4 *int
	var ptr5 **int

	a2 = 3000
	/*指针ptr4的地址*/
	ptr4 = &a2
	/*指针ptr5的地址*/
	ptr5 = &ptr4
	/* 获取 ptr5 的值 */
	fmt.Printf("变量 a2 = %d\n", a2)
	fmt.Printf("指针变量 *ptr4 = %d\n", *ptr4)
	fmt.Printf("指向指针的指针变量 **ptr5 = %d\n", **ptr5)

	fmt.Println("===============Go语言指针作为函数参数===================================")
	/*
		Go 语言指针作为函数参数
			Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。
	*/

	var a3 int = 100
	var b3 int = 200
	fmt.Printf("交换前 a3 的值 : %d\n", a3)
	fmt.Printf("交换前 b3 的值 : %d\n", b3)

	swap(&a3, &b3)
	fmt.Printf("交换后 a3 的值 : %d\n", a3)
	fmt.Printf("交换后 b3 的值 : %d\n", b3)
}

/*
		/* 调用函数用于交换值
	   * &a 指向 a 变量的地址
	   * &b 指向 b 变量的地址
*/
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
