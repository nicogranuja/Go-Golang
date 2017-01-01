package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {

	//implementing go routine for concurrence 
	go my_slow_name("Nicolas")

	//this takes time... because the method is slow*
	//but it's implemented while the other one runs too
	fmt.Println("Im bored it is too slow!!")

	//wait for user input so program doesn't end 
	var wait string
	fmt.Scanln(&wait)
}

func my_slow_name(name string) {
	letters := strings.Split(name, "")

	//for each _ is index but we are not using it
	for _,letter := range(letters){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}