package day_3

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func RunPart1(logger *log.Logger) {
	rucksack, err := ioutil.ReadFile("./2022/day_3/day_3.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}
	trimmedRucksack := strings.TrimSpace(string(rucksack))

	totalSum := 0
	rucksackItems := strings.Split(trimmedRucksack, "\n")
	for i := 0; i < len(rucksackItems); i++ {
		uniqueMap := make(map[string]int)
		rucksackItem := rucksackItems[i]
		splitIndex := len(rucksackItem) / 2
		firstCompartment := rucksackItem[0:splitIndex]
		secondCompartment := rucksackItem[splitIndex:]
		for i := 0; i < len(firstCompartment); i++ {
			_, ok := uniqueMap[string(firstCompartment[i])]
			if !ok {
				uniqueMap[string(firstCompartment[i])] = 1
			}
		}
		for i := 0; i < len(secondCompartment); i++ {
			secondCompartmentChar := rune(secondCompartment[i])
			_, ok := uniqueMap[string(secondCompartment[i])]
			if ok {
				if unicode.IsLower(secondCompartmentChar) {
					totalSum = totalSum + int(secondCompartmentChar) - 96
				} else {
					totalSum = totalSum + int(secondCompartmentChar) - 38
				}
				break
			}
		}
		uniqueMap = make(map[string]int)
	}

	logger.Println(totalSum)
}
