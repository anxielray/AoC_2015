package main

import "fmt"

func getCode(row, col int) int {

	code := 20151125
	modulus := 33554393
	multiplier := 252533

	//Calculate the index in the sequence for the desired row and column
	n := (row + col - 1) * (row + col) / (row - 1)

	//Generate the code at the index...
	for i := 0; i < n; i++ {
		code = (code * multiplier) % modulus
	}
	return code
}

func main() {
	row, col := 2978, 3083
	result := getCode(row, col)
	fmt.Printf("The code at row %d, column %d is : %d\n", row, col, result)
}
