package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInstructionsExtended(t *testing.T) {
	calculations := parseInstructionsExtended("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	assert.Equal(t, [][]int{{2, 4}, {8, 5}}, calculations)
}

func TestParseInstructions(t *testing.T) {
	calculations := parseInstructions("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")

	assert.Equal(t, [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}, calculations)
}

func TestCalculate(t *testing.T) {
	sum := calculate([][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}})

	assert.Equal(t, 161, sum)
}
