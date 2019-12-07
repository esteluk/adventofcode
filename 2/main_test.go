package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 0, 0, 0, 99}, executeOpCode([]int{1, 0, 0, 0, 99}))
	assert.Equal([]int{2, 3, 0, 6, 99}, executeOpCode([]int{2, 3, 0, 3, 99}))
	assert.Equal([]int{2, 4, 4, 5, 99, 9801}, executeOpCode([]int{2, 4, 4, 5, 99, 0}))
	assert.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, executeOpCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}))
}
