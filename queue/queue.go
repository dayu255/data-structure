package queue

import "errors"

type Queue[T any] interface {
	Push(v T)
	Pop()
	Empty() bool
	Len() int
	Front() (T, error)
}

type sliceQueue[T any] struct {
	items []T
}

func NewQueue[T any]() Queue[T] {
	return &sliceQueue[T]{
		items: make([]T, 0),
	}
}

func (q *sliceQueue[T]) Push(v T) {
	q.items = append(q.items, v)
}

func (q *sliceQueue[T]) Pop() {
	if len(q.items) == 0 {
		return
	}
	q.items = q.items[1:]
}

func (q *sliceQueue[T]) Empty() bool {
	return len(q.items) == 0
}

func (q *sliceQueue[T]) Len() int {
	return len(q.items)
}

func (q *sliceQueue[T]) Front() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("Queue is empty.")
	}
	return q.items[0], nil
}
