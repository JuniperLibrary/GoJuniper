package main

import (
	"fmt"
	"sync"
)

// SafeMap 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// NewSafeMap 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set 设置键值对
func (safeMap *SafeMap[K, V]) Set(k K, v V) {
	safeMap.mutex.Lock()
	defer safeMap.mutex.Unlock()
	safeMap.data[k] = v
}

// Get 获取
func (safeMap *SafeMap[K, V]) Get(k K) (V, bool) {
	safeMap.mutex.RLock()
	defer safeMap.mutex.RUnlock()
	v, ok := safeMap.data[k]
	return v, ok
}

// Del 删除键
func (safeMap *SafeMap[K, V]) Del(k K) {
	safeMap.mutex.Lock()
	defer safeMap.mutex.Unlock()
	delete(safeMap.data, k)
}

func (m *SafeMap[K, V]) Keys() []K {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	// 创建字符串到整数的映射
	scores := NewSafeMap[string, int]()
	scores.Set("Alice", 95)
	scores.Set("Bob", 87)

	if score, exists := scores.Get("Alice"); exists {
		fmt.Printf("Alice's score: %d\n", score) // 输出: Alice's score: 95
	}

	fmt.Println("Keys:", scores.Keys()) // 输出: Keys: [Alice Bob]
}
