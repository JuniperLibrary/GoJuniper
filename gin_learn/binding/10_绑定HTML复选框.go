//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MyForm 用于绑定 HTML 复选框
type MyForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../../templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	router.POST("/", func(c *gin.Context) {
		var fakeForm MyForm
		if err := c.ShouldBind(&fakeForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
	})

	router.Run(":8080")
}

// 测试命令：
// # 选择所有三个颜色
// curl -X POST http://localhost:8080/ \
//   -d "colors[]=red&colors[]=green&colors[]=blue"
// # 返回: {"color":["red","green","blue"]}

// # 只选择一个颜色
// curl -X POST http://localhost:8080/ \
//   -d "colors[]=green"
// # 返回: {"color":["green"]}

// # 未选中任何复选框 - 列表为空
// curl -X POST http://localhost:8080/
// # 返回: {"color":[]}
