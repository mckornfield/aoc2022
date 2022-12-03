package day1

import (
	"math"
	"strconv"
	"strings"
)

// import (
// 	"github.com/mckornfield/aoc2022/shared"
// )

func getMaxCalories(calorieData string) int {
	runningSum := 0
	maxSum := 0
	for _, line := range strings.Split(calorieData, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			maxSum = int(math.Max(float64(maxSum), float64(runningSum)))
			runningSum = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		runningSum += val
	}
	return maxSum

}

func reconcileMaxSums(runningSum int, maxSums []int) []int {
	for index, val := range maxSums {
		if runningSum > val {
			maxSums = append(maxSums[:index+1], maxSums[index:]...)
			maxSums[index] = runningSum
			break
		}
	}
	// fmt.Print(maxSums)
	if len(maxSums) > 3 {
		maxSums = maxSums[:3]
	}
	// fmt.Print(maxSums)
	return maxSums
}
func getTopThreeMaxCalories(calorieData string) int {
	runningSum := 0
	maxSums := []int{0}
	for _, line := range strings.Split(calorieData, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			maxSums = reconcileMaxSums(runningSum, maxSums)
			runningSum = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		runningSum += val
	}
	maxSums = reconcileMaxSums(runningSum, maxSums)

	sumOfSums := 0
	for _, val := range maxSums {
		sumOfSums += val
	}
	return sumOfSums
}
