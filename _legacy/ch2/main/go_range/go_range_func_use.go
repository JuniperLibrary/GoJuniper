package main

/*
	Go 语言范围 range函数
		Go 语言中 range 关键字用于 for 循环中迭代数组array、切片slice 、通道channel或者集合 map的元素。
		在数组和切片中它返回元素的索引和索引对用的值，在集合中返回 key-value对。
		for 循环的 range 格式可以对 slice、map、数组、字符串等进行循环。
		格式如下
			for key , value := range oldMap{
				newMap[key] = value
			}

		以上代码中的 key 和 value 是可以省略。如果只想读取 key 格式如下：
			for key := range oldMap

		或者这样：
			for key,_ := range oldMap
		如果只想读取value ,格式如下：
			for _,value := range oldMap
*/
import "fmt"

// 声明一个包含 2 的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// 遍历 pow 切片，i 是索引，v 是值
	for i, v := range pow {
		// 打印 2 的 i 次方等于 v
		fmt.Printf("2**%d = %d\n", i, v)
	}
	/*
		2. 遍历字符串
			range 迭代字符串时，返回每个字符的索引和 Unicode 代码点（rune）
	*/
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}

	/*
		3. 映射 Map
			for 循环的 range 格式可以忽略key和value
	*/
	map1 := make(map[int]float32)

	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	/*
		4. 遍历 channel 通道
	*/

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

	//这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range 也可以用在 map 的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
