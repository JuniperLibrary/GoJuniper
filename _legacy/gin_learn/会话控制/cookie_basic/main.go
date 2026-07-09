package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CookieBasic 示例：服务端发送cookie给客户端，客户端请求时携带cookie
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()

	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			// 参数说明：
			//   name string, cookie的名称
			//   value string, cookie的值
			//   maxAge int, 单位为秒
			//   path, cookie所在目录
			//   domain string, 域名
			//   secure bool, 是否只能通过https访问
			//   httpOnly bool, 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})

	// 启动服务，默认端口8080
	r.Run(":8000")
}
