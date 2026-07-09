package main

import "fmt"

func main() {
	// 创建一个无缓冲的整数通道
	ch := make(chan int)

	// 创建一个缓冲区大小为 2 的整数通道
	ch2 := make(chan int, 2)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch) // 输出: 1

	ch2 <- 2
	ch2 <- 3
	fmt.Println(<-ch2) // 输出: 2
	fmt.Println(<-ch2) // 输出: 3
}
