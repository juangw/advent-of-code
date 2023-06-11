package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var repeatingStringLen = 14

func containsDuplicate(strings []string) bool {
	disct := make(map[string]bool, len(strings))

	for _, str := range strings {
		if disct[str] {
			return true
		} else {
			disct[str] = true
		}
	}
	return false

}

func main() {
	packets, err := ioutil.ReadFile("day_6.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}

	var packetLines = strings.Split(string(packets), "\n")
	var uniqueChars = make([]string, 0, repeatingStringLen)
	for i := 0; i < len(packetLines); i++ {
		for j := 0; j < len(packetLines[i]); j++ {
			var char = string(packetLines[i][j])
			var foundDuplicate = false

			if len(uniqueChars) < repeatingStringLen {
				uniqueChars = append(uniqueChars, char)
			}

			if len(uniqueChars) == repeatingStringLen {
				// iterate through uniqueChars for duplicates
				if containsDuplicate(uniqueChars) {
					foundDuplicate = true
				}
			}

			if foundDuplicate {
				uniqueChars = uniqueChars[1:]
			} else if len(uniqueChars) == repeatingStringLen {
				fmt.Println(j + 1)
				break
			}
		}
	}
	fmt.Println(uniqueChars)
}
