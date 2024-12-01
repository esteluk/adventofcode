package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	// assert := assert.New(t)

	// assert.Equal([]int{2, 0, 0, 0, 99}, executeOpCode([]int{1, 0, 0, 0, 99}))
	// assert.Equal([]int{2, 3, 0, 6, 99}, executeOpCode([]int{2, 3, 0, 3, 99}))
	// assert.Equal([]int{2, 4, 4, 5, 99, 9801}, executeOpCode([]int{2, 4, 4, 5, 99, 0}))
	// assert.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, executeOpCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}))
}

func TestExtendedExamples(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, executeOpCodeInput([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 8))
	assert.Equal(0, executeOpCodeInput([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 7))

	assert.Equal(1, executeOpCodeInput([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 7))
	assert.Equal(0, executeOpCodeInput([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 9))

	assert.Equal(1, executeOpCodeInput([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 8))
	assert.Equal(0, executeOpCodeInput([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 7))

	assert.Equal(1, executeOpCodeInput([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 7))
	assert.Equal(0, executeOpCodeInput([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 9))

	assert.Equal(1, executeOpCodeInput([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 17))
	assert.Equal(0, executeOpCodeInput([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 0))

	assert.Equal(1, executeOpCodeInput([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 17))
	assert.Equal(0, executeOpCodeInput([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 0))

	assert.Equal(999, executeOpCodeInput([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 7))
	assert.Equal(1000, executeOpCodeInput([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 8))
	assert.Equal(1001, executeOpCodeInput([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 9))
}

func BenchmarkProgram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run()
	}
}
