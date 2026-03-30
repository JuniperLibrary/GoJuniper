package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person 用于绑定查询字符串参数的结构体
type Person struct {
	// 姓名
	Name string `form:"name"`
	// 地址
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	// 使用 Any 允许 GET 和 POST 方法
	route.Any("/testing", startPage)
	route.Run(":8085")
}

// startPage 处理查询字符串绑定请求
// ShouldBindQuery 仅绑定 URL 查询字符串参数，完全忽略请求体
func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBindQuery(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Address,
	})
}

/*
测试命令：

1. GET 请求带查询参数
curl "http://localhost:8085/testing?name=appleboy&address=xyz"
返回: {"address":"xyz","name":"appleboy"}

2. POST 请求带查询参数 - 请求体被忽略，只绑定查询字符串
curl -X POST "http://localhost:8085/testing?name=appleboy&address=xyz" \
  -d "name=ignored&address=ignored"
返回: {"address":"xyz","name":"appleboy"}
*/
