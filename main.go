package main

import (
	"github.com/lunatictiol/dsa-prac-go/arrays"
)

func main() {
	//arrays.BasicsOfArray()

	var arr arrays.MyIntArray

	arr.New(5)

	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arr.Push(4)

	arr.PrintArray()
	// arr.Pop()
	// arr.Pop()
	// arr.Pop()
	// arr.Pop()

	arr.Delete(2)

	arr.PrintArray()

	//fmt.Println(arr.Get(0))
}
