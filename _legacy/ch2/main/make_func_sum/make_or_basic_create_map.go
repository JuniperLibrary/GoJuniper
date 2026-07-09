package main

import "fmt"

func main() {
	// 方式1：创建并初始化有数据的map（最常用的原始方式）
	userInfo := map[string]string{
		"name": "张三",
		"age":  "20",
		"city": "北京",
	}
	fmt.Println(userInfo) // 输出：map[age:20 city:北京 name:张三]

	// 方式2：创建空map（仅定义结构，无数据）
	emptyMap := map[int]string{}
	fmt.Println(emptyMap)      // 输出：map[]
	fmt.Println(len(emptyMap)) // 输出：0（长度为0）
	emptyMap[1] = "苹果"         // 后续可添加数据
	fmt.Println(emptyMap)      // 输出：map[1:苹果]

	// 注意：仅声明不初始化的map是nil，不能直接赋值（会panic）
	var nilMap map[string]int // 只是声明，未初始化，值为nil
	// nilMap["a"] = 100 // 运行报错：panic: assignment to entry in nil map
	fmt.Println(nilMap == nil) // 输出：true

	// 方式2：创建空map，不指定容量
	m1 := make(map[int]string)
	m1[1] = "香蕉"
	m1[2] = "橙子"
	fmt.Println(m1) // 输出：map[1:香蕉 2:橙子]

	// 方式2：创建空map，指定初始容量（推荐大数据场景）
	m2 := make(map[string]int, 10) // 初始容量10，减少后续扩容开销
	m2["apple"] = 5
	m2["banana"] = 8
	fmt.Println(m2) // 输出：map[apple:5 banana:8]
	//fmt.Println(cap(m2))   // 注意：map无法直接获取容量，len()获取当前元素数
	fmt.Println(len(m2)) // 输出：2
}
