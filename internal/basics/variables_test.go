package basics_test

import (
	"sync"
	"testing"

	"gojuniper/internal/basics"
)

func TestZeroValues(t *testing.T) {
	zv := basics.GetZeroValues()

	// 测试各种类型的零值
	if zv.Int != 0 {
		t.Errorf("Int zero value = %d, want 0", zv.Int)
	}
	if zv.Int64 != 0 {
		t.Errorf("Int64 zero value = %d, want 0", zv.Int64)
	}
	if zv.Float != 0.0 {
		t.Errorf("Float zero value = %f, want 0.0", zv.Float)
	}
	if zv.Bool != false {
		t.Errorf("Bool zero value = %v, want false", zv.Bool)
	}
	if zv.String != "" {
		t.Errorf("String zero value = %s, want empty", zv.String)
	}
	if zv.Pointer != nil {
		t.Errorf("Pointer zero value = %v, want nil", zv.Pointer)
	}
	if zv.Slice != nil {
		t.Errorf("Slice zero value = %v, want nil", zv.Slice)
	}
	if zv.Map != nil {
		t.Errorf("Map zero value = %v, want nil", zv.Map)
	}
	if zv.Chan != nil {
		t.Errorf("Chan zero value = %v, want nil", zv.Chan)
	}
	if zv.Func != nil {
		// Function values have no meaningful zero value to print; just fail if non-nil.
		t.Errorf("Func zero value is non-nil, want nil")
	}
}

func TestShortVariableDeclaration(t *testing.T) {
	x, name, flag := basics.DeclareWithShort()

	if x != 10 {
		t.Errorf("x = %d, want 10", x)
	}
	if name != "Go" {
		t.Errorf("name = %s, want Go", name)
	}
	if flag != true {
		t.Errorf("flag = %v, want true", flag)
	}
}

func TestReassignWithShort(t *testing.T) {
	x, y := basics.ReassignWithShort()

	if x != 20 {
		t.Errorf("x = %d, want 20", x)
	}
	if y != 30 {
		t.Errorf("y = %d, want 30", y)
	}
}

func TestSwapVariables(t *testing.T) {
	a, b := basics.SwapVariables()

	if a != 2 {
		t.Errorf("a = %d, want 2", a)
	}
	if b != 1 {
		t.Errorf("b = %d, want 1", b)
	}
}

func TestGlobalCounter(t *testing.T) {
	// 重置计数器
	basics.GlobalCounter = 0

	// 测试全局变量的读写
	if basics.GlobalCounter != 0 {
		t.Errorf("GlobalCounter = %d, want 0", basics.GlobalCounter)
	}

	basics.GlobalCounter = 100
	if basics.GlobalCounter != 100 {
		t.Errorf("GlobalCounter = %d, want 100", basics.GlobalCounter)
	}
}

func TestInternalVariables(t *testing.T) {
	// 重置内部变量
	basics.SetInternalCounter(0)
	basics.SetInternalFlag(false)

	// 测试内部变量的设置和获取
	basics.SetInternalCounter(42)
	if got := basics.GetInternalCounter(); got != 42 {
		t.Errorf("internalCounter = %d, want 42", got)
	}

	basics.SetInternalFlag(true)
	if got := basics.GetInternalFlag(); got != true {
		t.Errorf("internalFlag = %v, want true", got)
	}
}

func TestAppConfig(t *testing.T) {
	// 测试变量组
	if basics.AppConfig.Host != "localhost" {
		t.Errorf("AppConfig.Host = %s, want localhost", basics.AppConfig.Host)
	}
	if basics.AppConfig.Port != 8080 {
		t.Errorf("AppConfig.Port = %d, want 8080", basics.AppConfig.Port)
	}
	if basics.AppConfig.Mode != "debug" {
		t.Errorf("AppConfig.Mode = %s, want debug", basics.AppConfig.Mode)
	}
}

func TestInitVariables(t *testing.T) {
	// 设置一些非零值
	basics.GlobalCounter = 999
	basics.SetInternalCounter(888)
	basics.SetInternalFlag(true)

	// 调用初始化函数
	basics.InitVariables()

	// 验证被重置
	if basics.GlobalCounter != 0 {
		t.Errorf("GlobalCounter = %d, want 0", basics.GlobalCounter)
	}
	if basics.GetInternalCounter() != 0 {
		t.Errorf("internalCounter = %d, want 0", basics.GetInternalCounter())
	}
	if basics.GetInternalFlag() != false {
		t.Errorf("internalFlag = %v, want false", basics.GetInternalFlag())
	}
}

func TestSafeCounter(t *testing.T) {
	// 重置计数器
	basics.ResetSafeCounter()

	// 单线程测试
	basics.SafeIncrement()
	basics.SafeIncrement()
	basics.SafeIncrement()

	if got := basics.GetSafeCounter(); got != 3 {
		t.Errorf("safeCounter = %d, want 3", got)
	}

	basics.SafeDecrement()

	if got := basics.GetSafeCounter(); got != 2 {
		t.Errorf("safeCounter = %d, want 2", got)
	}
}

func TestSafeCounterConcurrent(t *testing.T) {
	// 重置计数器
	basics.ResetSafeCounter()

	// 并发测试
	const goroutines = 100
	const increments = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				basics.SafeIncrement()
			}
		}()
	}

	wg.Wait()

	expected := goroutines * increments
	if got := basics.GetSafeCounter(); got != expected {
		t.Errorf("safeCounter = %d, want %d", got, expected)
	}
}
