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

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
