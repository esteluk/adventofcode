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

	thrust := findMaxThrust(program)
	fmt.Println(thrust)
}

func findMaxThrust(program []int) int {
	var maxThrust = 0
	var optimumPhaseSettings []int
	for _, el := range findAllPermutations() {
		thrust := calculateThrust(program, el)
		if thrust > maxThrust {
			maxThrust = thrust
			optimumPhaseSettings = el
		}
	}

	fmt.Println(optimumPhaseSettings)

	return maxThrust
}

func calculateThrust(program []int, phaseSettings []int) int {
	var amps = [5]amplifier{}

	for i, s := range phaseSettings {
		amps[i] = amplifier{phaseSetting: s}
	}

	var input int = 0
	for i := 0; i < len(amps); i++ {
		input = amps[i].execute(program, input)
	}

	return input
}

var permutations [][]int

// Heap's algorithm

func findAllPermutations() [][]int {
	permutations = nil
	settings := []int{0, 1, 2, 3, 4}
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
	phaseSetting int
}

func (a amplifier) execute(program []int, input int) int {
	return executeOpCodeInput(program, []int{a.phaseSetting, input})
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

func executeOpCodeInput(instructions []int, input []int) int {
	var output = make([]int, cap(instructions))
	copy(output, instructions)

	var inputCount int = 0

	lastExit := 0
	var i = 0
	for i < len(instructions) {
		v := output[i]

		mode := v % 100

		if mode == 1 {
			output[output[i+3]] = val(output, i, 1) + val(output, i, 2)
			i += 4
		} else if mode == 2 {
			output[output[i+3]] = val(output, i, 1) * val(output, i, 2)
			i += 4
		} else if mode == 3 {
			output[output[i+1]] = input[inputCount]
			inputCount++
			i += 2
		} else if mode == 4 {
			lastExit = val(output, i, 1)
			i += 2
		} else if mode == 5 {
			if val(output, i, 1) != 0 {
				i = val(output, i, 2)
			} else {
				i += 3
			}
		} else if mode == 6 {
			if val(output, i, 1) == 0 {
				i = val(output, i, 2)
			} else {
				i += 3
			}
		} else if mode == 7 {
			if val(output, i, 1) < val(output, i, 2) {
				output[output[i+3]] = 1
			} else {
				output[output[i+3]] = 0
			}
			i += 4
		} else if mode == 8 {
			if val(output, i, 1) == val(output, i, 2) {
				output[output[i+3]] = 1
			} else {
				output[output[i+3]] = 0
			}
			i += 4
		} else if mode == 99 {
			break
		} else {
			fmt.Println("Fatal error")
			fmt.Println(v)
			fmt.Println(output)
			os.Exit(1)
		}
	}

	return lastExit
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
