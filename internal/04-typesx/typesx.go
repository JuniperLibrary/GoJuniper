// Package typesx 提供“类型系统”相关的练习：
// - struct 与方法（值接收者/指针接收者）
// - 组合（embedding）
// - 接口与实现（fmt.Stringer）
package typesx

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	// ErrInvalidID 表示 ID 不合法。
	ErrInvalidID = errors.New("id must be > 0")
	// ErrEmptyName 表示 name 为空或只有空白字符。
	ErrEmptyName = errors.New("name must not be empty")
)

// User 表示一个简单的用户模型。
type User struct {
	ID   int
	Name string
}

// NewUser 构造一个 User，并做最基本的参数校验。
func NewUser(id int, name string) (User, error) {
	if id <= 0 {
		return User{}, ErrInvalidID
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return User{}, ErrEmptyName
	}
	return User{ID: id, Name: name}, nil
}

// Greeting 返回一个简单问候语（值接收者：不会修改 User）。
func (u User) Greeting() string {
	return fmt.Sprintf("hello, %s", u.Name)
}

// String 让 User 实现 fmt.Stringer 接口，方便 fmt.Println/Printf 打印。
func (u User) String() string {
	return fmt.Sprintf("User{ID:%d, Name:%q}", u.ID, u.Name)
}

// SetName 修改 User 的 Name（指针接收者：会修改原对象）。
func (u *User) SetName(name string) error {
	if u == nil {
		return errors.New("nil receiver")
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyName
	}
	u.Name = name
	return nil
}

// Admin 演示 struct embedding：Admin “拥有” User 的字段与方法。
type Admin struct {
	User
	Level int
}

// IsSuper 判断管理员等级是否达到“超级管理员”。
func (a Admin) IsSuper() bool {
	return a.Level >= 10
}

// Shaper 定义几何形状的行为：计算面积和周长。
type Shaper interface {
	Area() float64
	Perimeter() float64
}

// Circle 表示圆形。
type Circle struct {
	Radius float64
}

// Area 计算圆的面积。
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆的周长。
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle 表示矩形。
type Rectangle struct {
	Width, Height float64
}

// Area 计算矩形面积。
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长。
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// TypeAssertString 对 interface{} 进行类型断言，尝试提取字符串值。
func TypeAssertString(val interface{}) (string, bool) {
	s, ok := val.(string)
	return s, ok
}

// TypeSwitch 使用类型选择判断 interface{} 的底层类型。
func TypeSwitch(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	default:
		return "unknown"
	}
}

// Reader 定义读取行为。
type Reader interface {
	Read() string
}

// Writer 定义写入行为。
type Writer interface {
	Write(string)
}

// ReadWriter 通过接口组合同时描述读写行为。
type ReadWriter interface {
	Reader
	Writer
}

// File 是一个空结构体，同时实现 Reader 和 Writer。
type File struct{}

func (f File) Read() string {
	return "hello"
}

func (f File) Write(data string) {
	// 模拟写入，实际不执行 I/O。
}

// ========== 补录：方法集（method set）规则 ==========

// Describer 是一个仅由指针接收者方法实现的接口。
//
// ⚠️ **注意（方法集是接口满足与否的决定因素）**：
// Go 规定——值类型 T 的方法集只包含【值接收者】方法；
// 指针类型 *T 的方法集包含【值接收者 + 指针接收者】所有方法。
// 因此：如果某个接口的方法只被 *T 的指针接收者实现，那么 *T 满足该接口，但 T（值）不满足。
// 这是 Go 接口最经典、最易错的坑，直接决定你写的类型能不能当作接口参数传递。
type Describer interface {
	Describe() string
}

// Describe 用【指针接收者】实现 Describer。
//
// ⚠️ **注意**：因为 Describe 是指针接收者，它只属于 *User 的方法集，
// 不属于 User（值）的方法集。所以 var d Describer = &u 合法，但 var d Describer = u 编译报错。
func (u *User) Describe() string {
	return fmt.Sprintf("user#%d name=%s", u.ID, u.Name)
}

