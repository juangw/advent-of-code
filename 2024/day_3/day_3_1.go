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

	for _, line := range corruptedMemory {
		totalMultiply += calculateMem(line)
	}

	logger.Println(totalMultiply)
}

func calculateMem(line string) int {
	result := 0
	mulRe := regexp.MustCompile(`mul\(([0-9]{0,3}),([0-9]{0,3})\)`)
	mulMatches := mulRe.FindAllStringSubmatch(line, -1)
	if len(mulMatches) > 0 {
		for _, match := range mulMatches {
			first := cast.ToInt(match[1])
			second := cast.ToInt(match[2])
			result += first * second
		}
	}
	return result
}
