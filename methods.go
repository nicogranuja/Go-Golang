package main

import "fmt"


type User struct{
	age int
	name, lastName string
}

func (this User) full_name() string {
	return this.name + " " + this.lastName
}

func (this *User) set_name(name string){
	this.name = name
}

func main() {
	nico := new(User)

	nico.name = "Nico"
	nico.lastName = "Mondragon"

	fmt.Println(nico.full_name())

	//passing the struct as a pointer for the set method
	nico.set_name("Alvarez")

	fmt.Println(nico.full_name())	
}