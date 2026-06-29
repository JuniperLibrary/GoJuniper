// Package channelsx 提供 channel 主题的基础练习：
// - 生成器（generator）
// - pipeline：输入 -> 处理 -> 输出
// - fan-in：合并多个 channel
// - 关闭 channel 与 range 读取
package channelsx

import (
	"context"
	"sync"
)

// Generate 返回一个 channel，依次发送 0..n-1，然后关闭。
func Generate(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()
	return out
}

// Square 从 in 读取整数，发送平方结果到输出 channel。
func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v * v:
			}
		}
	}()
	return out
}

// Merge 将多个输入 channel 的值合并到一个输出 channel。
func Merge(ctx context.Context, ins ...<-chan int) <-chan int {
	out := make(chan int)
	if len(ins) == 0 {
		close(out)
		return out
	}

	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, ch := range ins {
		// for range 变量在循环中会复用，这里拷贝一份避免 goroutine 捕获到同一个变量。
		ch := ch
		go func() {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}()
	}

	// 等所有输入 channel 都关闭后，再关闭 out，让下游 range 正常退出。
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// FibonacciSelect 使用 select + quit channel 生成斐波那契数列的前 n 项。
func FibonacciSelect(ctx context.Context, n int) ([]int, error) {
	fibCh := make(chan int)
	quit := make(chan int, 1)

	go func() {
		x, y := 0, 1
		for {
			select {
			case fibCh <- x:
				x, y = y, x+y
			case <-quit:
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		select {
		case v := <-fibCh:
			result = append(result, v)
		case <-ctx.Done():
			quit <- 0
			return nil, ctx.Err()
		}
	}
	quit <- 0
	return result, nil
}

// FibonacciChannel 生成斐波那契数列的前 n 项，通过 channel 逐个发送，完成后 close。
func FibonacciChannel(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		x, y := 0, 1
		for i := 0; i < n; i++ {
			select {
			case out <- x:
				x, y = y, x+y
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
