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

	var treeLinesSlice = strings.Split(string(treeLines), "\n")
	var visibleTrees = len(treeLinesSlice) * 2
	visibleTrees = visibleTrees + ((len(treeLinesSlice[0]) - 2) * 2)
	for row, treeLine := range treeLinesSlice {
		if row == len(treeLine)-1 || row == 0 {
			continue
		}
		for col, treeHeight := range treeLine {
			if col == len(treeLine)-1 || col == 0 {
				continue
			}
			var height = int(treeHeight - '0')
			fmt.Println(height)
		}
	}
	fmt.Println(visibleTrees)
}
