package day3

import (
	"github.com/juangw/advent-of-code/files"
	"log"
	"regexp"
)

func RunPart2(logger *log.Logger) {
	corruptedMemory := files.ReadFile("./2024/day_3/day_3.txt")

	totalMultiply := 0
	instructionsEnabled := true

	for _, line := range corruptedMemory {
		var lineTotal int
		lineTotal, instructionsEnabled = calculateMemWithEnablement(instructionsEnabled, line)
		totalMultiply += lineTotal
	}

	logger.Println(totalMultiply)
}

func calculateMemWithEnablement(instructionsEnabled bool, line string) (int, bool) {
	result := 0
	nextStartIndex := 0
	enablementRe := regexp.MustCompile(`do\(\).*?|don't\(\).*?`)
	enablementMatches := enablementRe.FindAllStringIndex(line, -1)
	for _, enablementMatch := range enablementMatches {
		startIndex, endIndex := enablementMatch[0], enablementMatch[1]
		if instructionsEnabled {
			result += calculateMem(line[nextStartIndex:startIndex])
		}
		action := line[startIndex:endIndex]
		switch action {
		case "do()":
			if instructionsEnabled == false {
				instructionsEnabled = true
			}
		case "don't()":
			if instructionsEnabled == true {
				instructionsEnabled = false
			}
		}
		nextStartIndex = endIndex
	}
	if instructionsEnabled {
		result += calculateMem(line[nextStartIndex:])
	}
	return result, instructionsEnabled
}
