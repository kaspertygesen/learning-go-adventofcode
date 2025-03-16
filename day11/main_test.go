package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBlinkRecursive(t *testing.T) {
	blinks := 6

	cache := make(map[pair]int)

	sum := 0

	for _, s := range []int{125, 17} {
		sum += blinkRecursive(s, blinks, cache)
	}

	assert.Equal(t, 22, sum)
}

func TestBlink6(t *testing.T) {
	stones := blink([]int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32})

	assert.Equal(t, []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}, stones)
}

func TestBlink5(t *testing.T) {
	stones := blink([]int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032})

	assert.Equal(t, []int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}, stones)
}

func TestBlink4(t *testing.T) {
	stones := blink([]int{512072, 1, 20, 24, 28676032})

	assert.Equal(t, []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}, stones)
}

func TestBlink3(t *testing.T) {
	stones := blink([]int{253, 0, 2024, 14168})

	assert.Equal(t, []int{512072, 1, 20, 24, 28676032}, stones)
}

func TestBlink2(t *testing.T) {
	stones := blink([]int{253000, 1, 7})

	assert.Equal(t, []int{253, 0, 2024, 14168}, stones)
}

func TestBlink1(t *testing.T) {
	stones := blink([]int{125, 17})

	assert.Equal(t, []int{253000, 1, 7}, stones)
}

func TestParseInput(t *testing.T) {
	input := parseInputFile("test_input.txt")

	assert.Equal(t, []int{125, 17}, input)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
