package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoopCount(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	count := countLoops(positions)

	assert.Equal(t, 6, count)
}

func TestContainsLoop97(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	(*positions)[9][7].symbol = '#'

	isLoop := containsLoop(positions)

	assert.True(t, isLoop)
}

func TestContainsLoop81(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	(*positions)[8][1].symbol = '#'

	isLoop := containsLoop(positions)

	assert.True(t, isLoop)
}

func TestContainsLoop77(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	(*positions)[7][7].symbol = '#'

	isLoop := containsLoop(positions)

	assert.True(t, isLoop)
}

func TestContainsLoop76(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	(*positions)[7][6].symbol = '#'

	isLoop := containsLoop(positions)

	assert.True(t, isLoop)
}

func TestContainsLoop63(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	(*positions)[6][3].symbol = '#'

	isLoop := containsLoop(positions)

	assert.True(t, isLoop)
}

func TestPatrolAlgorithm2(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input_2.txt"))

	patrol(positions)

	assert.True(t, (*positions)[1][4].visited)
	assert.False(t, (*positions)[1][5].visited)
	assert.True(t, (*positions)[9][4].visited)
}

func TestPatrolAlgorithm(t *testing.T) {
	positions := linkPositions(parseInputFile("test_input.txt"))

	patrol(positions)

	assert.True(t, (*positions)[1][4].visited)
	assert.True(t, (*positions)[1][5].visited)
	assert.True(t, (*positions)[1][8].visited)
	assert.True(t, (*positions)[3][8].visited)
	assert.True(t, (*positions)[4][6].visited)
	assert.True(t, (*positions)[9][7].visited)
}

func TestGraph(t *testing.T) {
	labMap := parseInputFile("test_input.txt")
	positions := linkPositions(labMap)

	assert.Equal(t, '#', (*positions)[0][4].symbol)
	assert.Equal(t, '.', (*positions)[0][4].down.symbol)
	assert.Equal(t, '.', (*positions)[5][4].symbol)
	assert.Equal(t, '^', (*positions)[5][4].down.symbol)
	assert.Equal(t, (*positions)[5][4], *(*positions)[5][4].down.up)
}
