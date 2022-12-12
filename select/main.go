package main

import (
	"fmt"
	"time"
)

// package level channels

var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "one"
}

func task2() {
	time.Sleep(2 * time.Second)
	chan2 <- "two"
}

func main() {

	go task1()
	go task2()

	// All the select does is wait for information to be received by a particular channel.
	// In this case channel one for the first case and channel two for the second case.
	for x := 0; x < 2; x++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received", msg2)
		}
	}
}
