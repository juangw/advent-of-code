package day_1

import (
	"github.com/juangw/advent-of-code/go/2023/utils"
	"log"
	"strings"
)

func RunPart1(logger *log.Logger) {
	lines := utils.ReadFile("./2023/day_1/day_1.txt", logger)
	linesSlice := strings.Split(string(lines), "\n")

	// iterate through each char and check if parseable as digit and
	// build calibration values for each line and sum to total
	total := 0
	for _, line := range linesSlice {
		lineDigits := make([]string, len(line))
		lineTotal := 0
		for _, char := range line {
			utils.toInt(string(char))
			lineDigits = append(lineDigits, string(char))
		}

		if len(lineDigits) == 1 {
			lineTotal = utils.toInt(lineDigits[0] + lineDigits[0])
		} else {
			lineTotal = utils.toInt(lineDigits[0] + lineDigits[len(lineDigits)-1])
		}
		total = total + lineTotal
	}

	logger.Println(total)
}
