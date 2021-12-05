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
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), "\n")

	hozPos := 0
	depth := 0
	aim := 0

	for _, val := range stringLines {
		input := strings.Split(val, " ")
		inputAsInt, err := strconv.Atoi(input[1])
		check(err)

		switch input[0] {
		case "forward":
			hozPos += inputAsInt

			if aim > 0 {
				depth += inputAsInt * aim
			} else if aim < 0 {
				depth -= inputAsInt * aim
			}

			break
		case "down":
			aim += inputAsInt
			break
		case "up":
			aim -= inputAsInt
			break
		}
	}

	fmt.Println(hozPos * depth)
}
