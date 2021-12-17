package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type digit struct {
	value         int
	pattern       string
	patterNumbers []int
}

func mainPart2() {
	// var segments = [7]string{"a", "b", "c", "d", "e", "f", "g"}
	// var segmentsVals = [7]int{1,	  2,   3,   4,   5,   6,   7}
	var digits = []digit{}

	digits = append(digits, digit{value: 0, pattern: "abcefg", patterNumbers: []int{1, 2, 3, 5, 6, 7}})
	digits = append(digits, digit{value: 1, pattern: "cf", patterNumbers: []int{3, 6}})
	digits = append(digits, digit{value: 2, pattern: "acdeg", patterNumbers: []int{1, 3, 4, 5, 7}})
	digits = append(digits, digit{value: 3, pattern: "acdfg", patterNumbers: []int{1, 3, 4, 6, 7}})
	digits = append(digits, digit{value: 4, pattern: "bcdf", patterNumbers: []int{2, 3, 6}})
	digits = append(digits, digit{value: 5, pattern: "abdfg", patterNumbers: []int{1, 2, 4, 6, 7}})
	digits = append(digits, digit{value: 6, pattern: "abdefg", patterNumbers: []int{1, 2, 4, 5, 6, 7}})
	digits = append(digits, digit{value: 7, pattern: "acf", patterNumbers: []int{1, 3, 6}})
	digits = append(digits, digit{value: 8, pattern: "abcdefg", patterNumbers: []int{1, 2, 3, 4, 5, 6, 7}})
	digits = append(digits, digit{value: 9, pattern: "abcdfg", patterNumbers: []int{1, 2, 3, 4, 6, 7}})

	for _, dig := range digits {
		total := 0
		for _, digv := range dig.patterNumbers {
			total += digv
		}
		fmt.Println(dig.value, ": ", total)
	}

	input := getInputValues("test-input.txt")
	output := getOutputValues("test-input.txt")

	total := 0
	for i, inputVal := range input {
		digits := decryptInput(inputVal)

		total += decryptOutput(output[i], digits)
	}

	fmt.Println(total)
}

func decryptOutput(output string, digits []digit) int {
	vals := strings.Split(output, " ")

	returnValue := ""
	for _, val := range vals {
		intVal := getValueFromPattern(digits, val, true)
		returnValue += fmt.Sprint(intVal)
	}

	data, err := strconv.Atoi(returnValue)
	check(err)
	return data
}

func decryptInput(input string) []digit {
	vals := strings.Split(input, " ")

	var digits = []digit{}

	for _, val := range vals {
		if len(val) == 2 && !digitsContains(digits, 1) {
			digits = append(digits, digit{value: 1, pattern: val})
		} else if len(val) == 4 && !digitsContains(digits, 4) {
			digits = append(digits, digit{value: 4, pattern: val})
		} else if len(val) == 3 && !digitsContains(digits, 7) {
			digits = append(digits, digit{value: 7, pattern: val})
		} else if len(val) == 7 && !digitsContains(digits, 8) {
			digits = append(digits, digit{value: 8, pattern: val})
		}
	}

	// Get the rest of the numbers
	for {
		for _, val := range vals {
			if getValueFromPattern(digits, val, false) != 0 {
				continue
			}

			if len(val) == 6 && val[5] == getPatternForNumber(digits, 7)[0] {
				digits = append(digits, digit{value: 9, pattern: val})
			} else if len(val) == 5 && val[4] == getPatternForNumber(digits, 4)[0] {
				digits = append(digits, digit{value: 5, pattern: val})
			} else if len(val) == 5 && val[4] == getPatternForNumber(digits, 1)[0] {
				digits = append(digits, digit{value: 2, pattern: val})
			} else if len(val) == 5 && val[4] == getPatternForNumber(digits, 7)[0] {
				digits = append(digits, digit{value: 3, pattern: val})
			} else if len(val) == 6 && val[1] == getPatternForNumber(digits, 7)[0] {
				digits = append(digits, digit{value: 6, pattern: val})
			} else if len(val) == 6 && val[1] == getPatternForNumber(digits, 1)[0] {
				digits = append(digits, digit{value: 0, pattern: val})
			}
		}
		if len(digits) == 10 {
			break
		}
	}

	return digits
}

func getValueFromPattern(digits []digit, pattern string, doPanic bool) int {
	pattern = sortString(pattern)
	for _, digit := range digits {
		p := sortString(digit.pattern)
		if p == pattern {
			return digit.value
		}
	}

	if doPanic {
		panic("Value not found")
	}

	return 0
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func getPatternForNumber(digits []digit, number int) string {
	for _, digit := range digits {
		if digit.value == number {
			return digit.pattern
		}
	}

	panic("Pattern not found")
}

func digitsContains(digits []digit, number int) bool {
	for _, dig := range digits {
		if dig.value == number {
			return true
		}
	}

	return false
}
