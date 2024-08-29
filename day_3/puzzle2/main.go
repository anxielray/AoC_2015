package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//open the file and read as a string...
	file, err := os.ReadFile("../directions.txt")
	if err != nil {
		log.Fatal(err)
	}
	directions := string(file)
	// Test examples
	fmt.Printf("Santa and Robo-Santa visited %v homes altogether!\n", countUniqueHouses(directions))
}

// Function to count the number of unique houses visited
func countUniqueHouses(directions string) int {
	var (
		santaVisitedHomes []string
		robotVisitedHomes []string
		santaStartingHome string
		robotStartingHome string
		santaDirection    string
		robotDirection    string
		Homes             int
		newHome           string
		visitedHomes      []string
	)

	//define every ones path..
	for i, char := range directions {
		if i%2 == 0 {
			robotDirection += string(char)
		} else {
			santaDirection += string(char)
		}
	}
	// Start at the origin (0, 0)
	x, y := 0, 0
	rx, ry := 0, 0

	// keep the coordinate in a slice...
	santaStartingHome = fmt.Sprintf("%d,%d", x, y)
	santaVisitedHomes = append(santaVisitedHomes, santaStartingHome)
	robotStartingHome = fmt.Sprintf("%d,%d", rx, ry)
	robotVisitedHomes = append(robotVisitedHomes, robotStartingHome)

	// Iterate over each direction character
	for _, santaPath := range santaDirection {
		switch santaPath {
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
		santaVisitedHomes = append(santaVisitedHomes, newHome)
		continue
	}

	for _, robotPath := range robotDirection {
		switch robotPath {
		case '^':
			ry += 1
		case 'v':
			ry -= 1
		case '>':
			rx += 1
		case '<':
			rx -= 1
		}

		// Mark the new position...
		newHome = fmt.Sprintf("%d,%d", rx, ry)

		//add all homes to visited homes...
		robotVisitedHomes = append(robotVisitedHomes, newHome)
		continue
	}

	visitedHomes = append(visitedHomes, santaVisitedHomes...)
	visitedHomes = append(visitedHomes, robotVisitedHomes...)

	//loop through the visited homes and count the occurrence of each visit...
	homesMap := make(map[string]int)
	for i, home := range visitedHomes {
		homesMap[home]++
		i++
		continue
	}
	Homes = len(homesMap)
	return Homes
}
