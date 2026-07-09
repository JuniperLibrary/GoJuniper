# Gin 框架数据校验详解

Gin 使用 **struct tag + `go-playground/validator`** 进行请求体数据校验。

---

## 一、基本流程

```
请求进来
  → ShouldBindJSON(&struct)
    → Gin 根据 struct tag 自动调用 go-playground/validator
      → 校验通过：数据填充到 struct
      → 校验失败：返回 error（ShouldBind）或 400（Bind）
```

---

## 二、定义结构体 + 校验标签

```go
type Login struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}
```

`binding:"required"` 表示该字段必填，为空则校验失败。

### 常用内置标签

| 标签 | 含义 |
|------|------|
| `required` | 必填 |
| `min=3` | 最小值/最小长度 |
| `max=10` | 最大值/最大长度 |
| `email` | 邮箱格式 |
| `gt=0` | 大于 0 |
| `gtfield=CheckIn` | 大于另一个字段（如退房 > 入住） |
| `oneof=a b c` | 枚举值 |

---

## 三、两组绑定方法

| 类型 | 方法 | 行为 |
|------|------|------|
| **Must Bind** | `Bind`、`BindJSON`、`BindQuery` | 校验失败直接返回 400，中断请求 |
| **Should Bind** | `ShouldBind`、`ShouldBindJSON`、`ShouldBindQuery` | 校验失败返回 error，由开发者自行处理 |

### 示例代码

```go
// JSON 绑定
c.ShouldBindJSON(&login)

// 表单绑定（自动推断 Content-Type）
c.Bind(&form)

// URI 绑定
c.ShouldBindUri(&login)

// Query 绑定
c.ShouldBindWith(&b, binding.Query)
```

---

## 四、Struct Tag 逐个拆解

以这个为例：

```go
CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
```

反引号里面是 **三组独立的 tag**，用空格分隔：

| Tag | 含义 | Java 对标 |
|-----|------|-----------|
| `form:"check_in"` | 表单字段名叫 `check_in`，映射到这个 Go 字段 | 类似 `@RequestParam("check_in")` |
| `binding:"required,bookabledate"` | 校验规则：必填 + 自定义校验器 `bookabledate` | 类似 `@NotBlank` + 自定义 Validator |
| `time_format:"2006-01-02"` | 时间解析格式（Go 的日期格式写法） | 类似 `@JsonFormat(pattern="yyyy-MM-dd")` |

再看 `json:"user"`：

| Tag | 含义 | Java 对标 |
|-----|------|-----------|
| `json:"user"` | JSON 里的字段名是 `user`，映射到这个 Go 字段 | 类似 `@JsonProperty("user")` |

---

## 五、Go vs Java 对照表

| Java（Spring Boot） | Go（Gin） | 作用 |
|---------------------|-----------|------|
| `@JsonProperty("name")` | `json:"name"` | JSON 字段名映射 |
| `@RequestParam("name")` | `form:"name"` | 表单字段名映射 |
| `@PathVariable("id")` | `uri:"id"` | URL 路径参数映射 |
| `@NotBlank` | `binding:"required"` | 必填校验 |
| `@Size(min=2, max=20)` | `binding:"min=2,max=20"` | 长度校验 |
| `@Email` | `binding:"email"` | 邮箱格式校验 |
| `@JsonFormat(pattern=...)` | `time_format:"2006-01-02"` | 日期格式 |
| 自定义 Validator | `binding:"bookabledate"` + `RegisterValidation(...)` | 自定义校验 |

### 核心区别

- Java 用**注解**（`@` 符号，写在字段上方）
- Go 用**结构体标签**（反引号，写在字段后面）
- Java 有"约定大于配置"，`@JsonProperty` 可以省略
- Go 必须显式声明每个 tag
- Go 一个 struct 可以通过多个 tag 同时支持 JSON、表单、URI 三种输入方式

---

## 六、自定义校验器

### 注册自定义校验器

```go
import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

// 1. 定义校验函数：拒绝过去的日期
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
    date, ok := fl.Field().Interface().(time.Time)
    if ok {
        return !time.Now().After(date) // 过去的日期返回 false
    }
    return true
}

// 2. 注册到 validator 引擎
v := binding.Validator.Engine().(*validator.Validate)
v.RegisterValidation("bookabledate", bookableDate)
```

### 在 struct 中使用

```go
type Booking struct {
    CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
    CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}
```

### 调用示例

```go
func getBookable(c *gin.Context) {
    var b Booking
    if err := c.ShouldBindWith(&b, binding.Query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
}
```

### 测试

```bash
# 日期在未来，校验通过
curl "http://localhost:8085/bookable?check_in=2118-04-16&check_out=2118-04-17"
# Output: {"message":"Booking dates are valid!"}

# 退房早于入住，gtfield 校验失败
curl "http://localhost:8085/bookable?check_in=2118-03-10&check_out=2118-03-09"
# Output: {"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}
```

---

## 七、项目中的实际代码

完整示例位于：`_legacy/gin_learn/data_model/json_parse_bind.go`

该文件包含四种绑定方式的完整示例：

1. JSON 数据解析和绑定（`ShouldBindJSON`）
2. 表单数据解析和绑定（`Bind`）
3. URI 数据解析和绑定（`ShouldBindUri`）
4. 自定义验证器（`RegisterValidation` + `bookabledate`）
