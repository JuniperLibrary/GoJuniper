# Go 基础：切片与映射（slice 与 map）

这个文档对应 [internal/collections](file:///e:/dingchuan/GoJuniper/internal/collections)。

配合阅读：
- 实现代码：[collections.go](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go)
- 测试代码：[collections_test.go](file:///e:/dingchuan/GoJuniper/internal/collections/collections_test.go)

---

## 1. slice 是什么

`slice`（切片）是 Go 里最常用的“可变长度序列”。你可以把它理解成：
- 逻辑上：一段连续的元素序列
- 内部：指向底层数组的一段“视图”（包含指针、长度 len、容量 cap）

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

[UniqueInts](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go#L9-L21) 就用了这种思路：`out := make([]int, 0, len(xs))`。

---

## 2. map 是什么

`map[K]V` 是哈希表（键值对）。你最常用的是：

```go
m := map[string]int{}
m["a"]++
v := m["a"]
```

### 2.1 map 的“零值”是 nil

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

### 2.2 读 map 的 “comma ok” 写法

```go
v, ok := m["key"]
```

在 [UniqueInts](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go#L10-L19) 里就用 `ok` 判断是否出现过。

### 2.3 map 遍历顺序不固定

对 map 做 `for range` 的顺序是不保证的。你如果需要稳定顺序，就要把 key 拿出来排序。

[MapKeysSorted](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go#L32-L40) 做的就是这个事：先收集 key，再 `sort.Strings(keys)`。

---

## 3. 去重与计数：两个最常见题型

### 3.1 去重（Unique）

去重典型写法：
- 用 `map[T]struct{}` 记录是否出现过
- `struct{}` 是零大小类型，不占额外存储（语义也更清晰：只关心存在与否）

对应代码：[UniqueInts](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go#L9-L21)

### 3.2 计数（Frequency）

计数典型写法：
- value 是 `int`
- `out[s]++` 利用了“未存在时读零值”的特性

对应代码：[Frequency](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go#L23-L30)

---

## 4. 泛型（初学者视角）

[MapKeysSorted](file:///e:/dingchuan/GoJuniper/internal/collections/collections.go#L32-L40) 用了泛型：

```go
func MapKeysSorted[V any](m map[string]V) []string
```

你只要理解一件事：
- `V any` 表示 value 的类型不重要，任何类型都可以
- 我们只关心 key 是 `string`，所以函数可以通用

---

## 5. 练习清单（建议你按顺序做）

1. 写 `UniqueStrings([]string) []string`（保持第一次出现的顺序）
2. 写 `FrequencyInts([]int) map[int]int`
3. 写 `TopKWords(text string, k int) []string`（先分词，再计数，再排序取前 k）
4. 修改 `MapKeysSorted`：支持 `map[int]V`（提示：需要把 key 转成 slice 再排序）

---

## 6. 自测命令

```bash
go test ./internal/collections -shuffle on
```

