package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	instructions := parseInputFile("input.txt")

	calculations := parseInstructions(instructions)

	sum := calculate(calculations)

	fmt.Printf("Part 1: %d\n", sum)

	calculations2 := parseInstructionsExtended(instructions)

	sum2 := calculate(calculations2)

	fmt.Printf("Part 2: %d\n", sum2)
}

func parseInstructionsExtended(instructions string) [][]int {
	pattern, err := regexp.Compile(`(mul\(\d+?,\d+?\))|(don't\(\))|(do\(\))`)

	if err != nil {
		panic(err)
	}

	validInstructions := pattern.FindAllString(instructions, -1)

	var calculations [][]int

	do := true
	for _, v := range validInstructions {
		if v == "don't()" {
			do = false
			continue
		} else if v == "do()" {
			do = true
			continue
		}

		if do {
			numbers := strings.Split(v[4:len(v)-1], ",")

			integers := make([]int, 2)

			integers[0] = parseInt(numbers[0])
			integers[1] = parseInt(numbers[1])

			calculations = append(calculations, integers)
		}
	}

	return calculations
}

func calculate(calculations [][]int) int {
	var sum int

	for _, v := range calculations {
		sum += v[0] * v[1]
	}

	return sum
}

func parseInstructions(instructions string) [][]int {
	pattern, err := regexp.Compile(`mul\(\d+?,\d+?\)`)

	if err != nil {
		panic(err)
	}

	validInstructions := pattern.FindAllString(instructions, -1)

	var calculations [][]int

	for _, v := range validInstructions {
		numbers := strings.Split(v[4:len(v)-1], ",")

		integers := make([]int, 2)

		integers[0] = parseInt(numbers[0])
		integers[1] = parseInt(numbers[1])

		calculations = append(calculations, integers)
	}

	return calculations
}

func parseInputFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func parseInt(str string) int {
	if v, err := strconv.ParseInt(str, 10, 64); err == nil {
		return int(v)
	} else {
		panic(err)
	}
}
