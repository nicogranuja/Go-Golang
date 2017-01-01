package main

import("fmt")

func main() {
	//defining channel
	channel := make(chan string)

	//anonymous function in a go routine
	go func(channel chan string){
		//infinite loop*
		for{
			var name string
			fmt.Scanln(&name)

			//sending to channel
			channel <- name
		}
	}(channel)//using the channel we created


	//receiving from channel
	//as soon as channel gives the data it stops the infinite loop
	msg := <- channel
	fmt.Println("The message " + msg)
}