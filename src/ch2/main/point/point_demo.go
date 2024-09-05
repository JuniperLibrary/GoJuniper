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
}
