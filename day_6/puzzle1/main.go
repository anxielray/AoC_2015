package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//open the file and pass the file content...
	file, err := os.ReadFile("../directions.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}
	instructions := string(file)
	fmt.Printf("The number of light bulbs lit is %d\n", computeFunctions(instructions))
}

//function to compute the functions together...
func computeFunctions(instructions string) int {
	grid := makeGrid() //make a grid...
	myGridMap := createLightMap(grid)
	for _, instructionLine := range strings.Split(instructions, "\n") {
		switch strings.Fields(instructionLine)[0] {
		case "toggle":
			myGridMap = toggle(instructionLine, myGridMap)
		case "turn":
			if strings.Fields(instructionLine)[1] == "on" {
				myGridMap = turnOnGrid(instructionLine, myGridMap)
			} else if strings.Fields(instructionLine)[1] == "off" {
				myGridMap = turnOffGrid(instructionLine, myGridMap)
			}
		default:
			fmt.Println("Invalid instruction:", instructionLine)
		}
	}
	return countLihtBulbs(myGridMap)
}

//function to create a 2D grid of 1000x1000 cells...
func makeGrid() [][]string {
	var (
		grid [][]string
		row  []string
	)
	for x := 0; x < 1000; x++ {
		row = make([]string, 1000)
		for y := 0; y < 1000; y++ {
			row[y] = fmt.Sprintf("%s, %s", strconv.Itoa(x), strconv.Itoa(y))
		}
		grid = append(grid, row)
	}
	return grid
}

//handling the toggle instruction...
func turnOnGrid(instruction string, myGridMap map[string]bool) map[string]bool {
	startingCoordintate := strings.Fields(instruction)[2]
	lastCoordinate := strings.Fields(instruction)[4]
	startX, _ := strconv.Atoi(strings.Split(startingCoordintate, ",")[0])
	startY, _ := strconv.Atoi(strings.Split(startingCoordintate, ",")[1])
	endX, _ := strconv.Atoi(strings.Split(lastCoordinate, ",")[0])
	endY, _ := strconv.Atoi(strings.Split(lastCoordinate, ",")[1])

	//lop over the map and update their boolean values...
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			cell := fmt.Sprintf("%s, %s", strconv.Itoa(i), strconv.Itoa(j))
			myGridMap = lightTheBulb(cell, myGridMap)
		}
	}
	return myGridMap
}

//creat a map to store the lit and unlit bulbs...
func createLightMap(grid [][]string) map[string]bool {
	lightMap := make(map[string]bool)
	for _, row := range grid {
		for _, cell := range row {
			lightMap[cell] = false
		}
	}
	return lightMap
}

//function that will find the key in the map and update the value to true...
func lightTheBulb(coordinate string, myGridMap map[string]bool) map[string]bool {
	myGridMap[coordinate] = true
	return myGridMap
}

//create a function that will be turning off the light bulbs...
func turnOffTheBulb(coordinate string, myGridMap map[string]bool) map[string]bool {
	myGridMap[coordinate] = false
	return myGridMap
}

//turning off the bulbs in the grid...
func turnOffGrid(instruction string, myGridMap map[string]bool) map[string]bool {
	startingCoordintate := strings.Fields(instruction)[2]
	lastCoordinate := strings.Fields(instruction)[4]
	startX, _ := strconv.Atoi(strings.Split(startingCoordintate, ",")[0])
	startY, _ := strconv.Atoi(strings.Split(startingCoordintate, ",")[1])
	endX, _ := strconv.Atoi(strings.Split(lastCoordinate, ",")[0])
	endY, _ := strconv.Atoi(strings.Split(lastCoordinate, ",")[1])

	//lop over the map and update their boolean values...
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			cell := fmt.Sprintf("%s, %s", strconv.Itoa(i), strconv.Itoa(j))
			myGridMap = turnOffTheBulb(cell, myGridMap)
		}
	}
	return myGridMap
}

func toggle(instruction string, myGridMap map[string]bool) map[string]bool {
	startingCoordintate := strings.Fields(instruction)[1]
	lastCoordinate := strings.Fields(instruction)[3]
	startX, _ := strconv.Atoi(strings.Split(startingCoordintate, ",")[0])
	startY, _ := strconv.Atoi(strings.Split(startingCoordintate, ",")[1])
	endX, _ := strconv.Atoi(strings.Split(lastCoordinate, ",")[0])
	endY, _ := strconv.Atoi(strings.Split(lastCoordinate, ",")[1])

	//lop over the map and update their boolean values...
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			cell := fmt.Sprintf("%s, %s", strconv.Itoa(i), strconv.Itoa(j))
			myGridMap[cell] = !myGridMap[cell]
		}
	}
	return myGridMap
}

//create a funtion to count the numer of light bulbs that are lit...
func countLihtBulbs(myGridMap map[string]bool) int {
	//loop through the map...
	count := 0
	for _, value := range myGridMap {
		if value {
			count++
		}
	}
	return count
}
