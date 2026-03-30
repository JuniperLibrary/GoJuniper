//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StructA 嵌套结构体 A
type StructA struct {
	FieldA string `form:"field_a"`
}

// StructB 包含嵌套结构体
type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

// StructC 包含嵌套结构体指针
type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

// StructD 包含匿名内联结构体
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func main() {
	router := gin.Default()

	// 嵌套结构体
	router.GET("/getb", func(c *gin.Context) {
		var b StructB
		c.Bind(&b)
		c.JSON(http.StatusOK, gin.H{
			"a": b.NestedStruct,
			"b": b.FieldB,
		})
	})

	// 嵌套结构体指针
	router.GET("/getc", func(c *gin.Context) {
		var b StructC
		c.Bind(&b)
		c.JSON(http.StatusOK, gin.H{
			"a": b.NestedStructPointer,
			"c": b.FieldC,
		})
	})

	// 匿名内联结构体
	router.GET("/getd", func(c *gin.Context) {
		var b StructD
		c.Bind(&b)
		c.JSON(http.StatusOK, gin.H{
			"x": b.NestedAnonyStruct,
			"d": b.FieldD,
		})
	})

	router.Run(":8080")
}

// 测试命令：
// # 嵌套结构体 - fields from StructA are bound alongside StructB's own fields
// curl "http://localhost:8080/getb?field_a=hello&field_b=world"
// # 返回: {"a":{"FieldA":"hello"},"b":"world"}

// # 嵌套结构体指针 - works the same way, Gin allocates the pointer automatically
// curl "http://localhost:8080/getc?field_a=hello&field_c=world"
// # 返回: {"a":{"FieldA":"hello"},"c":"world"}

// # 匿名内联结构体 - fields are bound by their form tags as usual
// curl "http://localhost:8080/getd?field_x=hello&field_d=world"
// # 返回: {"d":"world","x":{"FieldX":"hello"}}
