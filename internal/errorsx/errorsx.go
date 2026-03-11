// Package errorsx 是“错误处理”主题的练习：
// - errors.Is / errors.Join
// - fmt.Errorf("%w") 的包裹错误
package errorsx

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrNotPositive 表示解析后得到的整数不是正数。
var ErrNotPositive = errors.New("value must be positive")

// ParsePositiveInt 将字符串解析为正整数。
// - 格式不对时：返回带上下文的错误（使用 %w 包裹原始错误）
// - n <= 0 时：返回 ErrNotPositive
func ParsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
	}
	if n <= 0 {
		return 0, ErrNotPositive
	}
	return n, nil
}

// Join 将两个 error 合并成一个：
// - 只要其中一个是 nil，就返回另一个
// - 两个都非 nil，则使用 errors.Join
func Join(a, b error) error {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return errors.Join(a, b)
}
