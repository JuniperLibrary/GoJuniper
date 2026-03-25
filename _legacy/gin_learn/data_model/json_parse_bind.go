package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

/*
	模型绑定和验证
		要将请求体绑定到类型中，请使用模型绑定。我们目前支持绑定 JSON、XML、YAML 和标准表单值（foo=bar&boo=baz）。
		请注意，你需要在所有要绑定的字段上设置相应的绑定标签。例如，从 JSON 绑定时，设置 json:"fieldname"。
		此外，Gin 提供了两组绑定方法：
			- 类型 Must bind：
				- 方法 ： Bind、BindJSON、BindXML、BindQuery、BindYAML
				- 行为： 这些方法底层使用  MustBindWith 。
						请求将使用 c.AbortWithError(400, err).SetType(ErrorTypeBind) 中止。
						这会将响应状态码设置为 400，并将 Content-Type 头设置为 text/plain; charset=utf-8。
						注意，如果你在此之后尝试设置响应码，将会出现警告
						[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。
						如果你希望更好地控制行为，请考虑使用 ShouldBind 等效方法。
			- 类型 - Should bind
				- 方法： ShouldBind 、ShouldBindJson、ShouldBindXML、ShouldBindQuery、ShouldBindYAML
				- 行为： 这些方法底层使用 ShouldBindWith。如果存在绑定错误，错误会被返回，由开发者负责适当地处理请求和错误。
		使用 Bind 方法时，Gin会尝试根据Content-Type 头来推断绑定器。如果你确定要绑定的内容类型，可以使用 MustBindWith 或 ShouldBindWith。

		你还可以指定特定字段为必填。如果一个字段标记了 binding:"required" 并且在绑定时值为空，将返回错误。

*/

// Login 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {

	/*
		1.Json 数据解析和绑定
	*/
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	jsonRouter := gin.Default()
	// JSON绑定
	jsonRouter.POST("/loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var login Login
		// 将request的body中的数据，自动按照json格式解析到login变量中
		if err := c.ShouldBindJSON(&login); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	jsonRouter.Run(":8080")

	/*
		2. 表单数据解析和绑定
	*/
	formRouter := gin.Default()
	formRouter.GET("/loginForm", func(c *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	formRouter.Run(":8080")

	/*
		3.URI数据解析和绑定
	*/
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	uriRouter := gin.Default()
	// JSON绑定
	// localhost:8000/root/admin
	uriRouter.GET("/:user/:password", func(c *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	uriRouter.Run(":8000")

	/*
		4. 自定义验证器
			Gin 使用 go-playground/validator 进行字段级验证。除了内置的验证器（如 required、email、min、max），你还可以注册自己的自定义验证函数。
			下面的示例注册了一个 bookabledate 验证器，用于拒绝过去的日期，确保预订的入住和退房日期始终在未来。
	*/

	validatorRouter := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("bookabledate", bookableDate)
	}
	validatorRouter.GET("/bookable", getBookable)
	validatorRouter.Run(":8000")
	/*
		# Both dates are in the future and check_out > check_in
		curl "http://localhost:8085/bookable?check_in=2118-04-16&check_out=2118-04-17"
		# Output: {"message":"Booking dates are valid!"}

		# check_out is before check_in -- fails gtfield validation
		curl "http://localhost:8085/bookable?check_in=2118-03-10&check_out=2118-03-09"
		# Output: {"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}
	*/
}

func getBookable(context *gin.Context) {
	var b Booking
	if err := context.ShouldBindWith(&b, binding.Query); err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
