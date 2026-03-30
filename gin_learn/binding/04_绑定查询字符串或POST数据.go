//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Person 用于绑定查询字符串或 POST 数据
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// ShouldBind 会根据 HTTP 方法和 Content-Type 自动选择绑定引擎
	// GET 请求: 使用查询字符串绑定
	// POST/PUT 请求: 根据 Content-Type (JSON/XML/表单) 选择不同的绑定器
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

// 测试命令：
// # GET 请求带查询字符串参数
// curl "http://localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"
// # 返回: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}

// # POST 请求带表单数据
// curl -X POST http://localhost:8085/testing \
//   -d "name=appleboy&address=xyz&birthday=1992-03-15"
// # 返回: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}

// # POST 请求带 JSON body
// curl -X POST http://localhost:8085/testing \
//   -H "Content-Type: application/json" \
//   -d '{"name":"appleboy","address":"xyz","birthday":"1992-03-15"}'
// # 返回: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}
