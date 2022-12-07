package day7

import (
	"testing"

	"github.com/mckornfield/aoc2022/shared"
)

const sampleInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

const errorMessageFmt = "Expected folder deletion count to be %d but was %d"

func TestProblemInput(t *testing.T) {
	input, err := shared.ReadFile("day7_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	result := getFolderDeletionCount(input)
	expectedResult := 1778099
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}
func TestProblemInputPt2(t *testing.T) {
	input, err := shared.ReadFile("day7_input.txt")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	result := getDirectoryToDelete(input)
	expectedResult := 1623571
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}

func TestSampleInput(t *testing.T) {
	result := getFolderDeletionCount(sampleInput)
	expectedResult := 95437
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}

func TestSampleInputPt2(t *testing.T) {
	result := getDirectoryToDelete(sampleInput)
	expectedResult := 24933642
	if result != expectedResult {
		t.Errorf(errorMessageFmt, expectedResult, result)
	}
}
