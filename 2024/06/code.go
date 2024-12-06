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
	parsed := strings.Split(strings.TrimSpace(input), "\n")
	var board [][]int = make([][]int, len(parsed))
	var guard actor
	visitedPositions := make(map[coord]bool)
	for y, t := range parsed {
		board[y] = make([]int, len(t))
		for x, char := range t {
			if char == 46 {
				board[y][x] = 0
			} else if char == 35 {
				board[y][x] = 1
			} else if char == 94 {
				board[y][x] = 0
				guard = actor{c: coord{xPos: x, yPos: y}, dir: UP}
			}
		}
	}

	visitedPositions[guard.c] = true

	for inBounds := true; inBounds; inBounds = guard.isInBounds(board) {
		guard = walk(guard, board)
		if guard.isInBounds(board) {
			visitedPositions[guard.c] = true
		}
	}
	if part2 {
		return "not implemented"
	}

	return len(visitedPositions)
}

func walk(guard actor, board [][]int) actor {
	plusX, plusY := guard.dir.val(1)
	targetX := plusX + guard.c.xPos
	targetY := plusY + guard.c.yPos

	if targetY < 0 || targetY >= len(board) || targetX < 0 || targetX > len(board[targetY]) {
		// Moving out of bounds!
		return actor{c: coord{xPos: targetX, yPos: targetY}, dir: guard.dir}
	}

	if board[targetY][targetX] == 0 {
		// We can move here
		return actor{c: coord{xPos: targetX, yPos: targetY}, dir: guard.dir}
	} else if board[targetY][targetX] == 1 {
		newDir := guard.dir.turn()
		newMoveX, newMoveY := newDir.val(1)
		targetX = newMoveX + guard.c.xPos
		targetY = newMoveY + guard.c.yPos
		return actor{c: coord{xPos: targetX, yPos: targetY}, dir: newDir}
	}
	return guard
}

type coord struct {
	xPos int
	yPos int
}

type direction int32

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

func (a direction) val(distance int) (int, int) {
	switch a {
	case UP:
		return 0, -distance
	case RIGHT:
		return distance, 0
	case DOWN:
		return 0, distance
	case LEFT:
		return -distance, 0
	default:
		return 0, 0
	}
}

func (a direction) turn() direction {
	switch a {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		return a
	}
}

type actor struct {
	c   coord
	dir direction
}

func (a actor) isInBounds(board [][]int) bool {
	return a.c.xPos >= 0 && a.c.xPos < len(board[0]) && a.c.yPos >= 0 && a.c.yPos < len(board)
}
