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

	fmt.Printf("Part 1: %d\n", numSafeReports)

	numSafeReportsWithDampener := evaluateSafetyWithDampener(reports)

	fmt.Printf("Part 2: %d\n", numSafeReportsWithDampener)
}

func evaluateSafetyWithDampener(reports [][]int) int {
	var sum int

	for _, r := range reports {
		var safe = evaluateReport(r)

		if !safe {
			for i, _ := range r {
				if i == 0 {
					s := r[1:]

					safe = evaluateReport(s)
				} else if i == len(r)-1 {
					s := r[:len(r)-1]

					safe = evaluateReport(s)
				} else {
					var s []int
					s = append(s, r[:i]...)
					s = append(s, r[i+1:]...)

					safe = evaluateReport(s)
				}

				if safe {
					break
				}
			}
		}

		if safe {
			sum++
		}
	}

	return sum
}

func evaluateSafety(reports [][]int) int {
	var sum int

	for _, r := range reports {
		safe := evaluateReport(r)

		if safe {
			sum++
		}
	}

	return sum
}

func evaluateReport(report []int) bool {
	asc := report[0]-report[1] < 0

	var prev int
	safe := true
	for i, l := range report {
		if i != 0 {
			diff := prev - l

			safe = safe && ((asc && diff < 0 && diff >= -3) || (!asc && diff > 0 && diff <= 3))
		}

		prev = l
	}

	return safe
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
