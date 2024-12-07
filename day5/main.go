package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := parseInputFile("input.txt")

	validUpdates := validUpdates(rules, updates)

	part1 := sumMiddle(validUpdates)

	fmt.Printf("Part 1: %d\n", part1)

	// invalidUpdates := invalidUpdates(rules, updates)

	// fmt.Print(invalidUpdates)

	part2 := 0

	fmt.Printf("Part 2: %d\n", part2)
}

// func correctInvalidUpdates(updates [][]int) {
// 	for _, u := range updates {
// 		slices.SortFunc(u, func(a, b int) int {
// 			return -1
// 		})
// 	}
// }

func invalidUpdates(rules [][]int, updates [][]int) [][]int {
	var invalidUpdates [][]int

	for _, u := range updates {
		valid := true

		for uIndex, n := range u {

			for _, r := range rules {
				if r[0] == n {
					if ruIndex := slices.Index(u, r[1]); ruIndex > -1 {
						if uIndex > ruIndex {
							valid = false
							break
						}
					}
				} else if r[1] == n {
					if ruIndex := slices.Index(u, r[0]); ruIndex > -1 {
						if uIndex < ruIndex {
							valid = false
							break
						}
					}
				}
			}

			if !valid {
				invalidUpdates = append(invalidUpdates, u)
				break
			}
		}
	}

	return invalidUpdates
}

func sumMiddle(updates [][]int) int {
	sum := 0

	for _, u := range updates {
		sum += u[(len(u)-1)/2]
	}

	return sum
}

func validUpdates(rules [][]int, updates [][]int) [][]int {
	var validUpdates [][]int

	for _, u := range updates {
		valid := true

		for uIndex, n := range u {

			for _, r := range rules {
				if r[0] == n {
					if ruIndex := slices.Index(u, r[1]); ruIndex > -1 {
						if uIndex > ruIndex {
							valid = false
							break
						}
					}
				} else if r[1] == n {
					if ruIndex := slices.Index(u, r[0]); ruIndex > -1 {
						if uIndex < ruIndex {
							valid = false
							break
						}
					}
				}
			}

			if !valid {
				break
			}
		}

		if valid {
			validUpdates = append(validUpdates, u)
		}
	}

	return validUpdates
}

func parseInputFile(path string) (rules, updates [][]int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		numbers := strings.Split(line, "|")

		var rule []int

		if number, err := parseInt(numbers[0]); err == nil {
			rule = append(rule, number)
		}

		if number, err := parseInt(numbers[1]); err == nil {
			rule = append(rule, number)
		}

		rules = append(rules, rule)
	}

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, ",")

		var update []int

		for _, v := range numbers {
			if number, err := parseInt(v); err == nil {
				update = append(update, number)
			}
		}

		updates = append(updates, update)
	}

	return
}

func parseInt(str string) (i int, err error) {
	if lInt, err := strconv.ParseInt(str, 10, 64); err == nil {
		return int(lInt), nil
	}

	return 0, err
}
