package day8

import (
	"fmt"
	"strings"
)

func printGrid(grid [][]int) {
	for _, val := range grid {
		fmt.Println(val)
	}
}

func isEdge(grid [][]int, i, j int) bool {
	return i == 0 || j == 0 || len(grid)-1 == i || len(grid[0])-1 == j

}

func areTreesShorterInDirection(grid [][]int, i, j, iDirection, jDirection int) bool {
	currentComparison := []int{i + iDirection, j + jDirection}
	for {
		currentVal := grid[i][j]
		i_1, j_1 := currentComparison[0], currentComparison[1]
		treeThatMustBeShorter := grid[i_1][j_1]
		if treeThatMustBeShorter >= currentVal {
			return false
		}
		if isEdge(grid, currentComparison[0], currentComparison[1]) {
			return true
		}
		currentComparison = []int{i_1 + iDirection, j_1 + jDirection}

	}
}

func isTreeVisible(grid [][]int, i, j int) bool {
	if isEdge(grid, i, j) {
		return true
	}
	for _, comparisonDirection := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		isVisible := areTreesShorterInDirection(grid, i, j, comparisonDirection[0], comparisonDirection[1])
		if isVisible {
			// fmt.Println("Tree is visible", grid[i][j], i, j)
			return true
		}
	}
	return false

}

func safeIndex(grid [][]int, row, column int) int {
	if row < 0 || row > len(grid)-1 {
		return -1
	}
	if column < 0 || column > len(grid[0])-1 {
		return -1
	}
	return grid[row][column]
}

func buildGrid(input string) [][]int {
	grid := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		rowSize := len(line)
		row := make([]int, rowSize)
		for j, c := range line {
			val := int(c - '0')
			row[j] = val
		}
		grid = append(grid, row)
	}
	return grid
}

func getVisibleTreeCount(input string) int {
	grid := buildGrid(input)
	visibleCount := 0
	for i, row := range grid {
		for j := range row {
			isVisible := isTreeVisible(grid, i, j)
			if isVisible {
				visibleCount++
			}
		}
	}
	return visibleCount
}

func getScenicScore(grid [][]int, i, j int) int {
	score := 1
	for _, comparisonDirection := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		i_del, j_del := comparisonDirection[0], comparisonDirection[1]
		indexCheck := safeIndex(grid, i+i_del, j+j_del)
		if indexCheck == -1 {
			continue
		}
		scoreInDirection := getScenicScoreInDirection(grid, i, j, i_del, j_del)
		score *= scoreInDirection
	}
	return score

}

func getScenicScoreInDirection(grid [][]int, i, j, iDirection, jDirection int) int {
	currentComparison := []int{i + iDirection, j + jDirection}
	scoreCount := 1
	for {
		currentVal := grid[i][j]
		i_1, j_1 := currentComparison[0], currentComparison[1]
		treeThatMustBeShorter := grid[i_1][j_1]
		if treeThatMustBeShorter >= currentVal {
			return scoreCount
		}
		if isEdge(grid, currentComparison[0], currentComparison[1]) {
			return scoreCount
		}
		currentComparison = []int{i_1 + iDirection, j_1 + jDirection}
		scoreCount++
	}
}

func getHighestScenicScore(input string) int {
	grid := buildGrid(input)
	highestScenicScore := 0
	for i, row := range grid {
		for j := range row {
			scenicScore := getScenicScore(grid, i, j)
			if highestScenicScore < scenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	return highestScenicScore
}
