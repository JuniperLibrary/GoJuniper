//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PersonWithDefaults 用于绑定表单字段的默认值
type PersonWithDefaults struct {
	Name      string    `form:"name,default=William"`
	Age       int       `form:"age,default=10"`
	Friends   []string  `form:"friends,default=Will;Bill"`                         // multi/csv: 在默认值中使用分号
	Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"` // ssv: 使用空格分隔
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`   // csv: 使用分号
}

func main() {
	r := gin.Default()
	r.POST("/person", func(c *gin.Context) {
		var req PersonWithDefaults
		// 根据 Content-Type 自动推断绑定器
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, req)
	})
	r.Run(":8080")
}

// 测试命令：
// # 不带任何请求体发送 POST 请求，将返回默认值
// curl -X POST http://localhost:8080/person
// # 返回: {"Name":"William","Age":10,"Friends":["Will","Bill"],"Addresses":["foo","bar"],"LapTimes":[1,2,3]}

// # 带部分参数
// curl -X POST http://localhost:8080/person -d "name=John&age=25"
// # 返回: {"Name":"John","Age":25,"Friends":["Will","Bill"],"Addresses":["foo","bar"],"LapTimes":[1,2,3]}
