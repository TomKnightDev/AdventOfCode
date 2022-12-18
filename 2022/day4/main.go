package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	lines := strings.Split(string(bytes), "\n")

	total := 0
	for _, l := range lines {
		pairs := strings.Split(l, ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")

		f1, _ := strconv.Atoi(first[0])
		f2, _ := strconv.Atoi(first[1])
		s1, _ := strconv.Atoi(second[0])
		s2, _ := strconv.Atoi(second[1])

		if (f1 >= s1 && f1 <= s2) || (s1 >= f1 && s1 <= f2) {
			total++
		}
	}

	fmt.Println(total)
}
