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

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	stringLines := strings.Split(string(data), "\n")

	length := len(stringLines[0])
	var oxygen = stringLines
	var co2 = stringLines

	//Oxygen
	oxRating := 0
	for i := 0; i < length; i++ {
		mostCommonBit, _ := getCommonBits(oxygen, i)
		oxygen = getAllLinesWithSameBit(oxygen, i, mostCommonBit)

		if len(oxygen) == 1 {
			oxRating = convertToInt(oxygen[0])
		}
	}

	//CO2
	co2Rating := 0
	for i := 0; i < length; i++ {
		_, leastCommonBit := getCommonBits(co2, i)
		co2 = getAllLinesWithSameBit(co2, i, leastCommonBit)

		if len(co2) == 1 {
			co2Rating = convertToInt(co2[0])
		}
	}

	fmt.Println(oxRating * co2Rating)
}

func getCommonBits(vals []string, index int) (mostCommon, leastCommon int) {
	total := 0
	for _, val := range vals {
		runes := []rune(val)
		lineVal, err := strconv.Atoi(string(runes[index]))
		check(err)
		total += lineVal
	}

	if total >= len(vals)/2 {
		mostCommon = 1
		leastCommon = 0
	} else {
		mostCommon = 0
		leastCommon = 1
	}

	return
}

func getAllLinesWithSameBit(vals []string, index int, bit int) []string {
	var valsToReturn = []string{}

	for _, val := range vals {
		runes := []rune(val)
		bitAtIndex := string(runes[index])
		if bitAtIndex == fmt.Sprint(bit) {
			valsToReturn = append(valsToReturn, val)
		}
	}

	return valsToReturn
}

func convertToInt(val string) int {
	runes := []rune(val)
	valToReturn := 0
	count := 1
	for i := len(runes) - 1; i >= 0; i-- {
		if string(runes[i]) == fmt.Sprint(1) {
			valToReturn += int(math.Pow(float64(count), 1))
		}

		count = count * 2
	}

	return valToReturn
}
