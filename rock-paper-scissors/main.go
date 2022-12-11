package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Always type constants in all UPPER case letters
const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	// First option and way of doing...
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}
	fmt.Println()
	fmt.Println("Player chose", playerChoice, "and value is", playerValue)
	fmt.Println("Computer chose", computerValue)
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
