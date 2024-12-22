package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

func main() {
	labMap := parseInputFile("input.txt")

	positions := linkPositions(labMap)

	patrol(positions)

	part1 := countVisitedPositions(positions)

	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part 2: %d\n", 0)
}

func countVisitedPositions(positions *[][]position) int {
	count := 0

	for i := 0; i < len(*positions); i++ {
		for j := 0; j < len((*positions)[i]); j++ {
			if (*positions)[i][j].visited {
				count++
			}
		}
	}

	return count
}

func patrol(positions *[][]position) {
	direction := up

	var current *position

	for i := 0; i < len(*positions); i++ {
		for j := 0; j < len((*positions)[i]); j++ {
			if (*positions)[i][j].symbol == '^' {
				current = &(*positions)[i][j]
			}
		}
	}

	for {
		current.visited = true

		exit := false

		switch direction {
		case up:
			if n := newDirection(current.up); n == 0 {
				current = current.up
			} else if n == 1 {
				current = current.right
				direction = right
			} else {
				exit = true
			}
		case right:
			if n := newDirection(current.right); n == 0 {
				current = current.right
			} else if n == 1 {
				current = current.down
				direction = down
			} else {
				exit = true
			}
		case down:
			if n := newDirection(current.down); n == 0 {
				current = current.down
			} else if n == 1 {
				current = current.left
				direction = left
			} else {
				exit = true
			}
		case left:
			if n := newDirection(current.left); n == 0 {
				current = current.left
			} else if n == 1 {
				current = current.up
				direction = up
			} else {
				exit = true
			}
		}

		if exit {
			break
		}
	}
}

func newDirection(next *position) int {
	if next == nil {
		return -1
	}
	if next.symbol == '#' {
		return 1
	}

	return 0
}

func linkPositions(labMap *[]string) *[][]position {
	positions := make([][]position, len(*labMap))

	for i := range positions {
		positions[i] = make([]position, len((*labMap)[0]))
	}

	for i, row := range *labMap {
		for j, p := range row {
			positions[i][j] = position{
				symbol: p,
				x:      j,
				y:      i,
			}
		}
	}

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions[i]); j++ {
			p := &positions[i][j]

			if i > 0 {
				p.up = &positions[i-1][j]
			}
			if j < len(positions[i])-1 {
				p.right = &positions[i][j+1]
			}
			if i < len(positions)-1 {
				p.down = &positions[i+1][j]
			}
			if j > 0 {
				p.left = &positions[i][j-1]
			}
		}
	}

	return &positions
}

func parseInputFile(path string) *[]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var labMap []string

	for scanner.Scan() {
		line := scanner.Text()

		labMap = append(labMap, line)
	}

	return &labMap
}

type position struct {
	symbol rune

	up    *position
	right *position
	down  *position
	left  *position

	visited bool

	x int
	y int
}