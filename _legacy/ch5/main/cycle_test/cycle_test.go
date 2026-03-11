package cycle_test

import (
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func length(s string) int {
	println("call length")
	return len(s)
}

func TestForLoop(t *testing.T) {
	s := "abcd"

	for i, n := 0, length(s); i < n; i++ {
		println(i, s[i])
	}
}

func TestForLoop1(t *testing.T) {
	var a int = 15
	var b int

	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		t.Logf("a 的值为：%d\n", a)
	}

	for a > b {
		a++
		t.Logf("a 的值为：%d\n", a)
	}

	for i, x := range numbers {
		t.Logf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

// 循环嵌套
// 输出2到100的素数
func TestForLoopFor(t *testing.T) {
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			t.Logf(" %d 是素数\n", i)
		}
	}
}
