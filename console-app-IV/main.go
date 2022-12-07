package main

// Provided idea and solution to 58. The do while loop in go using the for loop

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
	fmt.Println("Q - Quit the program")

	char := ' '

	for ok := true; ok; ok = char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))

		// if value char exist in the make it's ok and continue
		// The same example, none are better than the other, is just which is easier to read and understand
		if _, ok := coffees[i]; ok {
			fmt.Printf("You chose %s\n", coffees[i])
		}
	}

	fmt.Println("Program exiting...")
}
