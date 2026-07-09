package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleWare 权限验证中间件
// 检查客户端是否携带有效的cookie
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				// 验证通过，继续处理后续请求
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

// CookieMiddleware 示例：模拟实现权限验证中间件
// 有2个路由，login和home
// login用于设置cookie
// home是访问查看信息的请求
// 在请求home之前，先跑中间件代码，检验是否存在cookie
func main() {
	// 1.创建路由
	r := gin.Default()

	// 登录路由，用于设置cookie
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})

	// 首页路由，使用AuthMiddleWare中间件进行权限验证
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	// 启动服务
	r.Run(":8000")
}
