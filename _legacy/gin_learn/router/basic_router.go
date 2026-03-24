package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	/*
		1.基本路由
	*/
	/*	// 创建默认的路由引擎
		basicRouter := gin.Default()

		// 注册 GET 请求路由，处理根路径请求
		basicRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello word")
		})

		// 注册 POST 请求路由
		basicRouter.POST("/xxxpost", func(context *gin.Context) {

		})

		// 注册 PUT 请求路由
		basicRouter.PUT("/xxxput")

		//监听端口默认为 8080
		basicRouter.Run(":8000")*/

	/*
		2. Restful 风格的API
			gin 支持 RestFul 风格的Api

				1.获取文章 /blog/getXxx   Get blog/Xxx
				2.添加 /blog/addXxx      POST blog/Xxx
				3.修改 /blog/updateXxx   PUT blog/Xxx
				4.删除 /blog/delXxxx     DELETE blog/Xxx
	*/

	/*
		3.API参数
			可以通过Context的Param方法来获取API参数
			localhost:8000/xxx/zhangsan
	*/
	/*	apiParamRouter := gin.Default()
		apiParamRouter.GET("/user/:name/*action", func(context *gin.Context) {
			name := context.Param("name")
			action := context.Param("action")
			// 截取
			action = strings.Trim(action, "/")
			context.String(http.StatusOK, name+" is "+action)
		})

		apiParamRouter.Run(":8000")*/

	/*
		4.URL参数
			URL参数可以通过DefaultQuery()或Query()方法获取
			DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
			API ? name=zs
	*/
	/*	urlRouter := gin.Default()
		urlRouter.GET("/user", func(context *gin.Context) {
			// 指定默认值
			name := context.DefaultQuery("name", "zhangsan")
			context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
		})
		urlRouter.Run(":8001")*/

	/*
		5.表单参数
			表单传输为post请求，http常见的传输格式为四种：
				application/json
				application/x-www-form-urlencoded
				application/xml
				multipart/form-data
			表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
	*/
	/*	formRouter := gin.Default()
		formRouter.POST("/form", func(context *gin.Context) {
			types := context.DefaultPostForm("type", "post")
			username := context.PostForm("username")
			password := context.PostForm("password")
			// context.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
			context.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		})
		formRouter.Run(":8080")*/

	/*
		6.上传单个文件
			multipart/form-data格式用于文件上传
			gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
	*/
	/*	singleUpload := gin.Default()
		singleUpload.MaxMultipartMemory = 8 << 20
		singleUpload.POST("/upload", func(c *gin.Context) {
			file, err := c.FormFile("file")
			if err != nil {
				c.String(500, "上传图片出错")
			}
			// c.JSON(200, gin.H{"message": file.Header.Context})
			c.SaveUploadedFile(file, file.Filename)
			c.String(http.StatusOK, file.Filename)
		})
		singleUpload.Run()*/

	/*
		7.上传特定文件
			有的用户上传文件需要限制上传文件的类型以及上传文件的大小，但是gin框架暂时没有这些函数(也有可能是我没找到)，
			因此基于原生的函数写法自己写了一个可以限制大小以及文件类型的上传函数
	*/

	/*	limitUpload := gin.Default()
		limitUpload.POST("/upload", func(c *gin.Context) {
			_, headers, err := c.Request.FormFile("file")
			if err != nil {
				log.Printf("Error when try to get file: %v", err)
			}
			//headers.Size 获取文件大小
			if headers.Size > 1024*1024*2 {
				fmt.Println("文件太大了")
				return
			}
			//headers.Header.Get("Content-Type")获取上传文件的类型
			if headers.Header.Get("Content-Type") != "image/png" {
				fmt.Println("只允许上传png图片")
				return
			}
			// 保存文件
			c.SaveUploadedFile(headers, "./video/"+headers.Filename)
			c.String(http.StatusOK, headers.Filename)
		})
		limitUpload.Run()*/

	/*
		8.上传多个文件
	*/
	// 默认使用了2个中间件，默认使用了2个中间件Logger(), Recovery()
	multipleUpload := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	multipleUpload.MaxMultipartMemory = 8 << 20
	multipleUpload.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	multipleUpload.Run(":8000")

	/*
		9.Router Group
			routes group是为了管理一些相同的URL
	*/
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	routerGroup := gin.Default()
	// 路由组 1，处理Get请求
	v1 := routerGroup.Group("/v1")
	// {} 书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}

	// 路由组 2，处理Post请求
	v2 := routerGroup.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("submit", submit)
	}
	routerGroup.Run(":8000")

	/*
		10.路由的原理

	*/
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	// 获取表单参数
	name := c.DefaultPostForm("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
