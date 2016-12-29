package main

import "fmt"

func main() {
	slice := []int {1,2,3,4}
	//getting the length of the slice and the capacity of it times 2 as a good practice
	//so that if more data is appended the array doesn't have to be re-built
	slice_copy := make([]int, len(slice), cap(slice)*2)
	//function copies the min number of elements on both arrays
	//the values of slice onto slice_copy
	copy(slice_copy, slice)

	fmt.Println(slice)
	fmt.Println(slice_copy)

}