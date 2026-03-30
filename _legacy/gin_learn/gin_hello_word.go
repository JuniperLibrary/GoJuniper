package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		Gin web 框架

		Gin 是 Go 语言的一个 Web 框架，Gin 框架的官方网址为 https://github.com/gin-gonic/gin。
		go get -u github.com/gin-gonic/gin
	*/

	// 1.创建一个默认的路由引擎
	r := gin.Default()
	// 2。绑定路由规则 执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong\n",
		})

		c.String(http.StatusOK, "hello World!")
	})

	// 3. 监听端口 默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8080")
}
