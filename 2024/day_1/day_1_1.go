package day_1

import (
	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"log"
	"sort"
	"strings"
)

func RunPart1(logger *log.Logger) {
	lines := files.ReadFile("./2024/day_1/day_1.txt")

	total := 0
	totalLocations := len(lines)
	leftSlice := make([]int, totalLocations)
	rightSlice := make([]int, totalLocations)
	for lineIndex, line := range lines {
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			logger.Fatal("Invalid input string")
		}

		leftSlice[lineIndex] = cast.ToInt(split[0])
		rightSlice[lineIndex] = cast.ToInt(split[1])
	}
	sort.Sort(sort.IntSlice(leftSlice))
	sort.Sort(sort.IntSlice(rightSlice))

	for index := range leftSlice {
		difference := leftSlice[index] - rightSlice[index]
		if difference < 0 {
			difference = -difference
		}
		total += difference
	}

	logger.Println(total)
}
