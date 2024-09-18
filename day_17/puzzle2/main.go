package main

import (
	"fmt"
)

func main() {
	// Input: Container capacities
	containers := []int{50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 5}
	targetVolume := 150

	// Part 1: Find all valid combinations
	validCombinations := findCombinations(containers, targetVolume)

	if len(validCombinations) == 0 {
		fmt.Printf("No combinations can fit exactly %d liters\n", targetVolume)
		return
	}

	// Part 2: Find the minimum number of containers used in any valid combination
	minContainers, countMin := findMinContainerCombinations(validCombinations)

	// Output the results
	fmt.Printf("Part 1: Total combinations to store %d liters: %d\n", targetVolume, len(validCombinations))
	fmt.Printf("Part 2: Minimum containers used: %d, Ways to store %d liters using minimum containers: %d\n", minContainers, targetVolume, countMin)
}

// findCombinations returns all combinations of containers that sum to the target volume
func findCombinations(containers []int, target int) [][]int {
	var result [][]int
	helper(containers, target, []int{}, 0, &result)
	return result
}

// helper is a recursive function that finds all valid combinations
func helper(containers []int, target int, current []int, index int, result *[][]int) {
	if target == 0 {
		// If we reach the target volume, add the current combination to the result
		*result = append(*result, append([]int{}, current...))
		return
	}

	if target < 0 || index >= len(containers) {
		// If the target becomes negative or we have checked all containers, return
		return
	}

	// Include the current container in the combination
	helper(containers, target-containers[index], append(current, containers[index]), index+1, result)

	// Exclude the current container from the combination and move to the next one
	helper(containers, target, current, index+1, result)
}

// findMinContainerCombinations finds the minimum number of containers used and how many combinations use that number
func findMinContainerCombinations(combinations [][]int) (int, int) {
	if len(combinations) == 0 {
		// If there are no valid combinations, return (0,0)
		return 0, 0
	}

	minContainers := len(combinations[0])
	countMin := 0

	// Iterate through all valid combinations to find the minimum number of containers
	for _, combo := range combinations {
		if len(combo) < minContainers {
			minContainers = len(combo)
			countMin = 1 // Reset count for new minimum found
		} else if len(combo) == minContainers {
			countMin++ // Increment count for matching minimum length
		}
	}

	return minContainers, countMin
}
