# OpenCode + oh-my-opencode 入门教程

OpenCode 是一个开源的 AI 编程代理（AI coding agent），支持在终端中与 AI 交互完成代码相关任务。oh-my-opencode 是其增强插件，将单个 AI 代理升级为多智能体协作团队。

本教程将从零开始，带你完成 **安装 → 配置 → 使用 → 进阶** 的全流程。

---

## 目录

- [第一章：安装 OpenCode](#第一章安装-opencode)
- [第二章：配置 AI 模型](#第二章配置-ai-模型)
- [第三章：基本使用](#第三章基本使用)
- [第四章：安装 oh-my-opencode](#第四章安装-oh-my-opencode)
- [第五章：oh-my-opencode 使用](#第五章oh-my-opencode-使用)
- [第六章：实战案例](#第六章实战案例)
- [附录](#附录)

---

## 第一章：安装 OpenCode

### 1.1 系统要求

- macOS / Linux / Windows（推荐 WSL）
- 终端模拟器（推荐 [WezTerm](https://wezterm.org)、[Ghostty](https://ghostty.org) 或 [Kitty](https://sw.kovidgoyal.net/kitty/)）

### 1.2 一键安装（推荐）

这是最简单的方法，适用于 macOS 和 Linux：

```bash
curl -fsSL https://opencode.ai/install | bash
```

安装完成后验证：

```bash
opencode --version
```

如果输出类似 `1.3.6` 的版本号，说明安装成功。

### 1.3 其他安装方式

#### 使用 npm

```bash
npm install -g opencode-ai
```

#### 使用 Homebrew（macOS / Linux）

```bash
brew install anomalyco/tap/opencode
```

#### 使用 Bun

```bash
bun install -g opencode-ai
```

#### Windows

推荐使用 WSL，然后按 Linux 方式安装。

也可直接用包管理器：

```bash
# Chocolatey
choco install opencode

# Scoop
scoop bucket add extras
scoop install extras/opencode
```

#### Arch Linux

```bash
paru -S opencode-bin
```

### 1.4 桌面应用

OpenCode 也提供桌面端，可从 [GitHub Releases](https://github.com/anomalyco/opencode/releases) 下载。

| 系统平台 | 下载文件 |
|----------|----------|
| macOS（苹果芯片） | `opencode-desktop-darwin-aarch64.dmg` |
| macOS（英特尔） | `opencode-desktop-darwin-x64.dmg` |
| Windows | `opencode-desktop-windows-x64.exe` |
| Linux | `.deb`、`.rpm` 或 `AppImage` |

---

## 第二章：配置 AI 模型

OpenCode 支持 75+ 家模型提供商，你可以使用免费模型，也可以接入付费 API。

### 2.1 使用免费模型（零成本起步）

OpenCode 内置了免费模型，无需 API Key 即可使用。

启动 OpenCode：

```bash
opencode
```

输入 `/models` 查看可用模型，选择带有 **Free** 标签的模型即可，例如：

- `opencode/glm-5` — GLM-5，中文能力强
- `opencode/big-pickle` — 免费 GLM-4.6

### 2.2 配置 API 密钥

如果你有付费订阅（如 Claude、ChatGPT），可以配置 API Key 获取更强的模型。

**方式一：交互式配置**

在 OpenCode TUI 中输入：

```
/connect
```

选择你的提供商（如 Anthropic、OpenAI），按提示完成认证。

**方式二：命令行配置**

```bash
opencode auth login
```

**方式三：环境变量**

在 `~/.bash_profile` 或 `~/.zshrc` 中添加：

```bash
# Anthropic
export ANTHROPIC_API_KEY="sk-ant-..."

# OpenAI
export OPENAI_API_KEY="sk-..."

# Google Gemini
export GOOGLE_GENERATIVE_AI_API_KEY="..."
```

### 2.3 常用提供商配置指南

#### Anthropic（Claude）

1. 输入 `/connect`，选择 `Anthropic`
2. 选择认证方式：
   - **Claude Pro/Max**：浏览器 OAuth 认证（需订阅）
   - **手动输入 API Key**：从 [console.anthropic.com](https://console.anthropic.com/) 获取

#### OpenAI（ChatGPT）

1. 输入 `/connect`，选择 `OpenAI`
2. 选择认证方式：
   - **ChatGPT Plus/Pro**：浏览器 OAuth 认证（需订阅）
   - **手动输入 API Key**：从 [platform.openai.com](https://platform.openai.com/) 获取

#### GitHub Copilot

1. 输入 `/connect`，选择 `GitHub Copilot`
2. 浏览器打开 GitHub 授权页面，完成设备授权

#### OpenCode Zen（官方推荐）

1. 输入 `/connect`，选择 `opencode`
2. 浏览器打开 [opencode.ai/auth](https://opencode.ai/auth)，注册获取 API Key
3. 粘贴到终端

#### 查看已配置的提供商

```bash
opencode providers list
```

---

## 第三章：基本使用

### 3.1 启动 OpenCode

进入你的项目目录：

```bash
cd /path/to/your/project
opencode
```

例如，创建一个测试项目：

```bash
mkdir my-project
cd my-project
opencode
```

### 3.2 初始化项目

在 OpenCode TUI 中输入：

```
/init
```

OpenCode 会分析项目结构，生成 `AGENTS.md` 文件，帮助 AI 理解你的项目。

> **建议**：将 `AGENTS.md` 提交到 Git。

### 3.3 用自然语言交互

初始化后，直接用自然语言描述需求：

```
在当前目录下创建一个 Express.js 服务器，支持 /hello 路由返回 JSON
```

AI 会自动创建文件、编写代码。

### 3.4 引用文件

使用 `@` 符号引用项目中的文件：

```
解释 @src/main.ts 中的认证逻辑
```

按 `@` 后可以模糊搜索项目文件。

### 3.5 Plan 模式与 Build 模式

| 模式 | 切换方式 | 说明 |
|------|----------|------|
| **Build 模式** | 默认 | AI 直接修改代码 |
| **Plan 模式** | 按 `Tab` | AI 只提供建议，不修改代码 |

**推荐工作流**：

1. 按 `Tab` 进入 Plan 模式
2. 描述需求，让 AI 制定方案
3. 审核方案，提供反馈
4. 按 `Tab` 切回 Build 模式
5. 输入"执行"让 AI 修改代码

### 3.6 撤销与重做

```
/undo    # 撤销上一次修改
/redo    # 重做已撤销的修改
```

> `/undo` 和 `/redo` 需要项目是 Git 仓库。

### 3.7 常用 Slash 命令速查

#### 核心配置

| 命令 | 说明 | 快捷键 |
|------|------|--------|
| `/connect` | 添加 AI 提供商 | — |
| `/init` | 初始化项目，生成 AGENTS.md | `Ctrl+X I` |
| `/models` | 切换模型 | `Ctrl+X M` |

#### 会话管理

| 命令 | 说明 | 快捷键 |
|------|------|--------|
| `/new` | 新建会话 | `/clear` / `Ctrl+X N` |
| `/sessions` | 列出历史会话 | `/resume` / `Ctrl+X L` |
| `/share` | 分享会话链接 | `Ctrl+X S` |
| `/compact` | 压缩上下文（节省 token） | `/summarize` / `Ctrl+X C` |

#### 编辑操作

| 命令 | 说明 | 快捷键 |
|------|------|--------|
| `/undo` | 撤销 | `Ctrl+X U` |
| `/redo` | 重做 | `Ctrl+X R` |

#### 辅助功能

| 命令 | 说明 | 快捷键 |
|------|------|--------|
| `/details` | 显示工具执行详情 | `Ctrl+X D` |
| `/thinking` | 显示推理过程 | — |
| `/theme` | 切换主题 | `Ctrl+X T` |
| `/help` | 帮助 | `Ctrl+X H` |
| `/editor` | 外部编辑器撰写消息 | `Ctrl+X E` |
| `/export` | 导出对话为 Markdown | `Ctrl+X X` |

#### 退出

| 命令 | 说明 | 快捷键 |
|------|------|--------|
| `/exit` | 退出 | `/quit` / `Ctrl+X Q` |

### 3.8 快捷键速查

| 快捷键 | 功能 |
|--------|------|
| `Tab` | 切换 Plan / Build 模式 |
| `@` | 模糊搜索项目文件 |
| `Ctrl+C` | 中断当前操作 |
| `Ctrl+D` | 退出 |
| `↑ / ↓` | 浏览历史消息 |

### 3.9 CLI 参数速查

```bash
# 启动交互式 TUI
opencode

# 指定模型
opencode -m anthropic/claude-sonnet-4-5

# 继续上次会话
opencode -c

# 非交互模式（单次提示）
opencode run "解释这个函数的作用"

# 指定项目目录
opencode -c /path/to/project

# 无插件模式（纯净）
opencode --pure
```

---

## 第四章：安装 oh-my-opencode

oh-my-opencode（简称 oMo）是 OpenCode 的增强插件，提供多智能体协作、多模型调度等高级功能。

### 4.1 核心特性

| 特性 | 说明 |
|------|------|
| `ultrawork` | 一个词激活所有 Agent 并行工作 |
| Discipline Agents | Sisyphus、Hephaestus、Prometheus 等专业 Agent |
| Hash-Anchored Edit | 哈希锚定编辑，消除编辑出错 |
| LSP + AST-Grep | IDE 级别代码分析 |
| 多模型路由 | 按任务类型自动选择最佳模型 |
| Claude Code 兼容 | 你的 hooks、commands、skills 全部可用 |

### 4.2 前提条件

- OpenCode 已安装（`opencode --version` ≥ 1.0.150）
- Node.js / npm 已安装

### 4.3 安装步骤

#### 步骤一：全局安装

```bash
npm install -g oh-my-opencode
```

#### 步骤二：了解你的订阅

安装器需要知道你有哪些 AI 订阅，以便配置最佳模型组合。

**你需要回答以下问题**：

| 问题 | 如果有 | 如果没有 |
|------|--------|----------|
| Claude Pro/Max？ | 普通 → `--claude=yes`，20x → `--claude=max20` | `--claude=no` |
| OpenAI/ChatGPT Plus？ | `--openai=yes` | `--openai=no` |
| Gemini？ | `--gemini=yes` | `--gemini=no` |
| GitHub Copilot？ | `--copilot=yes` | `--copilot=no` |
| OpenCode Go（$10/月）？ | `--opencode-go=yes` | `--opencode-go=no` |
| Z.ai Coding Plan？ | `--zai-coding-plan=yes` | `--zai-coding-plan=no` |

#### 步骤三：运行安装器

根据你的订阅情况，选择对应的命令执行：

**场景一：有 Claude + OpenAI（推荐）**

```bash
oh-my-opencode install --no-tui --claude=yes --openai=yes --gemini=no --copilot=no
```

**场景二：只有 Claude Max（20x 模式）**

```bash
oh-my-opencode install --no-tui --claude=max20 --openai=no --gemini=no --copilot=no
```

**场景三：只有 GitHub Copilot**

```bash
oh-my-opencode install --no-tui --claude=no --openai=no --gemini=no --copilot=yes
```

**场景四：使用 OpenCode Go（$10/月，性价比方案）**

```bash
oh-my-opencode install --no-tui --claude=no --openai=no --gemini=no --copilot=no --opencode-go=yes
```

**场景五：零成本方案（使用免费模型）**

```bash
oh-my-opencode install --no-tui --claude=no --openai=no --gemini=no --copilot=no
```

> **注意**：没有 Claude 订阅时，Sisyphus Agent（核心编排器）性能会大幅下降。

#### 步骤四：配置认证

根据你选择的订阅，配置对应的认证：

```bash
opencode auth login
```

在交互界面中选择提供商完成认证。

#### 步骤五：验证安装

```bash
# 检查插件是否注册
cat ~/.config/opencode/opencode.json
# 应包含 "oh-my-openagent" 在 plugin 数组中

# 运行诊断工具
oh-my-opencode doctor
```

### 4.4 推荐订阅方案

| 方案 | 费用 | 包含 | 适合 |
|------|------|------|------|
| 最佳体验 | ~$30/月 | Claude Pro + ChatGPT Plus | 专业开发者 |
| 性价比方案 | ~$11/月 | Kimi Code + GLM Coding Plan | 预算有限 |
| 零成本方案 | $0 | OpenCode Zen 免费模型 | 学习体验 |

---

## 第五章：oh-my-opencode 使用

### 5.1 ultrawork —— 一键启动

`ultrawork`（或简写 `ulw`）是 oMo 的核心功能。在 prompt 中包含这个词，所有 Agent 会自动激活：

```
帮我重构这个模块 ultrawork
```

```
修复所有 TypeScript 错误 ulw
```

Agent 会自动拆分任务、并行执行、持续工作直到完成。

### 5.2 Discipline Agents

oh-my-opencode 内置了多个专业 Agent：

| Agent | 角色 | 默认模型 | 做什么 |
|-------|------|----------|--------|
| **Sisyphus** | 主编排器 | Claude Opus 4.6 | 规划、调度、驱动任务 |
| **Hephaestus** | 深度工作者 | GPT-5.4 | 自主探索、研究、执行 |
| **Prometheus** | 战略规划师 | Claude Opus 4.6 | 面试式需求分析 |
| **Oracle** | 架构师 | GPT-5.4 | 架构决策、调试排错 |
| **Explore** | 快速搜索 | Grok Code Fast 1 | 代码库搜索 |
| **Librarian** | 文档搜索 | Claude Haiku 4.5 | 文档和代码搜索 |

### 5.3 Hash-Anchored Edit（哈希锚定编辑）

传统编辑工具的问题：AI 需要精确复述代码行，一旦文件变化就出错。

oMo 的解决方案：每行代码有哈希标识符：

```
11#VK| function hello() {
22#XJ|   return "world";
33#MB| }
```

AI 通过 `LINE#ID` 引用要修改的行。如果文件被修改，哈希不匹配，编辑会被拒绝。

**效果**：Grok Code Fast 1 的成功率从 6.7% → 68.3%。

### 5.4 实用命令

| 命令 | 说明 |
|------|------|
| `/init-deep` | 生成分层 AGENTS.md 文件 |
| `/start-work` | 启动 Prometheus 面试式规划 |
| `/ulw-loop` | Ralph Loop，自循环直到完成 |
| `oh-my-opencode doctor` | 诊断工具，检查配置状态 |

### 5.5 内置 MCP 服务

oMo 内置了以下 MCP 服务，无需额外配置：

| MCP | 用途 |
|-----|------|
| Exa | 网络搜索 |
| Context7 | 官方文档查询 |
| Grep.app | GitHub 代码搜索 |

### 5.6 配置文件

配置文件位置：`~/.config/opencode/oh-my-opencode.json`

示例配置：

```jsonc
{
  "$schema": "https://raw.githubusercontent.com/code-yeongyu/oh-my-openagent/dev/assets/oh-my-opencode.schema.json",
  "agents": {
    "sisyphus": { "model": "anthropic/claude-opus-4-6" },
    "hephaestus": { "model": "openai/gpt-5.4" },
    "explore": { "model": "github-copilot/grok-code-fast-1" }
  },
  "categories": {
    "visual-engineering": { "model": "google/gemini-3.1-pro" },
    "ultrabrain": { "model": "openai/gpt-5.4" },
    "quick": { "model": "openai/gpt-5.4-mini" }
  }
}
```

---

## 第六章：实战案例

### 案例一：创建一个简单的 Node.js API

```bash
# 1. 创建项目
mkdir my-api && cd my-api
npm init -y

# 2. 启动 OpenCode
opencode

# 3. 初始化
# （在 TUI 中输入）
/init

# 4. 创建 API
# （在 TUI 中输入）
创建一个 Express.js 服务，支持 /hello 路由返回 JSON { message: 'Hello World' }
```

### 案例二：大型重构（使用 ultrawork）

```bash
cd /path/to/your/project
opencode
```

在 TUI 中输入：

```
将整个项目从 REST API 迁移到 GraphQL ultrawork
```

Agent 会自动：
1. Sisyphus 分析项目结构
2. Prometheus 制定迁移方案
3. Hephaestus 深度执行代码修改
4. Oracle 处理架构问题
5. 并行处理，直到完成

### 案例三：调试排错

在 TUI 中输入：

```
这个测试为什么失败了？看 @src/auth/login.test.ts
```

或者使用 Prometheus 规划模式：

```
/start-work
```

然后按提示回答问题，AI 会制定调试方案。

### 案例四：代码审查

```
审查 @src/api/users.ts 的代码质量，检查安全性和性能问题
```

---

## 附录

### A. 常见问题

**Q：安装后 `oh-my-opencode` 命令找不到？**

```bash
# 检查全局 npm 包
npm ls -g oh-my-opencode

# 重新安装
npm install -g oh-my-opencode

# 检查 PATH
echo $PATH | grep npm
```

**Q：Sisyphus Agent 不工作？**

Sisyphus 强烈推荐 Claude Opus 4.6。没有 Claude 订阅时至少需要 OpenCode Go（$10/月）提供的 Kimi K2.5 或 GLM-5。

**Q：如何卸载 oh-my-opencode？**

1. 编辑 `~/.config/opencode/opencode.json`，从 `plugin` 数组中移除 `"oh-my-openagent"`
2. 删除配置：`rm -f ~/.config/opencode/oh-my-opencode.json`
3. 卸载包：`npm uninstall -g oh-my-opencode`

**Q：如何更新？**

```bash
npm update -g oh-my-opencode
opencode upgrade  # 更新 OpenCode 本体
```

**Q：AST-Grep 未安装？**

```bash
npm install -g @ast-grep/cli
```

### B. 配置文件模板

#### ~/.config/opencode/opencode.json

```jsonc
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["oh-my-openagent@latest"],
  "model": "anthropic/claude-sonnet-4-5",
  "autoupdate": true
}
```

#### ~/.config/opencode/oh-my-opencode.json

```jsonc
{
  "$schema": "https://raw.githubusercontent.com/code-yeongyu/oh-my-openagent/dev/assets/oh-my-opencode.schema.json",
  "agents": {
    "sisyphus": { "model": "anthropic/claude-opus-4-6" },
    "hephaestus": { "model": "openai/gpt-5.4" },
    "prometheus": { "model": "anthropic/claude-opus-4-6" },
    "oracle": { "model": "openai/gpt-5.4" },
    "explore": { "model": "github-copilot/grok-code-fast-1" },
    "librarian": { "model": "anthropic/claude-haiku-4-5" }
  }
}
```

### C. 参考链接

- OpenCode 官网：[https://opencode.ai](https://opencode.ai)
- OpenCode 文档：[https://opencode.ai/docs](https://opencode.ai/docs)
- OpenCode GitHub：[https://github.com/anomalyco/opencode](https://github.com/anomalyco/opencode)
- oh-my-opencode GitHub：[https://github.com/code-yeongyu/oh-my-openagent](https://github.com/code-yeongyu/oh-my-openagent)
- OpenCode Discord：[https://opencode.ai/discord](https://opencode.ai/discord)
- oh-my-opencode Discord：[https://discord.gg/PUwSMR9XNk](https://discord.gg/PUwSMR9XNk)
- 菜鸟教程参考：[https://www.runoob.com/ai-agent/opencode-coding-agent.html](https://www.runoob.com/ai-agent/opencode-coding-agent.html)
