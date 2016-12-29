package main

import "fmt"

func main() {
	//creating a slice with the make method
	//make (type, length, capacity)
	slice := make([]int, 3, 5)

	//cap function returns the capacity
	fmt.Println(cap(slice))

	slice = append(slice, 2)
	
	//capacity is 5
	fmt.Println(cap(slice))
	//length is 4
	fmt.Println(len(slice))

	fmt.Println(slice)


}