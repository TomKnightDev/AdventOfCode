package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lanternFish struct {
	age   int
	count int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), ",")
	var fish = [9]lanternFish{}

	for _, line := range stringLines {
		val, err := strconv.Atoi(line)
		check(err)
		fish[val].count++
	}

	// 80 day loop
	for i := 0; i < 256; i++ {
		zeroCount := fish[0].count

		fish[0].count = fish[1].count
		fish[1].count = fish[2].count
		fish[2].count = fish[3].count
		fish[3].count = fish[4].count
		fish[4].count = fish[5].count
		fish[5].count = fish[6].count
		fish[6].count = fish[7].count + zeroCount
		fish[7].count = fish[8].count
		fish[8].count = zeroCount
		// fmt.Println(fish)
	}

	total := 0
	for _, lf := range fish {
		total += lf.count
	}

	fmt.Print(total)
}

// func main() {
// 	data, err := os.ReadFile("input.txt")
// 	check(err)

// 	stringLines := strings.Split(string(data), ",")

// 	var lanternFish = []int8{}

// 	for _, line := range stringLines {
// 		val, err := strconv.Atoi(line)
// 		check(err)
// 		lanternFish = append(lanternFish, int8(val))
// 	}

// 	// 80 day loop
// 	for i := 0; i < 256; i++ {
// 		count := len(lanternFish)
// 		for lf := 0; lf < count; lf++ {
// 			lanternFish[lf] -= 1
// 			if lanternFish[lf] < 0 {
// 				lanternFish = append(lanternFish, 8)
// 				lanternFish[lf] = 6
// 			}
// 		}
// 	}

// 	fmt.Print(len(lanternFish))
// }
