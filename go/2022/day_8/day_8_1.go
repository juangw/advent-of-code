package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	treeLines, err := ioutil.ReadFile("day_8.txt")
	if err != nil {
		err = fmt.Errorf("unable to read file: %v", err)
		fmt.Println(err)
	}

	var splitTreeLinesSlice = strings.Split(string(treeLines), "\n")
	var treeLinesSlice = splitTreeLinesSlice[:len(splitTreeLinesSlice)-1]
	var visibleTreesLength = len(treeLinesSlice) * 2
	var visibleTreesWidth = (len(treeLinesSlice[0]) - 2) * 2
	var visibleTrees = visibleTreesLength + visibleTreesWidth
	var innerVisibleTreeMap = make(map[string]int)
	for row, treeLine := range treeLinesSlice {
		if row == len(treeLine)-1 || row == 0 {
			continue
		}
		for col, treeHeight := range treeLine {
			if col == len(treeLine)-1 || col == 0 {
				continue
			}

			var height = int(treeHeight - '0')

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
			for i := col; i < len(treeLine); i++ {
				if col == i {
					continue
				} else if int(treeLinesSlice[row][i]-'0') >= height {
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
			for j := row; j < len(treeLinesSlice); j++ {
				if row == j {
					continue
				} else if int(treeLinesSlice[j][col]-'0') >= height {
					break
				}

				if j == len(treeLinesSlice)-1 {
					innerVisibleTreeMap[fmt.Sprintf("%v,%v", row, col)] = 0
				}
			}
		}
	}
	fmt.Println(innerVisibleTreeMap)
	fmt.Println(visibleTrees + len(innerVisibleTreeMap))
}
