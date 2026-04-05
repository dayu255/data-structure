package main

import (
	"github.com/dayu255/data-structure/list"
	"fmt"
)

func main() {
	s := stack.NewStack[int]()

	s.Push(10)
	s.Push(20)

	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
}
