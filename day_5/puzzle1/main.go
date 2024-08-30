package main

import(
    "fmt"
    "os"
    "strings"
)

func main() {
    //Read from the file and extract the content as a string...
    file, err := os.ReadFile("../strings.txt")
    if err!= nil {
        fmt.Println("Error reading the file:", err)
        return
    }
    sentences := string(file)
    fmt.Printf("Santa has %v nice sentences\n", NaughtyAndNiceString(sentences))
}

func NaughtyAndNiceString(sentences string) int {
    var (
        niceNo int
        naughtyNo int
        vowels int
        duplicates []string
        // clean bool
    )
    //extract each sentence from the sentences...
    for _, sentence := range strings.Fields(sentences) {
        //check if the sentence has at least three vowels...
        for _, char := range sentence {
            switch char {
            case 'a', 'e', 'i', 'o', 'u':
                vowels++
            }
        }
        
        //If vowels < 3, the sentence is naughty...
        if vowels >= 3 {
            //check for duplicate concurrent letters...
            for i := 0; i < len(sentence)-1; i++ {
                if sentence[i]== sentence[i+1] {
                    duplicates = append(duplicates, sentence)
                    break
                }
            }
            if len(duplicates) == 0 {
                naughtyNo++
                vowels = 0
                continue
            }
                
            //check for the occurrence of the naughty combination words...
            for _, word := range duplicates  {
                if !strings.Contains(word, "ab") && !strings.Contains(word, "cd") && !strings.Contains(word, "xy") && !strings.Contains(word, "pq") {
                    niceNo++
                    vowels = 0
                    break
                } else if strings.Contains(word, "ab") || strings.Contains(word, "cd") || strings.Contains(word, "xy") || strings.Contains(word, "pq"){
                    naughtyNo++
                    vowels = 0
                    continue
                }
            }
            duplicates = []string{}
            continue
        }else {
            naughtyNo++
            vowels = 0
            continue

        }
    }
    return niceNo
}