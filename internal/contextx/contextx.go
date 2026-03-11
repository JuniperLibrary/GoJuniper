// Package contextx 提供 context 主题的基础练习：
// - 取消（cancel）
// - 超时（timeout）
// - “等待一段时间或被取消”的常见模式
package contextx

import (
	"context"
	"time"
)

// SleepOrDone 等待 d，或者在 ctx 被取消时提前返回。
// 返回值：
// - 正常睡眠结束：nil
// - 被取消/超时：ctx.Err()
func SleepOrDone(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
