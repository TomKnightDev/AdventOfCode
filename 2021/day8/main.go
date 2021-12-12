package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), "\n")

	count := 0
	for _, line := range stringLines {
		stringVals := strings.Split(line, " | ")
		for i, val := range stringVals {
			if i%2 == 0 {
				continue
			}

			vals := strings.Split(val, " ")

			for _, v := range vals {
				if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
