package stack

import (
	"errors"
	"github.com/dayu255/data-structure/linkedList"
)

type Stack[T any] interface {
	Push(v T)
	Pop()
	Empty() bool
	Len() int
	Top() (T, error)
}

type linkedStack[T any] struct {
	l linkedlist.LinkedList[T]
}

func NewStack[T any]() Stack[T] {
	return &linkedStack[T]{
		l: linkedlist.NewLinkedList[T](),
	}
}

func (s *linkedStack[T]) Push(v T) {
	s.l.Insert(0, v)
}

func (s *linkedStack[T]) Pop() {
	if s.Empty() {
		return
	}
	s.l.Delete(0)
}

func (s *linkedStack[T]) Empty() bool {
	return s.l.Empty()
}

func (s *linkedStack[T]) Len() int {
	return s.l.Len()
}

func (s *linkedStack[T]) Top() (T, error) {
	if s.Empty() {
		var zero T
		return zero, errors.New("Stack is empty.")
	}
	return s.l.At(0), nil
}
