package main

import (
    "fmt"
    "strconv"
    "strings"
)

// lookAndSay performs one iteration of the look-and-say process
func lookAndSay(s string) string {
    var result strings.Builder
    i := 0
    length := len(s)

    for i < length {
        count := 1
        for i+1 < length && s[i] == s[i+1] {
            i++
            count++
        }
        result.WriteString(strconv.Itoa(count))
        result.WriteByte(s[i])
        i++
    }

    return result.String()
}

// getLengthAfterIterations applies the look-and-say process `iterations` times
func getLengthAfterIterations(initialString string, iterations int) int {
    result := initialString
    for i := 0; i < iterations; i++ {
        result = lookAndSay(result)
    }
    return len(result)
}

func main() {
    inputString := "1113122113"
    iterations := 40
    lengthOfResult := getLengthAfterIterations(inputString, iterations)
    fmt.Printf("The length of the result after %d iterations is: %d\n", iterations, lengthOfResult)
}
