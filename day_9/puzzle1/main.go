package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
He can start and end at any two (different) locations he wants, but he must visit each location exactly once.
What is the shortest distance he can travel to achieve this?
*/

func main() {
	//extracting the location from the file...
	file, err := os.ReadFile("../locations.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	locations := string(file)
	createMap(locations)
	createMapOfLocations(slice(locations))
	// fmt.Printf("The shortest route is %v\n", computeFunctions(locations))
}

// function to compute functions together...
// func computeFunctions(locations string) int {
// 	var (
// 		key      string
// 		distance int
// 	)
// 	locateMap := createMap(locations)
// 	for _, location := range strings.Split(locations, "\n") {
// 		distance, _ = strconv.Atoi(strings.Fields(location)[4])
// 		key = fmt.Sprintf("%s-%s", strings.Fields(location)[0], strings.Fields(location)[2])

// 	}
// 	return 0
// }

//function to loop over a string and make a map with all the strings, that are not "to" and are not "="
//The map keys should be like "Tristram-Faerun" where the Tristram was the first in the line and the Faerun came at index 2
func createMap(s string) map[string]int {
	var (
		m        = make(map[string]int)
		splitted = strings.Split(s, "\n")
	)
	for _, word := range splitted {
		value, _ := strconv.Atoi(strings.Fields(word)[4])
		m[fmt.Sprintf("%s-%s", strings.Fields(word)[0], strings.Fields(word)[2])] = value
	}
	return m
}

//function to contain all the locations into the same slice...
func slice(locations string) []string {
	var (
		Locations []string
	)
	split := strings.Split(locations, "\n")
	for _, location := range split {
		Locations = append(Locations, strings.Fields(location)[0])
		Locations = append(Locations, strings.Fields(location)[2])
	}
	return Locations
}

//function to create a map to check fo the occurrences of the locations in the slice of strings...
func createMapOfLocations(locations []string) map[string]int {
	var (
		m = make(map[string]int)
	)
	for _, location := range locations {
		m[location]++
	}
	fmt.Println("Final map:")
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
	return m
}

//function to extract the key from a map into a slice of string...
func extractKeys(m map[string]int) []string {
    var keys []string
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}