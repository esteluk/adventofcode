package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, Fuel(12))
	assert.Equal(2, Fuel(14))
	assert.Equal(654, Fuel(1969))
	assert.Equal(33583, Fuel(100756))
}

func TestRecursiveFuel(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, FuelForFuel(14))
	assert.Equal(966, FuelForFuel(1969))
	assert.Equal(50346, FuelForFuel(100756))
}
