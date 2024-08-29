package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//open the file and read as a string...
	file, err := os.ReadFile("directions.txt")
	if err != nil {
		log.Fatal(err)
	}
	directions := string(file)
	// Test examples
	fmt.Printf("Santa visited %v homes more than once!\n", countUniqueHouses(directions)) // Output: 2
}

// Function to count the number of unique houses visited
func countUniqueHouses(directions string) int {
	var (
		visitedHomes []string
		startingHome string
		newHome      string
		// count        int
	)
	// Start at the origin (0, 0)
	x, y := 0, 0

	// keep the coordinate in a slice
	startingHome = fmt.Sprintf("%d,%d", x, y)
	visitedHomes = append(visitedHomes, startingHome)

	// Iterate over each direction character
	for _, direction := range directions {
		switch direction {
		case '^':
			y += 1
		case 'v':
			y -= 1
		case '>':
			x += 1
		case '<':
			x -= 1
		}

		// Mark the new position...
		newHome = fmt.Sprintf("%d,%d", x, y)

		//add all homes to visited homes...
		visitedHomes = append(visitedHomes, newHome)
		continue
	}

	//loop through the visited homes and count the occurrence of each visit...
	visitedHomesMap := make(map[string]int)
	for i, home := range visitedHomes {
		visitedHomesMap[home]++
		i++
		continue
	}

	//count the number of occurrences more than 1...
	//count = countValuesGreaterThanOne(visitedHomesMap)
	// Return the number of unique houses visited
//return count
	return len(visitedHomesMap)
}

// func countValuesGreaterThanOne(myMap map[string]int) int {
// 	count := 0

// 	// Loop over the map
// 	for _, value := range myMap {

// 		if value > 1 {
// 			count++
// 			continue
// 		}
// 		continue
// 	}

// 	return count
// }
