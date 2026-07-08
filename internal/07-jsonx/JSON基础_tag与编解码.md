# Go 基础：JSON（struct tag / 编解码）

这个文档对应 [internal/07-jsonx](file:///e:/dingchuan/GoJuniper/internal/07-jsonx)。

配合阅读：
- 实现代码：[jsonx.go](file:///e:/dingchuan/GoJuniper/internal/07-jsonx/jsonx.go)

> ⚠️ **注意**：本仓库 `07-jsonx` 目录下只有实现代码 `jsonx.go`，没有 `_test.go`。学习时直接看 `jsonx.go` 里的 `Person` 结构体与 `EncodePerson`/`DecodePerson` 即可，练习清单里的函数也写在同包内。

---

## 1. 为什么推荐用 struct 做 JSON

你当然可以用 `map[string]any`，但初学阶段更推荐 `struct`：
- 有明确字段（更不容易拼错 key）
- 可以用类型约束（例如 int/string）
- IDE 更好提示

> ⚠️ **注意**：`encoding/json` 只能编解码**导出字段**（首字母大写）。小写的 `name string` 会被彻底忽略——既不编码进 JSON，也从 JSON 解不出来，而且**不报错**。这是 JSON 模块最常见的“静默失效”陷阱。

**为什么**用 struct 更好？Java 里你大概会用 Jackson 的 `@JsonProperty("...")` 注解，Go 的 struct tag 扮演了类似的角色（如 `json:"name"`），但 Go 没有编译期的字段校验，拼错 tag 同样不会报错，只能靠测试发现。

---

## 2. Marshal 与 Unmarshal

最常见两个函数：

```go
data, err := json.Marshal(v)
err = json.Unmarshal(data, &v)
```

注意点：
- `Unmarshal` 的第二个参数必须是指针（因为要写回结果）

> ⚠️ **注意**：`json.Unmarshal(data, &v)` 传的是 `&v`（指针）。如果写成 `json.Unmarshal(data, v)` 会编译不过——`Unmarshal` 形参是 `interface{}` 但内部用反射取地址写回，传值会导致无法赋值。本仓库 `DecodePerson` 里的 `var p Person; json.Unmarshal(b, &p)` 就是正确写法。

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

> ⚠️ **注意**：tag 格式是反引号包裹的 ``` `json:"name"` ```，冒号后**不要有空格**的位置其实很宽松，但值里如果带选项要用逗号分隔，例如 ``` `json:"name,omitempty"` ```。本仓库 `Person` 用了 `json:"id"`、`json:"name"`，而 `Nickname` 用了 `json:"nickname,omitempty"`。

---

## 4. omitempty：空值不输出

```go
type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
```

当字段是“空值”（例如 `""`、`0`、`false`、`nil`）时，Marshal 会省略该字段。

> ⚠️ **注意（omitempty 对零值 struct 无效）**：`omitempty` 只对“空值”生效，而 `time.Time{}`、自定义 struct 等即使全是零值也**不算空值**（它们不是 `""`/`0`/`false`/`nil`），所以 `omitempty` 不会省略零值 struct 字段。

> ⚠️ **注意（用指针区分“未设置”与“零值”）**：对于 `string`，你没法区分“JSON 里没给这个字段”和“给了空字符串 `""`”。解决办法是改用指针 `*string`（如本仓库 `Person.Nickname *string`）：指针为 `nil` 表示“未设置/被 omitempty 省略”，指向 `""` 表示“显式给了空串”。这在 API 更新语义（PATCH）里非常关键。

---

## 5. 本模块在练什么

请先看实现再看练习：
- 实现：[jsonx.go](file:///e:/dingchuan/GoJuniper/internal/07-jsonx/jsonx.go)

重点关注：
- 结构体 tag 的使用
- JSON 输入输出的错误处理（永远检查 `err`）

> ⚠️ **注意**：`time.Time` 的 JSON 默认编码格式是 **RFC3339**（`2006-01-02T15:04:05Z07:00`）。如果你期望别的格式（如 `"2026-03-11"`），需要自己实现 `MarshalJSON`/`UnmarshalJSON`，否则直接 `Marshal` 出来的字符串格式可能和前端约定不符。

---

## 6. 初学者常见坑

1. Unmarshal 忘了传指针：`json.Unmarshal(data, v)`（会报错）
2. 字段没导出（小写开头）导致 Unmarshal/Marshal 无效
3. 用 `map[string]any` 后到处做类型断言，很痛苦

> ⚠️ **注意**：第 2 条是“静默”的——不会编译报错，运行也不 panic，只是字段始终是零值。排错时优先确认字段名是否大写开头。

---

## 7. 练习清单

1. 定义 `type Order struct { ID string; Amount int }`，写 Marshal/Unmarshal 测试
2. 给字段加 `omitempty`，观察输出差异
3. 写一个 `DecodeUser(r io.Reader) (User, error)`（从 Reader 解 JSON）

---

## 8. 自测命令

```bash
go test ./internal/07-jsonx -shuffle on
```

> ⚠️ **注意**：当前目录没有测试文件，该命令会提示 “no test files”。先补 `jsonx_test.go` 或把练习函数写好再运行。
