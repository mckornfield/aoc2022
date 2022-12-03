package day2

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `A Y
B X
C Z`

const errorMessageFmt = "Expected score to be %d but was %d"

func TestSampleInput(t *testing.T) {
	score := getTotalScore(sampleInput)
	expectedScore := 15
	if score != expectedScore {
		t.Fatalf(errorMessageFmt, expectedScore, score)
	}
}

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day2_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	score := getTotalScore(input)
	expectedScore := 12535
	if score != expectedScore {
		t.Fatalf(errorMessageFmt, expectedScore, score)
	}
}

func TestSampleInputPart2(t *testing.T) {
	score := getScoreBasedOnCorrectDirections(sampleInput)
	expectedScore := 12
	if score != expectedScore {
		t.Fatalf(errorMessageFmt, expectedScore, score)
	}
}

func TestProblemInputPart2(t *testing.T) {
	input, err := shared.ReadFile("day2_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	score := getScoreBasedOnCorrectDirections(input)
	expectedScore := 15457
	if score != expectedScore {
		t.Fatalf(errorMessageFmt, expectedScore, score)
	}
}
