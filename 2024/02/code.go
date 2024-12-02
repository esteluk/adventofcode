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
	// when you're ready to do part 2, remove this "not implemented" block
	parsed := strings.Split(strings.TrimSpace(input), "\n")
	var safeCount = 0

	if part2 {
	p2Outer:
		for _, t := range parsed {
			content := strings.Split(t, " ")
			var firstDiff = 0
		removeIndexLoop:
			for removed := -1; removed < len(content); removed++ {
				var newSlice []string
				if removed == -1 {
					newSlice = content
				} else {
					// println("Removing item at index ", removed)
					newSlice = RemoveIndex(content, removed)
					// println("Testing slice ", strings.Join(newSlice, " "))
				}
				for j := 1; j < len(newSlice); j++ {
					this, _ := strconv.Atoi(newSlice[j])
					prev, _ := strconv.Atoi(newSlice[j-1])
					diff := this - prev

					if diff == 0 || diff > 3 || diff < -3 {
						// println("Discarding ", strings.Join(newSlice, " "))
						continue removeIndexLoop
					}

					if j == 1 {
						firstDiff = diff
					} else {
						// We want to make sure that the diff has the same sign everywhere
						if (firstDiff > 0 && diff < 0) || (firstDiff < 0 && diff > 0) {
							// Try removing a different thing
							// println("Discarding ", strings.Join(newSlice, " "))
							continue removeIndexLoop
						}
					}

					if j == len(newSlice)-1 {
						// println(strings.Join(newSlice, " "))
						safeCount++
						continue p2Outer
					}
				}
			}

		}
		return safeCount
	} else {

	outer:
		for _, t := range parsed {
			content := strings.Split(t, " ")
			var firstDiff = 0
			for j := 1; j < len(content); j++ {
				this, _ := strconv.Atoi(content[j])
				prev, _ := strconv.Atoi(content[j-1])
				diff := this - prev

				if diff == 0 || diff > 3 || diff < -3 {
					continue outer
				}

				if j == 1 {
					firstDiff = diff
				} else {
					// We want to make sure that the diff has the same sign everywhere
					if (firstDiff > 0 && diff < 0) || (firstDiff < 0 && diff > 0) {
						continue outer
					}
				}

				if j == len(content)-1 {
					safeCount++
				}
			}
		}
		return safeCount
	}
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
