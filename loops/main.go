package main

import "fmt"

// Nested loops and the golang debugger, called delve.

// set breaker if you want and then start debugging

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(" i is ", i)
		for x := 1; x <= 3; x++ {
			fmt.Print("   x: ", x)
		}
		fmt.Println()
	}
}
