package lib

import (
	"fmt"
	"os"
)

// Coord represents a location in 2D space
type Coord struct {
	X, Y int
}

// Computer contains the running state of a executor of intcode
type Computer struct {
	Intcode      []int
	Pointer      int
	RelativeBase int
}

// MakeComputer returns a computer configured to run the given bitcode
func MakeComputer(intcode []int) *Computer {
	a := new(Computer)
	a.Pointer = 0
	a.RelativeBase = 0

	code := make([]int, len(intcode)*10)
	copy(code, intcode)
	a.Intcode = code

	return a
}

// Val Grabs the value of a given parameter at the computer's current pointer
func (a *Computer) Val(param int) int {
	v := a.Intcode[a.Pointer]
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
		return a.Intcode[a.Pointer+param]
	} else if mode == 2 {
		// Relative mode
		pos := a.Intcode[a.Pointer+param] + a.RelativeBase
		return a.Intcode[pos]
	}

	// Position mode
	return a.Intcode[a.Intcode[a.Pointer+param]]
}

// Index Grabs the index of a given parameter at the computer's current pointer
func (a *Computer) Index(param int) int {
	v := a.Intcode[a.Pointer]
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
		return a.Intcode[a.Pointer+param] + a.RelativeBase
	}

	// Position mode
	return a.Intcode[a.Pointer+param]
}

// Execute runs the computer until input is next required, returning the output until that point
func (a *Computer) Execute(input int) ([]int, bool) {
	var output []int
	inputConsumed := false
	executionComplete := false
	for a.Pointer < len(a.Intcode) {
		v := a.Intcode[a.Pointer]

		mode := v % 100

		if mode == 1 {
			a.Intcode[a.Index(3)] = a.Val(1) + a.Val(2)
			a.Pointer += 4
		} else if mode == 2 {
			a.Intcode[a.Index(3)] = a.Val(1) * a.Val(2)
			a.Pointer += 4
		} else if mode == 3 {
			if inputConsumed {
				// Wait to be executed again with fresh input
				break
			} else {
				a.Intcode[a.Index(1)] = input
				a.Pointer += 2
				inputConsumed = true
			}
		} else if mode == 4 {
			output = append(output, a.Val(1))
			a.Pointer += 2
		} else if mode == 5 {
			if a.Val(1) != 0 {
				a.Pointer = a.Val(2)
			} else {
				a.Pointer += 3
			}
		} else if mode == 6 {
			if a.Val(1) == 0 {
				a.Pointer = a.Val(2)
			} else {
				a.Pointer += 3
			}
		} else if mode == 7 {
			if a.Val(1) < a.Val(2) {
				a.Intcode[a.Index(3)] = 1
			} else {
				a.Intcode[a.Index(3)] = 0
			}
			a.Pointer += 4
		} else if mode == 8 {
			if a.Val(1) == a.Val(2) {
				a.Intcode[a.Index(3)] = 1
			} else {
				a.Intcode[a.Index(3)] = 0
			}
			a.Pointer += 4
		} else if mode == 9 {
			a.RelativeBase += a.Val(1)
			a.Pointer += 2
		} else if mode == 99 {
			executionComplete = true
			break
		} else {
			fmt.Println("Fatal error")
			fmt.Println(v)
			fmt.Println(a.Intcode)
			os.Exit(1)
		}
	}

	return output, executionComplete
}
