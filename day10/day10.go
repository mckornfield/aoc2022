package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func checkSignalStrength(registerVal, cycle int) int {
	// fmt.Println("Cycle", cycle, "Val", registerVal)
	if cycle != 0 && (cycle-20)%40 == 0 {
		val := registerVal * cycle
		return val
	}
	return 0
}

func getSumOfSixSignalStrengths(input string) int {
	signalStrengthSum := 0
	cycle := 1
	registerVal := 1
	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			cycle++
		} else if strings.HasPrefix(line, "addx") {
			splitLine := strings.Split(line, " ")
			val, err := strconv.Atoi(splitLine[1])
			if err != nil {
				panic(err)
			}
			cycle++
			signalStrengthSum += checkSignalStrength(registerVal, cycle)
			cycle++
			registerVal += val
		}
		signalStrengthSum += checkSignalStrength(registerVal, cycle)
	}

	return signalStrengthSum
}

func decidePixelToDraw(registerVal, currentCycle int) string {
	var sb strings.Builder
	modCycle := currentCycle % 40
	if modCycle-2 <= registerVal && modCycle >= registerVal {
		sb.WriteRune('#')
	} else {
		sb.WriteRune('.')
	}
	fmt.Println(currentCycle, registerVal, sb.String())
	if currentCycle != 0 && modCycle == 0 {
		sb.WriteRune('\n')
	}
	return sb.String()
}

func getPicture(input string) string {
	var sb strings.Builder
	cycle := 1
	registerVal := 1
	// sb.WriteString(decidePixelToDraw(registerVal, cycle))
	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			sb.WriteString(decidePixelToDraw(registerVal, cycle))
			cycle++
		} else if strings.HasPrefix(line, "addx") {
			splitLine := strings.Split(line, " ")
			val, err := strconv.Atoi(splitLine[1])
			if err != nil {
				panic(err)
			}
			sb.WriteString(decidePixelToDraw(registerVal, cycle))
			cycle++
			sb.WriteString(decidePixelToDraw(registerVal, cycle))
			cycle++
			registerVal += val
		}
		// sb.WriteString(decidePixelToDraw(registerVal, cycle))
	}

	return sb.String()
}
