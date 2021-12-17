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
	mainPart2()
	return
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

func getInputValues(fileName string) []string {
	stringLines := readFromFile(fileName)
	var vals = []string{}

	for _, line := range stringLines {
		stringVals := strings.SplitN(line, " | ", -1)
		vals = append(vals, stringVals[0])
	}

	return vals
}

func getOutputValues(fileName string) []string {
	stringLines := readFromFile(fileName)
	var vals = []string{}

	for _, line := range stringLines {
		stringVals := strings.SplitAfterN(line, " | ", 2)
		vals = append(vals, stringVals[1])

		// for _, val := range stringVals {
		// 	vals = strings.Split(val, " ")
		// }
	}

	return vals
}

func readFromFile(fileName string) []string {
	data, err := os.ReadFile(fileName)
	check(err)

	return strings.Split(string(data), "\n")
}
