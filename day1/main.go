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

	fmt.Print(pairs)
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
	return l.left - l.right
}
