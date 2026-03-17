package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		Go 语言接口
			接口 interface 时Go 语言中的一种类型。用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约。
			Go 语言提供了另外一种数据类型即接口。
			它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
			Go的接口设计简单却功能强大，是实现多态和解耦的重要工具。
			接口可一让我们将不同的类型绑定到一组公工的方法上，从而实现多态和灵活的设计。

		接口的特点
			1. 隐式实现
				Go 中没有关键字显式声明某个类型实现了某个接口。
				只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了接口
			2. 接口类型变量
				接口变量可以存储实现该接口的任意值
				接口变量实际上包含了两个部分：
					动态类型：存储实际的值类型
					动态值：存储具体的值
			3. 零值接口
				接口的零值是nil
				一个未初始化的接口变量其值为nil，且不包含任何动态类型或值
			4. 空接口
				定义为 interface{} 可以表示任何类型
			5. 接口的常见用法
				1. 多态：不同类型实现同一接口，实现多态行为
				2. 解耦：通过接口定义依赖关系，降低模块之间的耦合
				3. 泛化：使用空接口 interface {} 表示任意类型

		接口定义和实现
			接口使用关键字 interface 其中包含方法声明：
				// 定义接口
				type interface_name interface {
					method_name1 [return_type]
					method_name2 [return_type]
					method_name3 [return_type]
					...
					method_namen [return_type]
				}

				// 定义结构体
				type struct_name struct {
					// variables
				}

				// 实现接口方法
				func (struct_name_variable struct_name) method_name1() [return_type] {
					// 方法实现
				}
				...
				func (struct_name_variable struct_name) method_namen() [return_type] {
					// 方法实现
				}
	*/

	c := Circler{Radius: 5}
	var s Shaper = c // 接口变量可以存储实现了接口类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	/*
		空接口
			空接口 interface {} 是 Go 的特殊接口，表示所有类型的超集。
				任意类型都实现了接口
				常用于需要存储任意类型数据的场景，如泛型容器、通用参数
	*/
	printValuer(42)          // int
	printValuer("hello")     // string
	printValuer(3.14)        // float64
	printValuer([]int{1, 2}) // slice

	/*
		类型断言
			类型断言用于从接口类型中提取其底层值。
			基本语法：
				value := iface.(Type)
			iface 是接口变量
			Type 是要断言的具体类型
			如果类型不匹配，会出发panic
	*/

	var i interface{} = "hello"
	str := i.(string) // 类型断言
	fmt.Println(str)  // 输出：hello

	// 为了避免 panic 奔溃 可以使用带检查的类型断言
	s2, ok := i.(string)
	// ok 是一个布尔 ，如果断言失败 value为零值 ok为false
	fmt.Println(s2, ok)

	var i2 interface{} = 42
	if str, ok := i2.(string); ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("Not a string")
	}

	/*
		类型选择
			type switch 是 go中的语法结构，用于根据接口变量的具体类型执行不同的逻辑
	*/
	printTyper(42)
	printTyper("hello")
	printTyper(3.14)
	printTyper([]int{1, 2, 3})

	/*
		接口组合
			接口可以通过嵌套组合 实现更复杂的行为描述
	*/

	var rw ReaderAndWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello go")
}

/*
1. 定义一个简单的接口
Shaper 是一个接口，定义两个方法，Area 和 Perimeter；任意类型只要实现了这两个方法，就被认为实现了 Shaper 接口
*/
type Shaper interface {
	Area() float64
	Perimeter() float64
}

// Circler 定义一个结构体
type Circler struct {
	Radius float64
}

func (c Circler) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circler) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printValuer(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func printTyper(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}

type ReaderComplex interface {
	Read() string
}

type WriterComplex interface {
	Write(data string)
}

type ReaderAndWriter interface {
	ReaderComplex
	WriterComplex
}

type File struct {
}

func (f File) Read() string {
	return "hello"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)

}
