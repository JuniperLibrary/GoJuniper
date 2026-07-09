package main

import "fmt"

// Stack 泛型 Stack 实现
type Stack[T any] struct {
	elements []T
}

// Push 入栈
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop 出栈
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, true
}

// Peek 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// isEmpty 判断栈是否为空
func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}
func main() {
	// 整数栈
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())

	// 字符串 栈
	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop())
}
