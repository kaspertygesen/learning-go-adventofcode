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

func TestSumDistance(t *testing.T) {
	sum := sumDistance([]locationIdPair{{1, 3}, {2, 3}, {3, 3}, {3, 4}, {3, 5}, {4, 9}})

	assert.Equal(t, sum, 11)
}

func TestGroup(t *testing.T) {
	groups := group([]int{3, 4, 2, 1, 3, 3})

	assert.Equal(t, groups[1], 1)
	assert.Equal(t, groups[2], 1)
	assert.Equal(t, groups[3], 3)
	assert.Equal(t, groups[4], 1)
}

func TestSimilarityScore(t *testing.T) {
	score := similarityScore([]int{3, 4, 2, 1, 3, 3}, map[int]int{
		3: 3,
		4: 1,
		5: 1,
		9: 1,
	})

	assert.Equal(t, score, 31)
}
