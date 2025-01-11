package day_2

import (
	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"log"
	"strings"
)

func RunPart2(logger *log.Logger) {
	reports := files.ReadFile("./2024/day_2/day_2.txt")

	minDiff := 1
	maxDiff := 3

	safeReports := 0
	for _, reportLevel := range reports {
		stringLevels := strings.Split(reportLevel, " ")

		intLevels := make([]int, len(stringLevels))
		for stringLevelIndex, stringLevel := range stringLevels {
			intLevel := cast.ToInt(stringLevel)
			intLevels[stringLevelIndex] = intLevel
		}

		isSafe := isLevelSafe(intLevels, minDiff, maxDiff)
		if isSafe {
			safeReports++
		} else {
			// Brute force removing one level at a time
			for intLevelIndex := range intLevels {
				intLevelsCopy := make([]int, len(intLevels))
				copy(intLevelsCopy, intLevels)
				newIntLevels := append(intLevelsCopy[:intLevelIndex], intLevelsCopy[intLevelIndex+1:]...)
				isSafeWithRemoval := isLevelSafe(newIntLevels, minDiff, maxDiff)
				if isSafeWithRemoval {
					safeReports++
					break
				}
			}
		}
	}

	logger.Println(safeReports)
}
