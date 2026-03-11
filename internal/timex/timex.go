// Package timex 提供 time 主题的基础练习：
// - 解析与格式化 RFC3339 时间
// - duration 的计算
// - “同一天起点”这种常见业务需求
package timex

import "time"

// ParseRFC3339 解析 RFC3339/RFC3339Nano 格式的时间字符串。
func ParseRFC3339(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// Until 返回 deadline - now 的 duration（可能为负数）。
func Until(deadline, now time.Time) time.Duration {
	return deadline.Sub(now)
}

// StartOfDay 返回 t 所在日期的 00:00:00（保持 t 的 Location）。
func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
