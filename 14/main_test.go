package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)

	r := parseInput("input_example1")
	assert.Equal(6, len(r))
	assert.Equal(31, r.oreRequired(1))
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)

	r := parseInput("input_example2")

	assert.Equal(13312, r.oreRequired(1))
	assert.Equal(82892753, r.findFuelProduced(1000000000000))
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)

	r := parseInput("example3")

	assert.Equal(180697, r.oreRequired(1))
	assert.Equal(5586022, r.findFuelProduced(1000000000000))
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)

	r := parseInput("example4")

	assert.Equal(2210736, r.oreRequired(1))
	assert.Equal(460664, r.findFuelProduced(1000000000000))
}
