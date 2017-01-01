package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	//using the os package
	file,err := os.Open("./Hello.txt")

	if err != nil{
		fmt.Println("Error!")
	}
	//create scanner to read the file
	scanner := bufio.NewScanner(file)

	var i int

	//iterate through the scanner
	for scanner.Scan(){
		i++
		//save the line of text that comes from the scanner
		line := scanner.Text()

		fmt.Println(i)
		fmt.Println(line)
	}
}