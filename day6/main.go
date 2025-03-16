package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
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

	part2 := countLoops(positions)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func countLoops(positions [][]position) int {
	loopCount := 0

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions[i]); j++ {
			p := &positions[i][j]

			if p.visited && p.symbol == '.' {
				fmt.Printf("X: %d Y: %d\n", p.x, p.y)
				p.symbol = '#'
				isLoop := containsLoop(positions)

				if isLoop {
					loopCount++
				}
				p.symbol = '.'
			}
		}
	}

	return loopCount
}

func containsLoop(positions [][]position) bool {
	direction := up

	var current *position

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions[i]); j++ {
			if positions[i][j].symbol == '^' {
				current = &positions[i][j]
			}
		}
	}

	var steps []step

	for {
		steps = append(steps, step{x: current.x, y: current.y, direction: direction})
		exit := false

		switch direction {
		case up:
			if n := newDirection(current.up); n == 0 {
				current = current.up
			} else if n == 1 {
				direction = right
			} else {
				exit = true
			}
		case right:
			if n := newDirection(current.right); n == 0 {
				current = current.right
			} else if n == 1 {
				direction = down
			} else {
				exit = true
			}
		case down:
			if n := newDirection(current.down); n == 0 {
				current = current.down
			} else if n == 1 {
				direction = left
			} else {
				exit = true
			}
		case left:
			if n := newDirection(current.left); n == 0 {
				current = current.left
			} else if n == 1 {
				direction = up
			} else {
				exit = true
			}
		}

		if exit {
			break
		}

		var indexes []int
		for i := range steps {
			if current.x == steps[i].x && current.y == steps[i].y && direction == steps[i].direction {
				indexes = append(indexes, i)
			}
		}

		if len(indexes) > 1 {
			s1 := steps[indexes[len(indexes)-2]:indexes[len(indexes)-1]]
			s2 := steps[indexes[len(indexes)-1]:]

			if slices.Equal(s1, s2) {
				fmt.Printf("Loop length: %d %d\n", len(s1), len(s2))
				return true
			}
		}
	}

	return false
}

func equal(p1, p2 []step) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if p1[i].x != p2[i].x || p1[i].y != p2[i].y || p1[i].direction != p2[i].direction {
			return false
		}
	}
	return true
}

func countVisitedPositions(positions [][]position) int {
	count := 0

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions[i]); j++ {
			if positions[i][j].visited {
				count++
			}
		}
	}

	return count
}

func patrol(positions [][]position) {
	direction := up

	var current *position

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions[i]); j++ {
			if positions[i][j].symbol == '^' {
				current = &positions[i][j]
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
				direction = right
			} else {
				exit = true
			}
		case right:
			if n := newDirection(current.right); n == 0 {
				current = current.right
			} else if n == 1 {
				direction = down
			} else {
				exit = true
			}
		case down:
			if n := newDirection(current.down); n == 0 {
				current = current.down
			} else if n == 1 {
				direction = left
			} else {
				exit = true
			}
		case left:
			if n := newDirection(current.left); n == 0 {
				current = current.left
			} else if n == 1 {
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

func linkPositions(labMap []string) [][]position {
	positions := make([][]position, len(labMap))

	for i := range positions {
		positions[i] = make([]position, len(labMap[0]))
	}

	for i, row := range labMap {
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

	return positions
}

func parseInputFile(path string) []string {
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

	return labMap
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

type step struct {
	x int
	y int

	direction int
}
