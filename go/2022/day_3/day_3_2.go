package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var numberOfCompartments = 3

func chunkSlice(slice []string, numberOfCompartments int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += numberOfCompartments {
		end := i + numberOfCompartments

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func main() {
	rucksack, err := ioutil.ReadFile("day_3.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}
	trimmedRucksack := strings.TrimSpace(string(rucksack))

	var totalSum = 0
	rucksacks := strings.Split(trimmedRucksack, "\n")
	compartments := chunkSlice(rucksacks, numberOfCompartments)
	for i := 0; i < len(compartments); i++ {
		var uniqueMap = make(map[string]int)
		for j := 0; j < numberOfCompartments; j++ {
			var currentWordUniqueMap = make(map[string]int)
			for k := 0; k < len(compartments[i][j]); k++ {
				currentCharByte := compartments[i][j][k]
				compartmentChar := rune(currentCharByte)
				curr_val, ok := uniqueMap[string(currentCharByte)]
				_, currentWordOk := currentWordUniqueMap[string(currentCharByte)]
				if !currentWordOk {
					currentWordUniqueMap[string(currentCharByte)] = 1
				}
				if !ok {
					uniqueMap[string(currentCharByte)] = 1
				} else if !currentWordOk {
					uniqueMap[string(currentCharByte)] = uniqueMap[string(currentCharByte)] + 1
					curr_val = curr_val + 1
				}
				if curr_val == 3 {
					if unicode.IsLower(compartmentChar) {
						totalSum = totalSum + int(compartmentChar) - 96
					} else {
						totalSum = totalSum + int(compartmentChar) - 38
					}
					break
				}
			}
			currentWordUniqueMap = make(map[string]int)
		}
		uniqueMap = make(map[string]int)
	}
	fmt.Println(totalSum)
}
