package condition_test

import (
	"fmt"
	"testing"
)

// if
func TestCondition(t *testing.T) {
	a := 10

	if a < 20 {
		t.Logf("a 小于 20\n")
	}
	t.Logf("a 的值为：%d\n", a)
}

// if ... else
func TestConditions(t *testing.T) {
	a := 100
	if a < 20 {
		fmt.Println("a < 20")
	} else {
		fmt.Println("a > 20")
	}
	fmt.Printf("a 的值为：%d\n", a)
}

// 嵌套if
func TestConditionCover(t *testing.T) {
	a := 100
	b := 200

	if a == 100 {
		if b == 200 {
			t.Log("a 的值为", a, "b的值为", b)
		}
	}
}
