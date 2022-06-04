package main

import "fmt"

func main() {

	arr1 := [3]int{1, 2, 3}
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	// fmt.Println(arr)
	fmt.Println(arr1)

	// slice build on the construct of the arr
	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	// slice is not a fixed entity as an array
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	slice1 = append(slice, 4)
	fmt.Println(slice1)
}