// AssignDescriber 尝试把一个 User 值赋给 Describer 接口变量。
// 它故意返回能否成功，用来在测试里演示"值不满足、指针满足"的差异。
//
// ⚠️ **注意**：本函数无法直接写 `var d Describer = u`（编译错误），
// 所以改用反射判断——但结论不变：User 值的方法集不含 Describe，*User 才有。
func AssignDescriber(u User) (asPtr bool) {
	// 值类型 User 无法满足 Describer（Describe 是指针接收者），所以只有指针版本可行。
	var d Describer = &u // 合法：*User 满足 Describer
	_ = d
	return true
}

// ========== 补录：接口 nil 陷阱 ==========

// NilError 是一个演示用的自定义错误类型，实现了 error 接口。
//
// ⚠️ **注意**：用来演示"接口 nil 陷阱"的辅助类型。
type NilError struct {
	Msg string
}

func (e *NilError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Msg
}

// ReturnsNilError 返回一个【指针为 nil 的 *NilError】，但装入 error 接口后接口本身非 nil。
//
// ⚠️ **注意（接口 nil 陷阱，最常见面试题）**：
// 接口变量 = (具体类型, 具体值)。即使具体值是 nil（如 *NilError(nil)），
// 只要"类型"不为 nil（这里是 *NilError），整个接口就不等于 nil。
// 下面这个函数返回 `*NilError` 类型的 nil，装入 error 接口后，
// 接口的"动态类型"是 *NilError、"动态值"是 nil —— 接口整体 != nil！
// 调用方写 `if err != nil` 会【误判为 true】，导致 bug。正确做法：直接返回 nil（不加一层指针包装）。
func ReturnsNilError() error {
	var e *NilError // e 是 nil 指针，类型是 *NilError
	return e        // 装入 error 接口：类型=*NilError, 值=nil => 接口非 nil
}

// ReturnsRealNil 正确示范：需要"无错误"时直接返回 untyped nil。
func ReturnsRealNil() error {
	return nil // 接口类型和值都是 nil => 接口整体 == nil
}

// ========== 补录：类型断言 panic 对照 ==========

// MustTypeAssertString 直接对 interface{} 做类型断言（不带 ok）。
//
// ⚠️ **注意**：不带 ok 的类型断言 `val.(string)` 在底层实际类型不是 string 时会【直接 panic】。
// 它适合你"确信类型一定匹配"的场景；不确定时务必用 `s, ok := val.(string)`（见 TypeAssertString）。
// 这个函数是危险用法的演示：调用方负责保证类型正确，否则运行时崩溃。
func MustTypeAssertString(val interface{}) string {
	return val.(string) // 类型不符 => panic
}

// ========== 补录：空标识符 `_` 的用法 ==========

// DiscardValue 演示空标识符 `_` 忽略不需要的返回值。
//
// ⚠️ **注意（空标识符 `_` 三种常见用法）**：
//  1. 忽略某个返回值：`_, v := m[k]`（只要 value，丢弃 ok）
//  2. 忽略 import 的副作用：`import _ "pkg"`（只执行包的 init，不引用其标识符）
//  3. 在 for-range 里忽略索引或值：`for _, v := range xs`（只要值）
//
// 这里演示第 1 种：调用 MultiReturn 但只取第二个值，第一个用 `_` 丢弃。
func DiscardValue() string {
	_, v := MultiReturn() // 忽略第一个返回值，保留第二个
	return v
}

// MultiReturn 返回两个字符串，供 DiscardValue 演示丢弃。
func MultiReturn() (string, string) {
	return "ignored", "kept"
}

// ========== 补录：strings.Builder 高效拼接 ==========

// JoinWithBuilder 用 strings.Builder 拼接字符串切片。
//
// ⚠️ **注意（性能：别用 + 循环拼接）**：
// Go 的字符串是不可变的。`s += piece` 每次都分配新内存并复制，循环 N 次是 O(N²)。
// strings.Builder 复用内部 []byte 缓冲区，循环拼接是 O(N)，大数据量下差异巨大。
// 教材《Learning Go》也强调：拼接多个字符串优先用 strings.Builder（或 strings.Join）。
func JoinWithBuilder(parts []string) string {
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}
