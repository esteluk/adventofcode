package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert := assert.New(t)

	s :=
		`.#..#
.....
#####
....#
...##`

	test(assert, s, 3, 4)

	s =
		`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

	test(assert, s, 5, 8)

	s =
		`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

	test(assert, s, 1, 2)

	s =
		`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

	test(assert, s, 6, 3)

	s =
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

	test(assert, s, 11, 13)

}

func test(a *assert.Assertions, s string, x, y int) {

	m := parseAsteroids(s)
	opt := findOptimalObservatory(m)

	a.Equal(x, opt.x)
	a.Equal(y, opt.y)
}

func parseAsteroids(stringMap string) []asteroid {
	reader := strings.NewReader(stringMap)
	return parseAsteroidMap(reader)
}
