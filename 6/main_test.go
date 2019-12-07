package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert := assert.New(t)

	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}

	graph := generateGraph(input)

	assert.Equal(42, graph.orbitCount())
}

func TestOutOfOrder(t *testing.T) {
	assert := assert.New(t)

	input := []string{"COM)B", "C)D", "B)C", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}

	graph := generateGraph(input)

	assert.Equal(42, graph.orbitCount())
}

func TestOrbitalChanges(t *testing.T) {
	assert := assert.New(t)

	input := []string{"COM)B", "C)D", "B)C", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}

	graph := generateGraph(input)

	assert.Equal(4, graph.distanceBetween("YOU", "SAN"))
}
