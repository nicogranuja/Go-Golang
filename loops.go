package main

import "fmt"

func main() {
	//only for loop exists
	for i:=0;i<10;i++{
		fmt.Println("Hello World")
	}

	i:=10
	//it is like a while loop
	for i<10{
		fmt.Println("Hello World")
		i++
	}

	//declaring j default value is 0
	var j int
	//indefinite loop
	//continue word works the same
	for{
		fmt.Println(j)
		j++
		//get out of the indefinite loop
		if j>10{
			break;
		}
	}
}