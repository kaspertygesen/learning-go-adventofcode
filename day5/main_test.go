package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumMiddleCorrected(t *testing.T) {
	rules, updates := parseInputFile("test_input.txt")
	correctedUpdates := correctInvalidUpdates(rules, invalidUpdates(rules, updates))
	sum := sumMiddle(correctedUpdates)

	assert.Equal(t, 123, sum)
}

func TestCorrectInvalidUpdates(t *testing.T) {
	rules, updates := parseInputFile("test_input.txt")
	correctedUpdates := correctInvalidUpdates(rules, invalidUpdates(rules, updates))

	assert.Equal(t, []int{97, 75, 47, 61, 53}, correctedUpdates[0])
	assert.Equal(t, []int{61, 29, 13}, correctedUpdates[1])
	assert.Equal(t, []int{97, 75, 47, 29, 13}, correctedUpdates[2])
	assert.Len(t, correctedUpdates, 3)
}

func TestInvalidUpdates(t *testing.T) {
	updates := invalidUpdates(parseInputFile("test_input.txt"))

	assert.Equal(t, []int{75, 97, 47, 61, 53}, updates[0])
	assert.Equal(t, []int{61, 13, 29}, updates[1])
	assert.Equal(t, []int{97, 13, 75, 29, 47}, updates[2])
	assert.Len(t, updates, 3)
}

func TestSumMiddle(t *testing.T) {
	sum := sumMiddle(validUpdates(parseInputFile("test_input.txt")))

	assert.Equal(t, 143, sum)
}

func TestValidUpdates(t *testing.T) {
	updates := validUpdates(parseInputFile("test_input.txt"))

	assert.Equal(t, []int{75, 47, 61, 53, 29}, updates[0])
	assert.Equal(t, []int{97, 61, 53, 29, 13}, updates[1])
	assert.Equal(t, []int{75, 29, 13}, updates[2])
	assert.Len(t, updates, 3)
}

func TestInputParser(t *testing.T) {
	rules, updates := parseInputFile("test_input.txt")

	assert.Equal(t, []int{47, 53}, (*rules)[0])
	assert.Equal(t, []int{53, 13}, (*rules)[len(*rules)-1])

	assert.Equal(t, []int{75, 47, 61, 53, 29}, updates[0])
	assert.Equal(t, []int{97, 13, 75, 29, 47}, updates[len(updates)-1])
}
