package tests

import (
	"errors"
	"testing"

	"gojuniper/internal/04-typesx"
)

// typesx 包的测试聚焦于“类型系统 + 工程写法”的几个常见点：
// - 构造函数做参数校验，并返回可识别的哨兵错误
// - 指针接收者方法用于修改对象自身
// - embedding（组合）让外层类型复用内层字段与方法
func TestNewUser(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		_, err := typesx.NewUser(0, "alice")
		if !errors.Is(err, typesx.ErrInvalidID) {
			t.Fatalf("expected ErrInvalidID, got %v", err)
		}
	})

	t.Run("empty name", func(t *testing.T) {
		_, err := typesx.NewUser(1, "   ")
		if !errors.Is(err, typesx.ErrEmptyName) {
			t.Fatalf("expected ErrEmptyName, got %v", err)
		}
	})

	t.Run("trims name", func(t *testing.T) {
		u, err := typesx.NewUser(1, "  alice  ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.Name != "alice" {
			t.Fatalf("name=%q, want %q", u.Name, "alice")
		}
	})
}

func TestUser_SetName(t *testing.T) {
	u, err := typesx.NewUser(1, "alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// SetName 是指针接收者方法：会修改 u 本身。
	if err := u.SetName("  bob "); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Name != "bob" {
		t.Fatalf("name=%q, want %q", u.Name, "bob")
	}
}

func TestAdmin_Embedding(t *testing.T) {
	u, err := typesx.NewUser(1, "alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	a := typesx.Admin{User: u, Level: 10}

	// embedding 的效果：可以直接访问 a.Name / a.Greeting() 等。
	if a.Name != "alice" {
		t.Fatalf("name=%q, want %q", a.Name, "alice")
	}
	if !a.IsSuper() {
		t.Fatalf("expected super admin")
	}
}

func TestShaper(t *testing.T) {
	c := typesx.Circle{Radius: 5}
	r := typesx.Rectangle{Width: 3, Height: 4}

	var s typesx.Shaper

	s = c
	if s.Area() != 78.53981633974483 {
		t.Fatalf("Circle Area: got %v, want %v", s.Area(), 78.53981633974483)
	}
	if s.Perimeter() != 31.41592653589793 {
		t.Fatalf("Circle Perimeter: got %v, want %v", s.Perimeter(), 31.41592653589793)
	}

	s = r
	if s.Area() != 12 {
		t.Fatalf("Rectangle Area: got %v, want 12", s.Area())
	}
	if s.Perimeter() != 14 {
		t.Fatalf("Rectangle Perimeter: got %v, want 14", s.Perimeter())
	}
}

func TestTypeAssertString(t *testing.T) {
	got, ok := typesx.TypeAssertString("hello")
	if !ok || got != "hello" {
		t.Fatalf(`TypeAssertString("hello") = %q, %v; want "hello", true`, got, ok)
	}

	_, ok = typesx.TypeAssertString(42)
	if ok {
		t.Fatalf(`TypeAssertString(42) ok = %v; want false`, ok)
	}
}

func TestTypeSwitch(t *testing.T) {
	tests := []struct {
		val  interface{}
		want string
	}{
		{42, "int"},
		{"go", "string"},
		{3.14, "float64"},
		{true, "unknown"},
	}
	for _, tc := range tests {
		if got := typesx.TypeSwitch(tc.val); got != tc.want {
			t.Errorf("TypeSwitch(%v) = %q, want %q", tc.val, got, tc.want)
		}
	}
}

func TestFileReadWriter(t *testing.T) {
	var r typesx.Reader = typesx.File{}
	if got := r.Read(); got != "hello" {
		t.Fatalf("File.Read() = %q, want %q", got, "hello")
	}

	var w typesx.Writer = typesx.File{}
	w.Write("test")

	var rw typesx.ReadWriter = typesx.File{}
	if got := rw.Read(); got != "hello" {
		t.Fatalf("ReadWriter.Read() = %q, want %q", got, "hello")
	}
}

// ===== 补录（参考《Learning Go》第二版）：方法集 / 接口 nil 陷阱 / 类型断言 panic / 空标识符 / strings.Builder =====

func TestAssignDescriber(t *testing.T) {
	u := typesx.User{ID: 1, Name: "Alice"}
	// AssignDescriber 演示：*User 满足 Describer，User 值不满足（方法集差异）
	if !typesx.AssignDescriber(u) {
		t.Error("AssignDescriber should return true for valid pointer assignment")
	}
}

func TestReturnsNilError(t *testing.T) {
	// ReturnsNilError 返回含 nil 指针的非 nil 接口 —— 最隐蔽的接口坑
	err := typesx.ReturnsNilError()
	if err == nil {
		t.Fatal("ReturnsNilError() should return non-nil interface (contains nil *NilError)")
	}

	// 正确的做法：ReturnsRealNil 返回真正的 nil
	realNil := typesx.ReturnsRealNil()
	if realNil != nil {
		t.Error("ReturnsRealNil() should return nil")
	}
}

func TestMustTypeAssertString(t *testing.T) {
	// 安全形式
	result := typesx.MustTypeAssertString("hello")
	if result != "hello" {
		t.Errorf("MustTypeAssertString(\"hello\") = %q, want \"hello\"", result)
	}

	// 危险形式：类型不匹配 => panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustTypeAssertString(123) should panic for non-string")
		}
	}()
	typesx.MustTypeAssertString(123) // 这里会 panic
}

func TestDiscardValue(t *testing.T) {
	result := typesx.DiscardValue()
	if result != "kept" {
		t.Errorf("DiscardValue() = %q, want \"kept\"", result)
	}
}

func TestJoinWithBuilder(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  string
	}{
		{"empty", []string{}, ""},
		{"single", []string{"hello"}, "hello"},
		{"multiple", []string{"a", "b", "c"}, "abc"},
		{"with spaces", []string{"hello", " ", "world"}, "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := typesx.JoinWithBuilder(tt.parts)
			if got != tt.want {
				t.Errorf("JoinWithBuilder(%v) = %q, want %q", tt.parts, got, tt.want)
			}
		})
	}
}
