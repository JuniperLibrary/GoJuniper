# Go 基础：泛型（type parameters）

这个文档对应 [internal/genericsx](file:///e:/dingchuan/GoJuniper/internal/genericsx)。

配合阅读：
- 实现代码：[genericsx.go](file:///e:/dingchuan/GoJuniper/internal/genericsx/genericsx.go)
- 测试代码：[genericsx_test.go](file:///e:/dingchuan/GoJuniper/internal/genericsx/genericsx_test.go)

---

## 1. 为什么要学泛型

在没有泛型时，你会遇到两类选择：
- 写一堆重复函数（`MapInts`、`MapStrings`...）
- 或者用 `any`，到处做类型断言（不安全、体验差）

泛型提供了第三条路：
- 代码复用
- 仍然保持类型安全

---

## 2. 最小例子：T any

```go
func Identity[T any](v T) T {
	return v
}
```

你可以把 `[T any]` 理解成：
- 这个函数有一个类型参数 `T`
- `any` 表示 T 可以是任何类型

---

## 3. 典型三件套：Map / Filter / Reduce

本仓库的练习聚焦这三种模式：
- `Map`：把 `[]T` 变成 `[]U`
- `Filter`：保留满足条件的元素
- `Reduce`：把 `[]T` 聚合成一个值

建议学习顺序：
1. 先看测试，理解调用方式：[genericsx_test.go](file:///e:/dingchuan/GoJuniper/internal/genericsx/genericsx_test.go)
2. 再看实现，理解类型参数怎么传递：[genericsx.go](file:///e:/dingchuan/GoJuniper/internal/genericsx/genericsx.go)

---

## 4. 初学者常见坑

1. 把泛型当“运行期反射”理解（不是的，它仍然是静态类型检查）
2. 过早写复杂约束（初学先掌握 `any` 与最简单约束即可）
3. 为了泛型而泛型（能用普通函数更清晰时，就先用普通函数）

---

## 5. 练习清单

1. 写 `Contains[T comparable](xs []T, v T) bool`
2. 写 `GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T`
3. 写 `MapError[T any, U any](xs []T, f func(T) (U, error)) ([]U, error)`

---

## 6. 自测命令

```bash
go test ./internal/genericsx -shuffle on
```

