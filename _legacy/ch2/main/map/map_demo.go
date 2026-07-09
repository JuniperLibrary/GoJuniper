package main

import "fmt"

func main() {
	/*
		Go 语言 Map 集合
			Map 是一种无序的键值对的集合。
			Map 最重要的一点是通过key来快速检索数据，key类似与索引，指向数据的值。
			Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map是无序的，遍历Map时返回的键值对的顺序时不固定的。
			在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。
			Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，
			因此对 Map 的修改会影响到所有引用它的变量。

		1. 定义 Map
			可以使用内建函数 make 或使用 map 关键字来定义 Map:
				// 使用 make 函数
				map_variable := make(map[KeyType]ValueType, initialCapacity)
			其中 KeyType 是键的类型，ValueType 是值的类型，initialCapacity 是可选的参数，用于指定 Map 的初始容量。
			Map 的容量是指 Map 中可以保存的键值对的数量，当 Map 中的键值对数量达到容量时，
			Map 会自动扩容。如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。

	*/

	// 创建一个空的字符串到整数的映射
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m) // 输出: map[one:1 two:2]

	// 创建一个初始容量为 10 的映射
	m2 := make(map[string]int, 10)
	fmt.Println(m2) // 输出: map[]

	/*
		2. 也可以使用字面量创建 Map：
	*/

	m3 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(m3)
	// 获取键值对
	v1 := m3["apple"]
	fmt.Println(v1)
	v2, ok := m3["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
	fmt.Println(ok, v2)

	// 修改元素
	m3["apple"] = 5
	fmt.Println(m3)
	fmt.Println(len(m3))
	delete(m3, "apple")

	for k, v := range m3 {
		fmt.Println(k, v)
	}

	/*
		示例
	*/
	var siteMap map[string]string /*创建集合 */
	siteMap = make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	/*使用键输出地图值 */
	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}

	/*
		3. delete 函数
			delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
	*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")
	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
