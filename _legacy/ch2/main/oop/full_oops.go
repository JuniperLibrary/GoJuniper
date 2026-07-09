package main

import "fmt"

// Vehicle 基类
type Vehicle struct {
	Brand string
}

func (v *Vehicle) Start() {
	fmt.Println(v.Brand, "started")
}

// Car 派生类
type Car struct {
	Vehicle // 嵌入Vehicle
	Model   string
}

// Start 重写Start方法
func (c *Car) Start() {
	fmt.Println(c.Brand, c.Model, "car started")
}

func main() {
	/*
		完整的继承
	*/
	v := Vehicle{Brand: "Toyota"}
	c := Car{
		Vehicle: Vehicle{Brand: "Honda"},
		Model:   "Civic",
	}

	v.Start()         // Toyota started
	c.Start()         // Honda Civic car started
	c.Vehicle.Start() // Honda started
}
