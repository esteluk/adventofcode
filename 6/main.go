package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readFile()
	graph := generateGraph(input)

	fmt.Println(graph.orbitCount())
	fmt.Println(graph.distanceBetween("SAN", "YOU"))

}

func readFile() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var orbits []string

	for scanner.Scan() {
		orbits = append(orbits, scanner.Text())
	}
	return orbits
}

func generateGraph(input []string) orbitalSystem {
	bodies := map[string]mass{
		"COM": makeCom(),
	}
	system := orbitalSystem{bodies: bodies}

	for _, el := range input {
		spl := strings.Split(el, ")")
		system.add(spl[1], spl[0])
	}

	fmt.Println(system.bodies)

	return system
}

// Types

func makeCom() mass {
	return com{}
}

type mass interface {
	Distance(graph map[string]mass) int
	Orbits() string
	IsRoot() bool
}

type com struct{}

type body struct {
	name   string
	orbits string
	// distance int
}

func (b body) Name() string {
	return b.name
}

func (b body) Orbits() string {
	return b.orbits
}

func (b body) IsRoot() bool {
	return false
}

func (b body) Distance(graph map[string]mass) int {
	parent := graph[b.orbits]
	return parent.Distance(graph) + 1
}

func (c com) Orbits() string {
	return ""
}

func (c com) IsRoot() bool {
	return true
}

func (c com) Distance(graph map[string]mass) int {
	return 0
}

type orbitalSystem struct {
	bodies map[string]mass
	// Bodies must be unique on name
}

// func (s orbitalSystem) init() {
// 	// s.bodies = make(map[string]mass)
// 	s.bodies["COM"] = makeCom()
// 	fmt.Println(s.bodies)
// }

func (s orbitalSystem) add(object, orbitting string) {
	new := body{
		orbits: orbitting,
		// distance: about.Distance() + 1,
	}
	s.bodies[object] = new
}

func (s orbitalSystem) orbitCount() int {
	var count int = 0
	for _, v := range s.bodies {
		count += v.Distance(s.bodies)
	}
	return count
}

func (s orbitalSystem) distanceBetween(a, b string) int {
	aPath := s.pathFrom(a)
	bPath := s.pathFrom(b)

	fmt.Println(aPath)
	fmt.Println(bPath)

	for i := 0; i < len(aPath); i++ {
		for j := 0; j < len(bPath); j++ {
			if aPath[i] == bPath[j] {
				return i + j - 2
			}
		}
	}
	return -1
}

func (s orbitalSystem) pathFrom(object string) []string {
	var a mass = s.bodies[object]
	var name string = object
	var path []string

	for !a.IsRoot() {
		path = append(path, name)
		name = a.Orbits()
		a = s.bodies[a.Orbits()]
	}

	return path
}
