package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	input := parseInputFile("test_input.txt")

	count := play(input)

	assert.Equal(t, 480, count)
}

func TestParseInput(t *testing.T) {
	input := parseInputFile("test_input.txt")

	assert.Equal(t, 94, input[0].a.x)
	assert.Equal(t, 34, input[0].a.y)

	assert.Equal(t, 22, input[0].b.x)
	assert.Equal(t, 67, input[0].b.y)

	assert.Equal(t, 8400, input[0].x)
	assert.Equal(t, 5400, input[0].y)

	assert.Equal(t, 69, input[3].a.x)
	assert.Equal(t, 23, input[3].a.y)

	assert.Equal(t, 27, input[3].b.x)
	assert.Equal(t, 71, input[3].b.y)

	assert.Equal(t, 18641, input[3].x)
	assert.Equal(t, 10279, input[3].y)
}
