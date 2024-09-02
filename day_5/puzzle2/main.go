package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//open the file containing the stringsof santa...
	file, err := os.ReadFile("../strings.txt")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	sentences := string(file)
	//count the number of nice strings...
	fmt.Printf("This time santa got %v nice sentences.\n", betterModel(sentences))
}

func betterModel(sentences string) int {
	var niceCount int = 0
	for _, sentence := range strings.Fields(sentences) {
		if hasTwoAdjacentPairs(sentence) {
			if checkForweirdCombo(sentence) {
				niceCount++
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return niceCount
}

//Function to check if there is existence at least 2 2-character substring in the string...
func hasTwoAdjacentPairs(s string) bool {
	var sub string
	for i := 0; i < len(s)-1; i++ {
		//create a substring...
		sub = s[i : i+2]
		//check if this substring appears at least twice...
		if strings.Contains(s[(i+2):], sub) {
			return true
		} else {
			continue
		}
	}
	return false
}

//function too check for a weird combination of charcaters; like at least one letter which
//repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
func checkForweirdCombo(s string) bool {
	for i := 0; i < len(s); i++ {
		if (i+2) < len(s) && (s[i] == s[i+2]) {
			return true
		}
	}
	return false
}
