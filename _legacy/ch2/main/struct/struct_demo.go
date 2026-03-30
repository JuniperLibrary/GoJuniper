package main

import "fmt"

// user 结构体代表一个用户，包含姓名和年龄信息。
// 结构体(struct)可匿名嵌入其他类型。
type user struct {
	name string
	age  int
}

// manager 结构体代表一个管理者，它嵌入了 user 结构体，增加了职称信息。
// 通过嵌入 user 结构体，manager 继承了 user 的所有字段和方法。
type manager struct {
	user
	title string
}

type BookStore struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	// 初始化一个 manager 实例
	//var m manager
	m := manager{} // 使用默认值初始化manager类型的变量m

	// 设置 manager 实例的姓名、年龄和职称
	m.name = "Tom"
	m.age = 29
	m.title = "Go"

	// 输出 manager 实例的信息
	fmt.Println(m)

	fmt.Println("================Go 语言结构体======================")
	/*
			Go 语言结构体
				Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
				结构体是由一系列具有相同类型或者不同类型的数据构成的数据集合。
				结构体表示一项记录，比如保存图书馆的书籍记录，每本书有不同的属性。

			定义结构体
				结构体定义需要使用 type 和 struct 语句。
					struct 语句定义一个新的数据类型，结构提中有一个或多个成员。
					type   语句设定了结构体的名称。
			结构体的格式如下：
				type struct_variable_type struct{
					member definition
					member definition
					...
		   			member definition
				}
			一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
				variable_name := structure_variable_type {value1, value2...valuen}
				或
				variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
	*/
	// 创建一个新的结构体
	fmt.Println(BookStore{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(BookStore{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(BookStore{title: "Go 语言", author: "www.runoob.com"})

	fmt.Println("=================访问结构体成员==============")

	/*
		访问结构体成员
			如果要访问结构体成员，需要使用点号 . 操作符，格式为：
				结构体.成员名"
	*/

	var Book1 BookStore /* 声明 Book1 为 BookStore 类型 */
	var Book2 BookStore /* 声明 Book2 为 BookStore 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.bookId)

	fmt.Println("================结构体作为函数参数==============")
	/* 打印 Book1 信息 */
	printBookStore(Book1)

	/* 打印 Book2 信息 */
	printBookStore(Book2)

	fmt.Println("================结构体指针==================")
	/*
		结构体指针
			你可以定义指向结构体的指针类似与其他指针变量，格式如下：
				var struct_pointer *BookStore
			以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
				struct_pointer = &BookStore
			使用结构体指针访问结构体成员，使用 "." 操作符：
				struct_pointer.title
	*/

	/* 打印 Book1 信息 */
	printBookStoreWithPointer(&Book1)

	/* 打印 Book2 信息 */
	printBookStoreWithPointer(&Book2)

}

func printBookStore(book BookStore) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}

func printBookStoreWithPointer(book *BookStore) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
