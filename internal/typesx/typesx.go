// Package typesx 提供“类型系统”相关的练习：
// - struct 与方法（值接收者/指针接收者）
// - 组合（embedding）
// - 接口与实现（fmt.Stringer）
package typesx

import (
	"errors"
	"fmt"
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
