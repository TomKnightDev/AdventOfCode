package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	numbers [][]int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), "\n")
	drawnStrings := strings.Split(string(stringLines[0]), ",")

	var drawnNumbers = []int{}
	for _, v := range drawnStrings {
		num, err := strconv.Atoi(string(v))
		check(err)
		drawnNumbers = append(drawnNumbers, num)
	}

	var boards = []board{}

	for i := 0; i < len(stringLines); i++ {
		if i == 0 {
			continue
		}

		line := stringLines[i]

		if line == "" {
			continue
		}

		// Pass five lines to build board
		boards = append(boards, buildBoard(stringLines[i:i+5]))

		i += 4
	}

	winningValue := 0

	for i := 4; i < len(drawnNumbers); i++ {
		for _, board := range boards {
			if checkBoardForWin(board, drawnNumbers[0:i]) {
				winningValue = getWinningBoardValue(board, drawnNumbers[0:i])
				fmt.Print(winningValue * drawnNumbers[i-1])
				return
			}
		}
	}

}

// Sum all found numbers and deduct this from sum of all numbers
func getWinningBoardValue(b board, numbers []int) int {
	foundNumbers := 0
	allNumbers := 0
	for i, av := range numbers {
		for _, bvs := range b.numbers {
			for _, bv := range bvs {
				if i == 0 {
					allNumbers += bv
				}
				if av == bv {
					foundNumbers += av
				}
			}
		}
	}

	return allNumbers - foundNumbers
}

func checkBoardForWin(b board, numbers []int) bool {
	// Check rows first
	for _, row := range b.numbers {
		if contains(numbers, row) {
			return true
		}
	}

	// Check columns
	var cols = [5][5]int{}
	for j := 0; j < len(b.numbers); j++ {
		for i := 0; i < len(b.numbers); i++ {
			cols[j][i] = b.numbers[i][j]
		}
	}

	for _, col := range cols {
		if contains(numbers, col[:]) {
			return true
		}
	}

	return false
}

func contains(a []int, b []int) bool {
	var found = []int{}

	for _, av := range a {
		if len(found) == len(b) {
			break
		}

		for _, bv := range b {
			if av == bv {
				found = append(found, av)
				break
			}
		}
	}

	return len(found) == len(b)
}

func buildBoard(lines []string) board {
	var boardToReturn = board{}

	for _, line := range lines {
		nums := strings.Split(line, " ")

		lineArray := make([]int, 0)
		for _, val := range nums {
			if val == "" {
				continue
			}

			num, err := strconv.Atoi(val)
			check(err)
			lineArray = append(lineArray, num)
		}

		boardToReturn.numbers = append(boardToReturn.numbers, [][]int{lineArray}...)
	}

	return boardToReturn
}
