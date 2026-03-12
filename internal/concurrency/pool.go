// Package concurrency 提供并发主题的入门练习：
// - goroutine + channel
// - context 取消
// - worker pool（固定并发度执行一批任务）
package concurrency

import (
	"context"
	"errors"
	"sync"
)

// ErrInvalidConcurrency 表示并发度参数不合法。
var ErrInvalidConcurrency = errors.New("concurrency must be >= 1")

// Job 表示一个可取消的任务。
// 任务应当定期检查 ctx.Done()（或使用 ctx 感知的 API），以便在取消时尽快退出。
type Job func(context.Context) error

// Run 使用指定并发度执行 jobs。
// - 任何一个 job 返回非 nil error：尽快取消其它任务，并返回该错误
// - jobs 为空：直接返回 nil
func Run(ctx context.Context, concurrency int, jobs []Job) error {
	if concurrency < 1 {
		return ErrInvalidConcurrency
	}
	if len(jobs) == 0 {
		return nil
	}

	// 用派生 ctx 统一控制“整组任务”的取消；一旦任一任务失败，就 cancel 掉其它任务。
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// jobCh 是任务队列：生产者把任务塞进来，多个 worker 并发消费。
	jobCh := make(chan Job)
	// errCh 只需要记录“第一个错误”，所以容量为 1；
	// 写入使用 non-blocking，避免多个 worker 同时报错时相互阻塞。
	errCh := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for job := range jobCh {
				if job == nil {
					continue
				}
				if err := job(ctx); err != nil {
					select {
					case errCh <- err:
						cancel()
					default:
					}
					return
				}
				if ctx.Err() != nil {
					return
				}
			}
		}()
	}

	// 生产者：依次投递任务；遇到取消则停止投递；最后关闭 jobCh 让 worker 退出循环。
	go func() {
		defer close(jobCh)
		for _, job := range jobs {
			select {
			case <-ctx.Done():
				return
			case jobCh <- job:
			}
		}
	}()

	// 等待所有 worker 退出（要么把任务做完，要么被取消，要么某个任务报错后提前退出）。
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
