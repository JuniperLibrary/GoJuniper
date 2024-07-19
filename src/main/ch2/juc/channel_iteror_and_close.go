package main

import (
	"fmt"
)

// fibonacci2 计算斐波那契数列的前n个数字，并通过通道c返回结果。
// 参数n: 指定计算斐波那契数列的前n个数字。
// 参数c: 用于发送计算结果的通道。
func fibonacci(n int, c chan int) {
	// 初始化前两个斐波那契数。
	x, y := 0, 1
	// 循环计算并发送斐波那契数列的前n个数字。
	for i := 0; i < n; i++ {
		// 发送当前斐波那契数。
		c <- x
		// 更新x和y的值，为下一次计算做准备。
		x, y = y, x+y
	}
	// 斐波那契数列计算完成后，关闭通道。
	close(c)
}

func main() {
	// 创建一个容量为10的整型通道，用于接收斐波那契数列的前10个数字。
	c := make(chan int, 10)
	// 开启一个goroutine计算斐波那契数列，并将结果发送到通道c。
	go fibonacci(cap(c), c)
	// 遍历通道c，打印接收的每个斐波那契数。
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
