package day1

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

const errorMessageFmt = "Expected max cals to be %d but was %d"

func TestSampleInput(t *testing.T) {
	maxCals := getMaxCalories(sampleInput)
	expectedCals := 24000
	if maxCals != expectedCals {
		t.Errorf(errorMessageFmt, expectedCals, maxCals)
	}
}

func TestProblemInput(t *testing.T) {
	val, err := shared.ReadFile("day1_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	maxCals := getMaxCalories(val)
	expectedCals := 71506
	if maxCals != expectedCals {
		t.Errorf(errorMessageFmt, expectedCals, maxCals)
	}
}

func TestSampleInputPartTwo(t *testing.T) {
	maxCals := getTopThreeMaxCalories(sampleInput)
	expectedCals := 45000
	if maxCals != expectedCals {
		t.Errorf(errorMessageFmt, expectedCals, maxCals)
	}
}

func TestProblemInputPartTwo(t *testing.T) {

	val, err := shared.ReadFile("day1_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	maxCals := getTopThreeMaxCalories(val)
	expectedCals := 209603
	if maxCals != expectedCals {
		t.Errorf(errorMessageFmt, expectedCals, maxCals)
	}
}
