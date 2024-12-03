package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputParser(t *testing.T) {
	left, right := parseInputFile("test_input.txt")

	assert.Equal(t, []int{3, 4, 2, 1, 3, 3}, left)
	assert.Equal(t, []int{4, 3, 5, 3, 9, 3}, right)
}

func TestPairLocationIds(t *testing.T) {
	pairs := pairLocationIds([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})

	assert.Equal(t, []locationIdPair{{1, 3}, {2, 3}, {3, 3}, {3, 4}, {3, 5}, {4, 9}}, pairs)
}

func TestSumDistance(t *testing.T) {
	sum := sumDistance([]locationIdPair{{1, 3}, {2, 3}, {3, 3}, {3, 4}, {3, 5}, {4, 9}})

	assert.Equal(t, 11, sum)
}

func TestGroup(t *testing.T) {
	groups := group([]int{3, 4, 2, 1, 3, 3})

	assert.Equal(t, 1, groups[1])
	assert.Equal(t, 1, groups[2])
	assert.Equal(t, 3, groups[3])
	assert.Equal(t, 1, groups[4])
}

func TestSimilarityScore(t *testing.T) {
	score := similarityScore([]int{3, 4, 2, 1, 3, 3}, map[int]int{
		3: 3,
		4: 1,
		5: 1,
		9: 1,
	})

	assert.Equal(t, 31, score)
}
