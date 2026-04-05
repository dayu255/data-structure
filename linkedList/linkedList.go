package linkedlist

type LinkedList[T any] interface {
	Add(v T)
	Delete(i int)
	At(i int) T
	Empty() bool
	Len() int
	Insert(i int, v T)
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type HeadLinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() LinkedList[T] {
	return &HeadLinkedList[T]{
		head: nil,
	}
}

func (h *HeadLinkedList[T]) Add(v T) {
	newNode := &Node[T]{
		value: v,
		next:  nil,
	}

	if h.head == nil {
		h.head = newNode
		return
	}

	pos := h.head
	for pos.next != nil {
		pos = pos.next
	}

	pos.next = newNode
}

func (h *HeadLinkedList[T]) Delete(n int) {
	if n < 0 || h.head == nil {
		panic("Index out of range.")
	}

	if n == 0 {
		h.head = h.head.next
		return
	}

	pos := h.head
	var former *Node[T]

	for i := 0; i < n; i++ {
		former = pos
		pos = pos.next
		if pos == nil {
			panic("Index out of range.")
		}
	}

	former.next = pos.next
}

func (h *HeadLinkedList[T]) At(n int) T {
	if n < 0 || h.head == nil {
		panic("Index out of range.")
	}

	pos := h.head
	for i := 0; i < n; i++ {
		pos = pos.next
		if pos == nil {
			panic("Index out of range.")
		}
	}

	return pos.value
}

func (h *HeadLinkedList[T]) Empty() bool {
	return h.head == nil
}

func (h *HeadLinkedList[T]) Len() int {
	var n int
	pos := h.head
	for pos != nil {
		pos = pos.next
		n++
	}
	return n
}

func (h *HeadLinkedList[T]) Insert(n int, v T) {
	if n < 0 || (h.head == nil && n != 0) {
		panic("Index out of range.")
	}

	newNode := &Node[T]{
		value: v,
		next:  nil,
	}

	if n == 0 {
		newNode.next = h.head
		h.head = newNode
		return
	}

	pos := h.head
	for i := 0; i < n-1; i++ {
		pos = pos.next
		if pos.next == nil {
			panic("Index out of range.")
		}
	}

	newNode.next = pos.next
	pos.next = newNode
}
