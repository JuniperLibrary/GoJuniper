package main

import "fmt"

func MeforDemo() {
	// 基础 for 循环（类 C 语言风格）
	for i := 0; i < 5; i++ {
		fmt.Printf("当前数字：%d\n", i)
	}

	// while 风格循环（省略初始化 / 后置语句）
	count := 0
	for count < 3 {
		fmt.Println("count= %d,continue。。", count)
		count++
	}
	fmt.Println("循环结束")

	// 无限循环 省略所有条件
	num := 0
	// 无限循环
	for {
		num++
		fmt.Printf("num = %d\n", num)
		// 当num等于4时终止循环
		if num == 4 {
			break
		}
	}
	fmt.Println("无限循环已终止")

	// for-range 循环（遍历集合） 专门用于遍历数组、切片、字符串 map、channel
	fruits := []string{"apple", "peach", "pear"}
	for index, value := range fruits {
		fmt.Printf("fruits[%d] = %s\n", index, value)
	}

	// 只遍历 value
	fmt.Println("只遍历 value")

	for _, value := range fruits {
		fmt.Println(value)
	}

	// 遍历字符串
	str := "Golang language is very cool."
	// for-range 会自动处理 Unicode 字符 按字符 遍历 ，而不是字节
	for i, c := range str {
		fmt.Printf("索引：%d，字符：%c（Unicode码：%d）\n", i, c, c)
	}

	// 对比 普通 for 按照自己遍历 中文占 3字节 会乱码
	fmt.Println("\n按字节遍历（错误示例）：")
	for i := 0; i < len(str); i++ {
		fmt.Printf("索引：%d，字节：%c\n", i, str[i])
	}

	// 遍历 map
	score := map[string]int{
		"小明": 90,
		"小红": 95,
		"小刚": 88,
	}
	// 遍历map（key：键，value：值，顺序不固定）
	for name, s := range score {
		fmt.Printf("%s的分数：%d\n", name, s)
	}

	// 循环控制
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("偶数：%d\n", i)
	}

	for i := 1; i <= 10; i++ {
		if i > 5 {
			fmt.Printf("找到目标：%d", i)
			break
		}
	}

	// 嵌套循环
	for i := 1; i <= 9; i++ {
		for j := i; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
		print(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			print(i + 1)
		}
	}

	for i := 4; i >= 0; i-- {
		print(i)
	}

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range x {
		print(i, ":", v)
	}

}
