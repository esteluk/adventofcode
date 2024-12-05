package main

import (
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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var correctOrderSum = 0
	var incorrectlyOrdered [][]string
	var rules []rule

	for _, t := range lines {
		b := strings.Split(t, "|")
		if len(b) == 2 {
			newRule := rule{before: b[0], after: b[1]}
			rules = append(rules, newRule)
		}

		// This wasn't a rule, so process this as an update
		pages := strings.Split(t, ",")

		if validate(pages, rules) {
			middlePage := (len(pages) - 1) / 2
			midPageValue, _ := strconv.Atoi(pages[middlePage])

			correctOrderSum += midPageValue
		} else {
			incorrectlyOrdered = append(incorrectlyOrdered, pages)
		}
	}

	if part2 {
		var incorrectlyOrderSum = 0
		for _, t := range incorrectlyOrdered {
			slices.SortFunc(t, func(a, b string) int {
				index := slices.IndexFunc(rules, func(c rule) bool {
					return (c.before == a && c.after == b) || (c.before == b && c.after == a)
				})
				rule := rules[index]
				if rule.before == a {
					return -1
				} else {
					return 1
				}

			})
			middlePage := (len(t) - 1) / 2
			midPageValue, _ := strconv.Atoi(t[middlePage])
			incorrectlyOrderSum += midPageValue
		}

		return incorrectlyOrderSum
	} else {
		return correctOrderSum
	}
}

func validate(pages []string, rules []rule) bool {
	for _, rule := range rules {
		befIndex := slices.Index(pages, rule.before)
		aftIndex := slices.Index(pages, rule.after)
		if befIndex > -1 && aftIndex > -1 {
			if befIndex > aftIndex {
				return false
			}
		}
	}
	return true
}

type rule struct {
	before string
	after  string
}
