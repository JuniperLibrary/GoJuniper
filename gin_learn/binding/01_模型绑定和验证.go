package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 用于绑定 JSON/XML/表单数据的结构体
type Login struct {
	// 用户名，支持 form、json、xml 三种格式绑定
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	// 密码，必填字段
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 绑定 JSON 示例
	// 请求体: {"user": "manu", "password": "123"}
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "未授权"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	// 绑定 XML 示例
	// 请求体: <?xml version="1.0" encoding="UTF-8"?><root><user>manu</user><password>123</password></root>
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "未授权"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	// 绑定 HTML 表单示例
	// 请求体: user=manu&password=123
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 根据 Content-Type 自动推断使用哪种绑定器
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "未授权"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	// 启动服务器，监听 8080 端口
	router.Run(":8080")
}

/*
测试命令：

1. JSON 请求（缺少密码字段，会返回验证错误）
curl -X POST http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'

2. XML 请求
curl -X POST http://localhost:8080/loginXML \
  -H 'content-type: application/xml' \
  -d '<?xml version="1.0" encoding="UTF-8"?><root><user>manu</user><password>123</password></root>'

3. 表单请求
curl -X POST http://localhost:8080/loginForm \
  -d 'user=manu&password=123'
*/
