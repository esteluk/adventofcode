package main

import (
	"adventofcode/lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(shortestIntersect(readfile()))
}

func readfile() (string, string) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines[0], lines[1]
}

func shortestIntersect(wire1, wire2 string) int {
	path1 := buildPointArray(wire1)
	path2 := buildPointArray(wire2)

	intersections := lib.Intersect(path1, path2)
	var minDist int = -1
	for _, el := range intersections {
		p, _ := el.(point)
		dist := path1.pos(p) + path2.pos(p)
		if minDist == -1 && dist > 0 {
			minDist = dist
		} else {
			minDist = min(minDist, dist)
		}
	}
	return minDist
}

func closestIntersect(wire1, wire2 string) int {
	path1 := buildPointArray(wire1)
	path2 := buildPointArray(wire2)

	intersections := lib.Intersect(path1, path2)

	var minDist int = -1

	for _, el := range intersections {
		p, _ := el.(point)
		if minDist == -1 && distance(p) > 0 {
			minDist = distance(p)
		} else {
			minDist = min(minDist, distance(p))
		}
	}

	return minDist
}

func buildPointArray(wire string) pointSlice {
	var arr pointSlice
	arr = append(arr, point{x: 0, y: 0})
	splits := strings.Split(wire, ",")

	for _, el := range splits {
		for i := 0; i < count(el); i++ {
			x, y := direction(el)
			last := arr[len(arr)-1]
			arr = append(arr, point{x: last.x + x, y: last.y + y})
		}
	}

	return arr
}

func direction(base string) (int, int) {
	if strings.HasPrefix(base, "U") {
		return 0, 1
	} else if strings.HasPrefix(base, "D") {
		return 0, -1
	} else if strings.HasPrefix(base, "L") {
		return -1, 0
	} else if strings.HasPrefix(base, "R") {
		return 1, 0
	}

	os.Exit(1)
	return 0, 0
}

func count(base string) int {
	val, _ := strconv.Atoi(base[1:])
	return val
}

type point struct {
	x, y int
}

func distance(p point) int {
	return abs(p.x) + abs(p.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type pointSlice []point

func (slice pointSlice) pos(val point) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
