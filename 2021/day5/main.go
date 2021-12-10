package main

import (
	"fmt"
	"math"
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
			// fmt.Print(val)
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

	matchX := -1
	matchY := -1
	if startx == endx {
		matchX = startx
	} else if starty == endy {
		matchY = starty
	} else if startx == starty && endx == endy {
		for i := 0; i <= int(math.Abs(float64(startx)-float64(endx))); i++ {
			if !containsCoord(coords{x: i, y: i}) {
				coordDic = append(coordDic, coords{x: i, y: i})
			}
		}
		return
	} else if startx == endy && starty == endx {
		for i := 0; i <= int(math.Abs(float64(startx)-float64(starty))); i++ {
			if !containsCoord(coords{x: startx - i, y: i}) {
				coordDic = append(coordDic, coords{x: startx - i, y: i})
			}
		}
		return
	} else if math.Abs(float64(startx-endx)) == math.Abs(float64(starty-endy)) {
		count := int(math.Abs(float64(startx) - float64(endx)))

		xIncrease := startx < endx
		yIncrease := starty < endy

		for i := 0; i <= count; i++ {
			xCount := i
			if !xIncrease {
				xCount = -i
			}
			yCount := i
			if !yIncrease {
				yCount = -i
			}

			if !containsCoord(coords{x: startx + xCount, y: starty + yCount}) {
				coordDic = append(coordDic, coords{x: startx + xCount, y: starty + yCount})
			}
		}
		return
	} else if startx == starty || endx == endy {
		xIncrease := startx < endx
		yIncrease := starty < endy

		iCount := int(math.Abs(float64(startx) - float64(endx)))
		if iCount == 0 {
			iCount = int(math.Abs(float64(starty) - float64(endy)))
		}

		for i := 0; i <= iCount; i++ {
			xCount := i
			if !xIncrease {
				xCount = -i
			}
			yCount := i
			if !yIncrease {
				yCount = -i
			}

			if !containsCoord(coords{x: startx + xCount, y: starty + yCount}) {
				coordDic = append(coordDic, coords{x: startx + xCount, y: starty + yCount})
			}
		}

	} else {
		return
	}

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
	} else if matchX >= 0 {
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
