package main

import (
	"fmt"
)

// DivideError 定义了除法操作中除数为0时的错误信息。
// 它包含被除数和除数的值，以便提供更详细的错误上下文。
// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// Error 方法实现了error接口，返回除数为0时的错误信息。
// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// Divide 函数执行两个整数的除法操作。
// 它返回商和一个错误信息，当除数为0时，错误信息不为空。
// 参数:
//   varDividee - 被除数
//   varDivider - 除数
// 返回值:
//   商的结果
//   当除数不为0时，错误信息为空；当除数为0时，包含除数为0的错误信息。
// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {
	// 正常的除法操作
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 除数为0的情况
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
