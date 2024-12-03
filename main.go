package main

import (
	"fmt"

	"github.com/lunatictiol/dsa-prac-go/questions"
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
	// nums := []int{1, 1, 2}
	// a := questions.RemoveDuplicates2(nums)
	// fmt.Println(nums)
	// fmt.Println(a)

	// a := questions.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	// fmt.Println(a)

	//	questions.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	questions.Rotate(nums, 7)
	fmt.Println(nums)

}
