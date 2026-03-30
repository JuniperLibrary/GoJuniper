# 会话控制学习笔记

## 1. Cookie介绍

### 1.1 HTTP协议的无状态性
- HTTP是无状态协议，服务器不能记录浏览器的访问状态
- 服务器无法区分两次请求是否由同一个客户端发出

### 1.2 Cookie的定义
- Cookie是解决HTTP协议无状态的方案之一
- Cookie是服务器保存在浏览器上的一段信息
- 浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器
- Cookie由服务器创建，并发送给浏览器，最终由浏览器保存

### 1.3 Cookie的用途
- 测试服务端发送cookie给客户端，客户端请求时携带cookie
- 用于用户身份识别、会话跟踪、个性化设置等

## 2. Cookie的使用

### 2.1 Gin框架中设置Cookie
```go
c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
```
参数说明：
- `name`: cookie的名称
- `value`: cookie的值
- `maxAge`: 过期时间（秒）
- `path`: cookie所在目录
- `domain`: 域名
- `secure`: 是否只能通过https访问
- `httpOnly`: 是否允许通过js获取cookie

### 2.2 获取Cookie
```go
cookie, err := c.Cookie("key_cookie")
```

### 2.3 示例代码
- `cookie_basic/main.go`: 演示基本的Cookie设置和获取

## 3. Cookie练习

### 3.1 权限验证中间件
- 实现一个权限验证中间件，检查客户端是否携带有效的cookie
- 有2个路由：login（设置cookie）和home（验证cookie）
- 使用中间件在请求home之前验证cookie

### 3.2 示例代码
- `cookie_middleware/main.go`: 完整的权限验证中间件示例

## 4. Cookie的缺点

### 4.1 主要缺点
1. **不安全**：Cookie以明文形式存储和传输
2. **增加带宽消耗**：每个请求都会携带Cookie数据
3. **可能被禁用**：用户可以在浏览器中禁用Cookie
4. **有容量限制**：每个域名下的Cookie数量和大小有限制

## 5. Sessions

### 5.1 gorilla/sessions库
- 提供cookie和文件系统session后端
- 主要功能：
  - 简单的API：设置签名（及可选加密）的cookie
  - 内置后端：可将session存储在cookie或文件系统中
  - Flash消息：持久读取的session值
  - 切换session持久性（"记住我"）的便捷方法
  - 旋转身份验证和加密密钥的机制
  - 每个请求有多个session

### 5.2 Session基本操作
1. **初始化存储**：
```go
var store = sessions.NewCookieStore([]byte("something-very-secret"))
```

2. **获取session**：
```go
session, err := store.Get(r, "session-name")
```

3. **存储值**：
```go
session.Values["foo"] = "bar"
session.Values[42] = 43
session.Save(r, w)
```

4. **获取值**：
```go
foo := session.Values["foo"]
```

5. **删除session**：
```go
session.Options.MaxAge = -1
session.Save(r, w)
```

### 5.3 示例代码
- `sessions_example/main.go`: 使用gorilla/sessions的完整示例

## 6. 学习要点总结

### 6.1 Cookie vs Sessions
| 特性 | Cookie | Sessions |
|------|--------|----------|
| 存储位置 | 客户端（浏览器） | 服务器端 |
| 安全性 | 较低（明文） | 较高（服务器存储） |
| 容量限制 | 较小（约4KB） | 较大（服务器限制） |
| 性能 | 每次请求都携带 | 只需传递session ID |
| 实现复杂度 | 简单 | 相对复杂 |

### 6.2 最佳实践
1. **安全性**：
   - 使用HTTPS传输敏感Cookie
   - 设置HttpOnly防止XSS攻击
   - 设置Secure标志确保仅通过HTTPS传输

2. **性能优化**：
   - 合理设置Cookie过期时间
   - 避免在Cookie中存储大量数据
   - 使用Session存储复杂数据

3. **用户体验**：
   - 提供"记住我"功能
   - 处理Cookie被禁用的情况
   - 实现合理的会话超时机制

## 7. Shell 脚本使用指南

### 7.1 Cookie中间件Shell脚本
`cookie_middleware/` 目录包含三个shell脚本，用于在macOS上测试和运行示例：

#### 7.1.1 自动化测试脚本 (`test_cookie.sh`)
```bash
cd cookie_middleware
./test_cookie.sh
```
该脚本会：
1. 自动启动和停止服务器
2. 测试未授权访问（应返回401）
3. 登录并设置cookie
4. 使用cookie访问受保护端点（应返回200）
5. 显示测试结果

#### 7.1.2 启动服务器脚本 (`run_server.sh`)
```bash
cd cookie_middleware
./run_server.sh
```
该脚本会：
- 检查端口占用情况
- 提示是否强制重启
- 启动服务器（端口8000）

#### 7.1.3 手动测试脚本 (`manual_test.sh`)
```bash
cd cookie_middleware
./manual_test.sh
```
该脚本用于在服务器运行时手动测试各个端点。

## 8. 运行示例

### 8.1 运行Cookie示例
```bash
cd cookie_basic
go run main.go
# 访问 http://localhost:8000/cookie
```

### 8.2 运行Cookie中间件示例
```bash
cd cookie_middleware
go run main.go
# 访问 http://localhost:8000/home (会看到错误)
# 访问 http://localhost:8000/login (设置cookie)
# 再次访问 http://localhost:8000/home (访问成功)
```

### 8.3 运行Sessions示例
```bash
cd sessions_example
go run main.go
# 访问 http://localhost:8080/save 保存session
# 访问 http://localhost:8080/get 获取session
# 访问 http://localhost:8080/delete 删除session
```

## 9. 参考资料
- [Gin官方文档](https://gin-gonic.com/)
- [gorilla/sessions文档](https://www.gorillatoolkit.org/pkg/sessions)
- [topgoer.com会话控制教程](https://topgoer.com/gin%E6%A1%86%E6%9E%B6/%E4%BC%9A%E8%AF%9D%E6%8E%A7%E5%88%B6/)
- [curl官方文档](https://curl.se/docs/manpage.html)