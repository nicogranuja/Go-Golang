package main

import(
	"os"
	"fmt"
	"bufio"
)

func main() {
	executeReadFile()
	//thanks to recover we kept going
	fmt.Println("Never happening because of panic unless using recover")
}

func executeReadFile(){
	exec := readFile()
	fmt.Println(exec)
}

func readFile() bool{

	//using the os package
	file,err := os.Open("./Helloo.txt")

	/*function added to the stack to be executed
	no matter what before the return statement
	*/
	defer func(){
		file.Close()
		fmt.Println("Defer")

		//prevents the script from stopping altogether
		r := recover()

		fmt.Println(r)

	}()

	if err != nil{
		//way to handle and print an error
		//stops executing script
		panic(err)
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