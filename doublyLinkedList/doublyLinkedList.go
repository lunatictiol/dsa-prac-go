package doublylinkedlist

import (
	"fmt"
)

type node struct {
	value    int
	next     *node
	previous *node
}
type DoubleLinkedList struct {
	head   *node
	Length int
	tail   *node
}

func (list *DoubleLinkedList) Append(value int) {
	newNode := &node{value: value, next: nil, previous: nil}
	if list.head == nil {
		list.head = newNode
		list.tail = list.head

		return
	}
	newNode.previous = list.tail
	list.tail.next = newNode
	list.tail = newNode
	list.Length++
}

func (list *DoubleLinkedList) Prepend(value int) {

	newNode := &node{value: value, previous: nil, next: list.head}
	list.head.previous = newNode
	list.head = newNode

	list.Length++

}

func (list *DoubleLinkedList) Insert(index int, value int) {
	fmt.Printf("size:%v\n", list.Length)
	if index <= 0 {
		list.Prepend(value)
		return
	}
	if index >= list.Length {
		list.Append(value)
		return
	}
	newNode := &node{
		value:    value,
		next:     nil,
		previous: nil,
	}
	currentNode := list.head
	for current := 0; current < index-1; current++ {

		currentNode = currentNode.next
	}
	//fmt.Printf("current Node: %v,  previous:null, next:%v\n ", currentNode.value, currentNode.next.value)

	newNode.next = currentNode.next
	newNode.previous = currentNode
	currentNode.next.previous = newNode
	currentNode.next = newNode

}

func (list *DoubleLinkedList) Delete(index int) {
	if index == 0 {
		list.head = list.head.next
		list.Length--
		return
	}
	currentNode := list.head
	for current := 0; current < index; current++ {
		currentNode = currentNode.next
	}
	nodeTobeShiftedBack := currentNode.next
	nodeTobeShiftedForward := currentNode.previous
	nodeTobeShiftedBack.previous = nodeTobeShiftedForward

	nodeTobeShiftedForward.next = nodeTobeShiftedBack
	list.Length--

}

func (list *DoubleLinkedList) Print() {
	var current *node = list.head
	for current != nil {
		if current == list.head && list.head.next == nil {
			fmt.Printf("current Node: %v,  previous:null, next:null\n ", current.value)

		} else if current == list.head {
			fmt.Printf("current Node: %v,  previous:null, next:%v\n ", current.value, current.next.value)

		} else if current.next == nil {
			fmt.Printf("current Node: %v,  previous:%v, next: null\n", current.value, current.previous.value)

		} else {
			fmt.Printf("current Node: %v,  previous:%v, next: %v\n", current.value, current.previous.value, current.next.value)

		}
		current = current.next
	}
	fmt.Println()
}

func (list *DoubleLinkedList) Get(index int) any {
	currentNode := list.head
	for current := 0; current < index-1; current++ {
		currentNode = currentNode.next
	}
	return currentNode.next.value

}
