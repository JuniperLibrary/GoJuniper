//go:build ignore
// +build ignore

package main

import (
	"encoding"
	"strings"

	"github.com/gin-gonic/gin"
)

// Birthday 使用自定义反序列化器
type Birthday string

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
// 将日期格式从 "2000-01-01" 转换为 "2000/01/01"
func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

// 确保实现接口
var _ encoding.TextUnmarshaler = (*Birthday)(nil)

// ============ 使用 BindUnmarshaler 的示例 ============

// Birthday2 使用 BindUnmarshaler 自定义绑定方式
type Birthday2 string

// UnmarshalParam 实现 binding.BindUnmarshaler 接口
func (b *Birthday2) UnmarshalParam(param string) error {
	*b = Birthday2(strings.Replace(param, "-", "/", -1))
	return nil
}

// 确保实现接口
var _ interface {
	UnmarshalParam(string) error
} = (*Birthday2)(nil)

func main() {
	route := gin.Default()

	// 使用 encoding.TextUnmarshaler 示例
	route.GET("/test1", func(ctx *gin.Context) {
		var request struct {
			Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
			Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
			BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
		}
		_ = ctx.BindQuery(&request)
		ctx.JSON(200, request)
	})

	// 使用 BindUnmarshaler 示例
	route.GET("/test2", func(ctx *gin.Context) {
		var request struct {
			Birthday         Birthday2   `form:"birthday"`
			Birthdays        []Birthday2 `form:"birthdays" collection_format:"csv"`
			BirthdaysDefault []Birthday2 `form:"birthdaysDef,default=2020-09-01;2020-09-02" collection_format:"csv"`
		}
		_ = ctx.BindQuery(&request)
		ctx.JSON(200, request)
	})

	route.Run(":8088")
}

// 测试命令：
// # 使用 encoding.TextUnmarshaler
// curl 'localhost:8088/test1?birthday=2000-01-01&birthdays=2000-01-01,2000-01-02'
// # 返回: {"Birthday":"2000/01/01","Birthdays":["2000/01/01","2000/01/02"],"BirthdaysDefault":["2020/09/01","2020/09/02"]}

// # 使用 BindUnmarshaler
// curl 'localhost:8088/test2?birthday=2000-01-01&birthdays=2000-01-01,2000-01-02'
// # 返回: {"Birthday":"2000/01/01","Birthdays":["2000/01/01","2000/01/02"],"BirthdaysDefault":["2020/09/01","2020/09/02"]}
