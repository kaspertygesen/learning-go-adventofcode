package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	input := parseInputFile("test_input.txt")

	groups := group(input)

	assert.Equal(t, 1930, sum(groups))
}

func TestGroups(t *testing.T) {
	input := parseInputFile("test_input.txt")

	groups := group(input)

	assert.Len(t, groups, 11)
}

func TestParseInput(t *testing.T) {
	input := parseInputFile("test_input.txt")

	assert.Equal(t, 'R', (*input)[0][0].plant)
	assert.Equal(t, 'E', (*input)[9][9].plant)
	assert.Len(t, *input, 10)
}
