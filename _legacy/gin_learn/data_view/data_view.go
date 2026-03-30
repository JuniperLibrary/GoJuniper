package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	router := gin.Default()

	// 1. json
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world", "status": 200})
	})

	// 2. 结构体响应
	router.GET("/someStruct", func(c *gin.Context) {
		var person struct {
			Name    string `json:"name"`
			Age     int    `json:"age"`
			Message string `json:"message"`
			Number  int    `json:"number"`
		}
		person.Name = "Jack"
		person.Age = 18
		person.Message = "hello world"
		person.Number = 1812
		c.JSON(200, person)
	})

	// 3. XML
	router.POST("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "hello world", "status": 200})
	})

	// 4.YAML响应
	router.POST("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "hello world", "status": 200})
	})

	// 5.protobuf格式,谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		lable := "label"

		data := &protoexample.Test{
			Label: &lable,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	router.Run(":8080")

	/*
		2.HTML模版渲染
			gin 支持加载HTML模版，然后根据模版参数进行配置并返回响应的数据，本质上就是字符串替换
			LoadHTMLGlob() 方法可以加载模版文件
	*/
	htmlRouter := gin.Default()
	htmlRouter.LoadHTMLGlob("templates/*")
	htmlRouter.GET("/index", func(c *gin.Context) {
		c.HTML(200, "template.html", gin.H{
			"title": "Hello World",
			"ce":    "123123123",
		})
	})
	htmlRouter.Run(":8080")

	/*
		3.重定向
	*/
	redirectRouter := gin.Default()
	redirectRouter.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.5lmh.com")
	})
	redirectRouter.Run()

	/*
		4.同步异步
			goroutine 机制可以很方便地实现异步处理
			另外，在启动新的goroutine 不应该使用原始上下文，必须使用它的只读副本
	*/
	syncRouter := gin.Default()
	// 1.异步
	syncRouter.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Panicln("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	// 2. 同步
	syncRouter.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})
	syncRouter.Run(":8080")
}
