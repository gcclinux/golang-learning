package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		winner = DRAW
	case PAPER:
		computerChoice = "Computer chose PAPER"
		winner = PLAYERWINS
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		winner = COMPUTERWINS
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player Wins"
	} else {
		roundResult = "Computer Wins"
	}

	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
