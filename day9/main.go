package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	discMap := mapDisc(parseInputFile("input.txt"))

	compact(discMap)

	part1 := checksum(discMap)

	discMap2 := mapDisc(parseInputFile("input.txt"))

	compact2(discMap2)

	part2 := checksum(discMap2)

	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part 2: %d\n", part2)
}

func compact2(blocks *[]block) {
	swapF := reflect.Swapper(*blocks)

	for i := len(*blocks) - 1; i >= 0; i-- {
		b := &(*blocks)[i]

		if b.empty {
			continue
		}

		emptySize := 0
		for iEmpty := 0; iEmpty < len(*blocks); iEmpty++ {
			if iEmpty >= i {
				break
			}

			if (*blocks)[iEmpty].empty {
				emptySize++
			} else {
				emptySize = 0
			}

			if emptySize == b.fileSize {
				for j := 0; j <= b.fileSize; j++ {
					swapF(i-b.fileSize+1+j, iEmpty-b.fileSize+1+j)
				}

				iEmpty++
				break
			}
		}
	}
}

func discMapToString(discMap *[]block) string {
	var discMapStr string

	for _, b := range *discMap {
		if !b.empty {
			discMapStr += strconv.Itoa(b.id)
		} else {
			discMapStr += "."
		}
	}

	return discMapStr
}

func checksum(blocks *[]block) int {
	sum := 0

	for i, b := range *blocks {
		sum += i * b.id
	}

	return sum
}

func compact(blocks *[]block) {
	swapF := reflect.Swapper(*blocks)

	iEmpty := 0

	for i := len(*blocks) - 1; i >= 0; i-- {
		b := &(*blocks)[i]

		if b.empty {
			continue
		}

		for ; iEmpty < len(*blocks); iEmpty++ {
			if iEmpty < i && (*blocks)[iEmpty].empty {
				swapF(i, iEmpty)
				iEmpty++
				break
			}
		}

		if iEmpty >= i {
			break
		}
	}
}

func mapDisc(input string) *[]block {
	var discMap []block

	fileId := 0

	for i, b := range input {
		if i%2 == 0 {
			for i := 0; i < (int(b) - '0'); i++ {
				discMap = append(discMap, block{id: fileId, fileSize: (int(b) - '0')})
			}
			fileId++
		} else {
			for i := 0; i < (int(b) - '0'); i++ {
				discMap = append(discMap, block{empty: true})
			}
		}
	}

	return &discMap
}

func parseInputFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}

type block struct {
	id       int
	empty    bool
	fileSize int
}
