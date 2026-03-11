package main

import (
	"fmt"

	"gojuniper/internal/basics"
)

func main() {
	// cmd/hello 是一个最小可运行的入口，演示：
	// - 导入项目内部包（internal/*）
	// - 运行一个简单的 main
	fmt.Println("sum:", basics.Sum(1, 2))
	fmt.Println("fizzbuzz(15):", basics.FizzBuzz(15))
}
