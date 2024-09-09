package main

import (
	"fmt"
)

func main() {
	originalMap := map[string]int{
		"ii": 1,
		"kl": 2,
		"kr": 3,
		"kt": 4,
		"jk": 5,
		"jo": 6,
		"e":  7,
		"f":  8,
		"g":  9,
		"bs": 10,
		"bt": 11,
	}

	newMap := createMapWithZeroValues(originalMap)
	fmt.Println("Original Map:", originalMap)
	fmt.Println("New Map:", newMap)
}

