//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestHeader 用于绑定 HTTP 请求头
type TestHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

// AuthHeader 用于绑定带验证的请求头
type AuthHeader struct {
	Token string `header:"Authorization" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		h := TestHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})

	r.GET("/auth", func(c *gin.Context) {
		h := AuthHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": h.Token})
	})

	r.Run(":8080")
}

// 测试命令：
// # 带自定义请求头
// curl -H "Rate:300" -H "Domain:music" http://localhost:8080/
// # 返回: {"Domain":"music","Rate":300}

// # 缺少请求头 - 使用零值
// curl http://localhost:8080/
// # 返回: {"Domain":"","Rate":0}

// # 验证必需的 Authorization 请求头
// curl -H "Authorization: bearer-token-123" http://localhost:8080/auth
// # 返回: {"token":"bearer-token-123"}

// # 缺少必需的请求头
// curl http://localhost:8080/auth
// # 返回: {"error":"Key: 'AuthHeader.Token' Error:Field validation for 'Token' failed on the 'required' tag"}

// 注意：请求头名称不区分大小写，header:"Rate" 会匹配 Rate、rate 或 RATE
