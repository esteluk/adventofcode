package main

import (
	"adventofcode/lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	run()
}

func run() {
	program := readFile()

	thrust := findMaxThrust(program, true)
	fmt.Println(thrust)
}

func findMaxThrust(program []int, feedback bool) int {
	var maxThrust = 0
	var optimumPhaseSettings []int
	for _, el := range findAllPermutations(feedback) {
		thrust := calculateThrust(program, el, feedback)
		if thrust > maxThrust {
			maxThrust = thrust
			optimumPhaseSettings = el
		}
	}

	fmt.Println(optimumPhaseSettings)

	return maxThrust
}

func calculateThrust(program []int, phaseSettings []int, feedback bool) int {
	var amps = [5]amplifier{}

	for i, s := range phaseSettings {
		var p = make([]int, cap(program))
		copy(p, program)
		amps[i] = amplifier{
			phaseSetting:  s,
			intcode:       p,
			pointer:       0,
			isInitialised: false,
		}
	}

	var input int = 0
	if feedback {
		var complete bool = false
		for !complete {
			for i := 0; i < len(amps); i++ {
				output := amps[i].execute(input)
				if output > -1 {
					input = output
				} else {
					complete = true
				}
			}
		}
	} else {
		for i := 0; i < len(amps); i++ {
			input = amps[i].execute(input)
		}
	}

	return input
}

var permutations [][]int

// Heap's algorithm

func findAllPermutations(feedback bool) [][]int {
	permutations = nil
	settings := []int{0, 1, 2, 3, 4}

	if feedback {
		settings = []int{5, 6, 7, 8, 9}
	}
	generate(len(settings), settings)
	return permutations
}

func generate(k int, array []int) {
	if k == 1 {
		var c = make([]int, 5)
		copy(c, array)
		permutations = append(permutations, c)
	} else {
		generate(k-1, array)

		for i := 0; i < k-1; i++ {
			// generate(k-1, array)
			if k%2 == 0 {
				array[i], array[k-1] = array[k-1], array[i]
			} else {
				array[0], array[k-1] = array[k-1], array[0]
			}
			generate(k-1, array)
		}
	}
}

type amplifier struct {
	phaseSetting  int
	intcode       []int
	pointer       int
	isInitialised bool
}

func (a *amplifier) execute(input int) int {
	for a.pointer < len(a.intcode) {
		v := a.intcode[a.pointer]

		mode := v % 100

		if mode == 1 {
			a.intcode[a.intcode[a.pointer+3]] = val(a.intcode, a.pointer, 1) + val(a.intcode, a.pointer, 2)
			a.pointer += 4
		} else if mode == 2 {
			a.intcode[a.intcode[a.pointer+3]] = val(a.intcode, a.pointer, 1) * val(a.intcode, a.pointer, 2)
			a.pointer += 4
		} else if mode == 3 {
			if a.isInitialised {
				a.intcode[a.intcode[a.pointer+1]] = input
			} else {
				a.intcode[a.intcode[a.pointer+1]] = a.phaseSetting
				a.isInitialised = true
			}

			a.pointer += 2
		} else if mode == 4 {
			exit := val(a.intcode, a.pointer, 1)
			a.pointer += 2
			return exit
		} else if mode == 5 {
			if val(a.intcode, a.pointer, 1) != 0 {
				a.pointer = val(a.intcode, a.pointer, 2)
			} else {
				a.pointer += 3
			}
		} else if mode == 6 {
			if val(a.intcode, a.pointer, 1) == 0 {
				a.pointer = val(a.intcode, a.pointer, 2)
			} else {
				a.pointer += 3
			}
		} else if mode == 7 {
			if val(a.intcode, a.pointer, 1) < val(a.intcode, a.pointer, 2) {
				a.intcode[a.intcode[a.pointer+3]] = 1
			} else {
				a.intcode[a.intcode[a.pointer+3]] = 0
			}
			a.pointer += 4
		} else if mode == 8 {
			if val(a.intcode, a.pointer, 1) == val(a.intcode, a.pointer, 2) {
				a.intcode[a.intcode[a.pointer+3]] = 1
			} else {
				a.intcode[a.intcode[a.pointer+3]] = 0
			}
			a.pointer += 4
		} else if mode == 99 {
			return -1
		} else {
			fmt.Println("Fatal error")
			fmt.Println(v)
			fmt.Println(a.intcode)
			os.Exit(1)
		}
	}

	fmt.Println("Default exit")
	return -1
}

func readFile() []int {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(lib.ScanCommaSeparated)

	var opcodes []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			opcodes = append(opcodes, i)
		}
	}
	return opcodes
}

func val(arr []int, pos, param int) int {

	v := arr[pos]
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

	if (v/div)%10 == 1 {
		return arr[pos+param]
	}
	return arr[arr[pos+param]]
}
