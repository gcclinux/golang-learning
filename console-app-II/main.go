package main

// My version of the 56. Solution to Challange

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// defer = as soon as
	// This function is to close the keyboard function that was just opened as soon as the above is ready
	// If the above throughs an error log it and ignore it by closing it.
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the Program")

	char := ' '
	var next int

	for next == 0 {
		// Works well on Unix & darwin but buggy on windows
		//char, _, err := keyboard.GetSingleKey()
		char, _, err = keyboard.GetKey() // Works well on Windows 10

		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))

		if char >= 49 && char <= 54 {
			fmt.Printf("You chose %s\n", coffees[i])
		} else if char == 81 || char == 113 {
			fmt.Println()
			fmt.Println("You choose to quit with", char)
			fmt.Println()
			next = int(char)
		}
	}

	fmt.Println("Program Exiting .....")
}
