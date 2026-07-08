// Package basics 提供变量练习示例
package basics

import "sync"

// ========== 包级变量（导出） ==========

// GlobalCounter 全局计数器（导出，可被其他包访问）
//
// ⚠️ 注意：包级变量在程序启动时初始化（零值或显式初值），生命周期贯穿整个程序。
// 多处并发读写会有 data race——这里没加锁，练习时不要用它做并发计数（见下方 SafeIncrement）。
var GlobalCounter int

// ========== 包级变量（未导出） ==========

var (
	// internalCounter 内部计数器（小写开头，仅包内可见）
	//
	// ⚠️ 注意：Go 的"可见性"由【标识符首字母大小写】决定，不是 pub/private 关键字。
	// 大写 = 导出（public），小写 = 包内私有（private）。这是 Go 唯一的作用域控制手段。
	internalCounter int
	// internalFlag 内部标志
	internalFlag bool
)

// ========== 变量组声明 ==========

var (
	// AppConfig 应用配置（包级变量组）
	//
	// ⚠️ 注意：可以在 var (...) 块里直接写一个【匿名结构体字面量】作为类型，
	// 不需要先 type 定义。这种"一次性"结构很常见，但外部包无法引用它的类型（匿名）。
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
//
// ⚠️ 注意：下面这些字段的"零值"是什么，决定了你声明变量后不用初始化也能用的安全边界：
//   - 指针/切片/map/chan/func/interface 的零值都是 nil（map 的 nil 不能写入！）
//   - 数值 0、字符串 ""、布尔 false
//
// 复习重点：Struct 整体可以用 ZeroValues{} 得到全零值，但含 slice/map 的 struct 不能用 == 比较。
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
//
// ⚠️ 注意：:= 是"短变量声明"，只能在【函数内部】使用，且编译器自动推断类型。
// 它等价于 var x int = 10 的简写。函数外必须用 var，不能用 :=。
func DeclareWithShort() (int, string, bool) {
	x := 10
	name := "Go"
	flag := true
	return x, name, flag
}

// ReassignWithShort 演示短变量声明的重新赋值特性
//
// ⚠️ 注意（Go 独有的微妙规则）：
//
//	x := 10
//	x, y := 20, 30
//
// 第二行 x 已经存在，但 y 是【新变量】，:= 允许这种"部分重新声明"——只要左侧至少有一个新变量。
	// x 会被重新赋值（20），y 被声明（30）。这容易和 Java 的变量重定义混淆，但行为不同：这里 x 是同一个变量被改值，不是新建绑定（Java 不允许同作用域重复声明同名变量）。
func ReassignWithShort() (int, int) {
	x := 10
	x, y := 20, 30
	return x, y
}

// ========== 变量初始化 ==========

// InitVariables 演示变量初始化
//
	// ⚠️ 注意：Go 没有 Java 那样的"静态初始化块/构造器依赖"复杂机制（Java 有 static{} 块和类加载顺序），
// 但有特殊的 init() 函数（本文件没写）。包级变量在这里手动重置，仅作演示。
func InitVariables() {
	GlobalCounter = 0
	internalCounter = 0
	internalFlag = false
}

// ========== 并发安全的计数器 ==========

var (
	// ⚠️ 注意：mu 和 safeCounter 是包级变量，多个 goroutine 共享。
	// 必须用 sync.Mutex 保护，否则 go test -race 会报 data race。
	mu          sync.Mutex
	safeCounter int
)

// SafeIncrement 安全地增加计数器
//
// ⚠️ 注意：defer mu.Unlock() 保证即使函数 panic 也会释放锁。
// 这是 Go 里"获取锁/释放锁"的标准写法——永远用 defer 释放，避免忘记 Unlock 导致死锁。
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
//
// ⚠️ 注意：Go 支持"平行赋值" a, b = b, a，一行完成交换，不需要临时变量。
// 等号右边会先全部求值，再赋给左边。这是 Go 比很多语言优雅的小语法糖。
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
