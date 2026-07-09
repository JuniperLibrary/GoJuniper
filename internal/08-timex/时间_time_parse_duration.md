# Go 基础：时间（time.Time / 解析 / Duration）

这个文档对应 [internal/08-timex](file:///e:/dingchuan/GoJuniper/internal/08-timex)。

配合阅读：
- 实现代码：[timex.go](file:///e:/dingchuan/GoJuniper/internal/08-timex/timex.go)

> ⚠️ **注意**：本仓库 `08-timex` 目录下只有实现代码 `timex.go`，没有 `_test.go`。学习时直接看 `timex.go` 里的 `ParseRFC3339`/`Until`/`StartOfDay` 即可，练习清单里的函数也写在同包内。

---

## 1. time.Time 是什么

`time.Time` 表示一个“时间点”（某一刻），不是时长。

它携带：
- 时间点（秒/纳秒）
- 时区信息（Location）

初学者先记住：同一个“时间点”在不同时区显示不同，但它指向的是同一刻。

> ⚠️ **注意（时区陷阱）**：`time.Time` 内部存的是“绝对时刻 + Location”。`time.Now()` 默认用 `time.Local`，而 `time.Parse(time.RFC3339, "2026-03-11T12:34:56Z")` 里带 `Z` 的会被解析成 **UTC**。`t.Hour()` 等方法是按 `t` 自带的 Location 显示的——同一个时刻在 UTC 和 Local 下 `Hour()` 可能差 8 小时。跨时区比较用 `.Equal()`，不要直接比 `.Hour()`。

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

> ⚠️ **注意（Duration 本质是 int64 纳秒）**：`time.Duration` 底层是 `int64`，单位是**纳秒**。千万别写 `d := 5 * time.Second` 后期望它是“5”——它的值是 `5_000_000_000`。要换算成秒用 `d / time.Second`，要构造时长务必乘 `time.Second` 这种常量，不要手写 `* 1_000_000_000`（易错且难读）。本仓库 `Until(deadline, now)` 直接返回 `deadline.Sub(now)` 的 `time.Duration`，天然带单位。

---

## 3. 解析时间：RFC3339

学习阶段推荐优先掌握 RFC3339（JSON/HTTP 很常见）：

```go
t, err := time.Parse(time.RFC3339, "2026-03-11T12:34:56Z")
```

对应练习请看：[timex.go](file:///e:/dingchuan/GoJuniper/internal/08-timex/timex.go)

> ⚠️ **注意（Go 没有 YYYY-MM-DD）**：Go 的时间格式化**不用占位符**，而用固定的“参考时间” `Mon Jan 2 15:04:05 MST 2006`（即 `2006-01-02 15:04:05`）。想格式化成 `"2026-03-11"` 要写 `t.Format("2006-01-02")`，写 `"YYYY-MM-DD"` 会得到字面量 `YYYY-MM-DD` 而非日期。本仓库 `ParseRFC3339` 直接用标准常量 `time.RFC3339Nano`，避免手写出错。

**为什么**用参考时间而不是 `YYYY-MM-DD`？这是 Go 的著名设计选择——用“一个具体的时刻”当模板，比记一堆占位符更符合直觉（也更好记：1月2日 下午3点4分5秒 2006年，恰好是 01 02 03 04 05 06）。

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

> ⚠️ **注意**：时间差返回的是 `time.Duration`，不是秒。想要“相差多少秒”用 `int(delta.Seconds())`。`time.After(d)`/`time.Sleep(d)` 接受的也是 `time.Duration`，传 `int` 会编译报错——这正是强制你用 `time.Second` 单位的好处。

---

## 5. StartOfDay（一天的开始）为什么要小心

一天的开始通常指“当地时区的 00:00:00”。

注意：
- 如果你用 UTC 算，结果可能和本地时区的“当天开始”不一致
- DST（夏令时）地区某些日期可能会出现一天不是 24 小时的情况（工程里更常见）

本仓库的练习函数用来理解“把时间截断到某天开始”的思路，先把概念吃透即可。

> ⚠️ **注意（保持 Location）**：本仓库 `StartOfDay` 用 `time.Date(y, m, d, 0, 0, 0, 0, t.Location())`，关键是**传入 `t.Location()`** 而不是 `time.UTC`。否则把本地时间截断到“当天开始”会变成 UTC 的 00:00，转换成当地时间后可能不是你想要的“零点”。处理“某天”类业务时，务必明确约定用哪个时区。

**常见坑**：DST 切换日（如春令时往前拨 1 小时），当天可能只有 23 小时，或在秋令时重复某一小时。`t.Add(24 * time.Hour)` 不保证落在“明天同一时刻”——需要“加一天”时用 `t.AddDate(0, 0, 1)` 更稳妥。

---

## 6. 练习清单

1. 写 `MustParseRFC3339(s string) time.Time`（解析失败就 panic，仅用于测试/小工具）
2. 写 `FormatRFC3339(t time.Time) string`
3. 写 `IsSameDay(a, b time.Time, loc *time.Location) bool`

---

## 7. 自测命令

```bash
go test ./internal/08-timex -shuffle on
```

> ⚠️ **注意**：当前目录没有测试文件，该命令会提示 “no test files”。先补 `timex_test.go` 或把练习函数写好再运行。
