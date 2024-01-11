package day_1

import (
	"log"
	"strings"
	"unicode"

	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
)

var numberAsLettersMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func RunPart2(logger *log.Logger) {
	lines := files.ReadFile("./2023/day_1/day_1.txt")

	// iterate through each char and check if parseable as digit and
	// build calibration values for each line and sum to total
	total := 0
	for _, line := range lines {
		lineDigits := []string{}
		lineTotal := 0
		charIndex := 0
		for charIndex < len(line) {
			char := rune(line[charIndex])
			if unicode.IsDigit(char) {
				lineDigits = append(lineDigits, string(char))
				charIndex += 1
			} else {
				found := false
				for k, v := range numberAsLettersMap {
					if strings.HasPrefix(line[charIndex:], k) {
						lineDigits = append(lineDigits, v)
						charIndex += 1
						found = true
						break
					}
				}
				if found == false {
					charIndex += 1
				}
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
