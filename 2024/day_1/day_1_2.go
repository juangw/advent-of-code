package day_1

import (
    "github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
    "strings"
	"log"
)

func RunPart2(logger *log.Logger) {
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

    similarityScoreMap := make(map[int]int)
    for _, leftValue:= range leftSlice {
        similarityScore := 0

        if _, ok := similarityScoreMap[leftValue]; ok {
            total += similarityScoreMap[leftValue]
            continue
        }

        for _, rightValue := range rightSlice {
            if leftValue == rightValue {
                similarityScore += leftValue
            }
        }

        total += similarityScore
        similarityScoreMap[leftValue] = similarityScore
    }

	logger.Println(total)
}

