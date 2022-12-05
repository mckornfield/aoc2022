package day5

import (
	"fmt"
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

const errorMessageFmt = "Expected message to be %s but was %s"

func TestSampleInput(t *testing.T) {
	tops := rearrangeCratesAndGetTopsOfStacks(sampleInput)
	expectedTops := "CMZ"
	if tops != expectedTops {
		t.Errorf(errorMessageFmt, expectedTops, tops)
	}
}

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day5_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	tops := rearrangeCratesAndGetTopsOfStacks(input)
	expectedTops := "ZBDRNPMVH"
	if tops != expectedTops {
		t.Errorf(errorMessageFmt, expectedTops, tops)
	}
}

func TestCreateStacks(t *testing.T) {
	stacks := createStacks([]string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	})
	if len(stacks) != 3 {
		t.Error("Expected length of stacks to be 3")
	}
	if stacks[0].Peek() != "N" {
		t.Errorf("Expected a value of %s but got %s", "N", stacks[0].Peek())
	}

	if stacks[1].Peek() != "D" {
		t.Errorf("Expected a value of %s but got %s", "D", stacks[1].Peek())
	}

	if stacks[2].Peek() != "P" {
		t.Errorf("Expected a value of %s but got %s", "P", stacks[2].Peek())
	}
}

func TestMoveStacks(t *testing.T) {
	stacks := createStacks([]string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	})
	stacks = performMoves(stacks, "move 1 from 2 to 1")
	val := stacks[0].Peek()
	fmt.Println(stacks)
	if val != "D" {
		t.Errorf("New top of stack 1 should be D")
	}
	stacks = performMoves(stacks, "move 2 from 2 to 3")
	val = stacks[2].Peek()
	fmt.Println(stacks)
	if val != "M" {
		t.Errorf("New top of stack 3 should be M")
	}
}

func TestSampleInputPt2(t *testing.T) {
	tops := rearrangeCratesAndGetTopsOfStacksImproved(sampleInput)
	expectedTops := "MCD"
	if tops != expectedTops {
		t.Errorf(errorMessageFmt, expectedTops, tops)
	}
}

func TestProblemInputPt2(t *testing.T) {
	input, err := shared.ReadFile("day5_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	tops := rearrangeCratesAndGetTopsOfStacksImproved(input)
	expectedTops := "WDLPFNNNB"
	if tops != expectedTops {
		t.Errorf(errorMessageFmt, expectedTops, tops)
	}
}
