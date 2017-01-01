package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	//this way loads the whole fail into memory
	//not so efficient 
	//look in the same folder
	//save contents into the variable
	file_data,err := ioutil.ReadFile("./Hello.txt")

	if err != nil{
		fmt.Println("Error!")
	}
	//print the variable
	fmt.Println(string (file_data))
}