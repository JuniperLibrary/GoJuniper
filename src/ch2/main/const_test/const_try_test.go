package main_test

import "testing"

// 定义一周中三天的常量，从星期一到星期三
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

// 定义文件权限的常量，分别为可读、可写和可执行
const (
	Readable = 1 << iota
	Writeable
	Executable
)

// TestConstantTry 测试定义的星期常量是否正确
func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

// TestConstantTryByte 测试文件权限常量以及位运算的使用
func TestConstantTryByte(t *testing.T) {
	// a变量用于演示位运算中与操作的效果
	//a := 7 //0111
	a := 1 //0001
	t.Log(Readable, Writeable, Executable)
	// 检查a是否具有指定的文件权限
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
