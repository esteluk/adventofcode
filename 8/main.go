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
	fmt.Println(len(layers[0]))

	var fewestZeroCount int = 99999
	var bestLayerIndex int = -1
	for i, layer := range layers {
		var count = 0
		for _, j := range layer {
			if j == 0 {
				count++
			}
		}
		if count < fewestZeroCount {
			fewestZeroCount = count
			bestLayerIndex = i
		}
	}

	fmt.Println("Best layer at", bestLayerIndex, "with zeros:", fewestZeroCount)

	bestLayer := layers[bestLayerIndex]
	var oneCount, twoCount int = 0, 0
	for _, v := range bestLayer {
		if v == 1 {
			oneCount++
		} else if v == 2 {
			twoCount++
		}
	}

	fmt.Println(oneCount, twoCount)
	fmt.Println("Number of one digits * number of two digits:", oneCount*twoCount)
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
