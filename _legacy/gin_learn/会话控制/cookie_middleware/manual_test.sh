#!/bin/bash

# 手动测试 cookie_middleware 的 shell 脚本
# 在服务器运行时使用

set -e

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

echo -e "${BLUE}=== Cookie 中间件手动测试 ===${NC}"
echo -e "${YELLOW}请确保服务器已在运行: ./run_server.sh${NC}"
echo ""

# 清理旧的 cookie 文件
rm -f $COOKIE_JAR

# 1. 测试未授权访问
echo -e "${YELLOW}1. 测试未授权访问 /home...${NC}"
echo -e "curl $SERVER_URL/home"
curl -s -w "\nHTTP 状态码: %{http_code}\n" "$SERVER_URL/home"
echo ""

# 2. 登录并设置 cookie
echo -e "${YELLOW}2. 登录并设置 cookie...${NC}"
echo -e "curl -c $COOKIE_JAR $SERVER_URL/login"
curl -s -c $COOKIE_JAR "$SERVER_URL/login"
echo ""
echo -e "${GREEN}Cookie 已保存到 $COOKIE_JAR${NC}"
echo ""

# 3. 使用 cookie 访问 /home
echo -e "${YELLOW}3. 使用 cookie 访问 /home...${NC}"
echo -e "curl -b $COOKIE_JAR $SERVER_URL/home"
curl -s -b $COOKIE_JAR "$SERVER_URL/home"
echo ""
echo ""

# 4. 显示 cookie 内容
echo -e "${YELLOW}4. Cookie 文件内容:${NC}"
if [ -f "$COOKIE_JAR" ]; then
    cat $COOKIE_JAR
else
    echo -e "${RED}Cookie 文件不存在${NC}"
fi
echo ""

echo -e "${GREEN}=== 测试完成 ===${NC}"
echo -e "${BLUE}提示: 使用 'rm $COOKIE_JAR' 清除 cookie${NC}"