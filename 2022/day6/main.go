package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to read file: " + err.Error())
	}

	input := string(bytes)

	for i := 3; i < len(input); i++ {
		b := make(map[string]interface{})

		for j := 0; j < 14; j++ {
			_, found := b[string(input[i+j])]
			if found {
				break
			}
			b[string(input[i+j])] = nil
		}

		if len(b) == 14 {
			fmt.Println(b)
			fmt.Println(i + 14)
			break
		}
	}
}
