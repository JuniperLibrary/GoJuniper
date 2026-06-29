# Go 基础：泛型（type parameters）

这个文档对应 [internal/14-genericsx](file:///e:/dingchuan/GoJuniper/internal/14-genericsx)。

配合阅读：
- 实现代码：[genericsx.go](file:///e:/dingchuan/GoJuniper/internal/14-genericsx/genericsx.go)
- 测试代码：[genericsx_test.go](file:///e:/dingchuan/GoJuniper/internal/14-genericsx/genericsx_test.go)

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
1. 先看测试，理解调用方式：[genericsx_test.go](file:///e:/dingchuan/GoJuniper/internal/14-genericsx/genericsx_test.go)
2. 再看实现，理解类型参数怎么传递：[genericsx.go](file:///e:/dingchuan/GoJuniper/internal/14-genericsx/genericsx.go)

---

## 4. GetLargest：有约束的泛型（对应 Rust PartialOrd）

除了 `any`，Go 泛型还支持带约束的类型参数。`cmp.Ordered` 是 Go 1.21+ 内置的约束，表示"可排序类型"（int、string、float64、rune 等）：

```go
func GetLargest[T cmp.Ordered](xs []T) (T, bool) {
    if len(xs) == 0 {
        var zero T
        return zero, false
    }
    largest := xs[0]
    for _, v := range xs[1:] {
        if v > largest {
            largest = v
        }
    }
    return largest, true
}
```

这直接对应 Rust 中 `fn get_largest<T: PartialOrd>(list: &[T]) -> &T` 的模式：

```rust
// Rust 实现
fn get_largest<T: PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];
    for item in list {
        if item > largest {
            largest = item;
        }
    }
    largest
}
```

测试用例如下：

```go
// Go：整数列表
numbers := []int{34, 50, 25, 100, 65}
got, _ := genericsx.GetLargest(numbers) // 100

// Go：rune 列表
chars := []rune{'y', 'm', 'a', 'q'}
got, _ = genericsx.GetLargest(chars) // 'y'
```

```rust
// Rust：整数列表
let numbers = vec![34, 50, 25, 100, 65];
let result = get_largest(&numbers); // 100

// Rust：字符列表
let chars = vec!['y', 'm', 'a', 'q'];
let result = get_largest(&chars); // 'y'
```

---

## 5. 初学者常见坑

1. 把泛型当“运行期反射”理解（不是的，它仍然是静态类型检查）
2. 过早写复杂约束（初学先掌握 `any` 与最简单约束即可）
3. 为了泛型而泛型（能用普通函数更清晰时，就先用普通函数）

---

## 6. 练习清单

1. 写 `Contains[T comparable](xs []T, v T) bool`（对应 Rust `slice.contains(&v)`）
2. 写 `GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T`
3. 写 `MapError[T any, U any](xs []T, f func(T) (U, error)) ([]U, error)`
4. 参考 `GetLargest` 写 `GetSmallest[T cmp.Ordered](xs []T) (T, bool)`

---

## 7. 自测命令

```bash
go test ./internal/14-genericsx -shuffle on
```

