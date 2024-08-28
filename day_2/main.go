package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//open the file and read as a string...
	file, err := os.ReadFile("dimensions.txt")
	if err != nil {
		log.Fatal(err)
	}
	//calculate the total surface area required for the wrapping paper...
	dimensions := string(file)
	wrapperPaperSA(dimensions)
	fmt.Printf("The elves need %v square feet of wrapping paper.\n", wrapperPaperSA(dimensions))
}

func wrapperPaperSA(dimensions string) int {
	var (
		length            int
		width             int
		height            int
		wantedWrapperArea int
		smallestArea      int
	)

	//Seperating the dimensions in terms of their seperation...
	for i, dimension := range strings.Fields(dimensions) {
		length, _ = strconv.Atoi(string(strings.Split(dimension, "x")[0]))
		width, _ = strconv.Atoi(string(strings.Split(dimension, "x")[1]))
		height, _ = strconv.Atoi(string(strings.Split(dimension, "x")[2]))
		surfaceArea := (2*(length*width) + 2*(width*height) + 2*(height*length))
		area1 := int(length) * int(width)
		area2 := int(width) * int(height)
		area3 := int(height) * int(length)

		smallestArea = area1
		if area2 < smallestArea {
			smallestArea = area2
		}
		if area3 < smallestArea {
			smallestArea = area3
		}
		wantedWrapperArea += surfaceArea + smallestArea
		i++
		continue
	}
	return wantedWrapperArea
}
