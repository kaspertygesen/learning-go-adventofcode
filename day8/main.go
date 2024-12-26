package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	locations := parseInputFile("input.txt")
	markAntinodes(locations)

	part1 := countAntinodes(locations)

	locations2 := parseInputFile("input.txt")
	markAntinodes2(locations2)

	part2 := countAntinodes(locations2)

	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part 2: %d\n", part2)
}

func markAntinodes2(locations *[][]location) {
	for y := range *locations {
		for x := range (*locations)[y] {
			location1 := &(*locations)[y][x]

			if location1.frequency == '.' {
				continue
			}

			for y1 := range *locations {
				for x1 := range (*locations)[y1] {
					location2 := &(*locations)[y1][x1]

					if location2.x != location1.x && location2.y != location1.y && location2.frequency == location1.frequency {
						location2.isAntinode = true

						vector1 := vector{x: location2.x - location1.x, y: location2.y - location1.y}
						vector2 := vector{x: vector1.x * -1, y: vector1.y * -1}

						for i := 1; ; i++ {
							if ax, ay := location2.x+(vector1.x*i), location2.y+(vector1.y*i); ax >= 0 && len((*locations)[y]) > ax && ay >= 0 && len(*locations) > ay {
								(*locations)[ay][ax].isAntinode = true
							} else {
								break
							}
						}

						for i := 1; ; i++ {
							if ax, ay := location1.x+(vector2.x*i), location1.y+(vector2.y*i); ax >= 0 && len((*locations)[y]) > ax && ay >= 0 && len(*locations) > ay {
								(*locations)[ay][ax].isAntinode = true
							} else {
								break
							}
						}
					}
				}
			}
		}
	}
}

func countAntinodes(locations *[][]location) int {
	count := 0

	for y := range *locations {
		for x := range (*locations)[y] {
			location := (*locations)[y][x]

			if location.isAntinode {
				count++
			}
		}
	}

	return count
}

func markAntinodes(locations *[][]location) {
	for y := range *locations {
		for x := range (*locations)[y] {
			location1 := &(*locations)[y][x]

			if location1.frequency == '.' {
				continue
			}

			for y1 := range *locations {
				for x1 := range (*locations)[y1] {
					location2 := &(*locations)[y1][x1]

					if location2.x != location1.x && location2.y != location1.y && location2.frequency == location1.frequency {
						vector1 := vector{x: location2.x - location1.x, y: location2.y - location1.y}
						vector2 := vector{x: vector1.x * -1, y: vector1.y * -1}

						if ax, ay := location2.x+vector1.x, location2.y+vector1.y; ax >= 0 && len((*locations)[y]) > ax && ay >= 0 && len(*locations) > ay {
							(*locations)[ay][ax].isAntinode = true
						}

						if ax, ay := location1.x+vector2.x, location1.y+vector2.y; ax >= 0 && len((*locations)[y]) > ax && ay >= 0 && len(*locations) > ay {
							(*locations)[ay][ax].isAntinode = true
						}
					}
				}
			}
		}
	}
}

func parseInputFile(path string) *[][]location {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var antennaMap [][]location

	y := 0

	for scanner.Scan() {
		var locations []location

		line := scanner.Text()

		for i, l := range line {
			locations = append(locations, location{x: i, y: y, frequency: l, isAntinode: false})
		}

		antennaMap = append(antennaMap, locations)

		y++
	}

	return &antennaMap
}

type location struct {
	x int
	y int

	frequency rune

	isAntinode bool
}

type vector struct {
	x int
	y int
}
