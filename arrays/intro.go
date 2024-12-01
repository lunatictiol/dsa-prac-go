package arrays

import "fmt"

func BasicsOfArray() {
	//slice
	arr := []string{"a", "b", "c"}
	//adds element at the end of the slice
	arr = append(arr, "d")

	fmt.Println(arr)
	// dynamically creating slice

	arr2 := make([]int, 5)
	arr2 = append(arr2, 2)
	arr2[0] = 1
	fmt.Println(arr2)
}
