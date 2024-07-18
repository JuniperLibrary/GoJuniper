package main

func MeSwitchDemo() {
	x := 0
	switch {
	case x > 0:
		print(x)
	case x < 0:
		print(-x)
	default:
		print(0)
	}
}
