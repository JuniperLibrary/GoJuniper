package main

import (
	"fmt"
	"time"
)

func main() {
	go task(1) // 启动一个协程
	go task(2)

	time.Sleep(time.Second * 6)
}

func task(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : %d\n", id, i)
		fmt.Println("------------------")
		time.Sleep(time.Second)
	}
}
