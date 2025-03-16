package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumE(t *testing.T) {
	input := parseInputFile("test_input_e.txt")

	addCoordinates(input)

	groups := group(input)

	sum := sum2(groups)

	assert.Equal(t, 236, sum)
}

func TestFencesE(t *testing.T) {
	input := parseInputFile("test_input_e.txt")

	addCoordinates(input)

	groups := group(input)

	for _, g := range groups {
		edgeGroups := edgeGroups(g)
		fences := fences(edgeGroups)

		if g[0].plant == "E" {
			assert.Equal(t, 12, fences)
		}
		if g[0].plant == "X" && g[0].y == 1 {
			assert.Equal(t, 4, fences)
		}
		if g[0].plant == "X" && g[0].y == 3 {
			assert.Equal(t, 4, fences)
		}
	}
}

func TestEdgeGroupsE(t *testing.T) {
	input := parseInputFile("test_input_e.txt")

	addCoordinates(input)

	groups := group(input)

	assert.Len(t, groups, 3)

	for _, g := range groups {
		edgeGroups := edgeGroups(g)
		if g[0].plant == "E" {
			assert.Len(t, edgeGroups, 9)

			// horizontal
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 0, position: 0}], []int{0, 1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 2, position: 0}], []int{1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 0, position: 2}], []int{1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 2, position: 2}], []int{1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 0, position: 4}], []int{1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 2, position: 4}], []int{0, 1, 2, 3, 4})

			// vertical
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 3, position: 0}], []int{0, 1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 1, position: 0}], []int{1, 3})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 1, position: 4}], []int{0, 2, 4})
		}
		if g[0].plant == "X" && g[0].y == 1 {

			assert.Len(t, edgeGroups, 4)

			// horizontal
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 0, position: 1}], []int{1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 2, position: 1}], []int{1, 2, 3, 4})

			// vertical
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 3, position: 1}], []int{1})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 1, position: 4}], []int{1})
		}
		if g[0].plant == "X" && g[0].y == 3 {
			assert.Len(t, edgeGroups, 4)

			// horizontal
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 0, position: 3}], []int{1, 2, 3, 4})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 2, position: 3}], []int{1, 2, 3, 4})

			// vertical
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 3, position: 1}], []int{3})
			assertEdgeGroup(t, edgeGroups[edgeGroup{direction: 1, position: 4}], []int{3})
		}
	}
}

func assertEdgeGroup(t *testing.T, actual []int, expected []int) {
	slices.Sort(actual)

	assert.Equal(t, actual, expected)
}

func TestSum2(t *testing.T) {
	input := parseInputFile("test_input.txt")

	addCoordinates(input)

	groups := group(input)

	sum := sum2(groups)

	assert.Equal(t, 1206, sum)
}

// func TestGroupsEdges(t *testing.T) {
// 	input := parseInputFile("test_input.txt")

// 	addCoordinates(input)

// 	groups := group(input)

// 	assert.Len(t, groups, 11)
// 	for _, g := range groups {
// 		for _, r := range g {
// 			if r.x == 0 && r.y == 0 {
// 				assert.Len(t, r.edges, 2)
// 				assert.Contains(t, r.edges, edge{direction: left, position: 0})
// 				assert.Contains(t, r.edges, edge{direction: up, position: 0})
// 			}
// 			if r.x == 1 && r.y == 0 {
// 				assert.Len(t, r.edges, 1)
// 				assert.Contains(t, r.edges, edge{direction: up, position: 0})
// 			}
// 			if r.x == 2 && r.y == 0 {
// 				assert.Len(t, r.edges, 1)
// 				assert.Contains(t, r.edges, edge{direction: up, position: 0})
// 			}
// 			if r.x == 3 && r.y == 0 {
// 				assert.Len(t, r.edges, 2)
// 				assert.Contains(t, r.edges, edge{direction: up, position: 0})
// 				assert.Contains(t, r.edges, edge{direction: right, position: 3})
// 			}
// 			if r.x == 3 && r.y == 1 {
// 				assert.Len(t, r.edges, 1)
// 				assert.Contains(t, r.edges, edge{direction: right, position: 3})
// 			}
// 			if r.x == 3 && r.y == 2 {
// 				assert.Len(t, r.edges, 1)
// 				assert.Contains(t, r.edges, edge{direction: down, position: 2})
// 			}
// 			if r.x == 4 && r.y == 2 {
// 				assert.Len(t, r.edges, 3)
// 				assert.Contains(t, r.edges, edge{direction: right, position: 4})
// 				assert.Contains(t, r.edges, edge{direction: down, position: 2})
// 				assert.Contains(t, r.edges, edge{direction: up, position: 2})
// 			}
// 			if r.x == 2 && r.y == 3 {
// 				assert.Len(t, r.edges, 3)
// 				assert.Contains(t, r.edges, edge{direction: right, position: 2})
// 				assert.Contains(t, r.edges, edge{direction: down, position: 3})
// 				assert.Contains(t, r.edges, edge{direction: left, position: 2})
// 			}
// 		}
// 	}
// }

func TestSum(t *testing.T) {
	input := parseInputFile("test_input.txt")

	groups := group(input)

	assert.Equal(t, 1930, sum(groups))
}

func TestGroups(t *testing.T) {
	input := parseInputFile("test_input.txt")

	groups := group(input)

	assert.Len(t, groups, 11)
}

func TestParseInput(t *testing.T) {
	input := parseInputFile("test_input.txt")

	assert.Equal(t, "R", input[0][0].plant)
	assert.Equal(t, "E", input[9][9].plant)
	assert.Len(t, input, 10)
}
