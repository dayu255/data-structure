package main

import (
	"github.com/dayu255/data-structure/stack"
	"github.com/dayu255/data-structure/linkedList"
	"fmt"
)

func main() {
	s := stack.NewStack[int]()

	s.Push(10)
	s.Push(20)

	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())

	l := linkedlist.NewLinkedList[int]()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	fmt.Println(l.At(0))
	fmt.Println(l.At(1))
	fmt.Println(l.At(2))
	fmt.Println(l.At(5))
}
