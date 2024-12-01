package questions

//consider two arrays
// [a,b,c,d]
//[x,y,z,a]
//write a function to which return true if there is a common element

//brute force approch
// for each element in array loop through second array to se if the similar element exists
//time complexity O(a*b) a-> size of first array, b-> size of second array.

// using maps
// complexity O(a+b)
func checkIfSimilarElementExists(arr []string, arr2 []string) bool {

	//create a string bool map of first array
	var arrMap = make(map[string]bool)
	for _, a := range arr {
		//check if the key already exists in the map
		_, exists := arrMap[a]
		//if it doesnt create a key
		if !exists {
			arrMap[a] = true
		}

	}
	// for each walue of arr2 check if a key exists in the map
	for _, b := range arr2 {
		_, exists := arrMap[b]
		if exists {
			//if it does return true
			return true
		}
	}
	//no key is found
	return false

}

// func main() {
// 	arr1 := []string{"a", "b", "c", "d", "x"}
// 	arr2 := []string{"v", "n", "m", "t", "o"}
// 	exists := checkIfSimilarElementExists(arr1, arr2)
// 	if exists {
// 		fmt.Println("exists")
// 	} else {
// 		fmt.Println("doesnt exists")
// 	}
// }
