package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputParser(t *testing.T) {
	reports := parseInputFile("test_input.txt")

	assert.Equal(t, []int{7, 6, 4, 2, 1}, reports[0])
	assert.Equal(t, []int{1, 3, 6, 7, 9}, reports[5])
}

func TestEvaluateSafety(t *testing.T) {
	numSafeReports := evaluateSafety([][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	})

	assert.Equal(t, 2, numSafeReports)
}

func TestEvaluateSafetyWithDampener(t *testing.T) {
	numSafeReports := evaluateSafetyWithDampener([][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{67, 67, 69, 72, 75, 78, 81},
		{67, 69, 72, 75, 78, 82},
		{67, 69, 75, 72, 75, 76},
	})

	assert.Equal(t, 7, numSafeReports)
}
