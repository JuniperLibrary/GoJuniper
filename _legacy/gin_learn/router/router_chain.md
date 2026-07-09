# Gin 路由原理详解

## 1. 核心架构
HTTP 请求 → Router（路由器）→ Handlers Chain（处理链）→ Response
           ↑                    ↑
        路由匹配            中间件 + 业务逻辑


## 2. 路由注册流程

从你的代码可以看到：
```go
// 1. 创建路由引擎
r := gin.Default()

// 2. 注册路由
r.GET("/", handler)
r.POST("/login", loginHandler)

// 3. 启动服务
r.Run(":8080")
```
底层发生了什么：
```go
// GET 方法内部实现（简化版）
func (engine *Engine) GET(path string, handlers ...HandlersChain) {
    engine.handle("GET", path, handlers)
}

// handle 方法将路由添加到路由树
func (engine *Engine) handle(method, path string, handlers HandlersChain) {
    // 添加到路由树中
    engine.addRoute(method, path, handlers)
}

```

## 3. 路由树结构（Radix Tree)

Gin 使用 Radix Tree（基数树） 来存储和匹配路由，这是一种高效的前缀树：

路由注册：
  /user/login
  /user/profile
  /user/:id
  /api/v1/users

对应的路由树（简化）：
/
├── user/
│   ├── login      → handler1
│   ├── profile    → handler2
│   └── :id        → handler3  (参数节点)
└── api/
    └── v1/
        └── users  → handler4

特点：
1. ✅ 快速匹配：O(k) 复杂度，k 是路径长度
2. ✅ 支持参数：:id、*action
3. ✅ 内存友好：共享公共前缀

## 4. 路由匹配过程
当请求 GET /user/123 时：
1. 从根节点开始遍历路由树
2. 匹配 "user/" 前缀 ✓
3. 匹配 "123"，发现是参数节点 `:id` ✓
4. 提取参数 id = "123"
5. 找到对应的 Handler Chain
6. 执行中间件 + 业务逻辑

## 5. Handler Chain（处理链）
你代码中的路由其实是一个 责任链模式：

```go
r.GET("/login", Logger(), Recovery(), loginHandler)
//                ↑          ↑           ↑
//             中间件 1   中间件 2    最终处理器
```
执行流程：

```go
// 简化版执行流程
func (c *Context) Next() {
    c.index++
    for c.index < len(c.handlers) {
        c.handlers[c.index](c)  // 执行下一个 handler
        c.index++
    }
}

// 实际执行顺序：
Logger() → Recovery() → loginHandler → Response
```
## 6. 路由组（Router Group）
你的代码中已经使用了路由组：
```go
v1 := routerGroup.Group("/v1")
{
    v1.GET("/login", login)    // → /v1/login
    v1.GET("/submit", submit)  // → /v1/submit
}

v2 := routerGroup.Group("/v2")
{
    v2.POST("/login", login)   // → /v2/login
    v2.POST("/submit", submit) // → /v2/submit
}
```
原理：
```go

// Group 方法内部（简化版）
func (group *RouterGroup) Group(relativePath string) *RouterGroup {
    return &RouterGroup{
        basePath: group.basePath + relativePath,  // 拼接路径
        handlers: group.handlers,                  // 继承中间件
    }
}

```

效果：

/v1/login = 基础路径 /v1 + 路由路径 /login
可以添加组级中间件（如鉴权）

## 7. 参数提取

从你的代码可以看到参数使用：

```go
// URL 参数
// /user?name=zs
name := c.Query("name")

// 路由参数
// /user/:name
name := c.Param("name")

// 通配符
// /user/*action
action := c.Param("action")

```

底层实现：
```go
// 匹配时将参数存入 Params 字典
params := map[string]string{
    "name":   "zhangsan",
    "action": "login/profile",
}

// Param 方法从字典中取值
func (c *Context) Param(key string) string {
    return c.Params[key]
}

```

## 8. Context 的核心作用
gin.Context 是整个路由处理的核心载体：
```go
type Context struct {
    Request  *http.Request      // HTTP 请求
    Writer   http.ResponseWriter // HTTP 响应
    Params   map[string]string   // 路由参数
    handlers HandlersChain       // 处理链
    index    int                 // 当前执行到的 handler 索引
    Keys     map[string]any      // 上下文数据（跨中间件传递）
}

```

你的代码中的应用：

```go
r.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")           // 获取路由参数
    c.String(200, "hello %s", name)   // 写入响应
})

```
## 9. 完整请求流程图
客户端请求 :8080/user/login?name=zs
         ↓
    http.Server
         ↓
    gin.Engine (实现 http.Handler 接口)
         ↓
    查找路由树 (匹配 GET /user/login)
         ↓
    创建 Context (封装 Request/Response/Params)
         ↓
    执行 Handler Chain:
         ├─ Logger()        // 记录日志
         ├─ Recovery()      // 捕获 panic
         └─ loginHandler    // 业务逻辑
              ↓
         c.Query("name")   // 获取参数
         c.String()        // 写入响应
              ↓
    返回 HTTP Response

## 10. 为什么 Gin 快？
    Radix Tree 路由匹配 - 高效的树结构
    复用 Context - sync.Pool 对象池
    零分配优化 - 减少内存分配
    中间件链 - 灵活的管道设计
    这就是为什么你在 _legacy/gin_learn 中写的代码能够正常工作的底层原理！🎯