package contextx_test

import (
	"context"
	"testing"
	"time"

	"gojuniper/internal/contextx"
)

// 这组测试覆盖两个分支：
// - 正常等待到期返回 nil
// - ctx 被取消后提前返回 ctx.Err()
func TestSleepOrDone(t *testing.T) {
	t.Run("finishes on time", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		if err := contextx.SleepOrDone(ctx, 10*time.Millisecond); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("cancels", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := contextx.SleepOrDone(ctx, 2*time.Second)
		if err == nil {
			t.Fatalf("expected error")
		}
	})
}
