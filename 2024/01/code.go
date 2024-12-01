package main

import (
	"fmt"
	"slices"
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		parsed := strings.Split(strings.TrimSpace(input), "\n")
		length := len(parsed)
		arr1 := make([]int, length)
		map2 := make(map[int]int)
		for i, t := range parsed {
			content := strings.Split(t, "   ")
			arr1[i], _ = strconv.Atoi(content[0])
			b, _ := strconv.Atoi(content[1])
			count, ok := map2[b]
			if ok {
				map2[b] = count + 1
			} else {
				map2[b] = 1
			}
		}

		var similarity int
		for _, t := range arr1 {

			count, ok := map2[t]
			if ok {
				similarity += count * t
			}
		}

		return similarity
	} else {
		parsed := strings.Split(strings.TrimSpace(input), "\n")
		length := len(parsed)
		arr1 := make([]int, length)
		arr2 := make([]int, length)
		for i, t := range parsed {
			content := strings.Split(t, "   ")
			arr1[i], _ = strconv.Atoi(content[0])
			arr2[i], _ = strconv.Atoi(content[1])
		}
		slices.Sort(arr1)
		slices.Sort(arr2)
		zipped, _ := zip(arr1, arr2)
		var sum int
		for _, t := range zipped {
			sum += t.distance()
		}
		return sum
	}
}

type intTuple struct {
	a, b int
}

func (t *intTuple) distance() int {
	size := t.b - t.a
	if size > 0 {
		return size
	} else {
		return -size
	}
}

func zip(a, b []int) ([]intTuple, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([]intTuple, len(a), len(a))

	for i, e := range a {
		r[i] = intTuple{e, b[i]}
	}

	return r, nil
}
