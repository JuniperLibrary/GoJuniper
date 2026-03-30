package main

import "fmt"

// Animal 父结构体
type Animal struct {
	Name string
}

// Speak 父结构体的方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// Dog 子结构体
type Dog struct {
	Animal // 嵌入 Animal 结构体
	Breed  string
}

func main() {
	/*
		Go 继承
			在面向对象编程（OOP）中，继承是一种机制，允许一个类（子类）从另一个类（父类）继承属性和方法。通过继承，子类可以复用父类的代码，并且可以在不修改父类的情况下扩展或修改其行为。
			Go 语言并不是一种传统的面向对象编程语言，它没有类和继承的概念。
			Go 使用结构体（struct）和接口（interface）来实现类似的功能。

		Go 中的 "继承"
			Go 语言没有传统面向对象语言中的类(class)和继承(inheritance)概念，而是通过组合(composition)和接口(interface)来实现类似的功能。
	*/

	/*
		1. 组合（Composition）
			组合是 Go 中实现代码复用的主要方式。通过将一个结构体嵌入到另一个结构体中，子结构体可以"继承"父结构体的字段和方法。
	*/
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}
	dog.Speak() // 调用父结构体的方法
	fmt.Println("Breed:", dog.Breed)

	/*
		代码解释
			Animal 是父结构体，包含一个字段 Name 和一个方法 Speak。
			Dog 是子结构体，通过嵌入 Animal 结构体，继承了 Animal 的字段和方法。
			在 main 函数中，我们创建了一个 Dog 实例，并调用了 Speak 方法。
	*/

	/*
		2.接口 （Interface）
			接口是 Go 中实现多态的主要方式。通过定义接口，不同的结构体可以实现相同的方法，从而实现类似继承的多态行为。
	*/

	var speaker SpeakerOops

	dogInterface := DogOops{
		AnimalOops: AnimalOops{Name: "Buddy"},
		Breed:      "Golden Retriever",
	}

	speaker = &dogInterface
	speaker.Speak() // 通过接口调用方法

	/*
		Go 与经典继承的区别
			特性					经典继承					Go 的方式
			代码复用				通过继承					通过组合(嵌入结构体)
			多态					通过继承和方法重写			通过接口实现
			关系					"是一个"(is-a)关系		"有一个"(has-a)或"实现了"关系
			灵活性				继承关系固定				可以运行时组合
	*/
}

// SpeakerOops 定义接口
type SpeakerOops interface {
	Speak()
}

// AnimalOops 父结构体
type AnimalOops struct {
	Name string
}

// Speak SpeakOop 实现接口方法
func (a *AnimalOops) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// DogOops 子结构体
type DogOops struct {
	AnimalOops
	Breed string
}

/*
代码解释
SpeakerOops 是一个接口，定义了一个 Speak 方法。
AnimalOops 结构体实现了 SpeakOop 接口。
DogOops 结构体通过嵌入 AnimalOops 结构体，间接实现了 Speaker 接口。
在 main 函数中，我们将 Dog 实例赋值给 Speaker 接口，并通过接口调用 Speak 方法。
*/
