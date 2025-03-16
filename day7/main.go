package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	equations := parseInputFile("input.txt")

	validEquations := evaluateEquations(equations)

	part1 := sumResults(validEquations)

	fmt.Printf("Part 1: %d\n", part1)

	validEquations2 := evaluateEquations2(equations)

	part2 := sumResults(validEquations2)

	fmt.Printf("Part 2: %d\n", part2)
}

func sumResults(equations []equation) int64 {
	var sum int64 = 0

	for _, e := range equations {
		sum += e.result
	}

	return sum
}

func evaluateEquations2(equations []equation) []equation {
	var validEquations []equation

	maxLength := 0
	for _, e := range equations {
		if l := len(e.numbers); maxLength < l {
			maxLength = l
		}
	}

	var combinations []string
	generateCombinations2("", maxLength-1, "+*|", &combinations)

	for _, e := range equations {
		for _, c := range combinations {
			operators := c[:len(e.numbers)-1]

			result := e.numbers[0]

			for i, n := range e.numbers[1:] {
				switch operators[i] {
				case '+':
					result += n
				case '*':
					result *= n
				case '|':
					result = concatNumbers(result, n)
				}
			}

			if result == e.result {
				validEquations = append(validEquations, e)
				break
			}
		}
	}

	return validEquations
}

func concatNumbers(n1, n2 int64) int64 {
	return n1*int64(math.Pow10(int(math.Log10(float64(n2))+1))) + n2
}

func generateCombinations2(s string, d int, operators string, combinations *[]string) {
	for _, o := range operators {
		if d != 0 {
			generateCombinations(s+string(o), d-1, operators, combinations)
		}
	}

	if d == 0 {
		*combinations = append(*combinations, s)
	}
}

func evaluateEquations(equations []equation) []equation {
	var validEquations []equation

	maxLength := 0
	for _, e := range equations {
		if l := len(e.numbers); maxLength < l {
			maxLength = l
		}
	}

	var combinations []string
	generateCombinations("", maxLength-1, "+*", &combinations)

	for _, e := range equations {
		if !slices.ContainsFunc(e.numbers, func(n int64) bool { return n < 2 }) {
			var sum int64 = 0
			var product int64 = 1

			for _, n := range e.numbers {
				sum += n
				product *= n
			}

			if e.result < sum || e.result > product {
				//fmt.Print(e)
				//fmt.Println("Impossible")
				continue
			} else if e.result == sum || e.result == product {
				//fmt.Print(e)
				//fmt.Println("Valid")
				validEquations = append(validEquations, e)
				continue
			}
		}

		for _, c := range combinations {
			operators := c[:len(e.numbers)-1]

			result := e.numbers[0]

			for i, n := range e.numbers[1:] {
				switch operators[i] {
				case '+':
					result += n
				case '*':
					result *= n
				}
			}

			if result == e.result {
				validEquations = append(validEquations, e)
				break
			}
		}
	}

	return validEquations
}

func generateCombinations(s string, d int, operators string, combinations *[]string) {
	for _, o := range operators {
		if d != 0 {
			generateCombinations(s+string(o), d-1, operators, combinations)
		}
	}

	if d == 0 {
		*combinations = append(*combinations, s)
	}
}

func parseInputFile(path string) []equation {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var equations []equation

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")

		numbersStr := strings.Fields(parts[1])

		result, _ := parseInt64(parts[0])

		var numbers []int64

		for _, v := range numbersStr {
			if vInt, err := parseInt64(v); err == nil {
				numbers = append(numbers, vInt)
			}
		}

		equations = append(equations, equation{result: int64(result), numbers: numbers})
	}

	return equations
}

func parseInt64(str string) (i int64, err error) {
	if v, err := strconv.ParseInt(str, 10, 64); err == nil {
		return v, nil
	}

	return 0, err
}

type equation struct {
	result int64

	numbers []int64
}
