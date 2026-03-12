package syncx_test

import (
	"sync"
	"testing"

	"gojuniper/internal/syncx"
)

// syncx 包的测试覆盖两个典型并发安全点：
// - Mutex 保护共享变量（Counter）
// - Once 确保初始化只执行一次（OnceValue）
func TestCounter(t *testing.T) {
	var c syncx.Counter

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	if got := c.Value(); got != 100 {
		t.Fatalf("got=%d, want 100", got)
	}
}

func TestOnceValue(t *testing.T) {
	var ov syncx.OnceValue[int]

	calls := 0
	f := func() int {
		calls++
		return 42
	}

	if got := ov.Get(f); got != 42 {
		t.Fatalf("got=%d, want 42", got)
	}
	if got := ov.Get(f); got != 42 {
		t.Fatalf("got=%d, want 42", got)
	}
	if calls != 1 {
		t.Fatalf("calls=%d, want 1", calls)
	}
}
