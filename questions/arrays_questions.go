package questions

import (
	"fmt"
	"slices"
	"strings"
)

// O(n)
func ReverseStr(s string) {

	arr := strings.Split(s, "")
	println(len(arr))

	//O(n)
	slices.Reverse(arr)

	//0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	//s t r i n g   t o   r  e  v  e  r  s  e

	/*
	   	Reverse reverses the elements of the slice in place.
	   	built in function in go 1.2+
	        func Reverse[S ~[]E, E any](s S) {
	        	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	        		s[i], s[j] = s[j], s[i]
	   			here right hand side is evaluated first
	   			so the value of s[j] and s[i] are first fetched
	   			that will be e and s and are Simultaneously assigned to left hand side
	   			so there is no need to declare a third temp variable


	       	}
	       }
	*/
	fmt.Println(arr)

	//O(n)
	justString := strings.Join(arr, "")
	fmt.Println(justString)

}

//merger two sorted arrays
//my answer

func MergeTwoSortedArray(arr1, arr2 []int) []int {
	r := make([]int, len(arr1)+len(arr2))

	maxForLoop := min(len(arr1), len(arr2))
	i := 0
	j := 0
	println(maxForLoop)
	for k := 0; k < maxForLoop; k++ {

		if arr1[i] < arr2[j] {
			fmt.Printf("arr1 [%v] = %v\n", i, arr1[i])
			fmt.Printf("arr2 [%v] = %v\n", i, arr2[i])
			r[k] = arr1[i]
			fmt.Printf("r [%v] = %v\n", i, r[i])
			i++

		} else {
			r[k] = arr2[j]
			j++
		}
	}
	if maxForLoop == len(arr1) {
		r = append(r[:len(r)-1], arr2[maxForLoop:]...)
	} else if maxForLoop == len(arr2) {

		r = append(r[:len(r)-1], arr1[maxForLoop:]...)
	}

	return r

}

/*
Initialization:

Use two pointers (i for arr1 and j for arr2) to track the current index of both arrays.

Create a result slice to store the merged elements.

Comparison:

Compare the elements at the pointers and append the smaller one to the result.
Remaining Elements:
Once one array is exhausted, append the remaining elements of the other array to the result.

*/

func SolOfMergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(arr1)+len(arr2))

	// Merge elements until one array is exhausted
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	// Append the remaining elements from arr1
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	// Append the remaining elements from arr2
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

//two sum

func TwoSum(nums []int, target int) []int {
	r := make([]int, 2)
	mappedArray := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mappedArray[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]

		v, ok := mappedArray[temp]
		if ok && i != v {
			return []int{i, mappedArray[temp]}

		}
	}

	return r
}
