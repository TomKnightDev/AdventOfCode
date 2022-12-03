package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to read file")
	}

	lines := strings.Split(string(bytes), "\n")

	elves := []int{}
	currentTotal := 0

	for _, l := range lines {
		if l == "" {
			elves = append(elves, currentTotal)
			currentTotal = 0
			continue
		}

		v, err := strconv.Atoi(l)
		if err != nil {
			panic("failed to convert string to int")
		}

		currentTotal += v
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Println(elves[0] + elves[1] + elves[2])
}
