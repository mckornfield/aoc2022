package day

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

func TestProblemInput(t *testing.T) {
	_, err := shared.ReadFile("day_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}
