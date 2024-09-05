package main

// X 是一个基于int的自定义类型
type X int

// inc 方法用于增加X类型的值
// 它通过指针接收器来直接修改变量的值，而不是返回新的值
// 参数:
//   x *X - X类型的指针，用于直接修改变量的值
func (x *X) inc() {
	*x++
}

func main() {
	// 初始化X类型的变量x
	var x X
	// 调用inc方法增加x的值
	x.inc()
	// 输出x的值
	println(x)
}
