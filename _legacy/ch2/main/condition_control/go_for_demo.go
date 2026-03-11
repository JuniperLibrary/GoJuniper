package main

func MeforDemo() {
	for i := 0; i < 10; i++ {
		print(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			print(i + 1)
		}
	}

	for i := 4; i >= 0; i-- {
		print(i)
	}

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range x {
		print(i, ":", v)
	}

}
