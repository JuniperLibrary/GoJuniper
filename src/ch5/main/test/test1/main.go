package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成一个随机数，让用户去猜这个数是多少？
func main() {
	var number int

	// 通过rand函数的Seed方法，来设置总值，这里我们以当前时间来设置总值，并且用的纳秒，十分精确了
	rand.Seed(time.Now().UnixNano())
	// 随机数的范围是0-100，但不包括100
	number = rand.Intn(100)
	fmt.Printf("猜猜一个数字，数字的范围是：[0-100]\n")

	// 死循环
	for {
		var input int
		// Scanf表示让用户输入，Scanf从终端读取一个整数，并传值给input变量，
		// &表示获取到该变量的内存地址
		fmt.Scanf("%d\n", &input)
		var flag bool = false

		switch {
		case number > input:
			fmt.Printf("您输入的数字太小了\n")
		case number == input:
			fmt.Printf("您猜对了\n")
			flag = true
		case number < input:
			fmt.Printf("您输入的数字太大了\n")
		}

		if flag {
			break
		}
	}
}
