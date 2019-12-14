package main

import (
	"adventofcode/lib"
	"fmt"
	"os"
)

func main() {
	robot := makeRobot(lib.ReadIntcode("input"))
	robot.run()

	fmt.Println(len(robot.surface))

	robot.surface.paint()
}

// Type definitions

type robot struct {
	pointer      int
	intcode      []int
	relativeBase int

	position  coord
	direction direction
	surface   surface
}

func (a *robot) run() {
	currentColour := a.surface.colour(a.position)

	output := a.execute(int(currentColour))

	if len(output) == 2 {
		a.surface[a.position] = colour(output[0])
		a.direction = a.direction.turn(output[1])
		a.position = a.position.add(a.direction.vector())
		a.run()
	} else {
		return
	}
}

func makeRobot(intcode []int) *robot {
	r := new(robot)
	r.pointer = 0
	r.relativeBase = 0
	r.direction = up
	r.position = coord{x: 0, y: 0}
	r.surface = make(map[coord]colour)

	r.surface[coord{x: 0, y: 0}] = white

	code := make([]int, len(intcode)*10)
	copy(code, intcode)
	r.intcode = code

	return r
}

type coord struct {
	x, y int
}

func (c coord) add(new coord) coord {
	return coord{x: c.x + new.x, y: c.y + new.y}
}

type colour int

const (
	black colour = 0
	white colour = 1
)

type direction int

const (
	up    direction = 0
	right direction = 1
	down  direction = 2
	left  direction = 3
)

func (d direction) turn(val int) direction {
	if val == 0 {
		val = 3
	}
	return direction((int(d) + val) % 4)
}

func (d direction) vector() coord {
	if d == up {
		return coord{x: 0, y: 1}
	} else if d == right {
		return coord{x: 1, y: 0}
	} else if d == down {
		return coord{x: 0, y: -1}
	} else {
		return coord{x: -1, y: 0}
	}
}

type surface map[coord]colour

func (s surface) colour(c coord) colour {
	currentColour, ok := s[c]
	if !ok {
		currentColour = black
	}
	return currentColour
}

func (s surface) paint() {
	var minX, maxX, minY, maxY int = 0, 0, 0, 0
	for k := range s {
		minX = min(minX, k.x)
		maxX = max(maxX, k.x)
		minY = min(minY, k.y)
		maxY = max(maxY, k.y)
	}

	output := make([][]string, 1+maxY-minY)

	for i := 0; i < len(output); i++ {
		row := make([]string, 1+maxX-minX)

		for j := 0; j < len(row); j++ {
			coord := coord{x: minX + j, y: maxY - i}
			colour := s.colour(coord)

			l := " "
			if colour == white {
				l = "#"
			}
			row[j] = l
		}
		output[i] = row
		fmt.Println(row)
	}

}

func (a *robot) execute(input int) []int {
	var output []int
	inputConsumed := false
	for a.pointer < len(a.intcode) {
		v := a.intcode[a.pointer]

		mode := v % 100

		if mode == 1 {
			a.intcode[a.index(3)] = a.val(1) + a.val(2)
			a.pointer += 4
		} else if mode == 2 {
			a.intcode[a.index(3)] = a.val(1) * a.val(2)
			a.pointer += 4
		} else if mode == 3 {
			if inputConsumed {
				// Wait to be executed again with fresh input
				break
			} else {
				a.intcode[a.index(1)] = input
				a.pointer += 2
				inputConsumed = true
			}
		} else if mode == 4 {
			output = append(output, a.val(1))
			a.pointer += 2
		} else if mode == 5 {
			if a.val(1) != 0 {
				a.pointer = a.val(2)
			} else {
				a.pointer += 3
			}
		} else if mode == 6 {
			if a.val(1) == 0 {
				a.pointer = a.val(2)
			} else {
				a.pointer += 3
			}
		} else if mode == 7 {
			if a.val(1) < a.val(2) {
				a.intcode[a.index(3)] = 1
			} else {
				a.intcode[a.index(3)] = 0
			}
			a.pointer += 4
		} else if mode == 8 {
			if a.val(1) == a.val(2) {
				a.intcode[a.index(3)] = 1
			} else {
				a.intcode[a.index(3)] = 0
			}
			a.pointer += 4
		} else if mode == 9 {
			a.relativeBase += a.val(1)
			a.pointer += 2
		} else if mode == 99 {
			output = append(output, -1)
			break
		} else {
			fmt.Println("Fatal error")
			fmt.Println(v)
			fmt.Println(a.intcode)
			os.Exit(1)
		}
	}

	return output
}

func (a *robot) val(param int) int {
	v := a.intcode[a.pointer]
	div := 1

	if param == 1 {
		div = 100
	} else if param == 2 {
		div = 1000
	} else if param == 3 {
		div = 10000
	} else if param == 4 {
		div = 100000
	}

	mode := (v / div) % 10
	if mode == 1 {
		// Immediate mode
		return a.intcode[a.pointer+param]
	} else if mode == 2 {
		// Relative mode
		pos := a.intcode[a.pointer+param] + a.relativeBase
		return a.intcode[pos]
	}

	// Position mode
	return a.intcode[a.intcode[a.pointer+param]]
}

func (a *robot) index(param int) int {
	v := a.intcode[a.pointer]
	div := 1

	if param == 1 {
		div = 100
	} else if param == 2 {
		div = 1000
	} else if param == 3 {
		div = 10000
	} else if param == 4 {
		div = 100000
	}

	mode := (v / div) % 10
	if mode == 1 {
		// Immediate mode
		fmt.Println("Not possible to write instructions in immediate mode")
		os.Exit(1)
	} else if mode == 2 {
		// Relative mode
		return a.intcode[a.pointer+param] + a.relativeBase
	}

	// Position mode
	return a.intcode[a.pointer+param]
}
