package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to read file: " + err.Error())
	}

	lines := strings.Split(string(bytes), "\n")

	stacks := make(map[string][]string)

	movementsStart := 0

	// build stacks
	for i, line := range lines {
		if line != "" {
			continue
		}

		// iterate over runes in the stack ID line
		for ri, r := range lines[i-1] {
			if string(r) == " " {
				continue
			}

			crates := []string{}

			for j := i - 2; j >= 0; j-- {
				if string(lines[j][ri]) == " " {
					break
				}
				crates = append(crates, string(lines[j][ri]))
			}

			stacks[string(r)] = crates
		}

		movementsStart = i + 1
		break
	}

	// move crates
	for i := movementsStart; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		text := strings.Split(lines[i], " ")
		amount, from, to := text[1], text[3], text[5]

		count, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal("failed to convert to int: " + err.Error())
		}

		// for s := 0; s < count; s++ {
		crates := stacks[from][len(stacks[from])-count : len(stacks[from])]
		stacks[to] = append(stacks[to], crates...)
		stacks[from] = stacks[from][:len(stacks[from])-count]
		fmt.Println(stacks)
		// }
	}

	fmt.Println(stacks)

	topCrates := ""
	for i := 1; i <= len(stacks); i++ {
		topCrates += stacks[fmt.Sprint(i)][len(stacks[fmt.Sprint(i)])-1]
	}
	fmt.Println(topCrates)
}
