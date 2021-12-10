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

	stringLines := strings.Split(string(data), ",")

	var lanternFish = []int8{}

	for _, line := range stringLines {
		val, err := strconv.Atoi(line)
		check(err)
		lanternFish = append(lanternFish, int8(val))
	}

	// 80 day loop
	for i := 0; i < 256; i++ {
		count := len(lanternFish)
		for lf := 0; lf < count; lf++ {
			lanternFish[lf] -= 1
			if lanternFish[lf] < 0 {
				lanternFish = append(lanternFish, 8)
				lanternFish[lf] = 6
			}
		}
	}

	fmt.Print(len(lanternFish))
}
