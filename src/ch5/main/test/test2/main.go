package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 求数组所有元素之和
func main() {
	//a := [...]int{1, 5, 10}
	//
	//sum := 0
	//
	//for i := 0; i < len(a); i++ {
	//	sum += a[i]
	//}
	//fmt.Println(sum)

	//plus

	var arr [10]int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10000) //取值范围[1-1000)
	}
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
