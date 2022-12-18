package main

import (
	"fmt"
	"os"
	"strings"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	points := make(map[string]int)

	for i, l := range letters {
		points[l] = i + 1
		points[strings.ToUpper(l)] = i + 27
	}

	lines := strings.Split(string(bytes), "\n")
	total := 0

	totalLines := len(lines)

	for i := 0; i < totalLines; i += 3 {
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]

		for _, r := range line1 {
			if strings.Contains(line2, string(r)) {
				if strings.Contains(line3, string(r)) {
					total += points[string(r)]
					break
				}
			}
		}
	}

	// for _, l := range lines {
	// 	length := len(l)
	// 	c1 := l[0 : length/2]
	// 	c2 := l[length/2:]

	// 	score := make(map[string]int)

	// 	for _, v := range c1 {
	// 		for _, v2 := range c2 {
	// 			if v == v2 {
	// 				if _, found := score[string(v)]; !found {
	// 					score[string(v)] = points[string(v)]
	// 					total += points[string(v)]
	// 				}

	// 			}
	// 		}
	// 	}
	// }

	fmt.Println(total)
}
