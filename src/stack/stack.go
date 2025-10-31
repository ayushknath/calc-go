package stack

import (
	"fmt"
	"os"
)

// 'values' is the actual slice
// 'top' stores the index of the topmost element
type Stack[T any] struct {
	values []T
	top    int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{values: make([]T, 0), top: -1}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) PrintStack() {
	fmt.Printf("[ ")
	for _, v := range s.values {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("]\n")
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		fmt.Println("top: stack is already empty")
		os.Exit(0)
	}
	return s.values[s.top]
}

func (s *Stack[T]) Push(x T) {
	s.top++
	s.values = append(s.values, x)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		fmt.Println("pop: stack is already empty")
		os.Exit(0)
	}
	pop := s.values[s.top]
	s.values = s.values[:len(s.values)-1]
	s.top--
	return pop
}
