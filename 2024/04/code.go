package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var wordSearch = 0

	if part2 {
		for y := 1; y < len(lines); y++ {
			// column:
			for x := 1; x < len(lines[y]); x++ {
				if searchXMas(lines, x, y) {
					wordSearch++
				}
			}
		}
	} else {
		// row:
		for y, line := range lines {
			// column:
			for x, _ := range line {
				// directionLoop:
				for _, direction := range allDirections {
					for position := 0; position < 4; position++ {
						if searchDirection(lines, direction, x, y, position) {
							if position == 3 {
								wordSearch++
							}
							continue
						} else {
							break
						}
					}
				}
			}

		}
	}
	return wordSearch
}

func searchXMas(data []string, xVal int, yVal int) bool {
	if xVal == 0 || yVal == 0 || yVal >= len(data)-1 || xVal >= len(data[yVal])-1 {
		return false
	}
	if data[yVal][xVal] != 65 {
		return false
	}
	upRightChar := data[yVal-1][xVal+1]
	downRightChar := data[yVal+1][xVal+1]
	downLeftChar := data[yVal+1][xVal-1]
	upLeftChar := data[yVal-1][xVal-1]

	// A / MAS
	if !((upRightChar == 77 && downLeftChar == 83) || (upRightChar == 83 && downLeftChar == 77)) {
		return false
	}

	// A \ MAS
	if !((upLeftChar == 77 && downRightChar == 83) || (upLeftChar == 83 && downRightChar == 77)) {
		return false
	}

	// println("Center coord", xVal, yVal)
	return true
}

type direction int32

const (
	UP direction = iota
	UP_RIGHT
	RIGHT
	DOWN_RIGHT
	DOWN
	DOWN_LEFT
	LEFT
	UP_LEFT
)

var allDirections = []direction{UP, UP_RIGHT, RIGHT, DOWN_RIGHT, DOWN, DOWN_LEFT, LEFT, UP_LEFT}

func (a direction) val(distance int) (int, int) {
	switch a {
	case UP:
		return 0, -distance
	case UP_RIGHT:
		return distance, -distance
	case RIGHT:
		return distance, 0
	case DOWN_RIGHT:
		return distance, distance
	case DOWN:
		return 0, distance
	case DOWN_LEFT:
		return -distance, distance
	case LEFT:
		return -distance, 0
	case UP_LEFT:
		return -distance, -distance
	default:
		return 0, 0
	}
}

func searchDirection(data []string, direction direction, xVal int, yVal int, position int) bool {
	// X 88
	// M 77
	// A 65
	// S 83
	xDistance, yDistance := direction.val(position)
	xCoord := xVal + xDistance
	yCoord := yVal + yDistance
	if xCoord < 0 || yCoord < 0 || yCoord >= len(data) || xCoord >= len(data[yCoord]) {
		// Out of bounds
		return false
	}
	testChar := data[yCoord][xCoord]
	switch position {
	case 0:
		return testChar == 88
	case 1:
		return testChar == 77
	case 2:
		return testChar == 65
	case 3:
		return testChar == 83
	default:
		return false
	}
}
