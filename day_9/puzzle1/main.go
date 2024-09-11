package main

import (
	"fmt"
//	"math"
)

// Location distances
var distances = map[string]map[string]int{
	"Tristram":      {"AlphaCentauri": 34, "Snowdin": 100, "Tambi": 63, "Faerun": 108, "Norrath": 111, "Straylight": 89, "Arbre": 132},
	"AlphaCentauri": {"Tristram": 34, "Snowdin": 4, "Tambi": 79, "Faerun": 44, "Norrath": 147, "Straylight": 133, "Arbre": 74},
	"Snowdin":       {"Tristram": 100, "AlphaCentauri": 4, "Tambi": 105, "Faerun": 95, "Norrath": 48, "Straylight": 88, "Arbre": 7},
	"Tambi":         {"Tristram": 63, "AlphaCentauri": 79, "Snowdin": 105, "Faerun": 68, "Norrath": 134, "Straylight": 107, "Arbre": 40},
	"Faerun":        {"Tristram": 108, "AlphaCentauri": 44, "Snowdin": 95, "Tambi": 68, "Norrath": 11, "Straylight": 66, "Arbre": 144},
	"Norrath":       {"Tristram": 111, "AlphaCentauri": 147, "Snowdin": 48, "Tambi": 134, "Faerun": 11, "Straylight": 115, "Arbre": 135},
	"Straylight":    {"Tristram": 89, "AlphaCentauri": 133, "Snowdin": 88, "Tambi": 107, "Faerun": 66, "Norrath": 115, "Arbre": 127},
	"Arbre":         {"Tristram": 132, "AlphaCentauri": 74, "Snowdin": 7, "Tambi": 40, "Faerun": 144, "Norrath": 135, "Straylight": 127},
}

// Function to find all permutations of locations
func permute(locations []string) [][]string {
	var result [][]string
	var generate func([]string, int)
	generate = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				generate(arr, n-1)
				if n%2 == 1 {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				}
			}
		}
	}
	generate(locations, len(locations))
	return result
}

// Function to calculate the distance of a route
func routeDistance(route []string) int {
	totalDistance := 0
	for i := 0; i < len(route)-1; i++ {
		loc1, loc2 := route[i], route[i+1]
		totalDistance += distances[loc1][loc2]
	}
	return totalDistance
}

func main() {
	// List of locations
	locations := []string{"Tristram", "AlphaCentauri", "Snowdin", "Tambi", "Faerun", "Norrath", "Straylight", "Arbre"}

	// Generate all permutations of locations
	routes := permute(locations)

	// Initialize variables for longest route
	longestDistance := 0
	var longestRoute []string

	// Loop through each route and calculate the distance
	for _, route := range routes {
		distance := routeDistance(route)
		if distance > longestDistance {
			longestDistance = distance
			longestRoute = route
		}
	}

	// Output the result
	fmt.Printf("The longest route is %v with a distance of %d\n", longestRoute, longestDistance)
}