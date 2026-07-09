package basics

import "fmt"

func main() {

	fmt.Print("请输入你的名字: ")
	var name string
	_, _ = fmt.Scan(&name)

	fmt.Println("你的名字是:", name)

	fmt.Print("请输入你的年龄: ")

	var age int
	n, err := fmt.Scan(&age)

	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	fmt.Println("读取参数:", n)
	fmt.Println("年龄:", age)
}
