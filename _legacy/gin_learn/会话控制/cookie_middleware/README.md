# Cookie 中间件示例 - Shell 脚本使用指南

## 概述
此目录包含一个基于 Gin 框架的 Cookie 中间件示例，以及用于测试和运行的 shell 脚本。

## 文件说明
- `main.go` - Go 源代码
- `cookie_middleware` - 编译后的可执行文件
- `run_server.sh` - 启动服务器的脚本
- `test_cookie.sh` - 自动化测试脚本
- `manual_test.sh` - 手动测试脚本

## 快速开始

### 1. 启动服务器
```bash
./run_server.sh
```
服务器将在 http://localhost:8080 启动。

### 2. 运行自动化测试
在另一个终端中运行：
```bash
./test_cookie.sh
```
该脚本将：
- 自动启动和停止服务器
- 测试未授权访问
- 登录并设置 cookie
- 验证授权访问
- 显示测试结果

### 3. 手动测试
如果服务器已在运行，可以使用手动测试脚本：
```bash
./manual_test.sh
```

## API 端点

### GET /login
设置 cookie 并返回登录成功消息。

**响应示例：**
```
Login success!
```

### GET /home
需要有效的 cookie 才能访问。

**未授权响应 (401)：**
```json
{"error":"err"}
```

**授权响应 (200)：**
```json
{"data":"home"}
```

## 测试流程

### 自动化测试 (`test_cookie.sh`)
1. 启动服务器
2. 测试未授权访问 `/home` (应返回 401)
3. 访问 `/login` 设置 cookie
4. 使用 cookie 访问 `/home` (应返回 200)
5. 验证响应内容
6. 清理并停止服务器

### 手动测试 (`manual_test.sh`)
1. 测试未授权访问
2. 登录并保存 cookie
3. 使用 cookie 访问受保护端点
4. 查看 cookie 内容

## Cookie 参数说明
在 `main.go` 中设置的 cookie 参数：
- **名称**: `abc`
- **值**: `123`
- **过期时间**: 60 秒
- **路径**: `/`
- **域名**: `localhost`
- **Secure**: `false` (允许 HTTP)
- **HttpOnly**: `true` (防止 JavaScript 访问)

## 故障排除

### 端口被占用
如果端口 8080 被占用，`run_server.sh` 会提示是否强制重启。

### 服务器启动失败
确保已编译 Go 程序：
```bash
go build
```

### 测试失败
确保服务器正在运行，并且端口正确。

## 安全注意事项
- 此示例使用简单的 cookie 值进行演示
- 在生产环境中，应使用更安全的认证机制
- 考虑使用 HTTPS 和加密的 cookie