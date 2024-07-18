package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := 10, 2
	c, err := div(a, b)
	fmt.Println(c, err)

	x := 100
	test(x)

	splits := make([]int, 0, 5)
	for i := 0; i < 9; i++ {
		splits = append(splits, i)
	}
	fmt.Println(splits)

	maps := make(map[string]int)
	maps["a"] = 1
	maps["b"] = 2
	maps["c"] = 3
	maps["d"] = 4
	maps["e"] = 5
	item, ok := maps["a"]
	fmt.Println(item, ok)

	delete(maps, "a")
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("div err")
	}
	return a / b, nil
}

func test(x int) func() {
	return func() {
		print(x)
	}
}
