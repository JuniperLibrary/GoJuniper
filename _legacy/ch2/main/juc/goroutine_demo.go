package main

import (
	"fmt"
	"sync"
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

	fmt.Println("===============Channel 通道===============")
	/*
		2. Channel 通道
			Channel 是用于Goroutine 之间的数据传递
			通道可用于两个 Goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
			使用 make 函数创建一个channel 使用 <- 操作符发送和接受收据。如果未指定，则为双向通道。
				ch <- v     // 把 v 发送到通道 ch
				v := v-ch   // 从 ch 接受数据
							// 并把值赋给 v
			声明一个通道很见到,我们使用 chan 关键字即可，通道在使用前必须创建
				ch := make(chan int)
			注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

			以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
	*/

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	go sumChannel(s[:len(s)/2], c)
	go sumChannel(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) // 从通道 c 中接收

	fmt.Println("===============通道缓冲区==================")

	/*
		3.通道缓冲区
			通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
				ch := make(chan int ,100)
			带缓冲区的通道允许发送段的数据发送和接受端的数据获取处于异步状态，
			就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要去接收端去获取数据。

			不过由于缓冲区的的大小是有限的，所以还是必须有接收端来接受数据的，否则缓冲区一满，数据发送端就无法在发送数据了。
			注意：
				如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
				如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；
				如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。
				接收方在有值可以接收之前会一直阻塞。
	*/

	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	chBuffer := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	chBuffer <- 1
	chBuffer <- 2

	// 获取这两个数据
	fmt.Println(<-chBuffer)
	fmt.Println(<-chBuffer)

	fmt.Println("===============遍历通道与关闭通道==================")

	/*
		4.Go 遍历通道与关闭通道
			Go 通过 range 关键字来实现遍历读取到的数据，类似与数组或切片。
				v,ok := <- ch
			如果通道接受不到数据后 ok 就为了false，这时通道就可以使用 close() 函数来关闭。
	*/

	cBuffer := make(chan int, 10)

	go fibonacciBuffer(cap(cBuffer), cBuffer)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了
	for i := range cBuffer {
		fmt.Println(i)
	}

	fmt.Println("===============Select=======================")
	/*
		5. Select 语句
			select 语句使得一个 goroutine 可以等待多个通信操作。
			select 会阻塞，直到其中的某个 case 可以继续执行
	*/
	channelSelect := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// 接收
			fmt.Println(<-channelSelect)
		}
		// 退出
		quit <- 0
	}()
	fibonacciSelect(channelSelect, quit)

	fmt.Println("===============使用 WaitGroup =======================")
	/*
		6.使用 WaitGroup
			sync.WaitGroup 用于等待多个 WaitGroup 完成

			同步多个 Goroutine：
	*/
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg)
	}
	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("All workers finished")

	/*
			7.高级特性
				7.1 Buffered Channel:
					创建有缓冲的 Channel
						ch := make(chan int, 10)
		        7.2 Context:
					用于控制 Goroutine 的生命周期
						context.WaitCancel、 context.WithTimeout
		        7.3 Mutex 和 RWMutex:
					sync.Mutex 提供互斥锁，用于保护共享资源。
						var mmu.Lock() // 锁定
						// critical section 临界段
						mu.Unlock()u sync.Mutex // 释放锁

	*/

	/*
		8. 并发变成小结
			Go 语言通过 GoRoutine 和 Channel 提供了强大的并发支持，简化了传统线程模型的复杂性。
			配合调度器和同步工具，可以轻松实现高性能并发程序。
				- Goroutines 是轻量级线程，使用 go 关键字启动。
				- Channels 用于 goroutines 之间的通信。
				- Select 语句 用于等待多个 channel 操作。

			常见问题
				死锁 DeadLock：
					- 示例：所有 Goroutine 都在等待，但没有任何数据可用。
					- 解决：避免无限等待、正确关闭通道。
				数据竞争 Data Race：
					- 示例：多个 goroutine 访问同一变量，导致数据不一致。
					- 解决：使用锁、原子操作、无锁数据结构。
	*/
}

func sumChannel(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum的值发送到 通道 c
}

func fibonacciBuffer(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/**
 * fibonacciSelect 使用 select 监听两个通道，生成斐波那契数列
 * @param c int 类型通道，用于发送生成的斐波那契数值
 * @param quit int 类型通道，用于接收退出信号
 */
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// 发送斐波那契数列
		case c <- x:
			// 计算下一个斐波那契数列
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}
