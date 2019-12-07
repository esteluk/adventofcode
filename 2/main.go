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
	fmt.Println(findInitial())
}

func findInitial() (noun, verb int) {
	program := readFile()

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			if output(n, v, program) == 19690720 {
				return n, v
			}
		}
	}
	return -1, -1
}

func output(noun, verb int, program []int) int {
	var execution = make([]int, cap(program))
	copy(execution, program)

	execution[1] = noun
	execution[2] = verb

	output := executeOpCode(execution)
	return output[0]
}

func readFile() []int {
	file, err := os.Open("input.txt")
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

func executeOpCode(instructions []int) []int {
	var output = make([]int, cap(instructions))
	copy(output, instructions)

	var i = 0
	for i < len(instructions) {
		v := output[i]

		if v == 1 {
			output[output[i+3]] = output[output[i+1]] + output[output[i+2]]
		} else if v == 2 {
			output[output[i+3]] = output[output[i+1]] * output[output[i+2]]
		} else if v == 99 {
			break
		} else {
			fmt.Println("Fatal error")
			fmt.Println(output)
			os.Exit(1)
		}
		i += 4
	}

	return output
}
