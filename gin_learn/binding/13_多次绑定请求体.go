//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// FormA 表单结构体 A
type FormA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

// FormB 表单结构体 B
type FormB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/bind", func(c *gin.Context) {
		objA := FormA{}
		objB := FormB{}

		// ShouldBindBodyWith 会读取请求体并将其存储在上下文中
		// 这样可以多次绑定同一请求体
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched formA", "foo": objA.Foo})
			return
		}

		// 此时会重用上下文中存储的请求体
		if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched formB", "bar": objB.Bar})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "request body did not match any known format"})
	})

	router.Run(":8080")
}

// 测试命令：
// # 请求体匹配 FormA
// curl -X POST http://localhost:8080/bind \
//   -H "Content-Type: application/json" \
//   -d '{"foo":"hello"}'
// # 返回: {"foo":"hello","message":"matched formA"}

// # 请求体匹配 FormB
// curl -X POST http://localhost:8080/bind \
//   -H "Content-Type: application/json" \
//   -d '{"bar":"world"}'
// # 返回: {"bar":"world","message":"matched formB"}

// # 请求体不匹配任何一个结构体
// curl -X POST http://localhost:8080/bind \
//   -H "Content-Type: application/json" \
//   -d '{"invalid":"data"}'
// # 返回: {"error":"request body did not match any known format"}

// 注意：ShouldBindBodyWith 会在绑定前将请求体存储到上下文中，这对性能有轻微影响
// 因此仅在需要多次绑定请求体时使用。对于不读取请求体的格式（如 Query、Form），可以多次调用 ShouldBind
