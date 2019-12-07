package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	// R75,D30,R83,U83,L12,D49,R71,U7,L72
	// U62,R66,U55,R34,D71,R55,D58,R83
	assert := assert.New(t)

	assert.Equal(6, closestIntersect("R8,U5,L5,D3", "U7,R6,D4,L4"))
	assert.Equal(159, closestIntersect("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"))
	assert.Equal(135, closestIntersect("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"))

	assert.Equal(30, shortestIntersect("R8,U5,L5,D3", "U7,R6,D4,L4"))
	assert.Equal(610, shortestIntersect("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"))
	assert.Equal(410, shortestIntersect("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"))
}
