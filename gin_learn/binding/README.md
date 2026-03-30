# Gin 数据绑定学习目录

本目录包含 Gin 框架数据绑定模块的完整示例代码，源自 [Gin 官方文档](https://gin-gonic.com/zh-cn/docs/binding/)。

## 文件清单

| 序号 | 文件名 | 官网章节名称 | 功能描述 |
|:---:|--------|-------------|----------|
| 01 | `01_模型绑定和验证.go` | 模型绑定和验证 | 演示 JSON、XML、表单 binding，使用 binding 标签进行字段验证 |
| 02 | `02_自定义验证器.go` | 自定义验证器 | 演示注册自定义验证函数，实现 gtfield、required 等验证规则 |
| 03 | `03_仅绑定查询字符串.go` | 仅绑定查询字符串 | 演示使用 ShouldBindQuery 仅绑定 URL 查询参数，忽略请求体 |
| 04 | `04_绑定查询字符串或POST数据.go` | 绑定查询字符串或 POST 数据 | 演示 ShouldBind 根据 HTTP 方法和 Content-Type 自动选择绑定器 |
| 05 | `05_绑定表单字段默认值.go` | 绑定表单字段的默认值 | 演示使用 form 标签的 default 选项设置字段默认值 |
| 06 | `06_数组集合格式.go` | 数组的集合格式 | 演示使用 collection_format 控制切片解析格式 (csv/ssv/tsv/pipes) |
| 07 | `07_绑定URI.go` | 绑定 URI | 演示使用 ShouldBindUri 绑定 URL 路径参数 |
| 08 | `08_绑定自定义反序列化器.go` | 绑定自定义反序列化器 | 演示实现 encoding.TextUnmarshaler 和 BindUnmarshaler 接口自定义反序列化 |
| 09 | `09_绑定请求头.go` | 绑定请求头 | 演示使用 ShouldBindHeader 绑定 HTTP 请求头 |
| 10 | `10_绑定HTML复选框.go` | 绑定 HTML 复选框 | 演示绑定同 name 属性的多个复选框值到切片 |
| 11 | `11_表单绑定.go` | Multipart/Urlencoded 绑定 | 演示绑定表单提交数据 (multipart/form-data 或 application/x-www-form-urlencoded) |
| 12 | `12_嵌套结构体绑定.go` | 使用自定义结构体绑定表单数据 | 演示嵌套结构体、指针结构体、匿名内联结构体的绑定 |
| 13 | `13_多次绑定请求体.go` | 将请求体绑定到不同的结构体 | 演示使用 ShouldBindBodyWith 多次绑定同一请求体到不同结构体 |
| 14 | `14_自定义标签绑定.go` | 使用自定义结构体标签绑定表单数据 | 演示创建自定义绑定器使用不同的结构体标签 (如 url 标签) |

## 运行方式

```bash
# 进入目录
cd gin_learn/binding

# 运行示例
go run 01_模型绑定和验证.go
go run 02_自定义验证器.go
# ... 以此类推
```

## 核心概念总结

### Binding 方法分类

| 类型 | 方法 | 行为 |
|------|------|------|
| Must Bind | `Bind`, `BindJSON`, `BindXML`, `BindQuery` | 绑定失败时自动调用 c.AbortWithError(400, err) |
| Should Bind | `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery` | 绑定失败时返回错误，由开发者自行处理 |

### 结构体标签

| 标签 | 用途 |
|------|------|
| `json:"field"` | 绑定 JSON 请求体 |
| `xml:"field"` | 绑定 XML 请求体 |
| `form:"field"` | 绑定表单/查询字符串 |
| `uri:"field"` | 绑定 URL 路径参数 |
| `header:"field"` | 绑定请求头 |
| `binding:"required"` | 必填验证 |
| `default=value` | 默认值 |
| `time_format` | 时间格式 |
| `collection_format` | 集合格式 (csv/ssv/tsv/pipes/multi) |

## 依赖说明

本目录使用独立的 go.mod 文件管理依赖：

```bash
cd gin_learn/binding
go mod tidy
```

需要 Go 1.22+ 版本。

## 学习建议

1. **从01开始**：按序号顺序学习，逐步深入
2. **动手实践**：运行每个示例，理解其作用
3. **查看文档**：结合 [Gin 官方文档](https://gin-gonic.com/zh-cn/docs/binding/) 学习
4. **尝试修改**：修改代码中的结构体和验证规则加深理解

## 相关资源

- [Gin Web Framework 官方文档](https://gin-gonic.com/zh-cn/)
- [GitHub - gin-gonic/gin](https://github.com/gin-gonic/gin)
- [go-playground/validator 文档](https://github.com/go-playground/validator)