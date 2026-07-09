#!/bin/bash

# 测试 cookie_middleware 示例的 shell 脚本
# 适用于 macOS

set -e  # 遇到错误时退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 配置
SERVER_PORT=8000
SERVER_URL="http://localhost:$SERVER_PORT"
COOKIE_JAR="cookies.txt"

# 清理函数
cleanup() {
    echo -e "${YELLOW}正在清理...${NC}"
    if [ ! -z "$SERVER_PID" ]; then
        kill $SERVER_PID 2>/dev/null || true
        echo -e "${YELLOW}已停止服务器 (PID: $SERVER_PID)${NC}"
    fi
    rm -f $COOKIE_JAR
}

# 设置退出时清理
trap cleanup EXIT

# 检查是否存在可执行文件
if [ ! -f "./cookie_middleware" ]; then
    echo -e "${RED}错误: 找不到 cookie_middleware 可执行文件${NC}"
    echo -e "${BLUE}请先运行: go build${NC}"
    exit 1
fi

echo -e "${BLUE}=== 测试 Cookie 中间件示例 ===${NC}"

# 1. 启动服务器
echo -e "${YELLOW}1. 启动服务器...${NC}"
./cookie_middleware &
SERVER_PID=$!
echo -e "${GREEN}服务器已启动 (PID: $SERVER_PID)${NC}"

# 等待服务器启动
echo -e "${YELLOW}等待服务器启动...${NC}"
sleep 2

# 检查服务器是否正在运行
if ! kill -0 $SERVER_PID 2>/dev/null; then
    echo -e "${RED}错误: 服务器启动失败${NC}"
    exit 1
fi

# 2. 测试未授权访问 /home
echo -e "${YELLOW}2. 测试未授权访问 /home...${NC}"
response=$(curl -s -w "%{http_code}" -o response.txt "$SERVER_URL/home")
http_code=${response: -3}
response_body=$(cat response.txt)
rm -f response.txt

echo -e "HTTP 状态码: $http_code"
echo -e "响应内容: $response_body"

if [ "$http_code" -eq 401 ]; then
    echo -e "${GREEN}✓ 正确: 未授权访问返回 401${NC}"
else
    echo -e "${RED}✗ 错误: 期望 401，但得到 $http_code${NC}"
fi

# 3. 登录并设置 cookie
echo -e "${YELLOW}3. 访问 /login 设置 cookie...${NC}"
login_response=$(curl -s -c $COOKIE_JAR "$SERVER_URL/login")
echo -e "登录响应: $login_response"

# 检查 cookie 文件
if [ -f "$COOKIE_JAR" ]; then
    echo -e "${GREEN}✓ Cookie 已保存${NC}"
    echo -e "Cookie 内容:"
    cat $COOKIE_JAR
else
    echo -e "${RED}✗ 错误: Cookie 文件未创建${NC}"
fi

# 4. 使用 cookie 访问 /home
echo -e "${YELLOW}4. 使用 cookie 访问 /home...${NC}"
auth_response=$(curl -s -b $COOKIE_JAR -w "%{http_code}" -o response.txt "$SERVER_URL/home")
auth_http_code=${auth_response: -3}
auth_body=$(cat response.txt)
rm -f response.txt

echo -e "HTTP 状态码: $auth_http_code"
echo -e "响应内容: $auth_body"

if [ "$auth_http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ 正确: 授权访问返回 200${NC}"
else
    echo -e "${RED}✗ 错误: 期望 200，但得到 $auth_http_code${NC}"
fi

# 5. 验证响应内容
if echo "$auth_body" | grep -q '"data":"home"'; then
    echo -e "${GREEN}✓ 正确: 响应包含预期数据${NC}"
else
    echo -e "${RED}✗ 错误: 响应内容不符合预期${NC}"
fi

# 总结
echo -e "${BLUE}=== 测试完成 ===${NC}"
echo -e "${GREEN}所有测试执行完毕！${NC}"

# 注意：服务器将在脚本退出时通过 trap 自动清理