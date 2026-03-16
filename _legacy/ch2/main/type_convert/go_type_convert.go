package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Go 语言类型转换
			类型转换用于将一种数据类型的变量转换为另外一种类型的变量。
			Go 语言类型转换基本格式如下：
				type_name(expression)
			type_name 为类型，expression 为表达式。
	*/

	/*
		1. 数值类型转换
			将整数型转换为浮点型
	*/
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	/*
		2. 字符串类型转换
			将一个字符串转换成另一个类型，
	*/
	var str string = "10"
	num, err := strconv.Atoi(str)
	fmt.Println(num, err)
	/*
		以上代码将字符串变量 str 转换为整型变量 num。
		注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，
		我们可以使用空白标识符 _ 来忽略这个错误
	*/

	str2 := "123"
	num2, err2 := strconv.Atoi(str2)
	if err2 != nil {
		fmt.Println("转换错误:", err2)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str2, num2)
	}

	num3 := 123
	str3 := strconv.Itoa(num3)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num3, str3)

	str4 := "3.14"
	num4, err4 := strconv.ParseFloat(str4, 64)
	if err4 != nil {
		fmt.Println("转换错误:", err4)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str4, num4)
	}

	num5 := 3.14
	str5 := strconv.FormatFloat(num5, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num5, str5)

	/*
		3. 接口类型转换
			接口类型转换有两种情况：类型断言和类型转换。
	*/

	/*
		3.1 类型断言
				类型断言用于将接口类型转换为指定类型，其语法为：
					value.(type)
					或者
					value.(T)
				其中 value 是接口类型的变量，type 或 T 是要转换成的类型。
				如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功
	*/
	var i interface{} = "Hello world"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}

	/*
		3.2 类型转换
			类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：
				T(value)
			T 是目标接口类型，value 是要转换的值。
			在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。
	*/

	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}
	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)
	// 修改 StringWriter 的字段
	sw.str = "Hello, World"
	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)

	/*
		解析：
			1. 定义接口和结构体：
				1.1 Writer 接口定义了 Write 方法。
				1.2 StringWriter 结构体实现了 Write 方法。
			2. 类型转换：
				2.1 将 StringWriter 实例赋值给 Writer 接口变量 w。
				2.2 使用 w.(*StringWriter) 将 Writer 接口类型转换为 StringWriter 类型。
			3.访问字段：
				3.1 修改 StringWriter 的字段 str，并打印其值。
	*/

	/*
		4. 空接口类型
			空接口 interface{} 可以持有任何类型的值。在实际应用中，空接口经常被用来处理多种类型的值
	*/

	printValue(42)
	printValue("hello")
	printValue(3.14)

	// 在这个例子中，printValue 函数接受一个空接口类型的参数，并使用类型断言和类型选择来处理不同的类型。
}

// Writer 定义一个接口 Writer
type Writer interface {
	Write([]byte) (int, error)
}

// StringWriter 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}
