package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	secretKey := "bgvyzdsv" // Updated secret key
	number := findAdventCoin(secretKey)
	fmt.Printf("The lowest number for the secret key '%s' with six leading zeroes is %d.\n", secretKey, number)
}

// findAdventCoin finds the lowest number that, when appended to the secret key,
// produces an MD5 hash that starts with at least six zeroes.
func findAdventCoin(secretKey string) int {
	number := 0
	for {
		// Create the input string by combining the secret key and the current number
		input := secretKey + strconv.Itoa(number)

		// Compute the MD5 hash of the input string
		hash := md5.Sum([]byte(input))
		hashHex := fmt.Sprintf("%x", hash)

		// Check if the hash starts with at least six zeroes
		if strings.HasPrefix(hashHex, "000000") {
			return number
		}

		// Increment the number and try again
		number++
	}
}
