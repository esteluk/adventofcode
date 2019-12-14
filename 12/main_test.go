package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeUntilInitialStateOne(t *testing.T) {
	assert := assert.New(t)

	system := readMoons("test_input_1")

	assert.Equal(2772, system.findPeriod())
}

func TestTimeUntilInitialStateTwo(t *testing.T) {
	assert := assert.New(t)

	system := readMoons("test_input_2")

	assert.Equal(4686774924, system.findPeriod())
}

func TestTimeUntilInitialStateResult(t *testing.T) {
	assert := assert.New(t)

	system := readMoons("input")

	assert.Equal(326489627728984, system.findPeriod())
}
