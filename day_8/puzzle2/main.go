package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Extract the content from the file
	file, err := os.ReadFile("../strings.txt") // Ensure the file path is correct
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	content := string(file)
	fmt.Printf("The difference of the equation is %v\n", computeDifference(content))
}

// Function to compute the difference between encoded characters and original code characters
func computeDifference(content string) int {
	totalCodeChars := 0
	totalEncodedChars := 0

	// Process each line in the content
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}

		// Count the number of characters in code
		codeChars := len(line)

		// Encode the string
		encodedString := encodeString(line)
		encodedChars := len(encodedString)

		totalCodeChars += codeChars
		totalEncodedChars += encodedChars
	}
	return totalEncodedChars - totalCodeChars
}

// Function to encode the string according to the specified rules
func encodeString(line string) string {
	// Start with surrounding double quotes
	encoded := "\""

	// Process escape sequences
	for _, char := range line {
		switch char {
		case '"':
			encoded += "\\\"" // Escape double quotes
		case '\\':
			encoded += "\\\\" // Escape backslashes
		default:
			encoded += string(char) // Add other characters as is
		}
	}

	encoded += "\"" // Close the surrounding double quotes
	return encoded
}
