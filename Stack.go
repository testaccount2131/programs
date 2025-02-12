package main

import (
	"fmt"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	topIndex := len(s.items) - 1
	topItem := s.items[topIndex]
	s.items = s.items[:topIndex]
	return topItem, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Peek at the top element
	if top, ok := stack.Peek(); ok {
		fmt.Println("Top element:", top)
	}

	// Pop elements from the stack
	for !stack.IsEmpty() {
		if item, ok := stack.Pop(); ok {
			fmt.Println("Popped:", item)
		}
	}

	// Check if the stack is empty
	fmt.Println("Is stack empty?", stack.IsEmpty())
}
