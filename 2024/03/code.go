package main

import (
	"strconv"
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

	if part2 {
		var sum = 0

		combined := strings.Join(parsed, "")
		doSplit := strings.Split(combined, "do()")

		for _, canDo := range doSplit {
			canDoSplit := strings.Split(canDo, "don't()")
			split := strings.Split(canDoSplit[0], "mul(")
			for _, possible := range split {

				index := strings.Index(possible, ")")
				if index == -1 {
					continue
				}
				csv := strings.Split(possible[:index], ",")
				if len(csv) != 2 {
					continue
				}
				firstNum, err := strconv.Atoi(csv[0])
				secondNum, err2 := strconv.Atoi(csv[1])
				if err == nil || err2 == nil {
					// println("Found", strings.Join(csv, ","))

					sum += firstNum * secondNum
				}
			}
		}
		return sum
	} else {
		var sum = 0

		for _, t := range parsed {
			// var testString = t
			// var lastResult = 1
			// for ok := true; ok; ok = lastResult != 0 {
			// 	lastResult, testString = firstMult(testString)
			// 	sum += lastResult
			// }

			split := strings.Split(t, "mul(")
			for _, possible := range split {
				index := strings.Index(possible, ")")
				if index == -1 {
					continue
				}
				csv := strings.Split(possible[:index], ",")
				if len(csv) != 2 {
					continue
				}
				firstNum, err := strconv.Atoi(csv[0])
				secondNum, err2 := strconv.Atoi(csv[1])
				if err == nil || err2 == nil {
					// println("Found", strings.Join(csv, ","))

					sum += firstNum * secondNum
				}
			}
		}
		return sum
	}
}

func firstMult(input string) (int, string) {
	// println("Testing string", input)
	originalString := strings.Clone(input)
	index := strings.Index(input, "mul(")
	if index < 0 {
		return 0, input
	}

	// println("Mul found at index", index)
	// println("Looking for close bracket in", input[index:])
	nextCloseBracketIndex := strings.Index(input[index+4:], ")")
	nextMulIndex := strings.Index(input[index+4:], "mul(")

	// println("Close bracket at index", nextCloseBracketIndex)

	if nextCloseBracketIndex > 11 || (nextMulIndex >= 0 && nextMulIndex < nextCloseBracketIndex) {
		// println("Next string", originalString[index+4:])
		return firstMult(originalString[index+4:])
	}

	if nextCloseBracketIndex > 0 && nextCloseBracketIndex < 11 {
		// Possibly a valid instruction in this range
		// println(input)
		// println(index)
		// println(nextCloseBracketIndex)
		split := strings.Split(originalString[index+4:index+4+nextCloseBracketIndex], ",")
		// println(strings.Join(split, ","))
		// println("Original string", originalString)

		if len(split) != 2 {
			return firstMult(originalString[index+4+nextCloseBracketIndex+4:])
		}
		firstNum, err := strconv.Atoi(split[0])
		secondNum, err2 := strconv.Atoi(split[1])
		if err == nil || err2 == nil {
			println("Found", strings.Join(split, ","))
			// Probably a valid command
			nextString := originalString[index+nextCloseBracketIndex+4:]
			// println(nextString)

			return firstNum * secondNum, nextString
		}
	}

	nextString := originalString[index+1:]
	// println("Next string", nextString)

	return firstMult(nextString)
}
