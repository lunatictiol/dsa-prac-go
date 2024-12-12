package stack

import (
	"fmt"
	"slices"
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
	s.top.next = newNode
	s.top = newNode
}

func (s *Stack) Pop() {
	if s.bottom == nil {
		fmt.Println("stack is empty")
		return
	}
	if s.bottom == s.top {
		s.bottom = nil
		return
	}
	var current *element = s.bottom
	for current.next != s.top {
		current = current.next
	}
	current.next = nil
	s.top = current
}

func (s *Stack) Peek() {
	fmt.Printf("Top element is %v\n", s.top.value)
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("Nothing to print. Stack is empty")
		return
	}
	var current *element = s.bottom
	temp := make([]int, 0)
	for current != nil {
		temp = append(temp, current.value)
		current = current.next
	}
	slices.Reverse(temp)
	fmt.Print("Top : ")
	fmt.Print(temp)
	fmt.Println(" : bottom")
}

func (s *Stack) IsEmpty() bool {
	return s.bottom == nil
}
