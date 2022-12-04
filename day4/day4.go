package day4

import (
	"strconv"
	"strings"
)

func getTopAndBottomOfRange(rangeNumbers string) (int, int) {
	nums := strings.Split(rangeNumbers, "-")
	lower, upper := nums[0], nums[1]
	lowerVal, err := strconv.Atoi(lower)
	if err != nil {
		panic(err)
	}
	upperVal, err2 := strconv.Atoi(upper)
	if err2 != nil {
		panic(err2)
	}
	return lowerVal, upperVal
}

func getOverlappingPairs(input string) int {
	pairs := 0
	for _, line := range strings.Split(input, "\n") {
		ranges := strings.Split(line, ",")
		firstRange, secondRange := ranges[0], ranges[1]
		firstLower, firstUpper := getTopAndBottomOfRange(firstRange)
		secondLower, secondUpper := getTopAndBottomOfRange(secondRange)
		if firstLower <= secondLower && firstUpper >= secondUpper {
			pairs += 1
		} else if secondLower <= firstLower && secondUpper >= firstUpper {
			pairs += 1
		}
	}
	return pairs
}

func numberContainedByOtherTwo(target, upper, lower int) bool {
	return target >= lower && upper <= lower
}

func getNumberset(bottom, top int) map[int]bool {
	set := make(map[int]bool)
	for i := bottom; i <= top; i++ {
		set[i] = true
	}
	return set
}

func setsOverlap(set1, set2 map[int]bool) bool {
	for i, _ := range set1 {
		if _, exists := set2[i]; exists {
			return true
		}
	}
	return false
}

func getAnyOverlappingPairs(input string) int {
	pairs := 0
	for _, line := range strings.Split(input, "\n") {
		ranges := strings.Split(line, ",")
		firstRange, secondRange := ranges[0], ranges[1]
		firstLower, firstUpper := getTopAndBottomOfRange(firstRange)
		secondLower, secondUpper := getTopAndBottomOfRange(secondRange)
		firstSet := getNumberset(firstLower, firstUpper)
		secondSet := getNumberset(secondLower, secondUpper)
		if setsOverlap(firstSet, secondSet) {
			pairs += 1
		}

	}
	return pairs
}
