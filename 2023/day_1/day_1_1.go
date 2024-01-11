package day_1

import (
	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"log"
	"unicode"
)

func RunPart1(logger *log.Logger) {
	lines := files.ReadFile("./2023/day_1/day_1.txt")

	// iterate through each char and check if parseable as digit and
	// build calibration values for each line and sum to total
	total := 0
	for _, line := range lines {
		lineDigits := []string{}
		lineTotal := 0
		for _, char := range line {
			if unicode.IsDigit(char) {
				lineDigits = append(lineDigits, string(char))
			}
		}

		if len(lineDigits) == 1 {
			lineTotal = cast.ToInt(lineDigits[0] + lineDigits[0])
		} else {
			lineTotal = cast.ToInt(lineDigits[0] + lineDigits[len(lineDigits)-1])
		}
		total = total + lineTotal
	}

	logger.Println(total)
}
