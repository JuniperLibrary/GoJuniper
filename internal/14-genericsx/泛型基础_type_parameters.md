# Go 基础：泛型（type parameters）

这个文档对应 [internal/14-genericsx](./)。

配合阅读：
- 实现代码：[genericsx.go](./genericsx.go)
- 测试代码：本目录下暂无 `*_test.go`，建议先读实现，后续补测试（见练习清单）

---

## 1. 为什么要学泛型

在没有泛型时，你会遇到两类选择：
- 写一堆重复函数（`MapInts`、`MapStrings`...）
- 或者用 `any`，到处做类型断言（不安全、体验差）

泛型提供了第三条路：
- 代码复用
- 仍然保持类型安全

> ⚠️ **注意**：泛型**不是运行期反射**——Go 的泛型是编译期单态化（编译器为用到的类型生成具体代码），所以没有运行时类型断言开销，类型错误在编译期就报出来。这和 Java 的泛型不同——Java 泛型是**类型擦除**（运行时无类型信息），而 Go 和 Java 都有编译期类型检查；Go 的 `fn largest[T cmp.Ordered]` 类似 Java 的 `<T extends Comparable<T>>` 在编译期就针对每个 `T` 生成/检查一份代码。

**为什么**：用 `any` 写 `Map` 虽然能复用，但返回值是 `any`，调用方必须类型断言，既不安全又啰嗦。泛型让“一次编写、多种类型、零运行时成本”成为可能。

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

> ⚠️ **注意**：类型参数写在**函数名后面的方括号** `[T any]` 里，位置在 `(参数)` 之前。初学者常把它写成 `func Identity(v T) T` 漏掉方括号——那样 `T` 会被当成普通类型名而找不到定义。调用时可省略（由编译器推断 `Identity(42)`），也可显式写明 `Identity[int](42)`。

---

## 3. 典型三件套：Map / Filter / Reduce

本仓库的练习聚焦这三种模式：
- `Map`：把 `[]T` 变成 `[]U`
- `Filter`：保留满足条件的元素
- `Reduce`：把 `[]T` 聚合成一个值

```go
func Map[A any, B any](xs []A, f func(A) B) []B {
	out := make([]B, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

func Filter[T any](xs []T, pred func(T) bool) []T {
	out := make([]T, 0, len(xs))
	for _, x := range xs {
		if pred(x) {
			out = append(out, x)
		}
	}
	return out
}

func Reduce[T any, Acc any](xs []T, acc Acc, f func(Acc, T) Acc) Acc {
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}
```

建议学习顺序：
1. 先看实现，理解类型参数怎么传递：[genericsx.go](./genericsx.go)
2. 再补测试，理解调用方式（本目录暂无测试，见练习清单）

**常见坑**：多个类型参数时用逗号分隔（`[A any, B any]`），每个都要单独声明约束；`Map` 的入参和出参类型不同（`A`→`B`），别误以为只能同类型。

---

## 4. GetLargest：有约束的泛型（对应 Java Comparable）

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

这直接对应 Java 中 `<T extends Comparable<T>> T getLargest(List<T> list)` 的模式：

```java
// Java 实现
public static <T extends Comparable<T>> T getLargest(List<T> list) {
    T largest = list.get(0);
    for (T item : list) {
        if (item.compareTo(largest) > 0) {
            largest = item;
        }
    }
    return largest;
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

```java
// Java：整数列表
List<Integer> numbers = List.of(34, 50, 25, 100, 65);
Integer result = getLargest(numbers); // 100

// Java：字符列表
List<Character> chars = List.of('y', 'm', 'a', 'q');
Character result2 = getLargest(chars); // 'y'
```

> ⚠️ **注意**：`cmp.Ordered` 是 **Go 1.21+** 才引入的预定义约束，表示“可比较大小”的类型。如果你写 `T any` 却在函数体里用 `v > largest`，**编译会直接报错**——因为 `any` 不保证支持 `>` 运算符。约束的本质是“告诉编译器 T 必须具备哪些能力”。同样地，本仓库的 `FindIndex`/`Contains` 用到 `T comparable`（可 `==` 比较），`Stack`/`SafeMap` 用到 `K comparable`，这些都是编译期约束，不满足就编译失败，而非运行时 panic。

> ⚠️ **注意**：**实例化时类型必须明确**。多数情况编译器能推断，但在链式调用或函数作为值时可能推断不出，需要显式写出，例如 `genericsx.Map[int, string](nums, strconv.Itoa)`。约束不匹配（比如给 `GetLargest` 传一个不可比较/不可排序的自定义 struct）会在编译期直接红掉，这正是泛型“类型安全”的价值——把错误挡在运行之前。

**常见坑**：把“约束”和“运行时判断”混淆。约束不参与运行，只在编译期检查；它也不会帮你做运行时类型转换。

---

## 5. 初学者常见坑

1. 把泛型当“运行期反射”理解（不是的，它仍然是静态类型检查）
2. 过早写复杂约束（初学先掌握 `any` 与最简单约束即可）
3. 为了泛型而泛型（能用普通函数更清晰时，就先用普通函数）

> ⚠️ **注意**：第 3 点特别重要——Go 泛型适合“同一份逻辑对多种类型都成立”的场景（如 `Map`/`Filter`/`SafeMap`）。如果逻辑只针对一种类型，**普通函数反而更清晰、编译更快**。不要为了“看起来高级”强行上泛型。这点和 Java 的泛型取舍一致：能具体就具体，需要抽象再抽象。

---

## 6. 练习清单

1. 写 `Contains[T comparable](xs []T, v T) bool`（对应 Java 的 `list.contains(v)`）
2. 写 `GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T`
3. 写 `MapError[T any, U any](xs []T, f func(T) (U, error)) ([]U, error)`
4. 参考 `GetLargest` 写 `GetSmallest[T cmp.Ordered](xs []T) (T, bool)`

---

## 7. 自测命令

```bash
go test ./internal/14-genericsx -shuffle on
```

> 当前目录暂未包含 `*_test.go`，运行上述命令时若提示 “no test files” 属正常。建议完成练习清单后补上测试，再执行该命令验证。
