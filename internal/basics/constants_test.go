package basics_test

import (
	"testing"

	"gojuniper/internal/basics"
)

func TestConstantBasic(t *testing.T) {
	// 测试单个常量
	if basics.Pi != 3.14159 {
		t.Errorf("Pi = %f, want 3.14159", basics.Pi)
	}
	if basics.AppName != "GoJuniper" {
		t.Errorf("AppName = %s, want GoJuniper", basics.AppName)
	}
	if basics.MaxRetries != 3 {
		t.Errorf("MaxRetries = %d, want 3", basics.MaxRetries)
	}
}

func TestConstantGroup(t *testing.T) {
	// 测试HTTP状态码常量组
	tests := []struct {
		name  string
		got   int
		want  int
	}{
		{"StatusOK", basics.StatusOK, 200},
		{"StatusCreated", basics.StatusCreated, 201},
		{"StatusNotFound", basics.StatusNotFound, 404},
		{"StatusError", basics.StatusError, 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestConstantConfig(t *testing.T) {
	host, port, timeout := basics.GetDefaultConfig()
	if host != "localhost" {
		t.Errorf("DefaultHost = %s, want localhost", host)
	}
	if port != 8080 {
		t.Errorf("DefaultPort = %d, want 8080", port)
	}
	if timeout != 30 {
		t.Errorf("TimeoutSec = %d, want 30", timeout)
	}
}

func TestUntypedConstant(t *testing.T) {
	// 未类型常量可以赋值给不同类型
	var intVal int = basics.UntypedInt
	var int64Val int64 = basics.UntypedInt
	var floatVal float64 = basics.UntypedInt

	if intVal != 100 {
		t.Errorf("intVal = %d, want 100", intVal)
	}
	if int64Val != 100 {
		t.Errorf("int64Val = %d, want 100", int64Val)
	}
	if floatVal != 100 {
		t.Errorf("floatVal = %f, want 100", floatVal)
	}
}

func TestTypedConstant(t *testing.T) {
	// 类型化常量有明确类型
	if basics.TypedInt != 42 {
		t.Errorf("TypedInt = %d, want 42", basics.TypedInt)
	}
	if basics.TypedString != "typed" {
		t.Errorf("TypedString = %s, want typed", basics.TypedString)
	}
}

func TestConstantExpression(t *testing.T) {
	secMin, secHour, secDay := basics.GetTimeConstants()

	if secMin != 60 {
		t.Errorf("SecondsPerMinute = %d, want 60", secMin)
	}
	if secHour != 3600 {
		t.Errorf("SecondsPerHour = %d, want 3600", secHour)
	}
	if secDay != 86400 {
		t.Errorf("SecondsPerDay = %d, want 86400", secDay)
	}
}

func TestGetConstantValues(t *testing.T) {
	name, pi, retries := basics.GetConstantValues()

	if name != "GoJuniper" {
		t.Errorf("AppName = %s, want GoJuniper", name)
	}
	if pi != 3.14159 {
		t.Errorf("Pi = %f, want 3.14159", pi)
	}
	if retries != 3 {
		t.Errorf("MaxRetries = %d, want 3", retries)
	}
}

func TestGetHTTPStatus(t *testing.T) {
	ok, created, notFound, err := basics.GetHTTPStatus()

	if ok != 200 {
		t.Errorf("StatusOK = %d, want 200", ok)
	}
	if created != 201 {
		t.Errorf("StatusCreated = %d, want 201", created)
	}
	if notFound != 404 {
		t.Errorf("StatusNotFound = %d, want 404", notFound)
	}
	if err != 500 {
		t.Errorf("StatusError = %d, want 500", err)
	}
}