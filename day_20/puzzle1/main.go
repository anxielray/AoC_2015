package main

import (
	"fmt"
)

func main() {
	puzzleInput := 34000000
	houseNumber := 0
	presentsThreshold := puzzleInput / 10
	houses := make([]int, presentsThreshold)

	for elf := 1; elf < presentsThreshold; elf++ {
		for house := elf; house < presentsThreshold; house += elf {
			houses[house] += elf * 10
		}
	}

	for i := 1; i < presentsThreshold; i++ {
		if houses[i] >= puzzleInput {
			houseNumber = i
			break
		}
	}

	fmt.Printf("The lowest house number to get at least %d presents is: %d\n", puzzleInput, houseNumber)
}
