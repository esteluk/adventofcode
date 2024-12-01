package main

import (
	"math"
)

func main() {

}

func findPossibleAnswers(lower, upper int) int {
	var count int = 0
	for i := lower; i <= upper; i++ {
		if isValidPassword(i) {
			count++
		}
	}
	return count
}

func isValidPassword(password int) bool {
	return isSixDigits(password) && isIdenticalAdjecentDigits(password) && isDigitsAscending(password)
}

func isSixDigits(password int) bool {
	return password > 99999 && password < 1000000
}

// func isWithinRange(password, lower, upper int) bool {
// 	return password >= lower && password <= upper
// }

func isIdenticalAdjecentDigits(password int) bool {
	for i := 0; i < 5; i++ {
		digit := digitAtPosition(password, i)
		next := digitAtPosition(password, i+1)
		again := digitAtPosition(password, i+2)
		prev := -1
		if i > 0 {
			prev = digitAtPosition(password, i-1)
		}

		if digit == next && digit != again {
			if prev >= 0 && digit != prev {
				return true
			} else if prev < 0 {
				return true
			}
		}
	}
	return false
}

func isDigitsAscending(password int) bool {
	var prev int = -1
	for i := 5; i >= 0; i-- {
		digit := digitAtPosition(password, i)
		if digit < prev {
			return false
		}
		prev = digit
	}
	return true
}

// Zero-indexed position
func digitAtPosition(password, position int) int {
	rem := password / int(math.Pow10(position))
	return rem % 10
}
