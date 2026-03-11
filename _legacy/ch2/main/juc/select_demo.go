package main

import "fmt"

// fibonacci2 生成斐波那契数列
// 参数:
// c - 用于发送斐波那契数的通道
// quit - 用于终止生成的通道
func fibonacci2(c, quit chan int) {
	x, y := 0, 1 // 初始化斐波那契数列的前两个数
	for {
		select {
		case c <- x: // 发送当前的斐波那契数
			x, y = y, x+y // 更新斐波那契数列的值
		case <-quit: // 接收到退出信号
			fmt.Println("quit")
			return // 终止函数执行
		}
	}
}

func main() {
	c := make(chan int)    // 创建用于接收斐波那契数的通道
	quit := make(chan int) // 创建用于发送退出信号的通道

	// 启动一个goroutine生成斐波那契数列
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 接收并打印斐波那契数
		}
		quit <- 0 // 生成退出信号
	}()

	fibonacci2(c, quit) // 调用斐波那契数列生成函数
}
