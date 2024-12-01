package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fuelRequired int = 0

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			fuelRequired += FuelForFuel(i)
		}
	}

	fmt.Println(fuelRequired)
}
