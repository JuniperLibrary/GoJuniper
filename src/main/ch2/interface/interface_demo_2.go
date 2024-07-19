package main

import "fmt"

// Shape 接口定义了计算形状面积的方法。
// 任何实现此接口的类型都可以计算其面积。
type Shape interface {
	area() float64
}

// Rectangle 结构体表示一个矩形。
// width 和 height 分别表示矩形的宽和高。
type Rectangle struct {
	width  float64
	height float64
}

// area 计算矩形的面积。
// 返回矩形的面积，即宽度乘以高度。
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Circle 结构体表示一个圆。
// radius 表示圆的半径。
type Circle struct {
	radius float64
}

// area 计算圆的面积。
// 返回圆的面积，即π乘以半径的平方。
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	// 使用 Rectangle 类型的实例初始化 Shape 接口变量
	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())

	// 使用 Circle 类型的实例初始化 Shape 接口变量
	s = Circle{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}
