package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	/*
		Go 并发
			并发是指程序同时执行多个任务的能力。
			Go 语言支持并发，通过 goroutines 和 channels 提供了一种简洁高效的方式来实现并发。
		Goroutines
			go 中的并发执行单元，类似于轻量级的线程
			Goroutine 的调度由 Go 运行时管理，用户无需手动分配线程。
			使用 go 关键字启动 Goroutine
			Goroutine 是非阻塞的，可以高效地运行成千上万个 Goroutine
		Channel
			Go 中用于在 Goroutine 之间通信的机制
			支持同步和数据共享，避免了显式的锁机制
			使用 chain 关键字，通过 <- 操作符来发送和接收数据
		Scheduler 调度器
			Go 的调度器基于 GMP 模型，调度器会将 Goroutine 分配到系统线程中执行，并通过 M 和 P的配合高效管理并发
				G ： Goroutine
				M ：系统进程 Machine
				P ：逻辑处理器 Processor
	*/

	// goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
	/**
	Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
	同一个程序中的所有 goroutine 共享同一个地址空间。
	*/
	go say("world")

	say("hello")

	fmt.Println("========================")

	/*
		输出是没有固定先后顺序，因为它们是两个 goroutine 在执行：
	*/
	go sayHello() // 启动 Goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("Main")
		time.Sleep(100 * time.Millisecond)
	}

	/*
		2. Channel 通道
			Channel 是用于Goroutine 之间的数据传递
			通道可用于两个 Goroutine 之间通过传递一个指定类型的值来同步运行和通讯。

	*/
}
