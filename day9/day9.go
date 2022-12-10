package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

var directionToCoordMapping = map[string]Coords{
	"R": {x: 1, y: 0},
	"U": {x: 0, y: 1},
	"D": {x: 0, y: -1},
	"L": {x: -1, y: 0},
}

func move(coords Coords, direction string, magntiude int) Coords {
	moveDirection := directionToCoordMapping[direction]
	newCoords := Coords{
		x: coords.x + moveDirection.x*magntiude,
		y: coords.y + moveDirection.y*magntiude,
	}
	return newCoords
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func areCoordsTouching(headCoords, tailCoords Coords) bool {
	return Abs(headCoords.x-tailCoords.x) < 2 && Abs(headCoords.y-tailCoords.y) < 2
}

func getDirectionFromCoords(target, current int) int {
	if target > current {
		return 1
	}
	return -1
}

func followTail(headCoords, tailCoords Coords, visistedSpots map[Coords]bool) Coords {
	for !areCoordsTouching(headCoords, tailCoords) {
		coordChange := Coords{0, 0}
		if headCoords.x != tailCoords.x { // Different column
			coordChange.x = getDirectionFromCoords(headCoords.x, tailCoords.x)
		}
		if headCoords.y != tailCoords.y { // Different row
			coordChange.y = getDirectionFromCoords(headCoords.y, tailCoords.y)
		}
		tailCoords = Coords{tailCoords.x + coordChange.x, tailCoords.y + coordChange.y}
		if visistedSpots != nil {
			fmt.Println(tailCoords)
			if _, ok := visistedSpots[tailCoords]; ok {
				fmt.Println(fmt.Sprintf("Already visited at %v", tailCoords))
			} else {
				visistedSpots[tailCoords] = true
			}

		}
	}
	return tailCoords
}

func getNumberOfTailPositions(input string) int {
	headCoords := Coords{x: 0, y: 0}
	tailCoords := Coords{x: 0, y: 0}
	tailPositions := map[Coords]bool{
		tailCoords: true,
	}

	for _, line := range strings.Split(input, "\n") {
		directionAndMagnitude := strings.Split(line, " ")
		direction := directionAndMagnitude[0]
		magntiude, err := strconv.Atoi(directionAndMagnitude[1])
		if err != nil {
			panic(err)
		}
		headCoords = move(headCoords, direction, magntiude)
		tailCoords = followTail(headCoords, tailCoords, tailPositions)
	}
	return len(tailPositions)
}

func printTail(rope []Coords) {

	grid := `..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
...........s..............
..........................
..........................
..........................
..........................
..........................`
	mapOfRope := make(map[Coords]int)
	for i, coords := range rope {

		// for coords.x > {
		// 	coords.x -=
		// }
		// for coord.y >  {

		// }
		mapOfRope[Coords{coords.x + 11, coords.y + 5}] = i
	}
	rows := strings.Split(grid, "\n")
	for j_rev, row := range rows {
		j := len(rows) - 1 - j_rev
		for i := range row {
			if val, ok := mapOfRope[Coords{i, j}]; ok {
				fmt.Print(val)
			} else if i == 11 && j == 5 {
				fmt.Print("s")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func getNumberOfTailPositionsForLongRope(input string) int {
	rope := make([]Coords, 10)
	for i := range rope {
		rope[i] = Coords{x: 0, y: 0}
	}

	tailPositions := map[Coords]bool{
		{0, 0}: true,
	}
	for _, line := range strings.Split(input, "\n") {
		directionAndMagnitude := strings.Split(line, " ")
		direction := directionAndMagnitude[0]
		magntiude, err := strconv.Atoi(directionAndMagnitude[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(directionAndMagnitude)
		for i := 0; i < magntiude; i++ {
			rope[0] = move(rope[0], direction, 1)
			tailIndex := len(rope) - 1
			for i := 0; i < tailIndex-1; i++ {
				rope[i+1] = followTail(rope[i], rope[i+1], nil)
			}
			rope[tailIndex] = followTail(rope[tailIndex-1], rope[tailIndex], tailPositions)
		}

	}
	// fmt.Println(tailPositions)
	return len(tailPositions)
}
