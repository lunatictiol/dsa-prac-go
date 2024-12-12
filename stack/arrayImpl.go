package stack

import (
	"fmt"
)

type StackArray struct {
	data []int
}

func (s *StackArray) Push(element int) {
	s.data = append(s.data, element)

}
func (s *StackArray) Pop() {
	if s.IsEmpty() {
		fmt.Println("nothing to pop stack is empty")
		return
	}
	s.data = append(s.data[:len(s.data)-1])
}

func (s *StackArray) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *StackArray) Peek() {
	fmt.Printf("Top is %v\n", s.data[len(s.data)-1])
}

func (s *StackArray) Print() {
	fmt.Print("Bottom : ")
	fmt.Print(s.data)
	fmt.Println(" : Top")
}
