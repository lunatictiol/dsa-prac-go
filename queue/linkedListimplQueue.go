package queue

import "fmt"

type element struct {
	value int
	next  *element
}
type Queue struct {
	top    *element
	bottom *element
}

func (q *Queue) EnQueueu(value int) {
	newNode := &element{value: value, next: nil}
	if q.top == nil {
		q.top = newNode
		q.bottom = q.top
		return
	}
	q.bottom.next = newNode
	q.bottom = newNode
}
func (q *Queue) DeQueueu() {
	if q.IsEmpty() {
		fmt.Println("No one in the queue")
		return
	}
	q.top = q.top.next
}

func (q *Queue) Print() {
	if q.IsEmpty() {
		fmt.Println("nothing to print queue is empty")
		return
	}
	var current *element = q.top
	fmt.Print("Top")
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println("End")
}

func (q *Queue) Peek() {
	if q.IsEmpty() {
		fmt.Println("no one in the queue")
		return
	}
	fmt.Printf("%v is in the front of the queue\n", q.top.value)
}

func (q *Queue) IsEmpty() bool {
	return q.top == nil
}
