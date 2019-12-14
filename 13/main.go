package main

import (
	"adventofcode/lib"
	"fmt"
	"time"
)

func main() {
	input := lib.ReadIntcode("input")

	arcade := makeArcade(input)

	arcade.play()

}

type tile int

const (
	empty  tile = 0
	wall   tile = 1
	block  tile = 2
	paddle tile = 3
	ball   tile = 4
)

func (t tile) print() string {
	if t == empty {
		return " "
	} else if t == wall {
		return "█"
	} else if t == block {
		return "#"
	} else if t == paddle {
		return "-"
	} else if t == ball {
		return "o"
	} else {
		return " "
	}
}

type surface map[lib.Coord]tile

func (s surface) tile(c lib.Coord) tile {
	currentTile, ok := s[c]
	if !ok {
		currentTile = empty
	}
	return currentTile
}

func (s surface) paddleLocation() lib.Coord {
	for k, v := range s {
		if v == paddle {
			return k
		}
	}

	return lib.Coord{X: -1, Y: -1}
}

func (s surface) ballLocation() lib.Coord {
	for k, v := range s {
		if v == ball {
			return k
		}
	}

	return lib.Coord{X: -1, Y: -1}
}

func (s surface) draw() {
	lib.CallClear()
	var minX, maxX, minY, maxY int = 0, 0, 0, 0
	for k := range s {
		minX = lib.Min(minX, k.X)
		maxX = lib.Max(maxX, k.X)
		minY = lib.Min(minY, k.Y)
		maxY = lib.Max(maxY, k.Y)
	}

	output := make([][]string, 1+maxY-minY)

	for i := 0; i < len(output); i++ {
		row := make([]string, 1+maxX-minX)

		for j := 0; j < len(row); j++ {
			coord := lib.Coord{X: minX + j, Y: maxY - i}

			tile := s.tile(coord)

			row[j] = tile.print()
		}
		output[i] = row
		fmt.Println(row)
	}
}

type arcade struct {
	computer *lib.Computer

	surface surface
	score   int
}

func makeArcade(intcode []int) *arcade {
	a := new(arcade)

	// Free play
	intcode[0] = 2

	a.computer = lib.MakeComputer(intcode)
	a.surface = make(map[lib.Coord]tile)

	return a
}

func (a *arcade) play() {
	complete := false
	input := 0
	for complete == false {
		a.run(input)
		a.surface.draw()
		fmt.Println("")
		fmt.Println(a.score)

		ballLoc := a.surface.ballLocation()
		paddleLoc := a.surface.paddleLocation()

		if ballLoc.X < paddleLoc.X {
			input = -1
		} else if ballLoc.X > paddleLoc.X {
			input = 1
		} else {
			input = 0
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (a *arcade) run(input int) bool {
	output, complete := a.computer.Execute(input)

	for i := 0; i < len(output)-2; i += 3 {

		if output[i] == -1 && output[i+1] == 0 {
			a.score = output[i+2]
		} else {
			c := lib.Coord{X: output[i], Y: output[i+1]}
			a.surface[c] = tile(output[i+2])
		}
	}

	return complete
}
