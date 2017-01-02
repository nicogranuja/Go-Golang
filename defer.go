package main

import(
	"os"
	"fmt"
	"bufio"
)

func main() {
	exec := readFile()
	fmt.Println(exec)
}

func readFile() bool{

	//using the os package
	file,err := os.Open("./Hello.txt")

	/*function added to the stack to be executed
	no matter what before the return statement
	*/
	defer func(){
		file.Close()
		fmt.Println("Defer")
	}()

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

	return true
}