package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	/*
		数据绑定
			Gin 提供了强大的绑定系统，可以将请求数据解析到Go结构体中并自动验证。
			你无需手动调用 c.PostForm() 或 读取 c.Request.Body ,只需定义一个带标签的结构体，让 Gin 来完成工作。


			Bind 与 ShouldBind
				Gin 提供了两组绑定方法：
					c.Bind、c.BindJSON等   				自动调用 c.AbortWithError(400,err)    希望Gin处理错误
					c.ShouldBind、c.ShouldBindJSON 等  	返回错误由你自己处理 					希望自定义错误响应
			大多数情况下，推荐使用 ShouldBind 以获得更好的错误处理控制
	*/

	router := gin.Default()

	router.POST("/login", func(context *gin.Context) {
		var form LoginForm

		// ShouldBind checks Context-Type to select a binding engine automatically
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "logged in"})
	})
	router.Run(":8080")

	/*
		绑定查询字符串或者Post数据
			ShouldBind 会根据HTTP方法和Context-Type 请求头自动选择绑定引擎：
				对于 GET 请求，使用查询字符串绑定（form 标签）。
				对于 POST/PUT 请求，它会检查 Content-Type——对 application/json 使用 JSON 绑定，
				对 application/xml 使用 XML 绑定，对 application/x-www-form-urlencoded 或 multipart/form-data 使用表单绑定。
			这意味着单个处理函数可以同时接受来自查询字符串和请求体的数据，无需手动选择数据源
	*/

	routeOr := gin.Default()
	routeOr.GET("/testing", startPageOr)
	routeOr.POST("/testing", startPageOr)
	routeOr.Run(":8085")

	/*
		# GET with query string parameters
		curl "http://localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"
		# Output: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}

		# POST with form data
		curl -X POST http://localhost:8085/testing \
		  -d "name=appleboy&address=xyz&birthday=1992-03-15"
		# Output: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}

		# POST with JSON body
		curl -X POST http://localhost:8085/testing \
		  -H "Content-Type: application/json" \
		  -d '{"name":"appleboy","address":"xyz","birthday":"1992-03-15"}'
		# Output: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}
	*/

}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPageOr(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person.Name, person.Address, person.Birthday)
	c.JSON(http.StatusOK, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday,
	})
}
