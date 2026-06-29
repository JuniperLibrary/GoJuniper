// Package basics 提供变量练习示例
package basics

import "sync"

// ========== 包级变量（导出） ==========

// GlobalCounter 全局计数器（导出，可被其他包访问）
var GlobalCounter int

// ========== 包级变量（未导出） ==========

var (
	// internalCounter 内部计数器（小写开头，仅包内可见）
	internalCounter int
	// internalFlag 内部标志
	internalFlag bool
)

// ========== 变量组声明 ==========

var (
	// AppConfig 应用配置（包级变量组）
	AppConfig = struct {
		Host string
		Port int
		Mode string
	}{
		Host: "localhost",
		Port: 8080,
		Mode: "debug",
	}
)

// ========== 零值演示 ==========

// ZeroValues 演示各种类型的零值
type ZeroValues struct {
	Int     int
	Int64   int64
	Float   float64
	Bool    bool
	String  string
	Pointer *int
	Slice   []int
	Map     map[string]int
	Chan    chan int
	Func    func()
}

// GetZeroValues 返回一个零值结构体，用于演示Go的零值特性
func GetZeroValues() ZeroValues {
	return ZeroValues{}
}

// ========== 短变量声明示例 ==========

// DeclareWithShort 使用短变量声明
func DeclareWithShort() (int, string, bool) {
	x := 10
	name := "Go"
	flag := true
	return x, name, flag
}

// ReassignWithShort 演示短变量声明的重新赋值特性
func ReassignWithShort() (int, int) {
	x := 10
	x, y := 20, 30
	return x, y
}

// ========== 变量初始化 ==========

// InitVariables 演示变量初始化
func InitVariables() {
	GlobalCounter = 0
	internalCounter = 0
	internalFlag = false
}

// ========== 并发安全的计数器 ==========

var (
	mu          sync.Mutex
	safeCounter int
)

// SafeIncrement 安全地增加计数器
func SafeIncrement() {
	mu.Lock()
	defer mu.Unlock()
	safeCounter++
}

// SafeDecrement 安全地减少计数器
func SafeDecrement() {
	mu.Lock()
	defer mu.Unlock()
	safeCounter--
}

// GetSafeCounter 获取安全计数器的值
func GetSafeCounter() int {
	mu.Lock()
	defer mu.Unlock()
	return safeCounter
}

// ResetSafeCounter 重置安全计数器
func ResetSafeCounter() {
	mu.Lock()
	defer mu.Unlock()
	safeCounter = 0
}

// ========== 变量交换 ==========

// SwapVariables 交换两个变量的值
func SwapVariables() (int, int) {
	a, b := 1, 2
	a, b = b, a
	return a, b
}

// ========== 内部变量访问（用于测试） ==========

// GetInternalCounter 获取内部计数器值（仅供测试使用）
func GetInternalCounter() int {
	return internalCounter
}

// SetInternalCounter 设置内部计数器值（仅供测试使用）
func SetInternalCounter(v int) {
	internalCounter = v
}

// GetInternalFlag 获取内部标志值（仅供测试使用）
func GetInternalFlag() bool {
	return internalFlag
}

// SetInternalFlag 设置内部标志值（仅供测试使用）
func SetInternalFlag(v bool) {
	internalFlag = v
}