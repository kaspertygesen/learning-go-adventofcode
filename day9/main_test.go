package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum2(t *testing.T) {
	discMap := mapDisc(parseInputFile("test_input.txt"))

	compact2(discMap)

	sum := checksum(discMap)

	assert.Equal(t, 2858, sum)
}

func TestCompact2(t *testing.T) {
	discMap := mapDisc(parseInputFile("test_input.txt"))

	compact2(discMap)

	discMapStr := discMapToString(discMap)

	assert.Equal(t, "00992111777.44.333....5555.6666.....8888..", discMapStr)
}

func TestChecksum(t *testing.T) {
	discMap := mapDisc(parseInputFile("test_input.txt"))

	compact(discMap)

	checksum := checksum(discMap)

	assert.Equal(t, 1928, checksum)
}

func TestCompact(t *testing.T) {
	discMap := mapDisc(parseInputFile("test_input.txt"))

	compact(discMap)

	discMapStr := discMapToString(discMap)

	assert.Equal(t, "0099811188827773336446555566..............", discMapStr)
}

func TestDiscMap(t *testing.T) {
	discMap := mapDisc(parseInputFile("test_input.txt"))

	discMapStr := discMapToString(discMap)

	assert.Equal(t, "00...111...2...333.44.5555.6666.777.888899", discMapStr)
}

func TestParseInput(t *testing.T) {
	input := parseInputFile("test_input.txt")

	assert.Equal(t, "2333133121414131402", input)
}
