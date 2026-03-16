package main

import "fmt"

/**
函数作为另外一个函数的实参
*/
// 定义一个计算函数类型
type calcFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func compute(a, b int, f calcFunc) int {
	return f(a, b)
}

func main() {
	// 将 add 函数 作为实参传入 compute
	result := compute(7, 8, add)
	fmt.Println(result)

	c1 := counter()
	fmt.Println(c1)
	fmt.Println(c1)

	c2 := counter()
	fmt.Println(c2)

	tom := Person{Name: "Tom", Age: 24}
	tom.GrowUp()
	fmt.Println(tom)
}

/*
返回一个闭包，该闭包会累积计数
*/
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) GrowUp() {
	p.Age++
}
