package basics_test

import (
	"testing"

	"gojuniper/internal/basics"
)

// basics 包的测试以“入门可读性”为第一目标：
// - 先看测试理解函数的输入输出与边界条件
// - 再去看实现（basics.go）对照理解写法
func TestSum(t *testing.T) {
	// 入门测试：验证最简单的纯函数输出。
	if got := basics.Sum(1, 2); got != 3 {
		t.Fatalf("Sum(1,2)=%d, want 3", got)
	}
}

func TestMax(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		_, ok := basics.Max(nil)
		if ok {
			t.Fatalf("expected ok=false")
		}
	})

	t.Run("non-empty slice", func(t *testing.T) {
		got, ok := basics.Max([]int{3, 1, 10, -2})
		if !ok {
			t.Fatalf("expected ok=true")
		}
		if got != 10 {
			t.Fatalf("Max=%d, want 10", got)
		}
	})
}

func TestFizzBuzz(t *testing.T) {
	got := basics.FizzBuzz(15)
	if len(got) != 15 {
		t.Fatalf("len=%d, want 15", len(got))
	}
	if got[2] != "Fizz" {
		t.Fatalf("got[2]=%q, want %q", got[2], "Fizz")
	}
	if got[4] != "Buzz" {
		t.Fatalf("got[4]=%q, want %q", got[4], "Buzz")
	}
	if got[14] != "FizzBuzz" {
		t.Fatalf("got[14]=%q, want %q", got[14], "FizzBuzz")
	}
}

func TestIsPrime(t *testing.T) {
	// 表驱动测试：把输入/期望值放到一个表里，循环执行。
	tests := []struct {
		n    int
		want bool
	}{
		{n: -1, want: false},
		{n: 0, want: false},
		{n: 1, want: false},
		{n: 2, want: true},
		{n: 3, want: true},
		{n: 4, want: false},
		{n: 97, want: true},
		{n: 99, want: false},
	}
	for _, tt := range tests {
		if got := basics.IsPrime(tt.n); got != tt.want {
			t.Fatalf("IsPrime(%d)=%v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestFactorialUint64(t *testing.T) {
	if _, err := basics.FactorialUint64(-1); err == nil {
		t.Fatalf("expected error")
	}
	if got, err := basics.FactorialUint64(0); err != nil || got != 1 {
		t.Fatalf("FactorialUint64(0)=%d,%v want 1,nil", got, err)
	}
	if got, err := basics.FactorialUint64(5); err != nil || got != 120 {
		t.Fatalf("FactorialUint64(5)=%d,%v want 120,nil", got, err)
	}
}

func TestFibonacciUint64(t *testing.T) {
	if _, err := basics.FibonacciUint64(-1); err == nil {
		t.Fatalf("expected error")
	}
	got, err := basics.FibonacciUint64(10)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	want := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "", want: ""},
		{in: "abc", want: "cba"},
		{in: "你好Go", want: "oG好你"},
	}
	for _, tt := range tests {
		if got := basics.ReverseString(tt.in); got != tt.want {
			t.Fatalf("ReverseString(%q)=%q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	if got := basics.CountWords("  hello   world \n go "); got != 3 {
		t.Fatalf("CountWords=%d, want 3", got)
	}
}
