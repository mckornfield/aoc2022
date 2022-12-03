package day3

import (
	"fmt"
	"strings"
)

func getPriorityFromChar(character rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	asciiNumber := int(character)
	lowercaseStart := int('a')
	uppercaseStart := int('A')
	if asciiNumber >= lowercaseStart {
		return asciiNumber - lowercaseStart + 1
	}
	return asciiNumber - uppercaseStart + 27

}

func convertStringToSet(rucksackString string) map[rune]int {
	rucksack := make(map[rune]int)
	for _, c := range rucksackString {
		if _, exists := rucksack[c]; !exists {
			rucksack[c] = 0
		}
		rucksack[c] += 1
	}
	return rucksack
}

func findCommonItem(rucksackA, rucksackB map[rune]int) rune {
	fmt.Println(rucksackA)
	fmt.Println(rucksackB)
	for key, _ := range rucksackA {
		if _, exists := rucksackB[key]; exists {
			return key
		}
	}
	panic("Couldn't find a common item")
}

func findCommonItemThree(rucksackA, rucksackB, rucksackC map[rune]int) rune {
	// fmt.Println(rucksackA)
	// fmt.Println(rucksackB)
	for key, _ := range rucksackA {
		if _, exists := rucksackB[key]; exists {
			if _, exists := rucksackC[key]; exists {
				return key
			}
		}
	}
	panic("Couldn't find a common item")
}

func getSumOfPriorities(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
		dividedSetSize := len(line) / 2
		fmt.Println(line[:dividedSetSize])
		fmt.Println(line[dividedSetSize:])
		rucksackA := convertStringToSet(line[:dividedSetSize])
		rucksackB := convertStringToSet(line[dividedSetSize:])
		commonItem := findCommonItem(rucksackA, rucksackB)
		sum += getPriorityFromChar(commonItem)
	}
	return sum
}

func getSumOfPrioritiesFromThreeElves(input string) int {
	sum := 0
	rucksackA := make(map[rune]int)
	rucksackB := make(map[rune]int)
	rucksackC := make(map[rune]int)
	rollingCounter := 0
	for _, line := range strings.Split(input, "\n") {
		if rollingCounter == 0 {
			rucksackA = convertStringToSet(line)
			rollingCounter++
		} else if rollingCounter == 1 {
			rucksackB = convertStringToSet(line)
			rollingCounter++
		} else if rollingCounter == 2 {
			rucksackC = convertStringToSet(line)
			commonItem := findCommonItemThree(rucksackA, rucksackB, rucksackC)
			sum += getPriorityFromChar(commonItem)
			rollingCounter = 0
		}

	}
	return sum
}
