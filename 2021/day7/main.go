package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type bestPos struct {
	pos  int
	cost int
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), ",")

	count := 0
	total := 0

	var intLines = []int{}
	for _, line := range stringLines {
		count++
		l, err := strconv.Atoi(line)
		check(err)
		total += l
		intLines = append(intLines, l)
	}

	var arrBestPos = []bestPos{}

	pos := total / count

	half := count / 2

	for i := 0; i < half; i++ {
		cost := getBestPos(intLines, pos+i)
		arrBestPos = append(arrBestPos, bestPos{pos: pos, cost: cost})
	}

	for i := 0; i < half; i++ {
		cost := getBestPos(intLines, pos-i)
		arrBestPos = append(arrBestPos, bestPos{pos: pos, cost: cost})
	}

	lowestCost := 0
	newPos := 0
	for _, bp := range arrBestPos {
		if lowestCost == 0 || bp.cost < lowestCost {
			lowestCost = bp.cost
			newPos = bp.pos
		}
	}

	fmt.Println("Lowest cost:", lowestCost, "New pos:", newPos)
}

func getBestPos(lines []int, pos int) int {
	totalCost := 0
	for _, line := range lines {
		c := int(math.Abs(float64(line - pos)))
		adjustment := 0
		for i := 1; i <= c; i++ {
			adjustment += i
		}

		totalCost += adjustment
	}

	return totalCost
}
