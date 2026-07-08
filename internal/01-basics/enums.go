// Package basics 提供枚举练习示例
package basics

import "fmt"

// ========== 基础 iota 枚举 ==========

const (
	// Weekday 周几枚举（基础用法）
	//
	// ⚠️ 注意：Go 没有 enum 关键字！"枚举"是用 const + iota 模拟出来的。
	// 本质就是一组有类型的整数常量。Sunday=0 是 iota 在该 const 块的"行号"。
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// ========== 跳过 0 的枚举 ==========

const (
	// 用 _ = iota 丢弃第 0 行，让真正的值从 1 开始。
	//
	// ⚠️ 注意：很多 Go 代码让枚举从 1 开始（而不是 0），目的是让"零值"代表"未设置/未知"状态。
	// 因为 int 的零值是 0，如果枚举从 0 开始，零值会和第一个枚举项冲突，难以区分"没赋值"和"取第一个"。
	_ = iota // 0 被丢弃
	// Status 状态枚举（从1开始）
	StatusPending   // 1
	StatusRunning   // 2
	StatusCompleted // 3
	StatusFailed    // 4
)

// ========== 位运算枚举（权限标志） ==========

const (
	// Permission 权限位标志
	//
	// ⚠️ 注意：1 << iota 利用位移生成 1, 2, 4, 8... 这种"每个位独立"的值，
	// 可以多位【组合】（OR）在一个 int 里，对应 Java 的 EnumSet 或位掩码（int 多标志位）。
	// 判断某位是否set：permissions & flag != 0（见 CheckPermission）。
	PermissionRead    = 1 << iota // 1 (二进制: 001)
	PermissionWrite               // 2 (二进制: 010)
	PermissionExecute             // 4 (二进制: 100)
)

const (
	// FileMode 文件模式标志
	//
	// ⚠️ 注意：这里 iota 乘 3（1 << (iota*3)）让每组权限占 3 个 bit，
	// 从而 user/group/other 可以共存而不冲突。这是"位字段布局"的典型用法。
	ModeUserRead   = 1 << (iota * 3) // 1   (用户读)
	ModeUserWrite                    // 8   (用户写)
	ModeUserExec                     // 64  (用户执行)
	ModeGroupRead                    // 512 (组读)
	ModeGroupWrite                   // 4096 (组写)
	ModeGroupExec                    // 32768 (组执行)
)

// ========== 自定义类型枚举 ==========

// Direction 方向枚举类型
//
// ⚠️ 注意：推荐用"自定义类型 + 底层 int"来定义枚举（type Direction int），
// 而不是裸 int 常量。这样 Direction 和 int 是不同类型，不能混用，类型安全更高。
type Direction int

const (
	// North 北
	North Direction = iota
	// South 南
	South
	// East 东
	East
	// West 西
	West
)

// String 实现 Stringer 接口
//
// ⚠️ 注意：Go 的"接口是隐式实现"的——只要类型有 String() string 方法，
// 就自动实现了 fmt.Stringer 接口，不需要写 implements。打印时 fmt.Println(d) 会调用它。
	// 这对应 Java 里接口是显式 implements，但 Go 是纯结构化的（duck typing）——只要方法签名匹配就自动满足接口。
// 这里的 [...]string{...}[d] 是用数组按枚举值索引取名字——⚠️ 前提是 d 不越界，否则 panic。
func (d Direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}

// ========== 另一个自定义类型枚举 ==========

// Color 颜色枚举类型
type Color int

const (
	// ColorNone 无颜色
	ColorNone Color = iota
	// ColorRed 红色
	ColorRed
	// ColorGreen 绿色
	ColorGreen
	// ColorBlue 蓝色
	ColorBlue
)

// String 实现 Stringer 接口
//
// ⚠️ 注意：用 switch 实现 String() 比数组索引更安全（default 分支兜底，不会越界 panic）。
// 枚举值很多或可能扩展时，优先用 switch 写法。
func (c Color) String() string {
	switch c {
	case ColorNone:
		return "None"
	case ColorRed:
		return "Red"
	case ColorGreen:
		return "Green"
	case ColorBlue:
		return "Blue"
	default:
		return "Unknown"
	}
}

// RGBA 返回颜色的RGBA值
//
// ⚠️ 注意：Go 约定"颜色分量"用 uint8 (0~255)。多个返回值用括号一次性声明。
func (c Color) RGBA() (r, g, b, a uint8) {
	a = 255
	switch c {
	case ColorRed:
		r = 255
	case ColorGreen:
		g = 255
	case ColorBlue:
		b = 255
	}
	return
}

// ========== 季节枚举 ==========

// Season 季节枚举类型
type Season int

const (
	// SeasonSpring 春季
	//
	// ⚠️ 注意：iota + 1 让枚举从 1 开始（跳过 0）。这是"从 1 起始"的另一种写法，
	// 与前面用 _ = iota 丢弃 0 行的效果相同，但更直观。
	SeasonSpring Season = iota + 1 // 从1开始
	// SeasonSummer 夏季
	SeasonSummer
	// SeasonAutumn 秋季
	SeasonAutumn
	// SeasonWinter 冬季
	SeasonWinter
)

// String 实现 Stringer 接口
//
// ⚠️ 注意：这里数组第 0 个元素是空字符串 ""，因为 Season 从 1 开始，第 0 位故意留空对齐。
func (s Season) String() string {
	return [...]string{"", "Spring", "Summer", "Autumn", "Winter"}[s]
}

// ========== 文件大小枚举 ==========

const (
	_  = iota             // 0 丢弃
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB = 1 << (10 * iota) // 1 << 20 = 1048576
	GB = 1 << (10 * iota) // 1 << 30 = 1073741824
	TB = 1 << (10 * iota) // 1 << 40
)

// ========== 辅助函数 ==========

// GetWeekdayValue 返回指定周几的枚举值
func GetWeekdayValue(day int) int {
	if day < 0 || day > 6 {
		return -1
	}
	return [...]int{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}[day]
}

// GetStatusName 返回状态名称
func GetStatusName(status int) string {
	switch status {
	case StatusPending:
		return "Pending"
	case StatusRunning:
		return "Running"
	case StatusCompleted:
		return "Completed"
	case StatusFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// CheckPermission 检查权限标志
//
// ⚠️ 注意：位标志判断的"黄金写法"是 permissions & flag != 0，
// 不是 == flag（后者要求其他位都为 0，组合权限时会误判）。
func CheckPermission(permissions, flag int) bool {
	return permissions&flag != 0
}

// CombinePermissions 组合权限标志
//
	// ⚠️ 注意：变参 ...int 把多个标志 OR 起来。对应 Java 里用 EnumSet 的 addAll 或位掩码 flags | flags2 | flags3。
func CombinePermissions(flags ...int) int {
	result := 0
	for _, flag := range flags {
		result |= flag
	}
	return result
}

// FormatFileSize 格式化文件大小
//
// ⚠️ 注意：用 switch {} 裸 switch 做"区间匹配"，从大到小判断（先 TB 再 GB...）。
// 顺序不能反（如果先判断 KB，大文件会被错误归类）。div 用 float64 避免整数除法丢精度。
func FormatFileSize(bytes int64) string {
	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}
