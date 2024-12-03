package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := parseInputFile("input.txt")

	numSafeReports := evaluateSafety(reports)

	fmt.Printf("Part 1: %d", numSafeReports)
}

func evaluateSafety(reports [][]int) int {
	var sum int

	for _, r := range reports {
		asc := r[0]-r[1] < 0

		var prev int
		safe := true
		for i, l := range r {
			if i != 0 {
				diff := prev - l

				if asc {
					safe = safe && diff < 0 && diff >= -3
				} else {
					safe = safe && diff > 0 && diff <= 3
				}
			}

			prev = l
		}

		if safe {
			sum++
		}
	}

	return sum
}

func parseInputFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		var levels []int

		for _, v := range numbers {
			if vInt, err := parseInt(v); err == nil {
				levels = append(levels, vInt)
			}
		}

		reports = append(reports, levels)
	}

	return reports
}

func parseInt(str string) (i int, err error) {
	if v, err := strconv.ParseInt(str, 10, 64); err == nil {
		return int(v), nil
	}

	return 0, err
}
