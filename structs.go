package main

import "fmt"


type User struct{
	age int
	name, lastName string

}

func main() {

	//1st way of creating a var of type User
	var nico User
	fmt.Println(nico)

	//2nd way defining the properties in disorder
	nico2 := User{age: 22, lastName:"Nico" }

	fmt.Println(nico2)

	//3rd way in order
	nico3 := User {22, "Nico", "Mondragon"}
	fmt.Println(nico3)	

	//4th way creates a pointer
	nico4 := new (User)
	fmt.Println(nico4)

	nico4.name = "Nicolas4"
	//access the data of the pointer
	//also less readable way: (*nico4).name
	fmt.Println(nico4.name)		 
}