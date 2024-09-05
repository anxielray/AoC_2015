package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

// Function to compute the difference between code characters and in-memory characters
func computeDifference(content string) int {
	totalCodeChars := 0
	totalMemoryChars := 0

	// Process each line in the content
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}

		// Count the number of characters in code
		codeChars := len(line)

		// Remove escape sequences to get the actual string in memory
		memoryString := processEscapes(line)
		memoryChars := len(memoryString)

		totalCodeChars += codeChars
		totalMemoryChars += memoryChars
	}
	println(totalCodeChars)
	println(totalMemoryChars)
	return totalCodeChars - totalMemoryChars
}

// Process the escape sequences and hexadecimal escapes in the string
func processEscapes(line string) string {
	// Convert hexadecimal escape sequences
	line = hexadecimalEscape(line)
	// Replace escaped quotes
	line = quotationEscape(line)
	// Remove escaped backslashes
	line = dealWithEscape(line)
	// Remove surrounding double quotes
	line = line[1 : len(line)-1]

	return line
}

// Function to deal with escaped backslashes
func dealWithEscape(line string) string {
	return strings.ReplaceAll(line, `\\`, `\`)
}

// Function to convert hexadecimal escape sequences to ASCII characters
func hexadecimalEscape(line string) string {
	pattern := regexp.MustCompile(`\\x([0-9a-fA-F]{2})`)
	return pattern.ReplaceAllStringFunc(line, func(match string) string {
		hex := match[2:] // Remove the "\x" prefix
		value, err := strconv.ParseInt(hex, 16, 8)
		if err != nil {
			return match
		}
		return string(rune(value))
	})
}

// Function to replace escaped quotes with actual quotes
func quotationEscape(line string) string {
	return strings.ReplaceAll(line, `\"`, `"`)
}
