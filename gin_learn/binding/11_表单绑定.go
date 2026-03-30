//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginForm 用于绑定 multipart/urlencoded 表单
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		// ShouldBind 自动根据 Content-Type 选择正确的绑定器
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	})

	router.Run(":8080")
}

// 测试命令：
// # Multipart 表单
// curl -X POST http://localhost:8080/login \
//   -F "user=user" -F "password=password"
// # 返回: {"status":"you are logged in"}

// # URL-encoded 表单
// curl -X POST http://localhost:8080/login \
//   -d "user=user&password=password"
// # 返回: {"status":"you are logged in"}

// # 错误的凭据
// curl -X POST http://localhost:8080/login \
//   -d "user=wrong&password=wrong"
// # 返回: {"status":"unauthorized"}

// # 缺少必填字段
// curl -X POST http://localhost:8080/login \
//   -d "user=user"
// # 返回: {"error":"Key: 'LoginForm.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
