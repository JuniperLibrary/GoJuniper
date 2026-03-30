package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking 包含绑定和验证数据的结构体
// 使用自定义验证器 bookabledate 验证日期必须是未来日期
type Booking struct {
	// 入住日期，必填，使用自定义验证器
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	// 退房日期，必填，必须大于入住日期，使用自定义验证器
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// bookableDate 自定义验证器函数
// 功能：拒绝过去的日期，确保预订的日期在未来
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		// 如果日期在过去，返回 false（验证失败）
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	route := gin.Default()

	// 注册自定义验证器 bookabledate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

// getBookable 处理预订日期查询请求
// 使用 ShouldBindWith 配合 binding.Query 仅绑定查询字符串
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message":   "预订日期有效！",
			"check_in":  b.CheckIn,
			"check_out": b.CheckOut,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

/*
测试命令：

1. 两个日期都在未来，且退房日期大于入住日期
curl "http://localhost:8085/bookable?check_in=2118-04-16&check_out=2118-04-17"
返回: {"message":"预订日期有效！"}

2. 退房日期在入住日期之前 - gtfield 验证失败
curl "http://localhost:8085/bookable?check_in=2118-03-10&check_out=2118-03-09"
返回: {"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}
*/
