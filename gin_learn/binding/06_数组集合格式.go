//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Filters 用于展示数组的集合格式
type Filters struct {
	Tags   []string `form:"tags" collection_format:"csv"`     // /search?tags=go,web,api
	Labels []string `form:"labels" collection_format:"multi"` // /search?labels=bug&labels=helpwanted
	IdsSSV []int    `form:"ids_ssv" collection_format:"ssv"`  // /search?ids_ssv=1 2 3
	IdsTSV []int    `form:"ids_tsv" collection_format:"tsv"`  // /search?ids_tsv=1	2	3
	Levels []int    `form:"levels" collection_format:"pipes"` // /search?levels=1|2|3
}

func main() {
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		var f Filters
		if err := c.ShouldBind(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, f)
	})
	r.Run(":8080")
}

// 测试命令：
// # CSV 格式：逗号分隔
// curl "http://localhost:8080/search?tags=go,web,api"
// # 返回: {"tags":["go","web","api"]}

// # Multi 格式：重复键
// curl "http://localhost:8080/search?labels=bug&labels=helpwanted"
// # 返回: {"labels":["bug","helpwanted"]}

// # SSV 格式：空格分隔
// curl "http://localhost:8080/search?ids_ssv=1%202%203"
// # 返回: {"ids_ssv":[1,2,3]}

// # TSV 格式：制表符分隔
// curl "http://localhost:8080/search?ids_tsv=1%091%093"
// # 返回: {"ids_tsv":[1,2,3]}

// # Pipes 格式：管道符分隔
// curl "http://localhost:8080/search?levels=1\|2\|3"
// # 返回: {"levels":[1,2,3]}
