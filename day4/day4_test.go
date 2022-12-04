package day4

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

const errorMessageFmt = "Expected assignment pairs to be %d but was %d"

func TestProblemSampleInput(t *testing.T) {
	pairs := getOverlappingPairs(sampleInput)
	expectedPairs := 2
	if pairs != expectedPairs {
		t.Fatalf(errorMessageFmt, expectedPairs, pairs)
	}
}

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day4_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	pairs := getOverlappingPairs(input)
	expectedPairs := 490
	if pairs != expectedPairs {
		t.Fatalf(errorMessageFmt, expectedPairs, pairs)
	}
}

func TestProblemSampleInputPt2(t *testing.T) {
	pairs := getAnyOverlappingPairs(sampleInput)
	expectedPairs := 4
	if pairs != expectedPairs {
		t.Fatalf(errorMessageFmt, expectedPairs, pairs)
	}
}

func TestProblemInputPart2(t *testing.T) {
	input, err := shared.ReadFile("day4_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	pairs := getAnyOverlappingPairs(input)
	expectedPairs := 921
	if pairs != expectedPairs {
		t.Fatalf(errorMessageFmt, expectedPairs, pairs)
	}
}
