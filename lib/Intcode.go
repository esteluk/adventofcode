package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadIntcode reads the given file and returns an array of intcode instructions
func ReadIntcode(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(ScanCommaSeparated)

	var opcodes []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			opcodes = append(opcodes, i)
		}
	}
	return opcodes
}
