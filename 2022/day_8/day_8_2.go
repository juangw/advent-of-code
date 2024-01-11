package day_8

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func RunPart2(logger *log.Logger) {
	treeLines, err := ioutil.ReadFile("./2022/day_8/day_8.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	splitTreeLinesSlice := strings.Split(string(treeLines), "\n")
	treeLinesSlice := splitTreeLinesSlice[:len(splitTreeLinesSlice)-1]
	scenicScoreMap := make(map[string]int)
	maxScenicScore := math.MinInt
	for row, treeLine := range treeLinesSlice {
		if row == len(treeLine)-1 || row == 0 {
			continue
		}
		for col, treeHeight := range treeLine {
			if col == len(treeLine)-1 || col == 0 {
				continue
			}

			height := int(treeHeight - '0')

			// scan left
			leftScenicScore := 0
			for i := col - 1; i >= 0; i-- {
				if int(treeLinesSlice[row][i]-'0') < height {
					leftScenicScore = leftScenicScore + 1
				} else {
					leftScenicScore = leftScenicScore + 1
					break
				}
			}

			// scan right
			rightScenicScore := 0
			for i := col + 1; i < len(treeLine); i++ {
				if int(treeLinesSlice[row][i]-'0') < height {
					rightScenicScore = rightScenicScore + 1
				} else {
					rightScenicScore = rightScenicScore + 1
					break
				}
			}

			// scan up
			upScenicScore := 0
			for j := row - 1; j >= 0; j-- {
				if int(treeLinesSlice[j][col]-'0') < height {
					upScenicScore = upScenicScore + 1
				} else {
					upScenicScore = upScenicScore + 1
					break
				}
			}

			// scan down
			downScenicScore := 0
			for j := row + 1; j < len(treeLinesSlice); j++ {
				if int(treeLinesSlice[j][col]-'0') < height {
					downScenicScore = downScenicScore + 1
				} else {
					downScenicScore = downScenicScore + 1
					break
				}
			}

			scenicScore := leftScenicScore * rightScenicScore * upScenicScore * downScenicScore
			scenicScoreMap[fmt.Sprintf("%v,%v", row, col)] = scenicScore
			maxScenicScore = max(maxScenicScore, scenicScore)
		}
	}

	logger.Println(maxScenicScore)
}
