package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := parseInputFile("input.txt")

	pairs := pairLocationIds(left, right)

	sum := sumDistance(pairs)

	fmt.Printf("Part 1: %d\n", sum)

	groups := group(right)

	score := similarityScore(left, groups)

	fmt.Printf("Part 2: %d\n", score)
}

func similarityScore(left []int, groups map[int]int) int {
	var score int

	for _, v := range left {
		score += v * groups[v]
	}

	return score
}

func group(arr []int) map[int]int {
	groups := make(map[int]int)

	for _, v := range arr {
		groups[v]++
	}

	return groups
}

func sumDistance(pairs []locationIdPair) int {
	var sum int

	for _, v := range pairs {
		sum += v.distance()
	}

	return sum
}

func pairLocationIds(left, right []int) []locationIdPair {
	sort.Ints(left[:])
	sort.Ints(right[:])

	var pairs []locationIdPair

	for i, v := range left {
		pairs = append(pairs, locationIdPair{v, right[i]})
	}

	return pairs
}

func parseInputFile(path string) (left, right []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		if lInt, err := parseInt(numbers[0]); err == nil {
			left = append(left, lInt)
		}

		if rInt, err := parseInt(numbers[1]); err == nil {
			right = append(right, rInt)
		}
	}

	return
}

func parseInt(str string) (i int, err error) {
	if lInt, err := strconv.ParseInt(str, 10, 64); err == nil {
		return int(lInt), nil
	}

	return 0, err
}

type locationIdPair struct {
	left  int
	right int
}

func (l *locationIdPair) distance() int {
	if d := l.left - l.right; d < 0 {
		return -d
	} else {
		return d
	}
}
