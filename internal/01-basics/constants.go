// Package basics 提供常量练习示例
package basics

// ========== 单个常量 ==========

// Pi 圆周率
const Pi = 3.14159

// AppName 应用名称（字符串常量）
const AppName = "GoJuniper"

// MaxRetries 最大重试次数（整型常量）
const MaxRetries = 3

// ========== 常量组 ==========

const (
	// HTTP 状态码
	StatusOK       = 200
	StatusCreated  = 201
	StatusNotFound = 404
	StatusError    = 500
)

const (
	// 配置相关常量
	DefaultHost = "localhost"
	DefaultPort = 8080
	TimeoutSec  = 30
)

// ========== 未类型常量 ==========

// UntypedInt 未类型整数常量，可以赋值给不同类型的整数变量
//
// ⚠️ 注意（Go 常量最反直觉的特性之一）：
// Go 的"未类型常量"（untyped constant）没有固定类型，赋值时会根据上下文适配。
// 比如 const X = 100 可以赋给 int、int64、float64 都行。
// 但一旦用在需要具体类型的表达式中，才会被定型（default type：整数→int，浮点→float64，字符串→string）。
const UntypedInt = 100

// UntypedFloat 未类型浮点常量
const UntypedFloat = 3.14

// ========== 类型化常量 ==========

const (
	// TypedInt 类型化常量，有明确类型
	//
	// ⚠️ 注意：类型化常量（TypedInt int = 42）不能再赋给 int64 等不同类型，会编译报错。
	// 未类型常量更灵活，类型化常量更安全（避免隐式转换 bug）。
	TypedInt int = 42
	// TypedString 类型化字符串常量
	TypedString string = "typed"
)

// ========== 常量表达式 ==========

const (
	// 可以在常量中使用表达式
	//
	// ⚠️ 注意：const 块里可以引用前面定义的常量做计算（常量表达式在编译期求值）。
	// SecondsPerHour 引用了 SecondsPerMinute，这叫"常量间的依赖"。
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * 60
	SecondsPerDay    = SecondsPerHour * 24
)

// GetConstantValues 返回常量值用于测试
func GetConstantValues() (string, float64, int) {
	return AppName, Pi, MaxRetries
}

// GetHTTPStatus 返回HTTP状态码常量
func GetHTTPStatus() (int, int, int, int) {
	return StatusOK, StatusCreated, StatusNotFound, StatusError
}

// GetDefaultConfig 返回默认配置常量
func GetDefaultConfig() (string, int, int) {
	return DefaultHost, DefaultPort, TimeoutSec
}

// GetTimeConstants 返回时间相关常量
func GetTimeConstants() (int, int, int) {
	return SecondsPerMinute, SecondsPerHour, SecondsPerDay
}
