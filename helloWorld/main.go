package main

import "fmt"

func main() {
	// Standard string varible....
	//var sayWhat string
	//sayWhat = "Hello World, through function & variable"

	// Short hand string varible.
	sayWhat := "Hello World, through function & short hand variable !!!"
	sayHelloWorld(sayWhat)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
