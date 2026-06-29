# 07 — JSON 编解码：struct tag / marshal / unmarshal

## 是什么

Go 标准库 `encoding/json` 提供了 Go 值 ↔ JSON 字节的互相转换（marshal / unmarshal）。你不需要手写解析器，只需要在 struct 字段上标注 `json:"字段名"` 的 tag，标准库自动完成映射。

核心函数只有两个：

```go
data, err := json.Marshal(v)     // Go 值 → JSON bytes
err = json.Unmarshal(data, &v)   // JSON bytes → Go 值（v 必须是指针）
```

## 怎么用

定义 struct 并标注 tag：

```go
type Person struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Nickname *string `json:"nickname,omitempty"`
}
```

编码：

```go
func EncodePerson(p Person) ([]byte, error) {
    return json.Marshal(p)
}
```

解码：

```go
func DecodePerson(b []byte) (Person, error) {
    var p Person
    if err := json.Unmarshal(b, &p); err != nil {
        return Person{}, err
    }
    return p, nil
}
```

`omitempty` 的效果：当 `Nickname` 为 nil（指针零值）时，编码结果中不会出现 `"nickname"` 字段。这让输出更干净，也避免泄漏空值。

## 为什么（对比 Rust）

| Go | Rust | 差异要点 |
|----|------|----------|
| `json.Marshal` / `json.Unmarshal` | `serde_json::to_string` / `serde_json::from_str` | 功能一致。但 serde 需要派生 `Serialize` / `Deserialize`，Go 直接反射，无需显式声明 |
| `json:"field_name"` | `#[serde(rename = "field_name")]` | 都是编译期/运行时告诉序列化框架字段名映射。Go 用反射在运行时解析 tag，serde 用过程宏在编译期生成代码 |
| `omitempty` | `#[serde(skip_serializing_if = "Option::is_none")]` | Go 的 omitempty 对零值（0、""、false、nil）全部跳过，控制粒度较粗；serde 可以精确控制某个字段的跳过条件 |
| `*string` + omitempty | `Option<String>` | Go 没有 `Option`，用指针模拟"可选字段"；Rust 用 `Option<T>` 更安全明确 |

Go 选择运行时反射而非编译期代码生成：好处是不需要额外工具链，坏处是性能略低且错误只能在运行时暴露（比如 tag 拼写错了不会编译报错）。

## 常见坑

1. **Unmarshal 忘了传指针**：`json.Unmarshal(data, v)` 编译能过但运行报错。必须传 `&v`。
2. **字段未导出（小写开头）**：Go 的反射只能看到大写开头的导出字段。小写字段在 marshal/unmarshal 时会被静默忽略，没有警告。
3. **`string` vs `[]byte` 混用**：`json.Marshal` 返回 `[]byte`，不要用 `string(data)` 去 Unmarshal（类型不匹配），但可以用 `[]byte(s)` 包装 string。
4. **未知字段的静默忽略**：Unmarshal 时输入 JSON 包含 struct 中不存在的字段，默认静默忽略。如果想检测，用 `json.NewDecoder` + `DisallowUnknownFields()`。
5. **`map[string]any` 后遗症**：用泛化 map 解析后，取字段要不断做类型断言（`v.(float64)`、`v.(string)`），代码冗长易错。尽量用 struct。

## 对应代码

- 实现：`internal/07-jsonx/jsonx.go`
- 测试：`internal/07-jsonx/jsonx_test.go`

```bash
go test ./internal/07-jsonx -shuffle on
```
