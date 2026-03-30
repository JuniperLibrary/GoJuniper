//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PersonURI 用于绑定 URI 路径参数
type PersonURI struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person PersonURI
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})

	route.Run(":8088")
}

// 测试命令：
// # 有效的 UUID - 绑定成功
// curl http://localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
// # 返回: {"name":"thinkerou","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3"}

// # 无效的 UUID - 绑定失败，返回验证错误
// curl http://localhost:8088/thinkerou/not-uuid
// # 返回: {"error":"Key: 'PersonURI.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"}
