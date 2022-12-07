package main

import (
	"bufio"
	"fmt"
	"myapp/eliza/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1) //delete new line from string for windows -1 = everywhere
		userInput = strings.Replace(userInput, "\n", "", -1)   //delete new line from string for none windows -1 = everywhere

		fmt.Println(doctor.Response(userInput))

		if userInput == "quit" {
			break
		}
	}
}
