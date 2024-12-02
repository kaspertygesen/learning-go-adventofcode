package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputParser(t *testing.T) {
	left, right := parseInputFile("test_input.txt")

	assert.Equal(t, left, []int{3, 4, 2, 1, 3, 3})
	assert.Equal(t, right, []int{4, 3, 5, 3, 9, 3})
}

func TestPairLocationIds(t *testing.T) {
	pairs := pairLocationIds([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})

	assert.Equal(t, pairs, []locationIdPair{{1, 3}, {2, 3}, {3, 3}, {3, 4}, {3, 5}, {4, 9}})
}
