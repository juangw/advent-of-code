package day3

import (
	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"log"
	"regexp"
)

func RunPart1(logger *log.Logger) {
	corruptedMemory := files.ReadFile("./2024/day_3/day_3.txt")

	totalMultiply := 0
	re := regexp.MustCompile(`mul\(([0-9]{0,3}),([0-9]{0,3})\)`)

	for _, line := range corruptedMemory {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			first := cast.ToInt(match[1])
			second := cast.ToInt(match[2])
			totalMultiply += first * second
		}
	}

	logger.Println(totalMultiply)
}
