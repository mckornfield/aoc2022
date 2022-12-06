package day6

type RollingSet map[rune]int

func (r *RollingSet) addItem(item rune) {
	if _, exists := (*r)[item]; !exists {
		(*r)[item] = 0
	}
	(*r)[item] += 1
}

func (r *RollingSet) removeItem(item rune) {
	(*r)[item] -= 1
	if (*r)[item] == 0 {
		delete(*r, item)
	}
}
func getIndexOfFourUnmatchedChars(input string) int {
	return getIndexOfUnmatchedChars(input, 4)
}

func getIndexOfFourteenUnmatchedChars(input string) int {
	return getIndexOfUnmatchedChars(input, 14)
}

func getIndexOfUnmatchedChars(input string, unmatchedCount int) int {
	runningSetOfChars := RollingSet{}
	runeSliceInput := []rune(input)
	for i, c := range runeSliceInput {
		oldestIndex := i - unmatchedCount
		if oldestIndex > -1 {
			runningSetOfChars.removeItem(runeSliceInput[oldestIndex])
		}
		runningSetOfChars.addItem(c)
		if len(runningSetOfChars) == unmatchedCount {
			return i + 1
		}
	}
	return -1
}
