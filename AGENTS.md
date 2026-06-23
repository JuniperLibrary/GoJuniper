# AGENTS.md — GoJuniper

Go 学习练习仓库。从基础到进阶，每个主题独立包 + 测试 + 文档。

## 项目结构

```
GoJuniper/
  cmd/                # 可执行程序入口
    hello/            # 最小 demo
    server/           # HTTP 服务 + flag 参数 + 优雅退出
  internal/           # 学习模块（每个 topic 配测试 + 知识文档）
    basics/            # 变量/控制流/基本类型
    funcsx/            # 函数进阶：闭包/defer/变参
    collections/       # slice + map 操作
    errorsx/           # 错误处理
    typesx/            # struct/方法/接口/组合
    iox/               # I/O 操作
    jsonx/             # JSON 编解码
    timex/             # 时间处理
    contextx/          # Context 取消/超时
    concurrency/       # 并发编程
    channelsx/         # Channel 模式
    syncx/             # 同步原语
    httpx/             # HTTP 服务
    genericsx/         # 泛型
  _legacy/            # 历史学习代码封存
  gin_learn/          # Gin web 框架学习
  docs/               # 学习笔记（按阶段组织）
  go.mod
  README.md
  AGENTS.md
```

## 常用命令

```bash
go test ./...                     # 运行所有测试
go test -shuffle on ./internal/basics/...  # 随机顺序跑单个模块测试
go run ./cmd/hello                # 运行 hello 示例
go run ./cmd/server -addr :8080   # 运行 http 服务
go vet ./...                      # 静态检查
gofmt -w .                        # 格式化代码
```

## 重要说明

- **Go 1.25.0** — module `gojuniper`
- **每个 internal 模块都有 `*_test.go`** — 通过测试驱动学习
- **知识文档在各模块目录下** — 以中文 *.md 形式存在，与代码同目录
- **`_legacy/` 被 Go 工具链自动忽略** — 旧代码存档，不影响构建
- **`gin_learn/` 是独立的 Gin 框架学习** — 与基础模块分离
- **无 CI / 无 Makefile** — `go test` 和 `go vet` 即验证手段
- **README.md 有完整的 Rust ↔ Go 概念对照表**

## 本仓库用途

渐进式 Go 学习：基本类型 → 函数 → 集合 → 错误处理 → 类型系统 → I/O → JSON → 时间 → Context → 并发 → Channel → 同步 → HTTP → 泛型。每个概念有可运行代码 + 测试 + 知识文档。

## 用户协作规范（首次对话请先阅读此处）

### 沟通风格
- 全程**中文**交流
- 语言简洁直接，不要客套/寒暄/表扬
- 不要一次性给出全部信息，要**步步引导**

### 教学模式（核心 — 不遵守会被纠正）
1. **先动手，后讲原理** — 先让用户运行现有代码、观察现象，再解释 Why
2. **一次只教一个概念** — 不超前引入未学过的知识点
3. **不要直接给答案** — 用提问引导用户自己发现模式
4. **每一步都要用户动手写/改代码**，不能只是阅读
5. 学习新的 module（如 `gin`）时，用 `go get` 添加依赖，让用户感受 Go 的工具链

### 文档输出与提交规范（每次必做）
1. **每一课都要输出**：
   - `docs/` 下的对应笔记文件（如 `docs/01-基础/01-变量与类型.md`）
   - 文件格式：是什么 → 怎么用 → 为什么 → 常见坑
   - 学习过程回顾总结（放在笔记末尾或单独说明）
2. 用 `todowrite` 管理学习进度，一课一个 todo，完成后立即标记
3. **每一课结束后必须提交代码**：用 `git commit` 提交当课的所有改动，commit message 格式为 `docs: 第X课 - <课程主题>`

### 学习路径编排
- 严格按照 TODO.md 的阶段路线图推进（阶段零 → 阶段五）
- 已完成的 `internal/` 模块按阶段跳学或复习，不允许跳阶段
- 语言基础未学完前（阶段五未完成前），不进入 FUTURE.md 的进阶内容
- 进阶学习（阶段六～九）参考 [docs/FUTURE.md](docs/FUTURE.md) 路线图
