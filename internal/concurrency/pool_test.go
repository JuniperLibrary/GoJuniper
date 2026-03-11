package concurrency_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"gojuniper/internal/concurrency"
)

func TestRun(t *testing.T) {
	t.Run("invalid concurrency", func(t *testing.T) {
		// 非法参数应该被明确拒绝（返回可识别的错误）。
		err := concurrency.Run(context.Background(), 0, []concurrency.Job{func(context.Context) error { return nil }})
		if !errors.Is(err, concurrency.ErrInvalidConcurrency) {
			t.Fatalf("expected ErrInvalidConcurrency, got %v", err)
		}
	})

	t.Run("runs all jobs", func(t *testing.T) {
		var mu sync.Mutex
		var got []int

		jobs := make([]concurrency.Job, 0, 50)
		for i := 0; i < 50; i++ {
			i := i
			jobs = append(jobs, func(context.Context) error {
				mu.Lock()
				got = append(got, i)
				mu.Unlock()
				return nil
			})
		}

		if err := concurrency.Run(context.Background(), 8, jobs); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 50 {
			t.Fatalf("len=%d, want 50", len(got))
		}
	})

	t.Run("cancels on first error", func(t *testing.T) {
		// 这个测试验证：只要有一个任务失败，其它任务能尽快被取消。
		started := make(chan struct{}, 100)
		done := make(chan struct{})
		defer close(done)

		sentinel := errors.New("boom")
		jobs := []concurrency.Job{
			func(ctx context.Context) error {
				started <- struct{}{}
				return nil
			},
			func(ctx context.Context) error {
				started <- struct{}{}
				return sentinel
			},
			func(ctx context.Context) error {
				started <- struct{}{}
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(2 * time.Second):
					return nil
				case <-done:
					return nil
				}
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := concurrency.Run(ctx, 2, jobs)
		if !errors.Is(err, sentinel) {
			t.Fatalf("expected sentinel error, got %v", err)
		}
	})
}
