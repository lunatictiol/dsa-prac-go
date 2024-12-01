package questions

import (
	"fmt"
	"time"
)

func bigO() {
	start := time.Now()
	arr := []string{"nemo", "main", "train", "odsjn"}

	for i := 0; i < len(arr); i++ {
		if "nemo" == arr[i] {
			fmt.Println("found")
		}
	}
	fmt.Println("not found")
	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
}
