package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInputFile("input.txt")

	var stones []int = input
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Printf("Part 1: %d\n", len(stones))

	cache := make(map[pair]int)

	sum := 0

	for _, stone := range input {
		sum += blinkRecursive(stone, 75, &cache)
	}

	part2 := sum

	fmt.Printf("Part 2: %d\n", part2)
}

type pair struct {
	n int
	l int
}

func blinkRecursive(n, remainingBlinks int, cache *map[pair]int) int {
	if remainingBlinks == 0 {
		return 1
	}

	r := 0

	if cr, ok := (*cache)[pair{n, remainingBlinks}]; ok {
		return cr
	} else if n == 0 {
		r = blinkRecursive(1, remainingBlinks-1, cache)
	} else if length := int(math.Log10(float64(n))) + 1; length%2 == 0 {
		nl := length / 2
		left := int(math.Floor(float64(n) / math.Pow10(nl)))
		right := n % int(math.Pow10(nl))

		r = blinkRecursive(left, remainingBlinks-1, cache) + blinkRecursive(right, remainingBlinks-1, cache)
	} else {
		r = blinkRecursive(n*2024, remainingBlinks-1, cache)
	}

	(*cache)[pair{n, remainingBlinks}] = r

	return r
}

func blink(stones []int) []int {
	var newStones []int

	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if str := strconv.Itoa(stone); len(str)%2 == 0 {
			sr, _ := strconv.ParseInt(str[len(str)/2:], 10, 64)
			sl, _ := strconv.ParseInt(str[:len(str)/2], 10, 64)

			newStones = append(newStones, int(sl))
			newStones = append(newStones, int(sr))
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	return newStones
}

func parseInputFile(path string) []int {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	contentStr := string(content)

	numbers := strings.Fields(contentStr)

	var stones []int

	for _, number := range numbers {
		if stone, err := parseInt(number); err == nil {
			stones = append(stones, stone)
		}

	}

	return stones
}

func parseInt(str string) (i int, err error) {
	if v, err := strconv.ParseInt(str, 10, 64); err == nil {
		return int(v), nil
	}

	return 0, err
}
