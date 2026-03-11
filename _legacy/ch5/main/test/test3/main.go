package main

import "fmt"

// 找出数组中和为给定值的两个元素的下标，比如数组：[1,3,5,7],找出两个元素之和等于8的下标分别是[0,4]和[1,2]
func main() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	var sum int = 12

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Printf("i=%d j=%d\n", i, j)
			}
		}
	}

}
