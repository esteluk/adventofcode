package main

import (
	"adventofcode/lib"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reactions := parseInput("input")
	fmt.Println(reactions.oreRequired(1))
	fmt.Println(reactions.findFuelProduced(1000000000000))
}

func (r reactions) findFuelProduced(oreQuantity int) int {
	oreForOneFuel := r.oreRequired(1)
	lowerBound := floor(oreQuantity, oreForOneFuel)
	upperBound := lowerBound * 2

	for upperBound-lowerBound > 1 {
		fmt.Println("upperbound", upperBound, "lowerbound", lowerBound)
		newSearch := (upperBound + lowerBound) / 2
		ore := r.oreRequired(newSearch)
		if ore > oreQuantity {
			upperBound = newSearch
		} else if ore < oreQuantity {
			lowerBound = newSearch
		} else {
			return newSearch
		}
		fmt.Println("upperbound", upperBound, "lowerbound", lowerBound)
	}

	return lowerBound
}

func floor(o, d int) int {
	return int(math.Floor(float64(o) / float64(d)))
}

func (r reactions) oreRequired(fuel int) int {

	required := requiredMap{"FUEL": fuel}

	for required.requiredReactions() {
		for k, v := range required {
			if k != "ORE" && v > 0 {

				reaction := r[k]
				multiplier := lib.Max(v/reaction.outputQuantity, 1)
				val := v - (reaction.outputQuantity * multiplier)
				if val == 0 {
					delete(required, k)
				} else {
					required[k] = val
				}

				for _, el := range reaction.inputs {
					required[el.compound] += el.quantity * multiplier
				}
			}
		}
	}

	fmt.Println(required)
	return required["ORE"]

}

type requiredMap map[string]int

func (m requiredMap) requiredReactions() bool {
	for k, v := range m {
		if k != "ORE" && v > 0 {
			return true
		}
	}

	return false
}

type reactions map[string]reaction

type reaction struct {
	outputKey      string
	outputQuantity int
	inputs         []input
}

type input struct {
	compound string
	quantity int
}

func parseCompound(s string) input {
	split := strings.Split(s, " ")
	int, _ := strconv.Atoi(split[0])
	return input{compound: split[1], quantity: int}
}

func parseInput(filename string) reactions {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var r reactions = make(map[string]reaction)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "=>")

		// Parse input
		inputs := strings.Split(split[0], ",")
		iArray := make([]input, len(inputs))

		for i, el := range inputs {
			iArray[i] = parseCompound(strings.TrimSpace(el))
		}

		output := parseCompound(strings.TrimSpace(split[1]))

		reaction := reaction{outputKey: output.compound, outputQuantity: output.quantity, inputs: iArray}
		r[output.compound] = reaction
	}

	return r
}
