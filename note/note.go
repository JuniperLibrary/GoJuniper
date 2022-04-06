package note

import "fmt"

//常量
const (
	//首字母大写 是一个公共的常量
	Version int = 100
	num     int = 100
)

func SayHello() {
	fmt.Print("不同包")
}

//变量和常量
func VariablesAndConstants() {
	//1
	var v1 int = 100
	fmt.Print(v1)
	//2
	var v2 int
	fmt.Print(v2)
	//3
	v3 := 100
	fmt.Print(v3)
	//4
	var v4 string
	fmt.Print(v4)
	//5
	v5 := "abc"
	fmt.Print(v5)
}

//基本数据类型
func BasicDataTypes() {
	fmt.Println("\n 整数型")
	var (
		n1 = 1
		//int8 最大128
		n2 int8 = 127
		n3 uint
	)
	fmt.Printf("n1=%v,type is %T\n", n1, n1)
	fmt.Printf("n2=%v,type is %T\n", n2, n2)
	fmt.Printf("n3=%v,type is %T\n", n3, n3)

	fmt.Println("\n 浮点型")
	var (
		f1         = 1.0
		f2 float32 = 1
		f3 float64
	)
	fmt.Printf("f1=%v,type is %T\n", f1, f1)
	fmt.Printf("f2=%v,type is %T\n", f2, f2)
	fmt.Printf("f3=%v,type is %T\n", f3, f3)

	fmt.Println("\n 数值型数据类型转换")
	n2 = int8(n3)
	fmt.Printf("n2=%v,type is %T\n", n2, n2)

	fmt.Println("\n 字符型")
	var (
		c1 byte
		c2      = '0'
		c3 rune = 23454
	)
	fmt.Printf("c1=%v,type is %T\n", c1, c1)
	fmt.Printf("c2=%v,type is %T\n", c2, c2)
	fmt.Printf("c3=%v,type is %T\n", c3, c3)
	c4 := 'A' - 'a'
	fmt.Printf("c4=%v,type is %T\n", c4, c4)
	c5 := "你好牛"
	fmt.Printf("c5%v,type is %T\n", c5, c5)

	fmt.Println("\n 布尔型")
	var (
		bool1 bool = true
	)
	fmt.Printf("boo1=%v,type is %T\n", bool1, bool1)

	fmt.Println("\n 字符串")
	var (
		s1 = "你好牛"
		s2 = " "
	)
	s3 := "暴力熊"
	fmt.Println(s1, s2, s3)
	fmt.Println("hello", s3)
	fmt.Println(len(s3))

}

//指针
//值拷贝
// func increase(n int) {
// 	//n= n+1
// 	n++
// 	fmt.Printf("n of address naddr=%v\n", &n) //n的取地址
// 	fmt.Printf("increase end of n=%v\n", n)   //n=2022
// }

// func Pointer() {
// 	var src = 2022
// 	increase(src)
// 	fmt.Printf("src of address naddr=%v\n", &src)      //取src地址
// 	fmt.Printf("increase called of end src=%v\n", src) //2022
// }

//如果想用引用传就可以--值传递
// func increaseUp(n *int) {
// 	//n= n+1
// 	*n++
// 	fmt.Printf("n of address naddr=%v\n", &n) //n的取地址
// 	fmt.Printf("increase end of n=%v\n", *n)   //n=2023
// }

// func Pointer1() {
// 	var src = 2022
// 	var ptr = &src
// 	increaseUp(ptr)
// 	fmt.Printf("src of address ptraddr=%v\n", &src)
// 	fmt.Printf("ptr of address ptraddr=%v\n", &ptr)      //取src地址
// 	fmt.Printf("increase called of end ptr=%v\n", *ptr) //2023
// }

func increaseUp(n *int) {
	//n= n+1
	*n++
	fmt.Printf("n of address naddr=%v\n", &n) //n的取地址
	fmt.Printf("increase end of n=%v\n", *n)  //n=2023
}

func Pointer1() {
	var src = 2022
	increaseUp(&src)
	fmt.Printf("src=%v,src of address ptraddr=%v\n", src, &src)
	var ptr = new(int)
	fmt.Printf("ptr=%v\n,ptr of address ptr_addr=%v\n,increse end of ptr=%v\n", ptr, &ptr,*ptr) //n的取地址
}

//fmt格式字符
func FmtVerbs(){
	fmt.Println("fmt格式字符")
	//1.
	fmt.Println("通用字符")
	 //%v 值
	 //%T 数据类型
	 
	//2.整数
	 //%d 十进制（没有前缀）
	 //%b 二进制（没有前缀）
	 //%o 八进制（没有前缀）
	 //%x 十六进制a-f（没有前缀）
	 //%X 十六进制A-F（没有前缀）
	 //%U U+四位16进制int32

}
