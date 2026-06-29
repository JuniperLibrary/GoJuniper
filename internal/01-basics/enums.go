// Package basics 提供枚举练习示例
package basics

import "fmt"

// ========== 基础 iota 枚举 ==========

const (
	// Weekday 周几枚举（基础用法）
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

// ========== 跳过 0 的枚举 ==========

const (
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
	PermissionRead   = 1 << iota // 1 (二进制: 001)
	PermissionWrite              // 2 (二进制: 010)
	PermissionExecute            // 4 (二进制: 100)
)

const (
	// FileMode 文件模式标志
	ModeUserRead   = 1 << (iota * 3) // 1   (用户读)
	ModeUserWrite                   // 8   (用户写)
	ModeUserExec                    // 64  (用户执行)
	ModeGroupRead                   // 512 (组读)
	ModeGroupWrite                  // 4096 (组写)
	ModeGroupExec                   // 32768 (组执行)
)

// ========== 自定义类型枚举 ==========

// Direction 方向枚举类型
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
	SeasonSpring Season = iota + 1 // 从1开始
	// SeasonSummer 夏季
	SeasonSummer
	// SeasonAutumn 秋季
	SeasonAutumn
	// SeasonWinter 冬季
	SeasonWinter
)

// String 实现 Stringer 接口
func (s Season) String() string {
	return [...]string{"", "Spring", "Summer", "Autumn", "Winter"}[s]
}

// ========== 文件大小枚举 ==========

const (
	_  = iota                      // 0 丢弃
	KB = 1 << (10 * iota)          // 1 << 10 = 1024
	MB = 1 << (10 * iota)          // 1 << 20 = 1048576
	GB = 1 << (10 * iota)          // 1 << 30 = 1073741824
	TB = 1 << (10 * iota)          // 1 << 40
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
func CheckPermission(permissions, flag int) bool {
	return permissions&flag != 0
}

// CombinePermissions 组合权限标志
func CombinePermissions(flags ...int) int {
	result := 0
	for _, flag := range flags {
		result |= flag
	}
	return result
}

// FormatFileSize 格式化文件大小
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