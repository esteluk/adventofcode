package main

import (
	"adventofcode/lib"
	"fmt"
	"os"
)

func main() {
	run()
}

func run() {
	input := lib.ReadIntcode("input")
	c := makeComputer(input)
	o := c.execute(2)
	fmt.Println(o)
}

type computer struct {
	pointer      int
	intcode      []int
	relativeBase int
}

func makeComputer(intcode []int) *computer {
	c := new(computer)
	c.pointer = 0
	c.relativeBase = 0

	code := make([]int, len(intcode)*10)
	copy(code, intcode)
	c.intcode = code

	return c
}

func (a *computer) execute(input int) []int {
	var output []int
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
			a.intcode[a.index(1)] = input
			a.pointer += 2
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

func (a *computer) val(param int) int {
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

func (a *computer) index(param int) int {
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
