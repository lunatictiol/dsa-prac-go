package questions

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
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

// remove dublicates
func RemoveDuplicates(nums []int) int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if !ok {
			m[nums[i]] = nums[i]
		}
	}

	v := make([]int, 0, len(m))
	for i, j := 0, nums[0]; i < len(m); i, j = i+1, j+1 {
		v = append(v, m[j])

	}
	for i := 0; i < len(v); i++ {
		nums[i] = v[i]
	}

	return len(nums)

}

func RemoveDuplicates2(nums []int) int {
	left := 0

	for right := 0; right < len(nums); right++ {
		fmt.Printf("left value: %v\n  ", nums[left])
		fmt.Printf("Right value: %v\n  ", nums[right])
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	return len(nums[:left+1])
}

func MovesZeors(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {

			nums[i], nums[j] = nums[j], nums[i]
			j++

		}

	}
	fmt.Println(nums)

}

func MaxSubArray(nums []int) int {
	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		currentSum := 0
		for j := i; j < len(nums); j++ {
			currentSum = currentSum + nums[j]
			maxSum = max(currentSum, maxSum)
		}

	}

	return maxSum

}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 1 && n == 0 {

		return
	}
	if m == 0 && n == 1 {
		nums1[0] = nums2[0]
		return
	}

	i, j := 0, 0

	// Merge elements until one array is exhausted
	for i < m && j < n {
		fmt.Printf(" i %v, nums1 :%v,j %v, nums2 : %v\n", i, nums1[i], j, nums2[j])
		if nums1[i] < nums2[j] && nums1[i] != nums2[j] {
			i++
		} else {
			nums1[i] = nums2[j]
			j++
		}

	}

	fmt.Println(nums1)
}

func Rotate(nums []int, k int) {

	// a := 0
	// index := len(nums) - 1
	// temp := make([]int, 0)
	// for a < k {
	// 	temp = append(temp, nums[index])

	// 	a++
	// 	index--
	// }
	// temp = append(temp, nums[:index]...)
	// copy(nums, temp)

	n := len(nums)
	k %= n // Handle cases where k > n
	temp := append(nums[n-k:], nums[:n-k]...)
	copy(nums, temp)

	fmt.Print(1000 % 100000)

}

func PlusOne(digits []int) {

	b := ""
	for _, v := range digits {
		if len(b) > 0 {
			b += ""
		}
		b += strconv.Itoa(v)
	}
	largeNumber := new(big.Int)
	oneBigInt := new(big.Int)
	oneBigInt.SetString("1", 10)
	largeNumber.SetString(b, 10)
	fmt.Println(largeNumber)
	largeNumber.Add(largeNumber, oneBigInt)
	value := largeNumber.String()
	arr := strings.Split(value, "")
	result := make([]int, 0)
	for _, v := range arr {
		s, _ := strconv.Atoi(v)
		result = append(result, s)
	}

	fmt.Println(result)

}

/*
k=  0   1     2 <len(arr)
 [ 1   2    3 ]
j=2 1 0
10^j * arr[k]
100 * 1 = 100
10 * 2 = 20
1 * 3 = 3
123 + 1
124


 100    10     1
100
20
3
123+1 124

k 0  1
  2  2  2   3
i 0  1  2  3


3


    j
0 , 0, 1, 1, 1, 2 ,2 ,3 ,3, 4
   i
i 1
j 1

*/

func SingleNumber(nums []int) {
	m := make(map[int]int)
	var k int
	//4 1 2 1 2
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)

		} else {
			m[v] = v
			k = v
		}
		fmt.Println(k)
	}

	//fmt.Println(m[k])

}

func SingleNumberWithXor(nums []int) {
	m := 2
	fmt.Println(m ^ 2)
	// for _, v := range nums {
	// 	fmt.Println(v)
	// 	m ^= v
	// 	fmt.Println(m)

	// }

}
