#!/bin/bash

# 启动 cookie_middleware 服务器的 shell 脚本
# 适用于 macOS

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

# 检查是否存在可执行文件
if [ ! -f "./cookie_middleware" ]; then
    echo -e "${RED}错误: 找不到 cookie_middleware 可执行文件${NC}"
    echo -e "${BLUE}请先运行: go build${NC}"
    exit 1
fi

# 检查服务器是否已经在运行
if lsof -Pi :$SERVER_PORT -sTCP:LISTEN -t >/dev/null ; then
    echo -e "${YELLOW}警告: 端口 $SERVER_PORT 已被占用${NC}"
    echo -e "${BLUE}是否强制重启服务器？ (y/n)${NC}"
    read -r response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
        # 查找并杀死占用端口的进程
        pid=$(lsof -t -i:$SERVER_PORT)
        kill -9 $pid 2>/dev/null || true
        echo -e "${GREEN}已停止占用端口的进程${NC}"
        sleep 1
    else
        echo -e "${RED}退出${NC}"
        exit 1
    fi
fi

echo -e "${BLUE}=== 启动 Cookie 中间件服务器 ===${NC}"
echo -e "${GREEN}服务器地址: $SERVER_URL${NC}"
echo -e "${YELLOW}按 Ctrl+C 停止服务器${NC}"
echo ""

# 启动服务器
./cookie_middleware