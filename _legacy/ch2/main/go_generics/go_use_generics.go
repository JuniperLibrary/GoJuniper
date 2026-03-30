package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	/*
		1.泛型语法详解
			类型参数声明
				泛型函数和类型通过类型参数列表来声明，语法类型为 [类型参数 约束]
				// 基本语法结构
				func 函数名[T 约束](参数 T) 返回值类型 {
					// 函数体
				}

				type 类型名[T 约束] struct {
					// 结构体字段
				}
			类型参数命名约定
				通常使用大写字母 T、K、V、E等
				T：表示 Type  类型
				K：表示 Key 键
				V：表示 Value 值
				E：表示 Element 元素
	*/

	/*
		2.约束 Constraints
			约束定义了类型参数必须满足的条件，是泛型的的核心概念。
				内置约束
					any 约束
						any 是空接口 interface{} 的别名，表示任何类型都可以。
	*/
	PrintAnyGenerics(42)
	PrintAnyGenerics("hello world")
	PrintAnyGenerics([]int{1, 2, 3})
	PrintAnyGenerics(3.14)

	/*
		3. comparable 约束
			comparable 表示类型支持 == 和 != 操作符
	*/
	// 使用示例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(FindIndex(numbers, 3)) // 输出: 2

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(FindIndex(names, "Bob")) // 输出: 1

	/*
		4.联合约束 Union Constraints
			使用  ｜ 与纳素u三
	*/
	// 使用示例
	fmt.Println(Add(10, 20))     // 输出: 30
	fmt.Println(Add(3.14, 2.71)) // 输出: 5.85
}

func PrintAnyGenerics[T any](value T) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// MaxInt 处理 int 类型的函数
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxFloat 处理 float64 类型的函数
func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Max 使用泛型的解决方案
func Max[T constraints.Ordered](a, b T) T {
	if b < a {
		return a
	}
	return b
}

func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// Number 数字类型约束
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}
