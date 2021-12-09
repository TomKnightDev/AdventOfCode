package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x     int
	y     int
	count int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var coordDic = []coords{}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), "\n")

	for _, line := range stringLines {
		getCoords(line)
	}

	// Find all counts greater than 1
	count := 0
	for _, val := range coordDic {
		if val.count > 0 {
			count++
		}
	}

	fmt.Println(count)
}

func getCoords(s string) {
	vals := strings.Split(s, " -> ")

	// Diff between the x values
	start := strings.Split(vals[0], ",")
	end := strings.Split(vals[1], ",")

	matchX := -1
	matchY := -1
	if start[0] == end[0] {
		v, err := strconv.Atoi(start[0])
		check(err)
		matchX = v
	} else if start[1] == end[1] {
		v, err := strconv.Atoi(start[1])
		check(err)
		matchY = v
	} else {
		return
	}

	// sort.Strings(start)
	startx, startxErr := strconv.Atoi(start[0])
	check(startxErr)
	starty, startyErr := strconv.Atoi(start[1])
	check(startyErr)

	// sort.Strings(end)
	endx, endxErr := strconv.Atoi(end[0])
	check(endxErr)
	endy, endyErr := strconv.Atoi(end[1])
	check(endyErr)

	if matchY >= 0 {
		if startx > endx {
			t := startx
			startx = endx
			endx = t
		}
		for i := startx; i <= endx; i++ {
			if !containsCoord(coords{x: i, y: matchY}) {
				coordDic = append(coordDic, coords{x: i, y: matchY})
			}
		}
	} else {
		if starty > endy {
			t := starty
			starty = endy
			endy = t
		}
		for i := starty; i <= endy; i++ {
			if !containsCoord(coords{x: matchX, y: i}) {
				coordDic = append(coordDic, coords{x: matchX, y: i})
			}
		}
	}
}

func containsCoord(coord coords) bool {
	for i, c := range coordDic {
		if c.x == coord.x && c.y == coord.y {
			coordDic[i].count++
			return true
		}
	}

	return false
}
