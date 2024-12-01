package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readFile()
	createImage(25, 6, input)
}

func createImage(width, height int, input string) {
	fmt.Println(input, len(input))

	pixelSize := width * height
	layerCount := len(input) / pixelSize

	var layers [][]int
	for i := 0; i < layerCount; i++ {
		layer := make([]int, pixelSize)
		for j := 0; j < pixelSize; j++ {
			index := i*pixelSize + j
			char := string(input[index])
			val, err := strconv.Atoi(char)
			if err == nil {
				layer[j] = val
			} else {
				fmt.Println(err)
			}
		}
		layers = append(layers, layer)
	}

	computedImage := make([][]int, height)
	for i := range computedImage {
		computedImage[i] = make([]int, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			computedImage[y][x] = pixel(layers, width, x, y)
		}
	}

	for _, l := range computedImage {
		fmt.Println(l)
	}

	/*
		[1 1 1 0 0 1 0 0 1 0 1 0 0 1 0 1 1 1 0 0 1 0 0 0 1]
		[1 0 0 1 0 1 0 1 0 0 1 0 0 1 0 1 0 0 1 0 1 0 0 0 1]
		[1 0 0 1 0 1 1 0 0 0 1 1 1 1 0 1 0 0 1 0 0 1 0 1 0]
		[1 1 1 0 0 1 0 1 0 0 1 0 0 1 0 1 1 1 0 0 0 0 1 0 0]
		[1 0 1 0 0 1 0 1 0 0 1 0 0 1 0 1 0 1 0 0 0 0 1 0 0]
		[1 0 0 1 0 1 0 0 1 0 1 0 0 1 0 1 0 0 1 0 0 0 1 0 0]
	*/

}

func pixel(layers [][]int, width, x, y int) int {
	for _, layer := range layers {
		val := layer[width*y+x]
		if val == 2 {
			continue
		}
		return val
	}
	return -1
}

func readFile() string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()

}
