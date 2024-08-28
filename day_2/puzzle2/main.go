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
	file, err := os.ReadFile("../dimensions.txt")
	if err != nil {
		log.Fatal(err)
	}
	dimensions := string(file)
	fmt.Printf("A length of %v is required by the elves\n", ribbon(dimensions))
}

func ribbon(dimensions string) int {
	var (
		length            int
		width             int
		height            int
		smallestPerimeter int
		volume            int
		wantedLength      int
	)

	//Seperating the dimensions in terms of their seperation...
	for i, dimension := range strings.Fields(dimensions) {
		length, _ = strconv.Atoi(string(strings.Split(dimension, "x")[0]))
		width, _ = strconv.Atoi(string(strings.Split(dimension, "x")[1]))
		height, _ = strconv.Atoi(string(strings.Split(dimension, "x")[2]))
		perim1 := 2 * (length + width)
		perim2 := 2 * (width + height)
		perim3 := 2 * (height + length)

		//calculate the volume of the present dimensions...
		volume = length * width * height

		//determine the smallest area to help s determine the smallest perimeter...
		smallestPerimeter = perim1
		if perim2 < smallestPerimeter {
			smallestPerimeter = perim2
		}
		if perim3 < smallestPerimeter {
			smallestPerimeter = perim3
		}
		wantedLength += smallestPerimeter + volume
		i++
		continue
	}
	return wantedLength
}
