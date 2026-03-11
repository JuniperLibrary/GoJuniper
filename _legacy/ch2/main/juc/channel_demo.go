package main

import "fmt"

// main 是程序的入口点。
// 它启动一个生产者和一个消费者goroutine，并等待消费者完成。
func main() {
	done := make(chan bool) // done 通道用于通知主goroutine消费者已完成。
	data := make(chan int)  // data 通道用于在生产者和消费者之间传递数据。

	// 启动消费者goroutine。
	go consumer(data, done)
	// 启动生产者goroutine。
	go producer(data)
	// 等待消费者完成。
	<-done
}

// producer 是一个goroutine，负责生成数据并发送到data通道。
// 参数data是用于发送数据的通道。
func producer(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
		fmt.Println("producer: ", i)
	}
	close(data) // 关闭data通道表示所有数据已经发送完毕。
}

// consumer 是一个goroutine，负责从data通道接收数据并打印。
// 参数data是用于接收数据的通道。
// 参数done是用于通知主goroutine消费者已完成的通道。
func consumer(data chan int, done chan bool) {
	for v := range data {
		fmt.Println("recv:", v)
	}
	done <- true // 发送true到done通道表示消费者已完成。
}
