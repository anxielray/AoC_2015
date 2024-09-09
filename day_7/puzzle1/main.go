package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//extract the instructions line by line...
	file, err := os.ReadFile("../bitwise-gates.txt")
	if err != nil {
		fmt.Printf("Error reading bitwise: %v\n", err)
		os.Exit(1)
	}
	bitWiseLogicGates := string(file)
	// fmt.Println("The final map is:")
	// for key, value := range computeFunctions(bitWiseLogicGates) {
	fmt.Printf("The final signal assigned to wire a is %v\n", computeFunctions(bitWiseLogicGates))
	// 	fmt.Printf("%s : %v\n", key, value)
	// }
}

/*
This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is
a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535).
A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from
one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have
a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y
to an AND gate, and then connect its output to wire z.

For example:

    123 -> x means that the signal 123 is provided to wire x.
    x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
    p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
    NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.

Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit
instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

UNDERSTANDING

- signal is the value of the bit
- wires are the lowercase alphabet charcters
- each wire carries only a single signal (value)
- one source but multiple distribution destinations.
- a gate is a vessel, it cannot pass a non-value signal.
- a gate will be defined by the bitwise operators.

GAP ANALYSIS

- create a map to hold the computed values for each wire
- start with wire a as the input signal, 0 as the initial value for all other wires
- follow the instructions to compute the values for each wire
- return the value of wire a
*/

func computeFunctions(bitWiseLogicGates string) int {
	//loop through the content to extract the instructions line by line...

	//create the map with all the wires in from the file...
	wireMap := createMapWithZeroValues(extractCharacterStrings(strings.Split(bitWiseLogicGates, "\n")))
	for _, line := range strings.Split(bitWiseLogicGates, "\n") {

		//check for the determiner of the bitwise operaation...
		if strings.Contains(line, "AND") {
			//perform the AND operation...
			wireMap = AND(line, wireMap)
		}
		if strings.Contains(line, "OR") {
			//perform the OR operation...
			wireMap = OR(line, wireMap)

		}
		if strings.Contains(line, "LSHIFT") {
			//perform the left-shift operation...
			// wireMap = AND(line, wireMap)

		}
		if strings.Contains(line, "RSHIFT") {
			//perform the right-shift operation...
			// wireMap = AND(line, wireMap)

		}
		if strings.Contains(line, "NOT") {
			//perform the NOT operation...
			wireMap = NOT(line, wireMap)
		}

		//when direct assignment is done...
		if len(strings.Fields(line)) == 3 {
			num, _ := strconv.Atoi(strings.Fields(line)[0])
			wireMap[strings.Fields(line)[2]] = num
		}

	}
	return wireMap["a"]
}

//create a function to perform the bitWise AND operation...
func AND(line string, wireMap map[string]int) map[string]int {
	valueA := strings.Fields(line)[0]
	valueB := strings.Fields(line)[2]
	benefactor := strings.Fields(line)[4]

	var (
		valueAint     int
		valueBint     int
		benefactorint int
	)
	//find the values(numbers) of the wires from the map
	for k, v := range wireMap {
		if k == valueA {
			valueAint = v
		} else if k == valueB {
			valueBint = v
		} else if k == benefactor {
			benefactorint = v
		}
	}
	benefactorint += bitwiseAnd(valueAint, valueBint)
	wireMap[benefactor] += benefactorint
	return wireMap
}

// Function to perform bitwise AND operation
func bitwiseAnd(a, b int) int {
	return a & b
}

//create a function to handle the bitWise OR operation...
func OR(line string, wireMap map[string]int) map[string]int {
	valueA := strings.Fields(line)[0]
	valueB := strings.Fields(line)[2]
	benefactor := strings.Fields(line)[4]

	var (
		valueAint     int
		valueBint     int
		benefactorint int
	)
	//find the values(numbers) of the wires from the map
	for k, v := range wireMap {
		if k == valueA {
			valueAint = v
		} else if k == valueB {
			valueBint = v
		} else if k == benefactor {
			benefactorint = v
		}
	}
	benefactorint += bitwiseOr(valueAint, valueBint)
	wireMap[benefactor] += benefactorint
	return wireMap
}

//function to handle the bitwise or...
func bitwiseOr(a, b int) int {
	return a | b
}

//create a function to handle the bitWise OR operation...
func NOT(line string, wireMap map[string]int) map[string]int {
	valueA := strings.Fields(line)[1]
	benefactor := strings.Fields(line)[3]

	var (
		valueAint     int
		benefactorint int
	)
	//find the values(numbers) of the wires from the map
	for k, v := range wireMap {
		if k == valueA {
			valueAint = v
		} else if k == benefactor {
			benefactorint = v
		}
	}
	benefactorint += bitwiseNot(valueAint)
	wireMap[benefactor] += benefactorint
	return wireMap
}

//function to perform the bitwise NOT operation...
func bitwiseNot(a int) int {
	return ^a & 0xFFFF
}

// Function to extract all lowercase character strings and their occurrences
func extractCharacterStrings(instructions []string) map[string]int {
	charMap := make(map[string]int)

	// Define regex patterns to find lowercase character strings (1 or 2 characters)
	pattern := regexp.MustCompile(`[a-z]{1,2}`)

	for _, instruction := range instructions {
		// Find all matches in the instruction
		matches := pattern.FindAllString(instruction, -1)
		for _, match := range matches {
			if len(match) > 0 {
				charMap[match]++
			}
		}
	}

	return charMap
}

// Function to create a new map with the same keys as the original map but with all values set to 0
func createMapWithZeroValues(original map[string]int) map[string]int {
	newMap := make(map[string]int)

	// Iterate over the original map and add each key to the new map with value 0
	for key := range original {
		newMap[key] = 0
	}

	return newMap
}
