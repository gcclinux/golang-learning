package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// package level variable = constant
var reader *bufio.Reader

// user defined type - more than 1 type in a single structure
type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {

	reader = bufio.NewReader(os.Stdin)

	// Original withou User defined structure
	// userName := readString("What is your name?")
	// age := readInt("How old are you?")

	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite Number:")

	/* Lesson 24
	// String interpolations removing spaces and adjusting output

	// Standards
	fmt.Println("Your Name is:", userName, ". You are", age, "years old.")

	// using concatinator, it's slower and need more memory
	fmt.Println("Your Name is: "+userName+". You are", age, "years old.")

	// using the Sprintf "S" String print "f" format
	fmt.Println(fmt.Sprintf("Your Name is: %s. You are %d years old.", userName, age))

	// using printf and even better than above ut will not return new line at the end so need to add \n
	fmt.Printf("Your Name is: %s. You are %d years old.\n", userName, age)
	*/

	// create our own varible type - user defined type

	fmt.Printf("Your Name is: %s. You are %d years old. And your favourite number is %g\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber)

}

func prompt() {
	fmt.Print("-> ")
}

// (s read string) return string
func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}
