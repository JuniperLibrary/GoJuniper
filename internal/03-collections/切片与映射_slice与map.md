# Go 基础：切片与映射（slice 与 map）

这个文档对应 [internal/03-collections](./)。

配合阅读：
- 实现代码：[collections.go](./collections.go)
- 测试代码：[collections_test.go](./collections_test.go)

---

## 1. slice 是什么

`slice`（切片）是 Go 里最常用的"可变长度序列"。你可以把它理解成：
- 逻辑上：一段连续的元素序列
- 内部：指向底层数组的一段"视图"（包含指针、长度 len、容量 cap）

你会经常用到三个操作：

```go
xs := []int{1, 2, 3}
_ = len(xs)
_ = cap(xs)
xs = append(xs, 4)
```

### 1.1 make：预分配容量

如果你预估最终要装多少元素，建议提前分配容量，减少扩容次数：

```go
out := make([]int, 0, 100) // len=0, cap=100
```

[UniqueInts](./collections.go#L9-L21) 就用了这种思路：`out := make([]int, 0, len(xs))`。

> ⚠️ **注意**：**`append` 的返回值必须接收**——写成 `s = append(s, x)`，不能只写 `append(s, x)` 然后丢弃结果。因为当底层数组容量不够时，Go 会分配新数组并复制，原切片变量不更新就会"丢了这次追加"。对应 Java 的 `ArrayList.add` 是原地改（自动扩容），Go 的 slice 是"描述符"，append 可能换底层数组，所以必须重新赋值。

> ⚠️ **注意**：**slice 直接赋值是共享底层数组**（`dst := src` 只是复制了"指针+len+cap"这个描述符），改一边会影响另一边。要真正独立拷贝得用 `copy(dst, src)`（见 `SliceCopyDemo`）。为什么这样设计？因为 slice 描述符很轻量，默认不拷贝数据——代价就是"看起来像值，实际是引用语义"，这是最常见坑之一。

---

## 2. map 是什么

`map[K]V` 是哈希表（键值对）。你最常用的是：

```go
m := map[string]int{}
m["a"]++
v := m["a"]
```

### 2.1 map 的"零值"是 nil

```go
var m map[string]int // nil
```

`nil map` 不能写入（会 panic），但可以读取（读到 value 的零值）：

```go
var m map[string]int
_ = m["a"] // 0
// m["a"] = 1 // panic
```

所以一般写法是用 `make` 初始化：

```go
m := make(map[string]int)
```

> ⚠️ **注意**：**`nil` map 可以读（读到零值）、但不能写（写会 panic）**。`var m map[string]int` 是 nil，必须先 `make` 或字面量 `map[string]int{}` 初始化才能写。对应 Java 的 `HashMap` 默认就是空的、可写（`new HashMap<>()`）；Go 的 nil map 是一个需要特别小心的存在（未初始化即不可用）。

> ⚠️ **注意**：**map 永远用 `make`，绝不用 `new`**。`make(map[K]V)` 返回已初始化、可直接写入的 map；而 `new(map[K]V)` 返回 `*map`，且指向的是 nil map，写入会 panic（见 `MakeVsNew`）。

### 2.2 读 map 的 "comma ok" 写法

```go
v, ok := m["key"]
```

在 [UniqueInts](./collections.go#L10-L19) 里就用 `ok` 判断是否出现过。

> ⚠️ **注意**：**读不存在的 key 不会报错，只会返回 value 类型的零值**。所以光看 `v := m["key"]` 无法区分"key 不存在"和"key 存在但值就是零值"。需要区分时用 `v, ok := m["key"]`，`ok` 才是"key 是否真的存在"的可靠信号。

### 2.3 map 遍历顺序不固定

对 map 做 `for range` 的顺序是不保证的。你如果需要稳定顺序，就要把 key 拿出来排序。

[MapKeysSorted](./collections.go#L32-L40) 做的就是这个事：先收集 key，再 `sort.Strings(keys)`。

> ⚠️ **注意**：**Go 的 map 遍历顺序是随机的**（每次运行可能不同，这是语言刻意设计的，防止你依赖顺序）。任何"需要稳定顺序"的场景，都必须把 key 取出排序（见 `MapKeysSorted`）。从 Java 的 `HashMap` 过来要注意：Java 的 `HashMap` 同样不保证顺序（Java 用 `LinkedHashMap` 才保插入顺序），但 Go 连"插入顺序"都不保留。

---

## 3. 去重与计数：两个最常见题型

### 3.1 去重（Unique）

去重典型写法：
- 用 `map[T]struct{}` 记录是否出现过
- `struct{}` 是零大小类型，不占额外存储（语义也更清晰：只关心存在与否）

对应代码：[UniqueInts](./collections.go#L9-L21)

> ⚠️ **注意**：**用 map 当 set 时，value 用 `struct{}{}` 最省内存**——空结构体占 0 字节，比 `map[T]bool` 省空间，语义上也更清楚（"只关心存在与否"）。判断存在用 `if _, ok := seen[v]; ok`，下划线丢弃 value 只看 `ok`。

### 3.2 计数（Frequency）

计数典型写法：
- value 是 `int`
- `out[s]++` 利用了"未存在时读零值"的特性

对应代码：[Frequency](./collections.go#L23-L30)

> ⚠️ **注意**：**map 读不存在的 key 得到零值**，所以 `out[s]++` 对没出现过的 key 会从 0 自增到 1，**不需要先判断 key 是否存在再初始化**。这是 Go map 零值特性带来的便利，也是写计数/去重时最省事的惯用法。

---

## 4. 泛型（初学者视角）

[MapKeysSorted](./collections.go#L32-L40) 用了泛型：

```go
func MapKeysSorted[V any](m map[string]V) []string
```

你只要理解一件事：
- `V any` 表示 value 的类型不重要，任何类型都可以
- 我们只关心 key 是 `string`，所以函数可以通用

> ⚠️ **注意**：**`[V any]` 是类型参数**，调用时由编译器根据实参推断，调用方可以传 `map[string]int`、`map[string]bool` 等任意 value 类型。这里只用到 key（固定 `string`），所以 value 被约束成 `any` 就能通用。对应 Java 的 `<V> List<String> mapKeysSorted(Map<String, V> m)`。

---

## 5. 练习清单（建议你按顺序做）

1. 写 `UniqueStrings([]string) []string`（保持第一次出现的顺序）
2. 写 `FrequencyInts([]int) map[int]int`
3. 写 `TopKWords(text string, k int) []string`（先分词，再计数，再排序取前 k）
4. 修改 `MapKeysSorted`：支持 `map[int]V`（提示：需要把 key 转成 slice 再排序）

---

## 6. 自测命令

```bash
go test ./internal/03-collections -shuffle on
```
