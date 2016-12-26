package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	age := 22
	//%v anything?, %d for digits, %t for booleans, %f for floats
	fmt.Printf("I am %d\n", age)

	//input output
	var age2 int
	fmt.Println("Insert your age")
	//the %d specifies the type of the variable scanned
	//the /n  makes sure it reads on the new line
	//&age2 references the memory and changes the variable
	fmt.Scanf("%d\n", &age2)
	//printing the age
	fmt.Printf("Your age is %d\n", age2)

	var name string
	fmt.Println("What is your name")
	//we have to access the variable by reference
	fmt.Scanf("%v\n", &name)
	//we can use %v or %s
	fmt.Printf("Your name is %v\n", name)


	//USING Bufio with os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert your name: ")
	//read until you find a \n
	//function ReadString returns two values
	text,err := reader.ReadString('\n')
	//if error is not equal to null
	//print error
	if(err != nil){
		fmt.Println(err)
	//we can concatenate strings with + or ,
	}else{
		fmt.Println("Hello "+ text)
	}


}