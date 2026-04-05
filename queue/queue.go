package queue

import (
	"github.com/dayu255/data-structure/linkedList"
)

type Queue[T any] interface {
	Push(v T)
	Pop()
	Empty() bool
	Len() int
	Front() (T, error)
}

type sliceQueue[T any] struct {
	l linkedlist.LinkedList[T]
}

func NewQueue[T any]() Queue[T] {
	return &sliceQueue[T]{
		linkedlist.NewLinkedList[T](),
	}
}

func (q sliceQueue[T]) Push(v T) {
	q.l.Add(v)
}

func (q *sliceQueue[T]) Pop() {
	q.l.Delete(0)
}

func (q *sliceQueue[T]) Empty() bool {
	return q.l.Empty()
}

func (q *sliceQueue[T]) Len() int {
	return q.l.Len()
}

func (q *sliceQueue[T]) Front() (T, error) {
	return q.l.At(0), nil
}
