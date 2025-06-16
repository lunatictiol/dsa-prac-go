package arrays

import (
	"fmt"
//	"slices"
)

type MyIntArray struct {
	length int
	data   []int
}

func (a *MyIntArray) New(length int) {
	a.length = 0
	a.data = make([]int, length)

}

func (a *MyIntArray) Get(index int) int {
	return a.data[index]
}

func (a *MyIntArray) Push(item int) int {
	a.data[a.length] = item
	a.length++
	return a.length
}
func (a *MyIntArray) Pop() {

	if a.length == 0 {
		panic("no element left in the array")

	}
	a.data = a.data[:a.length-1]
	a.length--
}
func (a *MyIntArray) Delete(index int) {
	// temp := a.data[:index]

	// temp = append(temp, a.data[index+1:]...)
	// a.data = temp
	// a.length--
	if index < 0 || index >= a.length {
		panic("index is invalid") // Handle invalid index case
	}

	 a.data = append(a.data[:index], a.data[index+1:]...)
        
	// a.data = slices.Delete(a.data,index,index)
	a.length--

}

func (a *MyIntArray) PrintArray() {
	for _,x:= range a.data {
		fmt.Print(x, " ")
	}
	fmt.Println()
}
