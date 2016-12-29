package main

import (
	"fmt"
)

func main() {
	/*
		1. Address in memory
		2. Instead of variable's value
			we get the address in memory
		3. X,Y => as123d => 6 (x and y point to the same address 
			and have a value of 6)
		4. X=> as123d => 5 (if we change x to 5)
			now Y changes to 5 too
		5. Syntax: *T => *int *string *Struct
		6. Zero value => nil	
	*/
	var x,y *int
	integer := 5

	x = &integer
	y = &integer
	//to deal with the variable's value
	//we use an asterisk again
	//by this time both have the same value
	fmt.Println(*x,*y)

	//now if we change *x value we will change *y
	*x = 6

	fmt.Println(*x,*y)
}