package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(dat), "\n")

	var intLines = []int{}
	increases := 0

	// Build array of integers
	for _, i := range stringLines {
		j, err := strconv.Atoi(i)
		check(err)
		intLines = append(intLines, j)
	}

	for i := 2; i < len(intLines)-1; i++ {
		previous := intLines[i-2] + intLines[i-1] + intLines[i]
		current := intLines[i-1] + intLines[i] + intLines[i+1]

		if current > previous {
			increases++
		}
	}

	fmt.Print(increases)
}
