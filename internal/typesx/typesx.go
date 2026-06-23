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
