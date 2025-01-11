package day_2

import (
	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"github.com/juangw/advent-of-code/math"
	"log"
	"strings"
)

func RunPart1(logger *log.Logger) {
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
		}
	}

	logger.Println(safeReports)
}

func isLevelSafe(intLevels []int, minDiff int, maxDiff int) bool {
	isLevelSafe := true
	isAscending := true
	for intLevelIndex, intLevel := range intLevels {
		if intLevelIndex == 0 {
			continue
		}

		// Determine if the levels are ascending or descending
		if intLevelIndex == 1 {
			if intLevel < intLevels[intLevelIndex-1] {
				isAscending = false
			} else {
				isAscending = true
			}
		}

		if math.AbsInt(intLevel-intLevels[intLevelIndex-1]) >= minDiff && math.AbsInt(intLevel-intLevels[intLevelIndex-1]) <= maxDiff {
			if isAscending && intLevel < intLevels[intLevelIndex-1] {
				isLevelSafe = false
				break
			} else if !isAscending && intLevel > intLevels[intLevelIndex-1] {
				isLevelSafe = false
				break
			}
			continue
		} else {
			isLevelSafe = false
			break
		}
	}

	return isLevelSafe
}
