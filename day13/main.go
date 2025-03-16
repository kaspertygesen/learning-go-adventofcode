package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := parseInputFile("input.txt")

	part1 := play(input)
	part2 := 0

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func play(machines []machine) int {
	count := 0

	for _, machine := range machines {
		for a := 0; a*machine.a.x < machine.x || a*machine.a.y < machine.y; a++ {
			ax := a * machine.a.x
			ay := a * machine.a.y

			if bx, by := (machine.x - ax), (machine.y - ay); bx%machine.b.x == 0 && by%machine.b.y == 0 && bx/machine.b.x == by/machine.b.y {
				b := bx / machine.b.x

				if a > 100 || b > 100 {
					break
				}

				count += a*3 + b
			}
		}

		//fmt.Printf("mX: %d mY: %d\n", machine.x, machine.y)
	}

	return count
}

var aRegex *regexp.Regexp = regexp.MustCompile(`Button\sA:\sX\+(?<x>\d+?),\sY\+(?<y>\d+)`)
var bRegex *regexp.Regexp = regexp.MustCompile(`Button\sB:\sX\+(?<x>\d+?),\sY\+(?<y>\d+)`)
var prizeRegex *regexp.Regexp = regexp.MustCompile(`Prize:\sX=(?<x>\d+?),\sY=(?<y>\d+)`)

func parseInputFile(path string) []machine {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var machines []machine

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if aMatches := aRegex.FindStringSubmatch(line); len(aMatches) > 0 {
			ax, _ := strconv.ParseInt(aMatches[1], 10, 64)
			ay, _ := strconv.ParseInt(aMatches[2], 10, 64)

			scanner.Scan()
			line = scanner.Text()
			bMatches := bRegex.FindStringSubmatch(line)

			bx, _ := strconv.ParseInt(bMatches[1], 10, 64)
			by, _ := strconv.ParseInt(bMatches[2], 10, 64)

			scanner.Scan()
			line = scanner.Text()
			pMatches := prizeRegex.FindStringSubmatch(line)

			px, _ := strconv.ParseInt(pMatches[1], 10, 64)
			py, _ := strconv.ParseInt(pMatches[2], 10, 64)

			machines = append(machines, machine{a: button{int(ax), int(ay)}, b: button{int(bx), int(by)}, x: int(px), y: int(py)})
		}
	}

	return machines
}

type machine struct {
	a button
	b button

	x int
	y int
}

type button struct {
	x int
	y int
}
