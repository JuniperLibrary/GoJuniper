package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	/*
			Go 语言递归函数
				递归时一种函数直接或间接调用的自身的变成技术。
				递归函数通常包含两个部分：
					1. 基准条件 Base Case ： 递归的终止条件，防止函数无限调用自身。
					2. 递归条件 Recursive Case: 这是函数调用自身的部分，用于将问题分解为更小的子问题。

				语法格式如下：
					func recursion() {
		  	 			recursion() // 函数调用自身
					}

					func main() {
						recursion()
					}
	*/

	/*
		1. 阶乘
			阶乘是一个正整数的乘积，表示为 n!。例如：
	*/
	fmt.Println(factorial(5))

	/*
		2. 斐波那契数列
	*/
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

	/*
		3. 求平方根
	*/
	x := 25.0
	result := sqrt(x)
	fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)

	/*
		4. 递归的常见应用
			递归在许多算法和数据结构中都有广泛应用，例如：
				树和图的遍历：如深度优先搜索（DFS）。
				分治算法：如归并排序、快速排序。
				动态规划：如斐波那契数列的计算。
	*/

	// 文件目录遍历
	walkDir(".", "")
}

// 递归函数计算阶乘
func factorial(n int) int {
	// 基准条件
	if n == 0 {
		return 1
	}
	// 递归条件
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}

func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	if diff := guess*guess - x; diff < epsilon && -diff > epsilon {
		return guess
	}
	newGuess := (guess + x/guess) / 2
	if newGuess < prevGuess {
		return guess
	}
	return sqrtRecursive(x, newGuess, guess, epsilon)
}

func walkDir(dir string, indent string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, entry := range entries {
		fmt.Println(indent + entry.Name())
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), indent+"  ")
		}
	}
}
