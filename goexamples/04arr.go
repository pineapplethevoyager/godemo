package main

import "fmt"

func main() {
	//array
	var arr [3]int
	arr[0] = 1
	fmt.Println(arr)

	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)

	//slice
	slice0 := arr2[:]
	slice0 = append(slice0, 4,5,6)//changes array and slice

	slice := []int{1,2,3}
	slice = append(slice, 4,5,6)
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2,s3,s4)
}
