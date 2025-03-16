package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatNumbers(t *testing.T) {
	assert.Equal(t, int64(11), concatNumbers(1, 1))
	assert.Equal(t, int64(1234), concatNumbers(123, 4))
}

func TestEvaluateEquations(t *testing.T) {
	equations := parseInputFile("test_input.txt")

	validEquations := evaluateEquations(equations)

	assert.Equal(t, 3, len(validEquations))
	assert.Equal(t, equation{result: 190, numbers: []int64{10, 19}}, validEquations[0])
	assert.Equal(t, equation{result: 3267, numbers: []int64{81, 40, 27}}, validEquations[1])
	assert.Equal(t, equation{result: 292, numbers: []int64{11, 6, 16, 20}}, validEquations[2])
}

func TestGenerateCombinations(t *testing.T) {
	var combinations []string

	generateCombinations("", 3, "+*", &combinations)

	assert.Equal(t, int(math.Pow(2, 3)), len(combinations))
}

func TestParseInput(t *testing.T) {
	equations := parseInputFile("test_input.txt")

	assert.Equal(t, equation{result: 190, numbers: []int64{10, 19}}, equations[0])
	assert.NotEqual(t, equation{result: 180, numbers: []int64{10, 19}}, equations[0])
	assert.NotEqual(t, equation{result: 190, numbers: []int64{11, 19}}, equations[0])
	assert.Equal(t, equation{result: 83, numbers: []int64{17, 5}}, equations[2])
	assert.Equal(t, equation{result: 292, numbers: []int64{11, 6, 16, 20}}, equations[8])
}
