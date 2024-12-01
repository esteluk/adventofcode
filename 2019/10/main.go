package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	aMap := readFile("input")
	fmt.Println(aMap)

	findOptimalObservatory(aMap)

	// testFoo()
}

func testFoo() {

	s :=
		`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

	asteroids := parseAsteroidMap(strings.NewReader(s))

	findOptimalObservatory(asteroids)

}

func findOptimalObservatory(asteroids []asteroid) asteroid {

	var bestCandidate asteroid
	var mostVisible int = 0
	var candidateTargets set

	for _, candidate := range asteroids {
		visible := set{asteroids: make(map[float64]asteroid)}
		// fmt.Println("Evaluating", candidate)
		// For each asteroids, find the count of visible asteroids
		for _, target := range asteroids {
			if target.x == candidate.x && target.y == candidate.y {
				continue
			}
			x := float64(target.x - candidate.x)
			y := float64(target.y - candidate.y)

			angle := math.Atan2(y, x)
			angle += (math.Pi / 2)
			if angle < 0 {
				angle += 2 * math.Pi
			}

			if existing, ok := visible.asteroids[angle]; ok {
				if target.distance(candidate) < existing.distance(candidate) {
					visible.asteroids[angle] = target
				}
			} else {
				visible.asteroids[angle] = target
			}
		}

		// fmt.Println("Can see", len(visible.asteroids), "bodies")
		// fmt.Println(visible.asteroids)

		if len(visible.asteroids) > mostVisible {
			bestCandidate = candidate
			candidateTargets = visible
			mostVisible = len(visible.asteroids)
		}
	}

	fmt.Println("Candidate", bestCandidate, "can see", mostVisible, "asteroids")

	sortedAngles := make([]float64, mostVisible)
	var i = 0
	for k := range candidateTargets.asteroids {
		sortedAngles[i] = k
		i++
	}

	sort.Float64s(sortedAngles)

	for i, el := range sortedAngles {
		fmt.Println("Asteroid", i+1, "to be destroyed is at", candidateTargets.asteroids[el], "bearing", el)
	}

	fmt.Println("200th asteroid is at bearing", sortedAngles[199], "which is asteroid", candidateTargets.asteroids[sortedAngles[199]])

	return bestCandidate
}

type set struct {
	asteroids map[float64]asteroid
}

type asteroid struct {
	x, y int
}

func (a *asteroid) distance(candidate asteroid) float64 {
	return math.Sqrt(math.Pow(float64(candidate.x-a.x), 2) + math.Pow(float64(candidate.y+a.y), 2))
}

func readFile(filename string) []asteroid {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return parseAsteroidMap(file)
}

func parseAsteroidMap(r io.Reader) []asteroid {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var asteroids []asteroid
	var line = 0
	for scanner.Scan() {
		b := scanner.Bytes()

		for i, el := range b {
			if el == 35 {
				asteroids = append(asteroids, asteroid{x: i, y: line})
			}
		}
		line++
	}
	return asteroids
}
