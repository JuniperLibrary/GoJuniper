package main

import (
	"fmt"
	"os"
)

// 功能实现
/**
main 函数不支持传入参数
在程序中直接通过os.Args获取命令行参数
*/
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Print("hello world,go", os.Args)
	}
	os.Exit(-1)
}
