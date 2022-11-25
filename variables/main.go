package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed radon number generator
	rand.Seed(time.Now().UnixNano())

	// one way - declare, then assign (two steps)
	// (1) var firstNumber int
	// (2) firstNumber = 2
	var firstNumber = rand.Intn(8) + 2

	// another way, declare type and name and assign value
	var secondNumber = rand.Intn(8) + 2

	// single line variable: declare, assign, let GO figure out type.
	subtraction := rand.Intn(8) + 2

	// Calculate the answer before the questions
	var answer = firstNumber*secondNumber - subtraction

	askQuestion(firstNumber, secondNumber, subtraction, answer)

}

func askQuestion(firstNumber, secondNumber, subtraction, answer int) {
	// Reader variable using bufio package and using os package
	reader := bufio.NewReader(os.Stdin)

	// need to display a welcome and instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 nd 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	fmt.Println("Multiple your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now Multiple the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now Divide the result by the original number you thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer

	fmt.Println("The answer is", answer)
}
