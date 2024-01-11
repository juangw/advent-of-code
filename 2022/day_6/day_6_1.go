package day_6

import (
	"io/ioutil"
	"log"
	"strings"
)

func RunPart1(logger *log.Logger) {
	packets, err := ioutil.ReadFile("./2022/day_6/day_6.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	packetLines := strings.Split(string(packets), "\n")
	uniqueChars := make([]string, 0, 4)
	for i := 0; i < len(packetLines); i++ {
		for j := 0; j < len(packetLines[i]); j++ {
			char := string(packetLines[i][j])
			foundDuplicate := false

			if len(uniqueChars) < 4 {
				uniqueChars = append(uniqueChars, char)
			}

			if len(uniqueChars) == 4 {
				// iterate through uniqueChars for duplicates
				if containsDuplicate(uniqueChars) {
					foundDuplicate = true
				}
			}

			if foundDuplicate {
				uniqueChars = uniqueChars[1:]
			} else if len(uniqueChars) == 4 {
				logger.Println(j + 1)
				break
			}
		}
	}
}
