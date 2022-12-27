package main

import "fmt"

func Main() {

	count := 1
	for i := 0; i <= 100; i = i + 2 {
		fmt.Print(count, " - > ")
		fmt.Println("i is", i)
		count++
	}
}
