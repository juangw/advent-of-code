package main

import (
	"fmt"
	"io/ioutil"
	"strings"
  "unicode"
)

func main() {
	rucksack, err := ioutil.ReadFile("day_3.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}
	trimmedRucksack := strings.TrimSpace(string(rucksack))

  var totalSum = 0
  rucksackItems := strings.Split(trimmedRucksack, "\n")
	for i := 0; i < len(rucksackItems); i++ {
	  uniqueMap := make(map[string]int)
    var rucksackItem = rucksackItems[i]
		var splitIndex = len(rucksackItem) / 2
		var firstCompartment = rucksackItem[0:splitIndex]
		var secondCompartment = rucksackItem[splitIndex:]
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
  fmt.Println(totalSum)
}
