package main

import (
	"fmt"
	"os"
	"strings"
)

// rock = 1
// paper = 2
// scissors = 3

// lose = 0
// draw = 3
// win = 6

// X = lose
// Y = draw
// Z = win

var outcomes = map[string]int{
	"A X": 3, // rock/lose = scissors
	"A Y": 4, // rock/draw = rock
	"A Z": 8, // rock/win = paper
	"B X": 1, // paper/lose = rock
	"B Y": 5, // paper/draw = paper
	"B Z": 9, // paper/win = scissors
	"C X": 2, // scissors/lose = paper
	"C Y": 6, // scissors/draw = scissors
	"C Z": 7, // scissors/win  = rock
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to read file")
	}

	lines := strings.Split(string(bytes), "\n")

	total := 0
	for _, game := range lines {
		total += outcomes[game]
	}

	fmt.Println(total)
}
