package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element
	}
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func performMoves(stacks []Stack, line string) []Stack {
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

func performMovesAdvanced(stacks []Stack, line string) []Stack {
	r := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	results := r.FindStringSubmatch(line)
	numberOfMoves, _ := strconv.Atoi(results[1])
	oldStackIndex, _ := strconv.Atoi(results[2])
	newStackIndex, _ := strconv.Atoi(results[3])
	tempStack := Stack{}
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
	stacks := make([]Stack, 5)
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
		sb.WriteString(stacks[i].Peek())
	}
	return sb.String()
}

func rearrangeCratesAndGetTopsOfStacksImproved(input string) string {
	stackLines := make([]string, 0)
	stacks := make([]Stack, 5)
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
		sb.WriteString(stacks[i].Peek())
	}
	return sb.String()
}

func createStacks(stackLines []string) []Stack {
	// "[X] [Y] [Z]"
	sampleBox := "[X] "
	stackCount := (len(stackLines[0]) + 1) / len(sampleBox)
	fmt.Println(len(stackLines))
	stacks := make([]Stack, stackCount)
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
