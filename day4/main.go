package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := parseInputFile("input.txt")

	views := views(input)

	part1 := search(views)

	fmt.Printf("Part 1: %d\n", part1)

	part2 := xmas(input)

	fmt.Printf("Part 2: %d\n", part2)
}

func xmas(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0

	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == 'A' && x > 0 && x < len(lines[0])-1 && y > 0 && y < len(lines)-1 {
				s1 := string(lines[y-1][x-1]) + "A" + string(lines[y+1][x+1])
				s2 := string(lines[y-1][x+1]) + "A" + string(lines[y+1][x-1])

				if (s1 == "MAS" || s1 == "SAM") && (s2 == "SAM" || s2 == "MAS") {
					sum++
				}
			}
		}
	}

	return sum
}

func search(views []string) int {
	var sum int

	for _, v := range views {
		i := 0
		for {
			j := strings.Index(v[i:], "XMAS")
			if j == -1 {
				break
			}

			sum++

			i += j + 1
		}

		i = 0
		for {
			j := strings.Index(v[i:], "SAMX")
			if j == -1 {
				break
			}

			sum++

			i += j + 1
		}
	}

	return sum
}

func views(input string) []string {
	var views []string
	views = append(views, input)

	lines := strings.Split(input, "\n")

	var verticalView []byte
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			verticalView = append(verticalView, lines[y][x])
		}
		verticalView = append(verticalView, "\n"[0])
	}

	views = append(views, string(verticalView))

	var diagonalLtrView []byte
	for x := 0; x < len(lines[0]); x++ {

		dx, dy := x, 0
		for dx < len(lines[0]) && dy < len(lines) {
			diagonalLtrView = append(diagonalLtrView, lines[dy][dx])

			dx++
			dy++
		}
		diagonalLtrView = append(diagonalLtrView, "\n"[0])
	}

	for y := 1; y < len(lines); y++ {

		dx, dy := 0, y
		for dx < len(lines[0]) && dy < len(lines) {
			diagonalLtrView = append(diagonalLtrView, lines[dy][dx])

			dx++
			dy++
		}
		diagonalLtrView = append(diagonalLtrView, "\n"[0])
	}

	views = append(views, string(diagonalLtrView))

	var diagonalRtlView []byte
	for x := len(lines[0]) - 1; x >= 0; x-- {

		dx, dy := x, 0
		for dx >= 0 && dy < len(lines) {
			diagonalRtlView = append(diagonalRtlView, lines[dy][dx])

			dx--
			dy++
		}
		diagonalRtlView = append(diagonalRtlView, "\n"[0])
	}

	for y := 1; y < len(lines); y++ {

		dx, dy := len(lines[0])-1, y
		for dx >= 0 && dy < len(lines) {
			diagonalRtlView = append(diagonalRtlView, lines[dy][dx])

			dx--
			dy++
		}
		diagonalRtlView = append(diagonalRtlView, "\n"[0])
	}

	views = append(views, string(diagonalRtlView))

	return views
}

func parseInputFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}
