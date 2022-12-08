package day8

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `30373
25512
65332
33549
35390`

const errorMessageFmt = "Expected visible tree count to be %d but was %d"

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day8_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	result := getVisibleTreeCount(input)
	expectedResult := 1798
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}

func TestProblemInputPt2(t *testing.T) {
	input, err := shared.ReadFile("day8_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	result := getHighestScenicScore(input)
	expectedResult := 259308
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}

func TestSampleInput(t *testing.T) {
	result := getVisibleTreeCount(sampleInput)
	expectedResult := 21
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}

func TestSampleInputPt2(t *testing.T) {
	result := getHighestScenicScore(sampleInput)
	expectedResult := 8
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}
