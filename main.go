package main

import (
	"github.com/lunatictiol/dsa-prac-go/linkedlist"
)

func main() {
	//arrays.BasicsOfArray()

	// var arr arrays.MyIntArray

	// arr.New(5)

	// arr.Push(1)
	// arr.Push(2)
	// arr.Push(3)
	// arr.Push(4)

	// arr.PrintArray()
	// // arr.Pop()
	// // arr.Pop()
	// // arr.Pop()
	// // arr.Pop()

	// arr.Delete(2)

	// arr.PrintArray()

	//questions.ReverseStr("string to reverse")
	// sortedArray := questions.SolOfMergeSortedArrays([]int{1, 12, 23, 34}, []int{2, 16, 26, 88})
	// fmt.Println(sortedArray)
	//a := questions.TwoSum([]int{3, 2, 4}, 6)
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// a := questions.RemoveDuplicates2(nums)
	// // fmt.Println(nums)
	// fmt.Println(a)

	// a := questions.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	// fmt.Println(a)

	//	questions.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// questions.Rotate(nums, 7)
	// fmt.Println(nums)
	//questions.PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
	//questions.SingleNumberWithXor([]int{4, 1, 2, 1, 2})
	list := &linkedlist.LinkedList{}

	list.Append(2)
	list.Append(3)
	list.Prepend(1)
	list.Append(4)
	list.Prepend(0)
	list.Insert(1, 7)
	list.Insert(3, 8)
	list.Insert(7, 10)
	//	n := list.Get(3)
	//fmt.Println(n)
	list.Print()
	list.Delete(0)
	list.Print()

}
