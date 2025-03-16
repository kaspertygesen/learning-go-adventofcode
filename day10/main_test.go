package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPaths(t *testing.T) {
	input := parseInputFile("test_input.txt")

	count := findPaths(input, false)

	assert.Equal(t, 81, count)
}

func TestFindPathsDistinct(t *testing.T) {
	input := parseInputFile("test_input.txt")

	count := findPaths(input, true)

	assert.Equal(t, 36, count)
}

func TestEvaluateTrailhead(t *testing.T) {
	input := parseInputFile("test_input.txt")

	tops := evaluateTrailhead(input, 2, 0, true)

	assert.Contains(t, tops, coordinate{4, 3})
	assert.Contains(t, tops, coordinate{0, 3})
	assert.Contains(t, tops, coordinate{5, 4})
	assert.Contains(t, tops, coordinate{4, 5})
	assert.Contains(t, tops, coordinate{1, 0})
	assert.Len(t, tops, 5)
}

func TestParseInput(t *testing.T) {
	input := parseInputFile("test_input.txt")

	assert.Equal(t, 8, input[0][0])
	assert.Equal(t, 1, input[5][3])
	assert.Equal(t, 2, input[7][7])
}
