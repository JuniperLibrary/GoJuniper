//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	CustomerTag   = "url"
	DefaultMemory = 32 << 20
)

// CustomerBinding 自定义绑定器，使用自定义标签
type CustomerBinding struct{}

// Name 返回绑定名称
func (customerBinding CustomerBinding) Name() string {
	return "form"
}

// Bind 绑定请求到对象，使用自定义标签
func (customerBinding CustomerBinding) Bind(req *http.Request, obj any) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(DefaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	// 使用 url 标签代替 form 标签
	if err := binding.MapFormWithTag(obj, req.Form, CustomerTag); err != nil {
		return err
	}
	return validate(obj)
}

// validate 执行结构体验证
func validate(obj any) error {
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(obj)
}

// FormA 外部类型，无法修改其标签
type FormA struct {
	FieldA string `url:"field_a"`
}

func main() {
	router := gin.Default()

	router.GET("/list", func(c *gin.Context) {
		var urlBinding = CustomerBinding{}
		var opt FormA
		// 使用自定义绑定器读取 url 标签
		if err := c.MustBindWith(&opt, urlBinding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"field_a": opt.FieldA})
	})

	router.Run(":8080")
}

// 测试命令：
// # 自定义绑定读取 "url" 结构体标签而不是 "form"
// curl "http://localhost:8080/list?field_a=hello"
// # 返回: {"field_a":"hello"}

// # 缺少参数 - 空字符串
// curl "http://localhost:8080/list"
// # 返回: {"field_a":""}

// 注意：自定义绑定需要实现 binding.Binding 接口
// binding.MapFormWithTag 辅助函数使用自定义标签将表单值映射到结构体字段
