package main

import "fmt"

//Interface defines empty methods for structs (not implemented)
type User interface{
	Permissions() int // 1-5
	Name() string
}


//by implementing the methods we are using the User interface*
type Admin struct{
	name string
}

type Editor struct{
	name string
}

//begin methods implemented
func (this Admin) Permissions() int {
	return 5
}

func (this Admin) Name() string {
	return this.name
}

func (this Editor) Permissions() int {
	return 3
}

func (this Editor) Name() string {
	return this.name
}
//end methods implemented


//function for the interface that takes a user struct
//we send admin that implements user
func Auth(user User) string{
	if user.Permissions() > 4 {
		return user.Name() + " has admin access."
	}
	return user.Name() +" does not have access."
}

func main() {
	//admin implements user
	admin := Admin{"Nicolas"}

	editor := Editor{"Mondragon"}

	fmt.Println(Auth(admin))
	fmt.Println(Auth(editor))

	//now with an array of structs
	users := []User{Admin{"Nicolas"},Editor{"Mondragon"}}

	//foreach in go? 
	for _, user := range users{
		fmt.Println(Auth(user))
	}

}