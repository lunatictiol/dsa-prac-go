package main

import "fmt"


func main (){
  arr:= []int{1,2,3,4}

  fmt.Println(len(arr))
  for x := range arr {
	  fmt.Println(x)
  }
}
