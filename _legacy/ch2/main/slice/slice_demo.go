package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main() {
	/*
		1.Go 语言切片 Slice
			Go 语言切片是对数组的抽象。
			Go 数组的长度不可改变，在特定场景下这样的结合就不太适合，Go中提供了一种灵活，功能强悍你的内置类型切片（“动态数组”）
			与数组相比切片的长度是不固定的，可以追加元素，再追加时可能使切片的容量增大。

		2.定义切片
			你可以声明一个未指定大小的数组来定义切片：
				var identifier []type
			切片不需要说明长度，或 使用 make() 函数来创建切片
				var slice1 []type = make([]type, len)
			也可以简写为：
				slice1 := make([]type,len)
			也可以指定容量，其中 capacity 为可选参数
				make([]type,length,capacity)
				这里 len 是数组的长度并且也是切片的初始长度。
		3. 切片初始化
			s := [] int{1,2,3}
			直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

			s := arr[:]
			初始化切片 s，是数组 arr 的引用。

			s := arr[startIndex:endIndex]
			将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

			s := arr[startIndex:]
			默认 endIndex 时将表示一直到arr的最后一个元素。

			s := arr[:endIndex]
			默认 startIndex 时将表示从 arr 的第一个元素开始。

			s1 := s[startIndex:endIndex]
			通过切片 s 初始化切片 s1。

			s :=make([]int,len,cap)
			通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。
	*/

	fmt.Println("===================len() 和 cap() 函数==============")
	/*
		4. len() 和 cap() 函数
			切片是可索引的，并且可以由 len() 方法获取长度。
			切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
	*/
	var numbers2 = make([]int, 3, 5)
	printSlice(numbers2)

	fmt.Println("=====================空 nil 切片===================")
	/*
		5. 空 nil 切片
			一个切片在未初始化之前默认为 nil，长度为 0
	*/
	var numbers3 []int
	printSlice(numbers3)
	if numbers3 == nil {
		fmt.Printf("numbers3 is nil\n")
	}

	fmt.Println("=====================切片截取===================")

	/*
		6. 切片截取
			可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
	*/

	numbers4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(numbers4)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers4)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers4[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers4[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers4[4:])

	numbers5 := make([]int, 0, 5)
	printSlice(numbers5)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers4[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers4[2:5]
	printSlice(number3)

	fmt.Println("=====================切片截取==end===================")

	fmt.Println("===================== append() 和 copy() 函数======================")

	/*
		7. append() 和 copy() 函数
			如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	*/

	var numbers9 []int
	printSlice(numbers9)

	/* 允许追加空切片 */
	numbers9 = append(numbers9, 0)
	printSlice(numbers9)

	/* 向切片添加一个元素 */
	numbers9 = append(numbers9, 1)
	printSlice(numbers9)

	/* 同时添加多个元素 */
	numbers9 = append(numbers9, 2, 3, 4)
	printSlice(numbers9)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers91 := make([]int, len(numbers9), (cap(numbers9))*2)

	/* 拷贝 numbers9 的内容到 numbers91 */
	copy(numbers91, numbers9)
	printSlice(numbers91)

	fmt.Println("===================== append() 和 copy() 函数====end==================")

	// 创建一个长度和容量均为 5 的整数切片
	slice := make([]int, 5)
	fmt.Println(slice) // 输出: [0 0 0 0 0]

	// 创建一个长度为 3，容量为 5 的整数切片
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2) // 输出: [0 0 0]

	/*
		Go语言数组声明需要制定元素的类型和个数 语法格式
			var arrayName [size]dataType

		以下定义了数组 balance 长度为 10 类型为 float32：var balance [10]float32
	*/

	var numbers [5]int
	fmt.Println(numbers)

	var numbers1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	numbers6 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers6)

	/*
		在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int 和 [10]int 是不同的类型。
	*/

	// 如果数组的长度不确定 可以使用 ... 来代替数组的长度 编译器会根据元素的个数自行推断
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)

	/*
		如果设置了数组的长度，我们还可以通过指定下标来初始化元素
	*/
	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance3)

	/*
		访问数组元素
	*/
	var n [10]int
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	/*
		示例
	*/

	balance4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出数组元素 */
	for i := 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance4[i])
	}

	balance5 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j := 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance5[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance6 := [5]float32{1: 2.0, 3: 7.0}
	for k := 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance6[k])
	}
}
