package main

import(
	"fmt"
)

func main() {
	//creating slice leaving brackets empty
	//slice created with null value "nil"
	matrix := [] int{1,2,3,4}

	fmt.Println(matrix)

	//creating array
	arr := [3] int {1,2,3}

	//converting to slice
	//slicing the array from index
	//prints [1 2] from "0 to 1" index
	slice := arr[:2]

	fmt.Println(slice)
}