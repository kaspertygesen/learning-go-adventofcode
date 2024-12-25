package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAntinodes(t *testing.T) {
	locations := parseInputFile("test_input.txt")

	markAntinodes(locations)

	count := countAntinodes(locations)

	for y := range *locations {
		for x := range (*locations)[y] {
			location := (*locations)[y][x]

			if location.isAntinode {
				fmt.Print("#")
			} else {
				fmt.Print(string(location.frequency))
			}
		}
		fmt.Print("\n")
	}

	assert.Equal(t, 14, count)
}

func TestParseInput(t *testing.T) {
	antennaMap := parseInputFile("test_input.txt")

	assert.Equal(t, location{x: 0, y: 0, frequency: '.', isAntinode: false}, (*antennaMap)[0][0])
	assert.Equal(t, location{x: 8, y: 1, frequency: '0', isAntinode: false}, (*antennaMap)[1][8])
	assert.Equal(t, location{x: 8, y: 8, frequency: 'A', isAntinode: false}, (*antennaMap)[8][8])
}
