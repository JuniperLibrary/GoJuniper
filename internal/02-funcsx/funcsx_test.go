package funcsx

import (
	"errors"
	"testing"
)

func TestCounter(t *testing.T) {
	c := Counter()
	if c() != 1 || c() != 2 || c() != 3 {
		t.Error("Counter closure should increment: 1,2,3")
	}
	// a fresh closure has its own captured state
	d := Counter()
	if d() != 1 {
		t.Error("fresh Counter should start at 1")
	}
}

func TestRecoverFromPanic(t *testing.T) {
	ok := func() {}
	if err := RecoverFromPanic(ok); err != nil {
		t.Errorf("no panic expected, got %v", err)
	}
	boom := func() { panic("kaboom") }
	err := RecoverFromPanic(boom)
	if err == nil {
		t.Fatal("expected recovered error")
	}
	if !errors.Is(err, errors.New("panicked: kaboom")) && err.Error() != "panicked: kaboom" {
		// errors.Is won't match a new instance; just check message
		if err.Error() != "panicked: kaboom" {
			t.Errorf("unexpected err message: %v", err)
		}
	}
}

func TestSumVariadic(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-1, 1}, 0},
	}
	for _, tt := range tests {
		if got := Sum(tt.nums...); got != tt.want {
			t.Errorf("Sum(%v)=%d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestDeferInspect(t *testing.T) {
	// defer registers in LIFO order; final log is "body" + reversed defers
	got := DeferInspect([]int{1, 2, 3})
	want := "bodydefer(3)defer(2)defer(1)"
	if got != want {
		t.Errorf("DeferInspect([1,2,3])=%q, want %q", got, want)
	}
}

func TestApplyFunc(t *testing.T) {
	add := func(a, b int) int { return a + b }
	if got := ApplyFunc(2, 3, add); got != 5 {
		t.Errorf("ApplyFunc=%d, want 5", got)
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 1}, {1, 1}, {5, 120}, {10, 3628800},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d)=%d, want %d", tt.n, got, tt.want)
		}
	}
}
