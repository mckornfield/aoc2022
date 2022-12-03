package day3

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

func TestPriorityCalc(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  int
	}{
		{name: "lower case a", input: 'a', want: 1},
		{name: "lower case p", input: 'p', want: 16},
		{name: "upper case A", input: 'A', want: 27},
		{name: "upper case L", input: 'L', want: 38},
		{name: "upper case P", input: 'P', want: 42},
	}
	for _, tc := range tests {
		result := getPriorityFromChar(tc.input)
		if tc.want != result {
			t.Errorf("%s: expected: %d, actual: %d", tc.name, tc.want, result)
		}
	}
}

const sampleInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

const errorMessageFmt = "Expected prioritySum to be %d but was %d"

func TestSampleInput(t *testing.T) {
	prioritySum := getSumOfPriorities(sampleInput)
	expectedSum := 157
	if prioritySum != expectedSum {
		t.Fatalf(errorMessageFmt, expectedSum, prioritySum)
	}
}

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day3_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	prioritySum := getSumOfPriorities(input)
	expectedSum := 7785
	if prioritySum != expectedSum {
		t.Fatalf(errorMessageFmt, expectedSum, prioritySum)
	}
}

func TestSampleInputPartTwo(t *testing.T) {
	prioritySum := getSumOfPrioritiesFromThreeElves(sampleInput)
	expectedSum := 70
	if prioritySum != expectedSum {
		t.Fatalf(errorMessageFmt, expectedSum, prioritySum)
	}
}

func TestPuzzleInputPartTwo(t *testing.T) {
	input, err := shared.ReadFile("day3_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	prioritySum := getSumOfPrioritiesFromThreeElves(input)
	expectedSum := 2633
	if prioritySum != expectedSum {
		t.Fatalf(errorMessageFmt, expectedSum, prioritySum)
	}
}
