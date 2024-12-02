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
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		if lInt, err := parse(numbers[0]); err == nil {
			left = append(left, lInt)
		}

		if rInt, err := parse(numbers[1]); err == nil {
			right = append(right, rInt)
		}
	}

	sort.Ints(left[:])
	sort.Ints(right[:])

	fmt.Print(left)
	fmt.Print(right)
}

func parse(str string) (i int, err error) {
	if lInt, err := strconv.ParseInt(str, 10, 64); err == nil {
		return int(lInt), nil
	}

	return 0, err
}
