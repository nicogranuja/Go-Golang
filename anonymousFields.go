package main

import "fmt"

type Human struct{
	name string
}

type Tutor struct{
	//anonymous field
	Human
	Dummy
}

type Dummy struct{
	name string
}

func (this Human) speak() string{
	return "bla bla bla"
}

//overwriting the func speak
func (this Tutor) speak() string{
	return "Hello Tutor"
}

func main() {
	//name would be ambiguous
	tutor := Tutor {Human{"Nico"}, Dummy{"Uriel"}}

	fmt.Println(tutor.Human.name, tutor.Dummy.name)

	
	fmt.Println(tutor.speak())
}