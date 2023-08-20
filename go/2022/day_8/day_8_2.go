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
			var leftScenicScore = 1
			for i := 0; i < col; i++ {
				if int(treeLinesSlice[row][i]-'0') < height {
					leftScenicScore = leftScenicScore + 1
				} else {
					break
				}
			}

			// scan right
			var rightScenicScore = 1
			for i := col; i < len(treeLine); i++ {
				if col == i {
					continue
				} else if int(treeLinesSlice[row][i]-'0') < height {
					rightScenicScore = rightScenicScore + 1
				} else {
					break
				}
			}

			// scan up
			var upScenicScore = 1
			for j := 0; j < row; j++ {
				if int(treeLinesSlice[j][col]-'0') < height {
					upScenicScore = upScenicScore + 1
				} else {
					break
				}
			}

			// scan down
			var downScenicScore = 1
			for j := row; j < len(treeLinesSlice); j++ {
				if row == j {
					continue
				} else if int(treeLinesSlice[j][col]-'0') < height {
					downScenicScore = downScenicScore + 1
				} else {
					break
				}
			}
			scenicScore := leftScenicScore * rightScenicScore * upScenicScore * downScenicScore
			scenicScoreMap[fmt.Sprintf("%v,%v", row, col)] = scenicScore
			maxScenicScore = max(maxScenicScore, scenicScore)
		}
	}

	// logger.Println(scenicScoreMap)
	logger.Println(maxScenicScore)
}
