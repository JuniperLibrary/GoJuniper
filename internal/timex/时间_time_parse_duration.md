# Go 基础：时间（time.Time / 解析 / Duration）

这个文档对应 [internal/timex](file:///e:/dingchuan/GoJuniper/internal/timex)。

配合阅读：
- 实现代码：[timex.go](file:///e:/dingchuan/GoJuniper/internal/timex/timex.go)
- 测试代码：[timex_test.go](file:///e:/dingchuan/GoJuniper/internal/timex/timex_test.go)

---

## 1. time.Time 是什么

`time.Time` 表示一个“时间点”（某一刻），不是时长。

它携带：
- 时间点（秒/纳秒）
- 时区信息（Location）

初学者先记住：同一个“时间点”在不同时区显示不同，但它指向的是同一刻。

---

## 2. time.Duration 是什么

`time.Duration` 表示时长（比如 2 秒、500 毫秒）。

```go
var d time.Duration = 2 * time.Second
```

常用单位：
- `time.Second`
- `time.Millisecond`
- `time.Minute`
- `time.Hour`

---

## 3. 解析时间：RFC3339

学习阶段推荐优先掌握 RFC3339（JSON/HTTP 很常见）：

```go
t, err := time.Parse(time.RFC3339, "2026-03-11T12:34:56Z")
```

对应练习请看：[timex.go](file:///e:/dingchuan/GoJuniper/internal/timex/timex.go)

---

## 4. 计算时间差

两个时间点差值：

```go
delta := t2.Sub(t1)
```

加减时长：

```go
t3 := t1.Add(30 * time.Minute)
```

---

## 5. StartOfDay（一天的开始）为什么要小心

一天的开始通常指“当地时区的 00:00:00”。

注意：
- 如果你用 UTC 算，结果可能和本地时区的“当天开始”不一致
- DST（夏令时）地区某些日期可能会出现一天不是 24 小时的情况（工程里更常见）

本仓库的练习函数用来理解“把时间截断到某天开始”的思路，先把概念吃透即可。

---

## 6. 练习清单

1. 写 `MustParseRFC3339(s string) time.Time`（解析失败就 panic，仅用于测试/小工具）
2. 写 `FormatRFC3339(t time.Time) string`
3. 写 `IsSameDay(a, b time.Time, loc *time.Location) bool`

---

## 7. 自测命令

```bash
go test ./internal/timex -shuffle on
```

