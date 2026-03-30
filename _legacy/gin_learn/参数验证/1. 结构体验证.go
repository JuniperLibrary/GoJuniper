package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PersonValid struct {
	Name     string `json:"name" binding:"required,gte=2,lte=100"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Birthday string `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	/*
		结构体验证
			用gin框架的数据验证，可以不用解析数据，减少if else，会简洁许多。
	*/
	r := gin.Default()
	r.GET("/5lmh", func(c *gin.Context) {
		var person PersonValid
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	r.Run()
}
