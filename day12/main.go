package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	input := parseInputFile("input.txt")

	groups := group(input)

	part1 := sum(groups)
	part2 := sum2(groups)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func sum2(groups [][]region) int {
	sum := 0

	for _, g := range groups {
		edgeGroups := edgeGroups(g)

		fences := fences(edgeGroups)

		sum += fences * len(g)
	}

	return sum
}

func fences(edgeGroups map[edgeGroup][]int) int {
	fences := 0
	for _, edges := range edgeGroups {
		slices.Sort(edges)

		last := -1
		for _, e := range edges {
			if last == -1 {
				fences++
			} else if last+1 < e {
				fences++
			}

			last = e
		}
	}

	return fences
}

func edgeGroups(group []region) map[edgeGroup][]int {
	edgeGroups := make(map[edgeGroup][]int)

	for _, r := range group {
		for _, e := range r.edges {
			eg := edgeGroup{direction: e.direction, position: e.position}

			//if _, ok := edgeGroups[eg]; ok {
			edgeGroups[eg] = append(edgeGroups[eg], e.value)
			//} else {
			//	edgeGroups[eg] = []int{e.value}
			//}
		}
	}

	return edgeGroups
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

func group(garden [][]region) [][]region {
	var groups [][]region

	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			r := &garden[y][x]

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

func search(garden [][]region, p string, x, y int, group *[]region) bool {
	r := &garden[y][x]

	if r.visited {
		return r.plant == p
	}

	if r.plant == p {
		r.visited = true
		fences := 4

		var edges []edge

		if nx := x - 1; nx >= 0 {
			if search(garden, p, nx, y, group) {
				fences--
			} else {
				edges = append(edges, edge{direction: left, position: x, value: y})
			}
		} else {
			edges = append(edges, edge{direction: left, position: x, value: y})
		}

		if ny := y - 1; ny >= 0 {
			if search(garden, p, x, ny, group) {
				fences--
			} else {
				edges = append(edges, edge{direction: up, position: y, value: x})
			}
		} else {
			edges = append(edges, edge{direction: up, position: y, value: x})
		}

		if nx := x + 1; nx < len(garden[0]) {
			if search(garden, p, nx, y, group) {
				fences--
			} else {
				edges = append(edges, edge{direction: right, position: x, value: y})
			}
		} else {
			edges = append(edges, edge{direction: right, position: x, value: y})
		}

		if ny := y + 1; ny < len(garden[0]) {
			if search(garden, p, x, ny, group) {
				fences--
			} else {
				edges = append(edges, edge{direction: down, position: y, value: x})
			}
		} else {
			edges = append(edges, edge{direction: down, position: y, value: x})
		}

		r.fences = fences
		r.edges = edges

		*group = append(*group, *r)

		return true
	} else {
		return false
	}
}

func addCoordinates(garden [][]region) {
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			garden[y][x].x = x
			garden[y][x].y = y
		}
	}
}

func parseInputFile(path string) [][]region {
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

	return garden
}

type region struct {
	plant   string
	visited bool
	fences  int

	x int
	y int

	edges []edge
}

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

type edge struct {
	direction int
	position  int
	value     int
}

type edgeGroup struct {
	direction int
	position  int
}
