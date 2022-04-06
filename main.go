package main

/*
1.在不同包 调用函数
 1.1在同一个包不同文件调用sayHello
	在项目根目录编译 go build 绝对路径
	快速运行        go run /Users/wanglufei/Documents/GoStudy
 1.2 在不同包不同文件调用sayHello
	在项目根目录编译 go build 绝对路径
	单个主文件 go build
*/
import (
	"gostudy/note"
)

func main()  {
	//fmt.Print("hello Go")
	//同包
	// sayHello()
	//不同包
	//note.SayHello()

	//fmt.Print(note.Version)

	//note.VariablesAndConstants()
	// note.BasicDataTypes()
	// note.Pointer()
	note.Pointer1()
}