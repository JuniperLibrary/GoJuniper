package main

import (
	"fmt"
	"sync"
)

/*
	核心使用要点（快速上手）
		定义业务结构体：比如 Order/User/Product，用唯一字段（如 OrderID/ID）作为集合的 uniqueKey；
		创建集合实例：set := NewGenericSet()；
		添加元素：set.Add(唯一Key, 结构体实例)；
		获取 / 操作元素：通过 Get 获取空接口值，再用 value, ok := elem.(Order) 安全断言还原为具体结构体；
		并发安全：代码中加了 sync.RWMutex，如果你的业务没有并发场景（比如单协程处理），可以直接移除锁，提升性能。
	总结
		空接口 + 结构体是 Go 中替代 “内置集合” 的最优方案，兼顾灵活性和自定义能力；
		核心是用唯一标识保证集合元素唯一性，用类型断言安全还原具体结构体类型；
		工具类已封装常用操作（Add/Remove/Get/ 遍历），可直接复用，也可扩展筛选、排序等功能。
*/

// GenericSet ------------------------------
// 1. 通用 Set 集合（支持任意类型元素）
// ------------------------------
type GenericSet struct {
	mu       sync.Mutex             // 并发安全
	elements map[string]interface{} // key:唯一标识。value 任意类型的元素
}

// NewGenericSet 创建 Set 实例
func NewGenericSet() *GenericSet {
	return &GenericSet{
		elements: make(map[string]interface{}),
	}
}

// Add 添加元素（uniqueKey 必须唯一，如 ID 转字符串）
func (s *GenericSet) Add(uniqueKey string, elem interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements[uniqueKey] = elem
}

// Remove 删除元素
func (s *GenericSet) Remove(uniqueKey string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.elements, uniqueKey)
}

// Get 获取元素（返回空接口 + 是否存在）
func (s *GenericSet) Get(uniqueKey string) (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	elem, ok := s.elements[uniqueKey]
	return elem, ok
}

// Exists 判断元素是否存在
func (s *GenericSet) Exists(uniqueKey string) bool {
	_, ok := s.Get(uniqueKey)
	return ok
}

// Len 获取集合长度
func (s *GenericSet) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.elements)
}

// ForEach 遍历集合（回调函数处理元素）
func (s *GenericSet) ForEach(fn func(key string, elem interface{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.elements {
		fn(k, v)
	}
}

// Order ------------------------------
// 2. 自定义业务结构体示例（比如订单）
// ------------------------------
type Order struct {
	OrderID string  // 唯一标识
	UserID  int     // 用户ID
	Amount  float64 // 订单金额
	Status  string  // 订单状态
}

func main() {
	// 1. 创建订单集合
	orderSet := NewGenericSet()

	// 2. 构造订单数据
	order1 := Order{
		OrderID: "ORD20260317001",
		UserID:  1001,
		Amount:  99.9,
		Status:  "PAID",
	}
	order2 := Order{
		OrderID: "ORD20260317002",
		UserID:  1002,
		Amount:  199.9,
		Status:  "UNPAID",
	}

	// 3. 添加订单到集合（用 OrderID 作为唯一 Key）
	orderSet.Add(order1.OrderID, order1)
	orderSet.Add(order2.OrderID, order2)
	// 重复添加同一订单（自动覆盖，保证唯一性）
	orderSet.Add(order1.OrderID, order1)

	// 4. 查看集合长度
	fmt.Println("订单集合长度：", orderSet.Len()) // 输出：2

	// 5. 获取单个订单（类型断言还原为 Order 结构体）
	elem, ok := orderSet.Get("ORD20260317001")
	if ok {
		// 安全类型断言：避免 panic
		if order, ok := elem.(Order); ok {
			fmt.Printf("获取到订单：%+v\n", order)
		}
	}

	// 6. 遍历所有订单，筛选出已支付的订单
	fmt.Println("\n已支付的订单：")
	orderSet.ForEach(func(key string, elem interface{}) {
		if order, ok := elem.(Order); ok {
			if order.Status == "PAID" {
				fmt.Printf("订单号：%s，金额：%.2f\n", order.OrderID, order.Amount)
			}
		}
	})

	// 7. 删除订单
	orderSet.Remove("ORD20260317002")
	fmt.Println("\n删除后集合长度：", orderSet.Len()) // 输出：1
}
