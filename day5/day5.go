package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mckornfield/aoc2022/shared"
)

func performMoves(stacks []shared.Stack, line string) []shared.Stack {
	r := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	results := r.FindStringSubmatch(line)
	fmt.Println(results)
	numberOfMoves, _ := strconv.Atoi(results[1])
	oldStackIndex, _ := strconv.Atoi(results[2])
	newStackIndex, _ := strconv.Atoi(results[3])
	fmt.Println(numberOfMoves, oldStackIndex, newStackIndex)
	for i := 0; i < numberOfMoves; i++ {
		// Off by one
		value, _ := stacks[oldStackIndex-1].Pop()
		stacks[newStackIndex-1].Push(value)
	}
	return stacks

}

func performMovesAdvanced(stacks []shared.Stack, line string) []shared.Stack {
	r := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	results := r.FindStringSubmatch(line)
	numberOfMoves, _ := strconv.Atoi(results[1])
	oldStackIndex, _ := strconv.Atoi(results[2])
	newStackIndex, _ := strconv.Atoi(results[3])
	tempStack := shared.Stack{}
	for i := 0; i < numberOfMoves; i++ {
		// Off by one
		value, _ := stacks[oldStackIndex-1].Pop()
		tempStack.Push(value)
	}
	for i := 0; i < numberOfMoves; i++ {
		value, _ := tempStack.Pop()
		stacks[newStackIndex-1].Push(value)
	}
	return stacks

}

func rearrangeCratesAndGetTopsOfStacks(input string) string {
	stackLines := make([]string, 0)
	stacks := make([]shared.Stack, 5)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "move") {
			performMoves(stacks, line)
			// Do move things
		} else if strings.HasPrefix(line, " 1") {
			fmt.Println(stackLines)
			stacks = createStacks(stackLines)
		} else {
			stackLines = append(stackLines, line)
		}
	}
	var sb strings.Builder
	for i := 0; i < len(stacks); i++ {
		sb.WriteString(stacks[i].Peek().(string))
	}
	return sb.String()
}

func rearrangeCratesAndGetTopsOfStacksImproved(input string) string {
	stackLines := make([]string, 0)
	stacks := make([]shared.Stack, 5)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "move") {
			performMovesAdvanced(stacks, line)
			// Do move things
		} else if strings.HasPrefix(line, " 1") {
			stacks = createStacks(stackLines)
		} else {
			stackLines = append(stackLines, line)
		}
	}
	var sb strings.Builder
	for i := 0; i < len(stacks); i++ {
		sb.WriteString(stacks[i].Peek().(string))
	}
	return sb.String()
}

func createStacks(stackLines []string) []shared.Stack {
	// "[X] [Y] [Z]"
	sampleBox := "[X] "
	stackCount := (len(stackLines[0]) + 1) / len(sampleBox)
	fmt.Println(len(stackLines))
	stacks := make([]shared.Stack, stackCount)
	for i := range stackLines { // go through in reverse
		line := stackLines[len(stackLines)-1-i]
		// Move in 3 character chunks
		for lineIndex := 0; lineIndex < len(line); lineIndex += 4 {
			possibleLetterChar := string(line[lineIndex+1])
			fmt.Println(possibleLetterChar)
			if possibleLetterChar != " " {
				stacks[lineIndex/4].Push(possibleLetterChar)
			}
		}

	}
	fmt.Println(stacks)
	return stacks

}
