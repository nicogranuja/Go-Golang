package main

import "fmt"

func main() {
	//the value is fixed  3 in this case cant change it
	//only one type
	//Zero default value for elements
	var arr [3] int
	//declaring array with initial values
	arr2 := [3] int {3,33,66}

	//printing arrays
	fmt.Println(arr2, arr,)

	//print length
	fmt.Println(len(arr2))

	//accessing the array elements with index
	for i:=0;i<len(arr2);i++{
		fmt.Println(arr2[i])
	}

	//multidimensional array
	var matrix [2][2] int
}