package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
