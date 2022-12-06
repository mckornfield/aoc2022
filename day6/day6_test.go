package day6

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `bvwbjplbgvbhsrlpgdmjqwftvncz`
const sampleInput1 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
const sampleInput2 = `nppdvjthqldpwncqszvftbrmjlhg`
const sampleInput3 = `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
const sampleInput4 = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

const errorMessageFmt = "Expected first marker to be %d but was %d"

var test_configs = []struct {
	input             string
	expectedNumber    int
	expectedPt2Number int
}{
	{input: sampleInput1, expectedNumber: 7, expectedPt2Number: 19},
	{input: sampleInput, expectedNumber: 5, expectedPt2Number: 23},
	{input: sampleInput2, expectedNumber: 6, expectedPt2Number: 23},
	{input: sampleInput3, expectedNumber: 10, expectedPt2Number: 29},
	{input: sampleInput4, expectedNumber: 11, expectedPt2Number: 26},
}

func TestSampleInput(t *testing.T) {
	for _, tc := range test_configs {
		marker := getIndexOfFourUnmatchedChars(tc.input)
		if tc.expectedNumber != marker {
			t.Errorf(errorMessageFmt, tc.expectedNumber, marker)
		}
	}
}

func TestSampleInputPart2(t *testing.T) {
	for _, tc := range test_configs {
		marker := getIndexOfFourteenUnmatchedChars(tc.input)
		if tc.expectedPt2Number != marker {
			t.Errorf(errorMessageFmt, tc.expectedPt2Number, marker)
		}
	}
}

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day6_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	marker := getIndexOfFourUnmatchedChars(input)
	expected := 1623
	if expected != marker {
		t.Errorf(errorMessageFmt, expected, marker)
	}

}

func TestProblemInputPt2(t *testing.T) {
	input, err := shared.ReadFile("day6_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	marker := getIndexOfFourteenUnmatchedChars(input)
	expected := 3774
	if expected != marker {
		t.Errorf(errorMessageFmt, expected, marker)
	}

}
