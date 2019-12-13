package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	system := readMoons("input")
	system.printDesc()

	for i := 0; i < 1000; i++ {

		fmt.Println("After", i+1, "steps")
		system = system.timeStep()
		system.printDesc()
		fmt.Println("")
	}

	fmt.Println("Total energy in system is", system.totalEnergy())
}

func readMoons(filename string) system {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result system

	for scanner.Scan() {
		result = append(result, parseMoon(scanner.Text()))
	}
	return result
}

type system []moon

func (s system) printDesc() {
	for _, el := range s {
		fmt.Println(el.desc())
	}
}

func (s system) totalEnergy() int {
	totalEnergy := 0

	for _, el := range s {
		potential := abs(el.position.x) + abs(el.position.y) + abs(el.position.z)
		kinetic := abs(el.velocity.x) + abs(el.velocity.y) + abs(el.velocity.z)

		totalEnergy += potential * kinetic
	}
	return totalEnergy
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (s system) timeStep() system {
	var newSystem system

	for _, el := range s {
		// Update velocities
		velX := el.velocity.x
		velY := el.velocity.y
		velZ := el.velocity.z

		for _, m := range s {
			if el == m {
				continue
			}
			if m.position.x > el.position.x {
				velX++
			} else if m.position.x < el.position.x {
				velX--
			}
			if m.position.y > el.position.y {
				velY++
			} else if m.position.y < el.position.y {
				velY--
			}
			if m.position.z > el.position.z {
				velZ++
			} else if m.position.z < el.position.z {
				velZ--
			}
		}

		moon := moon{
			position: vector{
				x: el.position.x + velX,
				y: el.position.y + velY,
				z: el.position.z + velZ,
			},
			velocity: vector{
				x: velX,
				y: velY,
				z: velZ,
			},
		}

		newSystem = append(newSystem, moon)
	}

	return newSystem
}

type moon struct {
	position vector
	velocity vector
}

func (m moon) desc() string {
	return fmt.Sprintf("pos=%s, vel=%s", m.position.desc(), m.velocity.desc())
}

func parseMoon(input string) moon {
	xR := regexp.MustCompile(`(?:x=)(-?\d+)`)
	yR := regexp.MustCompile(`(?:y=)(-?\d+)`)
	zR := regexp.MustCompile(`(?:z=)(-?\d+)`)

	x, _ := strconv.Atoi(xR.FindStringSubmatch(input)[1])
	y, _ := strconv.Atoi(yR.FindStringSubmatch(input)[1])
	z, _ := strconv.Atoi(zR.FindStringSubmatch(input)[1])

	pos := vector{
		x: x,
		y: y,
		z: z,
	}

	vel := vector{x: 0, y: 0, z: 0}

	return moon{position: pos, velocity: vel}
}

type vector struct {
	x, y, z int
}

func (v vector) desc() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", v.x, v.y, v.z)
}
