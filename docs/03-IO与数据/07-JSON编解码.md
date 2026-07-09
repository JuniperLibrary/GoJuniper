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

> ⚠️ **注意**：**只有导出的（大写开头）字段才会被编解码**。小写开头字段在 marshal/unmarshal 时会被静默忽略，连警告都没有。这是反射的限制，也是最常见的"字段丢了"根源。

> ⚠️ **注意**：`Unmarshal` 的第二个参数**必须传指针**（`&p`）。写成 `json.Unmarshal(b, p)` 编译能过（因为 `p` 是可寻址变量会被取地址？实际不会，且语义错误），运行时空结构体。指针才能让库把数据写回原变量。

### 为什么用 `*string` + omitempty

`*string` 模拟 Java 的 `Optional<String>`：Go 没有 `Optional`，用指针（nil 表示"无值"）。`omitempty` 对指针而言，**只有 `nil` 才会被省略**——如果指向空字符串 `""`，仍会出现在输出里（空字符串不是零值的指针）。

> ⚠️ **注意**：`omitempty` 对指针的"省略"语义是"指针为 nil 才省略"。它不同于某些人的直觉（"值为空就省略"）：一个 `*string` 指向 `""` 时不会被省略，只是指向 `nil` 时才会。

## 为什么（对比 Java）

| Go | Java | 差异要点 |
|----|------|----------|
| `json.Marshal` / `json.Unmarshal` | Jackson `objectMapper.writeValueAsString` / `readValue` | 功能一致。但 Jackson 通过注解（如 `@JsonProperty`）配置，Go 直接反射，无需显式声明 |
| `json:"field_name"` | `@JsonProperty("field_name")` | 都是告诉序列化框架字段名映射。Go 用反射在运行时解析 tag，Jackson 用注解（运行时反射） |
| `omitempty` | `@JsonInclude(Include.NON_NULL)` | Go 的 omitempty 对零值（0、""、false、nil）全部跳过，控制粒度较粗；Jackson 可精确控制（NON_NULL / NON_EMPTY） |
| `*string` + omitempty | `Optional<String>` / `@Nullable String` | Go 没有 `Optional`，用指针模拟"可选字段"；Java 用 `Optional` 或 `@Nullable` 更安全明确 |

Go 选择运行时反射而非编译期代码生成：好处是不需要额外工具链，坏处是性能略低且错误只能在运行时暴露（比如 tag 拼写错了不会编译报错）。Java 的 Jackson 同样基于运行时反射/注解。

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
