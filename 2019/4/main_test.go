package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert := assert.New(t)

	assert.False(isValidPassword(111111))
	assert.False(isValidPassword(223450))
	assert.False(isValidPassword(123789))

	assert.True(isValidPassword(112233))
	assert.False(isValidPassword(123444))
	assert.True(isValidPassword(111122))
	assert.False(isValidPassword(666888))

	assert.False(isValidPassword(666668))
}

func BenchmarkProgram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPossibleAnswers(245182, 790572)
	}
}
