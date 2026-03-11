# Go 基础：JSON（struct tag / 编解码）

这个文档对应 [internal/jsonx](file:///e:/dingchuan/GoJuniper/internal/jsonx)。

配合阅读：
- 实现代码：[jsonx.go](file:///e:/dingchuan/GoJuniper/internal/jsonx/jsonx.go)
- 测试代码：[jsonx_test.go](file:///e:/dingchuan/GoJuniper/internal/jsonx/jsonx_test.go)

---

## 1. 为什么推荐用 struct 做 JSON

你当然可以用 `map[string]any`，但初学阶段更推荐 `struct`：
- 有明确字段（更不容易拼错 key）
- 可以用类型约束（例如 int/string）
- IDE 更好提示

---

## 2. Marshal 与 Unmarshal

最常见两个函数：

```go
data, err := json.Marshal(v)
err = json.Unmarshal(data, &v)
```

注意点：
- `Unmarshal` 的第二个参数必须是指针（因为要写回结果）

---

## 3. struct tag：控制 JSON 字段名

```go
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
```

关键点：
- tag 的 key 是 `json`
- 值是 JSON 字段名

---

## 4. omitempty：空值不输出

```go
type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
```

当字段是“空值”（例如 `""`、`0`、`false`、`nil`）时，Marshal 会省略该字段。

---

## 5. 本模块在练什么

请先看测试再看实现：
- 测试：[jsonx_test.go](file:///e:/dingchuan/GoJuniper/internal/jsonx/jsonx_test.go)
- 实现：[jsonx.go](file:///e:/dingchuan/GoJuniper/internal/jsonx/jsonx.go)

重点关注：
- 结构体 tag 的使用
- JSON 输入输出的错误处理

---

## 6. 初学者常见坑

1. Unmarshal 忘了传指针：`json.Unmarshal(data, v)`（会报错）
2. 字段没导出（小写开头）导致 Unmarshal/Marshal 无效
3. 用 `map[string]any` 后到处做类型断言，很痛苦

---

## 7. 练习清单

1. 定义 `type Order struct { ID string; Amount int }`，写 Marshal/Unmarshal 测试
2. 给字段加 `omitempty`，观察输出差异
3. 写一个 `DecodeUser(r io.Reader) (User, error)`（从 Reader 解 JSON）

---

## 8. 自测命令

```bash
go test ./internal/jsonx -shuffle on
```

