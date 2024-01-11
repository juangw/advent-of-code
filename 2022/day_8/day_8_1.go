package day_8

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func RunPart1(logger *log.Logger) {
	treeLines, err := ioutil.ReadFile("./2022/day_8/day_8.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	splitTreeLinesSlice := strings.Split(string(treeLines), "\n")
	treeLinesSlice := splitTreeLinesSlice[:len(splitTreeLinesSlice)-1]
	visibleTreesLength := len(treeLinesSlice) * 2
	visibleTreesWidth := (len(treeLinesSlice[0]) - 2) * 2
	visibleTrees := visibleTreesLength + visibleTreesWidth
	innerVisibleTreeMap := make(map[string]int)
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
			for i := 0; i < col; i++ {
				if int(treeLinesSlice[row][i]-'0') >= height {
					break
				}

				if i == col-1 {
					innerVisibleTreeMap[fmt.Sprintf("%v,%v", row, col)] = 0
				}
			}

			// scan right
			for i := col + 1; i < len(treeLine); i++ {
				if int(treeLinesSlice[row][i]-'0') >= height {
					break
				}

				if i == len(treeLine)-1 {
					innerVisibleTreeMap[fmt.Sprintf("%v,%v", row, col)] = 0
				}
			}

			// scan up
			for j := 0; j < row; j++ {
				if int(treeLinesSlice[j][col]-'0') >= height {
					break
				}

				if j == row-1 {
					innerVisibleTreeMap[fmt.Sprintf("%v,%v", row, col)] = 0
				}
			}

			// scan down
			for j := row + 1; j < len(treeLinesSlice); j++ {
				if int(treeLinesSlice[j][col]-'0') >= height {
					break
				}

				if j == len(treeLinesSlice)-1 {
					innerVisibleTreeMap[fmt.Sprintf("%v,%v", row, col)] = 0
				}
			}
		}
	}

	logger.Println(visibleTrees + len(innerVisibleTreeMap))
}
