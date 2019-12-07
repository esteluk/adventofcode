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

	executeOpCodeInput(program, 5)
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

func executeOpCode(instructions []int) int {
	return executeOpCodeInput(instructions, 0)
}

func executeOpCodeInput(instructions []int, input int) int {
	var output = make([]int, cap(instructions))
	copy(output, instructions)

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
			output[output[i+1]] = input
			i += 2
		} else if mode == 4 {
			lastExit = val(output, i, 1)
			fmt.Println(lastExit)
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
