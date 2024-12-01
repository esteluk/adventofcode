package main

import (
	"adventofcode/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert := assert.New(t)
	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	c := makeComputer(program)
	output := c.execute(0)

	assert.Equal(program, output)
}

func TestSixteenDigit(t *testing.T) {
	assert := assert.New(t)
	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}

	c := makeComputer(program)
	output := c.execute(0)

	assert.Equal([]int{1219070632396864}, output)
}

func TestLargeNumber(t *testing.T) {
	assert := assert.New(t)
	program := []int{104, 1125899906842624, 99}

	c := makeComputer(program)
	output := c.execute(0)

	assert.Equal([]int{1125899906842624}, output)
}

func TestResult(t *testing.T) {
	assert := assert.New(t)
	input := lib.ReadIntcode("input")
	c := makeComputer(input)
	o := c.execute(1)

	assert.Equal([]int{4080871669}, o)
}

func TestResultPartTwo(t *testing.T) {
	assert := assert.New(t)
	input := lib.ReadIntcode("input")
	c := makeComputer(input)
	o := c.execute(2)

	assert.Equal([]int{75202}, o)
}
