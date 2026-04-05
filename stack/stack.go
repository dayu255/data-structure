package stack

import "errors"

type Stack[T any] interface {
	Push(v T)
	Pop()
	Empty() bool
	Len() int
	Top() (T, error)
}

type sliceStack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return &sliceStack[T]{
		items: make([]T, 0),
	}
}

func (s *sliceStack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *sliceStack[T]) Pop() {
	if s.Empty() {
		return
	}
	s.items = s.items[:len(s.items) - 1]
}

func (s *sliceStack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *sliceStack[T]) Len() int {
	return len(s.items)
}

func (s *sliceStack[T]) Top() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}
