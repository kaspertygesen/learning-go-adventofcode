package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := parseInputFile("input.txt")

	groups := group(input)

	part1 := sum(groups)
	part2 := 0

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func sum(groups [][]region) int {
	sum := 0

	for _, g := range groups {
		for _, r := range g {
			sum += r.fences * len(g)
		}
	}

	return sum
}

func group(garden *[][]region) [][]region {
	var groups [][]region

	for y := 0; y < len(*garden); y++ {
		for x := 0; x < len((*garden)[y]); x++ {
			r := &(*garden)[y][x]

			if r.visited {
				continue
			}

			group := []region{}
			search(garden, r.plant, x, y, &group)

			if len(group) > 0 {
				groups = append(groups, group)
			}
		}
	}

	return groups
}

func search(garden *[][]region, p string, x, y int, group *[]region) bool {
	r := &(*garden)[y][x]

	if r.visited {
		return r.plant == p
	}

	if r.plant == p {
		r.visited = true
		fences := 4

		if nx := x - 1; nx >= 0 {
			if search(garden, p, nx, y, group) {
				fences--
			}
		}

		if ny := y - 1; ny >= 0 {
			if search(garden, p, x, ny, group) {
				fences--
			}
		}

		if nx := x + 1; nx < len((*garden)[0]) {
			if search(garden, p, nx, y, group) {
				fences--
			}
		}

		if ny := y + 1; ny < len((*garden)[0]) {
			if search(garden, p, x, ny, group) {
				fences--
			}
		}

		r.fences = fences
		*group = append(*group, *r)

		return true
	} else {
		return false
	}
}

func parseInputFile(path string) *[][]region {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var garden [][]region

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var regions []region

		for _, p := range line {
			regions = append(regions, region{plant: string(p), visited: false})
		}

		garden = append(garden, regions)
	}

	return &garden
}

type region struct {
	plant   string
	visited bool
	fences  int
}
