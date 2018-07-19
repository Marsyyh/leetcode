package main

import (
	"fmt"

	"github.com/Marsyyh/go-datastructure/container/stack"
)

type MyQueue struct {
	value int
	test  []int
}

func MyQueueBuilder(value int, test []int) MyQueue {
	return MyQueue{value, test}
}

func main() {
	s := stack.New()
	s.Push(20)
	s.Push(50)
	fmt.Println(s)
	fmt.Println(s.Len())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s.Len())
	fmt.Println(s.IsEmpty())
	s.Pop()
	fmt.Println(s.IsEmpty())
	fmt.Println(s.IsEmpty())
}
