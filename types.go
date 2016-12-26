package main

import (
	"fmt"
	"strconv"
)

func main() {
	//converting int to string
	age := 22

	age_str := strconv.Itoa(age)
	//age has to be converted to string to be concatenated
	fmt.Println(age_str + " my age now")


	//Atoi returns two values
	//second is an error but since we are not using it we use _
	age_int,_ := strconv.Atoi(age_str)

	fmt.Println(age_int + 22)
}