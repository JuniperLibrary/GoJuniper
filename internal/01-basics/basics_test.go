package basics

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
		{100, 200, 300},
	}
	for _, tt := range tests {
		if got := Sum(tt.a, tt.b); got != tt.want {
			t.Errorf("Sum(%d,%d)=%d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
		ok   bool
	}{
		{"empty", nil, 0, false},
		{"single", []int{42}, 42, true},
		{"normal", []int{3, 9, 1, 7}, 9, true},
		{"negatives", []int{-5, -1, -9}, -1, true},
	}
	for _, tt := range tests {
		got, ok := Max(tt.xs)
		if ok != tt.ok || got != tt.want {
			t.Errorf("Max(%v)=(%d,%v), want (%d,%v)", tt.xs, got, ok, tt.want, tt.ok)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	want := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	if len(got) != len(want) {
		t.Fatalf("FizzBuzz(15) len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("FizzBuzz(15)[%d]=%q, want %q", i, got[i], want[i])
		}
	}
	if FizzBuzz(0) != nil {
		t.Error("FizzBuzz(0) should be nil")
	}
	if FizzBuzz(-3) != nil {
		t.Error("FizzBuzz(-3) should be nil")
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false}, {1, false}, {2, true}, {3, true},
		{4, false}, {17, true}, {18, false}, {97, true},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d)=%v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestFactorialUint64(t *testing.T) {
	tests := []struct {
		n     int
		want  uint64
		isErr bool
	}{
		{-1, 0, true},
		{0, 1, false},
		{1, 1, false},
		{5, 120, false},
		{20, 2432902008176640000, false},
		{21, 0, true}, // 21! overflows uint64
	}
	for _, tt := range tests {
		got, err := FactorialUint64(tt.n)
		if tt.isErr {
			if !errors.Is(err, ErrOverflow) && !errors.Is(err, ErrNegativeN) {
				t.Errorf("FactorialUint64(%d) err=%v, want overflow/negative error", tt.n, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("FactorialUint64(%d) unexpected err=%v", tt.n, err)
		}
		if got != tt.want {
			t.Errorf("FactorialUint64(%d)=%d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacciUint64(t *testing.T) {
	got, err := FibonacciUint64(7)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	want := []uint64{0, 1, 1, 2, 3, 5, 8}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("FibonacciUint64(7)[%d]=%d, want %d", i, got[i], want[i])
		}
	}
	if _, err := FibonacciUint64(-1); !errors.Is(err, ErrNegativeN) {
		t.Error("FibonacciUint64(-1) should return ErrNegativeN")
	}
	// overflow check
	if _, err := FibonacciUint64(95); !errors.Is(err, ErrOverflow) {
		t.Error("FibonacciUint64(95) should overflow uint64")
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"你好世界", "界世好你"}, // multi-byte UTF-8 preserved
	}
	for _, tt := range tests {
		if got := ReverseString(tt.in); got != tt.want {
			t.Errorf("ReverseString(%q)=%q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"one", 1},
		{"the quick brown fox", 4},
		{"  extra   spaces  ", 2},
	}
	for _, tt := range tests {
		if got := CountWords(tt.in); got != tt.want {
			t.Errorf("CountWords(%q)=%d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestIotaDemo(t *testing.T) {
	a, b, c := IotaDemo()
	if a != 0 || b != 1 || c != 2 {
		t.Errorf("IotaDemo()=(%d,%d,%d), want (0,1,2)", a, b, c)
	}
}

func TestSwapByPointer(t *testing.T) {
	x, y := 1, 2
	SwapByPointer(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after swap x=%d y=%d, want 2 1", x, y)
	}
}

func TestTypeConvertDemo(t *testing.T) {
	i, s := TypeConvertDemo(3.9)
	if i != 3 {
		t.Errorf("int part=%d, want 3", i)
	}
	if s != "3.9" {
		t.Errorf("string=%q, want \"3.9\"", s)
	}
}

func TestMakeVsNew(t *testing.T) {
	mk, nw := MakeVsNew()
	if mk == "" || nw == "" {
		t.Errorf("MakeVsNew() returned empty: make=%q new=%q", mk, nw)
	}
}

func TestGetZeroValues(t *testing.T) {
	z := GetZeroValues()
	var zero ZeroValues
	if !reflect.DeepEqual(z, zero) {
		t.Errorf("GetZeroValues()=%+v, want all zero %+v", z, zero)
	}
}

func TestDeclareWithShort(t *testing.T) {
	x, name, flag := DeclareWithShort()
	if x != 10 || name != "Go" || !flag {
		t.Errorf("DeclareWithShort()=(%d,%q,%v), want (10,Go,true)", x, name, flag)
	}
}

func TestReassignWithShort(t *testing.T) {
	x, y := ReassignWithShort()
	// x := 10; x, y := 20, 30  => x=20, y=30 (at least one new var on LHS)
	if x != 20 || y != 30 {
		t.Errorf("ReassignWithShort()=(%d,%d), want (20,30)", x, y)
	}
}

func TestSwapVariables(t *testing.T) {
	a, b := SwapVariables()
	if a != 2 || b != 1 {
		t.Errorf("SwapVariables()=(%d,%d), want (2,1)", a, b)
	}
}

func TestSafeCounter(t *testing.T) {
	ResetSafeCounter()
	for i := 0; i < 100; i++ {
		SafeIncrement()
	}
	if got := GetSafeCounter(); got != 100 {
		t.Errorf("GetSafeCounter()=%d, want 100", got)
	}
	SafeDecrement()
	if got := GetSafeCounter(); got != 99 {
		t.Errorf("after decrement=%d, want 99", got)
	}
}

func TestGetConstantValues(t *testing.T) {
	name, pi, retries := GetConstantValues()
	if name != "GoJuniper" || pi != 3.14159 || retries != 3 {
		t.Errorf("GetConstantValues()=(%q,%v,%v)", name, pi, retries)
	}
	if Pi != 3.14159 || AppName != "GoJuniper" || MaxRetries != 3 {
		t.Error("exported constants mismatch")
	}
}

func TestIotaEnums(t *testing.T) {
	if Sunday != 0 || Saturday != 6 {
		t.Errorf("Weekday iota wrong: %d..%d", Sunday, Saturday)
	}
	if StatusPending != 1 || StatusFailed != 4 {
		t.Errorf("Status enum should start at 1: %d..%d", StatusPending, StatusFailed)
	}
	if PermissionRead != 1 || PermissionWrite != 2 || PermissionExecute != 4 {
		t.Errorf("Permission bits wrong: %d %d %d", PermissionRead, PermissionWrite, PermissionExecute)
	}
	if North != 0 || West != 3 {
		t.Errorf("Direction wrong: %d..%d", North, West)
	}
	if North.String() != "North" || West.String() != "West" {
		t.Error("Direction.String() wrong")
	}
	if ColorRed.String() != "Red" {
		t.Error("ColorRed.String() wrong")
	}
	r, g, b, a := ColorBlue.RGBA()
	if r != 0 || g != 0 || b != 255 || a != 255 {
		t.Errorf("ColorBlue.RGBA()=(%d,%d,%d,%d), want (0,0,255,255)", r, g, b, a)
	}
	if SeasonSpring != 1 || SeasonWinter != 4 {
		t.Errorf("Season should start at 1: %d..%d", SeasonSpring, SeasonWinter)
	}
	if KB != 1024 || MB != 1048576 || GB != 1073741824 {
		t.Errorf("File size consts wrong: KB=%d MB=%d GB=%d", KB, MB, GB)
	}
}

func TestCheckCombinePermissions(t *testing.T) {
	combined := CombinePermissions(PermissionRead, PermissionWrite)
	if !CheckPermission(combined, PermissionRead) {
		t.Error("combined should have Read")
	}
	if CheckPermission(combined, PermissionExecute) {
		t.Error("combined should not have Execute")
	}
}

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{512, "512 B"},
		{1024, "1.00 KB"},
		{1048576, "1.00 MB"},
		{1073741824, "1.00 GB"},
	}
	for _, tt := range tests {
		if got := FormatFileSize(tt.bytes); got != tt.want {
			t.Errorf("FormatFileSize(%d)=%q, want %q", tt.bytes, got, tt.want)
		}
	}
}

func TestMathMaxUint64Ref(t *testing.T) {
	if math.MaxUint64 != 18446744073709551615 {
		t.Error("unexpected MaxUint64")
	}
}

// ========== 新增函数测试 ==========

func TestNumericLiteralsDemo(t *testing.T) {
	dec, bin, oct, cmplx := NumericLiteralsDemo()
	if dec != 1_000_000 {
		t.Errorf("decimal with underscore = %d, want 1000000", dec)
	}
	if bin != 0b1010 {
		t.Errorf("binary = %d, want 10", bin)
	}
	if oct != 0o777 {
		t.Errorf("octal = %d, want 511", oct)
	}
	if cmplx != (1 + 2i) {
		t.Errorf("complex = %v, want (1+2i)", cmplx)
	}
}

func TestRawStringDemo(t *testing.T) {
	s := RawStringDemo()
	// 原始字符串保留换行、制表符、反斜杠、引号
	if len(s) == 0 {
		t.Error("RawStringDemo() should not be empty")
	}
	// 验证包含关键特征
	wantSubstrs := []string{
		"原始字符串",
		"换行",
		"制表符",
		`C:\Users\Name`,  // 反斜杠未转义
		`"key": "value"`, // 双引号未转义
	}
	for _, sub := range wantSubstrs {
		if !contains(s, sub) {
			t.Errorf("RawStringDemo() missing %q", sub)
		}
	}
}

func TestFloatPitfallsDemo(t *testing.T) {
	moneyWrong, directEq := FloatPitfallsDemo()
	// 陷阱演示：两者都应该为 false（展示浮点误差导致的错误结果）
	if moneyWrong {
		t.Error("moneyWrong should be false (0.1+0.2 != 0.3)")
	}
	if directEq {
		t.Error("directEq should be false (0.1+0.2 != 0.3 with ==)")
	}

	// 验证正确做法：epsilon 比较
	if !FloatEpsilonEqual(0.1+0.2, 0.3, 1e-9) {
		t.Error("FloatEpsilonEqual(0.1+0.2, 0.3, 1e-9) should be true")
	}
	if !FloatEpsilonEqual(1.0, 1.0, 1e-12) {
		t.Error("FloatEpsilonEqual(1.0, 1.0, 1e-12) should be true")
	}
	if FloatEpsilonEqual(1.0, 1.0001, 1e-9) {
		t.Error("FloatEpsilonEqual(1.0, 1.0001, 1e-9) should be false")
	}
}

// contains 辅助：判断字符串是否包含子串（用于测试 RawStringDemo）
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > len(sub) && (s[:len(sub)] == sub || contains(s[1:], sub)))
}
