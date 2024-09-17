package main

import (
	"bufio"
	"fmt"
	"os"
)

const gridSize = 100

// Parses the input and returns a 2D slice representing the initial grid.
func parseInput() [][]rune {
	grid := make([][]rune, gridSize)
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		grid[i] = []rune(scanner.Text())
		i++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return grid
}

// Counts the number of 'on' neighbors around a given position in the grid.
func countOnNeighbors(grid [][]rune, x, y int) int {
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0

	for _, dir := range directions {
		neighborX, neighborY := x+dir[0], y+dir[1]
		if neighborX >= 0 && neighborX < gridSize && neighborY >= 0 && neighborY < gridSize {
			if grid[neighborX][neighborY] == '#' {
				count++
			}
		}
	}
	return count
}

// Simulates one step of the grid.
func simulateStep(grid [][]rune) [][]rune {
	newGrid := make([][]rune, gridSize)
	for i := range newGrid {
		newGrid[i] = make([]rune, gridSize)
		for j := 0; j < gridSize; j++ {
			onNeighbors := countOnNeighbors(grid, i, j)
			if grid[i][j] == '#' {
				// Light stays on if it has 2 or 3 neighbors that are on.
				if onNeighbors == 2 || onNeighbors == 3 {
					newGrid[i][j] = '#'
				} else {
					newGrid[i][j] = '.'
				}
			} else {
				// Light turns on if exactly 3 neighbors are on.
				if onNeighbors == 3 {
					newGrid[i][j] = '#'
				} else {
					newGrid[i][j] = '.'
				}
			}
		}
	}
	return newGrid
}

// Counts the total number of lights that are on in the grid.
func countLightsOn(grid [][]rune) int {
	count := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	// Parse the initial grid from the input
	grid := parseInput()

	// Simulate 100 steps
	for step := 0; step < 100; step++ {
		grid = simulateStep(grid)
	}

	// Count the lights that are on after 100 steps
	fmt.Println("Number of lights on after 100 steps:", countLightsOn(grid))
}
