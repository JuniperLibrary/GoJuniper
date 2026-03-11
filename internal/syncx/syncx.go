// Package syncx 提供 sync 主题的基础练习：
// - sync.Mutex 保护共享变量
// - sync.Once 确保初始化只执行一次
package syncx

import "sync"

// Counter 是一个线程安全的计数器。
type Counter struct {
	mu sync.Mutex
	n  int
}

// Inc 将计数加一。
func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

// Add 将计数加上 delta。
func (c *Counter) Add(delta int) {
	c.mu.Lock()
	c.n += delta
	c.mu.Unlock()
}

// Value 返回当前计数。
func (c *Counter) Value() int {
	c.mu.Lock()
	v := c.n
	c.mu.Unlock()
	return v
}

// OnceValue 演示 lazy init：第一次调用 f 计算值并缓存，后续直接返回缓存。
type OnceValue[T any] struct {
	once sync.Once
	v    T
}

// Get 返回缓存值；如果是第一次调用，会执行 f 计算。
func (o *OnceValue[T]) Get(f func() T) T {
	o.once.Do(func() {
		o.v = f()
	})
	return o.v
}
