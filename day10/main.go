package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	terrainMap := parseInputFile("input.txt")

	part1 := findPaths(terrainMap, true)

	part2 := findPaths(terrainMap, false)

	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part 2: %d\n", part2)
}

func findPaths(terrainMap [][]int, distrinct bool) int {
	count := 0

	for y := 0; y < len(terrainMap); y++ {
		for x := 0; x < len(terrainMap[y]); x++ {
			height := terrainMap[y][x]

			if height == 0 {
				tops := evaluateTrailhead(terrainMap, x, y, distrinct)

				count += len(tops)
			}
		}
	}

	return count
}

func evaluateTrailhead(terrainMap [][]int, x, y int, distrinct bool) []coordinate {
	height := terrainMap[y][x]

	var tops []coordinate

	if yu := y - 1; yu >= 0 {
		tops = append(tops, evaluateNext(terrainMap, height, x, yu, distrinct)...)
	}
	if yd := y + 1; yd < len(terrainMap) {
		tops = append(tops, evaluateNext(terrainMap, height, x, yd, distrinct)...)
	}

	if xr := x + 1; xr < len(terrainMap) {
		tops = append(tops, evaluateNext(terrainMap, height, xr, y, distrinct)...)
	}
	if xl := x - 1; xl >= 0 {
		tops = append(tops, evaluateNext(terrainMap, height, xl, y, distrinct)...)
	}

	if !distrinct {
		return tops
	}

	var topMap map[coordinate]bool = make(map[coordinate]bool)
	for _, t := range tops {
		topMap[t] = true
	}

	var distinctTops []coordinate
	for k := range topMap {
		distinctTops = append(distinctTops, k)
	}

	return distinctTops
}

func evaluateNext(terrainMap [][]int, height int, x, y int, distrinct bool) []coordinate {
	if nextHeight := &terrainMap[y][x]; *nextHeight == 9 && height == 8 {
		return []coordinate{{x, y}}
	} else if *nextHeight == height+1 {
		return evaluateTrailhead(terrainMap, x, y, distrinct)
	} else {
		return []coordinate{}
	}
}

func parseInputFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var terrainMap [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var numbers []int

		for _, v := range line {
			numbers = append(numbers, (int(v) - '0'))
		}

		terrainMap = append(terrainMap, numbers)
	}

	return terrainMap
}

type coordinate struct {
	x int
	y int
}
