# typesx：类型系统学习

本目录的知识文档：
- [结构体_方法_接口与组合.md](./结构体_方法_接口与组合.md)
- [类型断言与组合.md](./类型断言与组合.md)

实现与测试：
- [typesx.go](./typesx.go)
- 单元测试统一放在 [internal/16-tests/typesx_typesx_test.go](../16-tests/typesx_typesx_test.go)

---

## 学习内容

1. **结构体**：定义、创建、字段访问
2. **方法**：值接收者 vs 指针接收者
3. **接口**：隐式实现、小接口原则
4. **组合（Embedding）**：用组合代替继承
5. **类型断言**：从接口提取具体类型
6. **Type Switch**：处理多种类型

## 自测命令

```bash
go test ./internal/04-typesx -shuffle on
```
