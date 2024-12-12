package linkedlist

import (
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}
type LinkedList struct {
	head   *node
	Length int
	tail   *node
}

func (list *LinkedList) Append(value any) {
	newNode := &node{value: value, next: nil}
	if list.head == nil {
		list.head = newNode
		list.tail = list.head
		return
	}
	list.tail.next = newNode
	list.tail = newNode
	list.Length++
}

func (list *LinkedList) Prepend(value any) {

	newNode := &node{value: value, next: list.head}
	list.head = newNode
	list.Length++

}

func (list *LinkedList) Insert(index int, value any) {
	if index <= 0 {
		list.Prepend(value)
		return
	}
	if index >= list.Length {
		list.Append(value)
		return
	}
	newNode := &node{
		value: value,
		next:  nil,
	}
	currentNode := list.head
	for current := 0; current < index-1; current++ {
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next = newNode
	list.Length++

}

func (list *LinkedList) Delete(index int) {
	if index == 0 {
		list.head = list.head.next
		list.Length--
		return
	}
	currentNode := list.head
	for current := 0; current < index-1; current++ {
		currentNode = currentNode.next
	}

	currentNode.next = currentNode.next.next
	list.Length--

}

func (list *LinkedList) Print() {
	var current *node = list.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println()
}

func (list *LinkedList) Reverse() {
	current := list.head
	var prevNode *node
	var nodeToBeVisited *node
	for current != nil {
		nodeToBeVisited = current.next
		current.next = prevNode
		prevNode = current
		current = nodeToBeVisited
	}
	list.head = prevNode

}

func (list *LinkedList) Get(index int) any {
	currentNode := list.head
	for current := 0; current < index-1; current++ {
		currentNode = currentNode.next
	}
	return currentNode.next.value

}
