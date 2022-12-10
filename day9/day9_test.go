package day9

import (
	"github.com/mckornfield/aoc2022/shared"
	"testing"
)

const sampleInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

const errorMessageFmt = "Expected follow count to be %d but was %d"

// func TestProblemInput(t *testing.T) {
// 	input, err := shared.ReadFile("day9_input.txt")
// 	if err != nil {
// 		t.Fatalf("Error: %s", err)
// 	}
// 	result := getNumberOfTailPositions(input)
// 	expectedResult := 5960
// 	if result != expectedResult {
// 		t.Errorf(errorMessageFmt, expectedResult, result)
// 	}
// }

// func TestSampleInput(t *testing.T) {
// 	result := getNumberOfTailPositions(sampleInput)
// 	expectedResult := 13
// 	if result != expectedResult {
// 		t.Errorf(errorMessageFmt, expectedResult, result)
// 	}
// }

// func TestSampleInputPt2(t *testing.T) {
// 	result := getNumberOfTailPositionsForLongRope(sampleInput)
// 	expectedResult := 1
// 	if result != expectedResult {
// 		t.Errorf(errorMessageFmt, expectedResult, result)
// 	}
// }

func TestSampleInputPt2MoreUseful(t *testing.T) {
	largerInput := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	result := getNumberOfTailPositionsForLongRope(largerInput)
	expectedResult := 35
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}

func TestProblemInputPt2(t *testing.T) {
	input, err := shared.ReadFile("day9_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	result := getNumberOfTailPositionsForLongRope(input)
	expectedResult := 2297
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}
