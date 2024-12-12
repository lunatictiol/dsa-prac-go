package stack

import (
	"fmt"
)

type element struct {
	value int
	next  *element
}
type Stack struct {
	top    *element
	bottom *element
}

func (s *Stack) Push(value int) {
	newNode := &element{value: value, next: nil}
	if s.top == nil {
		s.top = newNode
		s.bottom = s.top
		return
	}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Pop() {
	s.top = s.top.next

}

func (s *Stack) Peek() {
	fmt.Printf("Top element is %v\n", s.top.value)
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("Nothing to print. Stack is empty")
		return
	}
	var current *element = s.top
	fmt.Print("Top ")
	for current != nil {
		fmt.Print(current.value, " ")

		current = current.next
	}
	fmt.Println(" bottom")

}

func (s *Stack) IsEmpty() bool {
	return s.bottom == nil
}
