package day2

import (
	"strings"
)

var yourChoicePoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var letterToThrow = map[string]string{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

// Opponent throw first, and then yours
var youWinCombos = map[string]bool{
	Rock + Paper:     true,
	Paper + Scissors: true,
	Scissors + Rock:  true,
}

func getMatchScore(opponentLetter string, yourLetter string) int {
	// A is Rock, B is Paper, C is Scissors for Opponent
	// X is Rock, Y is Paper, Z is Scissors for You
	// 1          2           3 are the mappings for your point
	score := yourChoicePoints[yourLetter]

	yourThrow := letterToThrow[yourLetter]
	opponentThrow := letterToThrow[opponentLetter]

	if yourThrow == opponentThrow {
		// Draw
		score += 3
	} else if _, ok := youWinCombos[opponentThrow+yourThrow]; ok { // You won
		score += 6

	} // You lose on fall through, no score change
	return score
}

func getTotalScore(playsAndStrat string) int {
	score := 0
	for _, line := range strings.Split(playsAndStrat, "\n") {
		items := strings.Split(line, " ")
		if len(items) != 2 {
			panic(line)
		}
		score += getMatchScore(items[0], items[1])
	}

	return score
}

var decisionPoints = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

const (
	Rock     string = "ROCK"
	Paper           = "PAPER"
	Scissors        = "SCISSORS"
)

var yourChoiceNamedPoints = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func getScoreBasedOnDesiredOutcome(opponentChoice string, yourDecision string) int {
	// A is Rock, B is Paper, C is Scissors for Opponent
	// X means lose, Y means draw, Z means win

	score := decisionPoints[yourDecision]

	youLose := yourDecision == "X"
	youTie := yourDecision == "Y"
	youWin := yourDecision == "Z"

	opponentIsRock := opponentChoice == "A"
	opponentIsPaper := opponentChoice == "B"
	opponentIsScissors := opponentChoice == "C"
	yourChoice := ""
	// X is Rock, Y is Paper, Z is Scissors for You
	// 1          2           3 are the mappings for your point
	if opponentIsPaper {
		if youWin {
			yourChoice = Scissors
		} else if youLose {
			yourChoice = Rock
		} else if youTie {
			yourChoice = Paper
		}

	} else if opponentIsRock {
		if youWin {
			yourChoice = Paper
		} else if youLose {
			yourChoice = Scissors
		} else if youTie {
			yourChoice = Rock
		}
	} else if opponentIsScissors {
		if youWin {
			yourChoice = Rock
		} else if youLose {
			yourChoice = Paper
		} else if youTie {
			yourChoice = Scissors
		}
	}
	score += yourChoiceNamedPoints[yourChoice]
	return score
}

func getScoreBasedOnCorrectDirections(playsAndStrat string) int {
	score := 0
	for _, line := range strings.Split(playsAndStrat, "\n") {
		items := strings.Split(line, " ")
		if len(items) != 2 {
			panic(line)
		}
		score += getScoreBasedOnDesiredOutcome(items[0], items[1])
	}

	return score
}
