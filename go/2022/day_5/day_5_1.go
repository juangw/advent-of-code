package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	supplyStacks, err := ioutil.ReadFile("day_5.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}

	supplyStacksSlice := strings.Split(string(supplyStacks), "\n\n")
	supplyStacksGraph, supplyStacksMoves := supplyStacksSlice[0], supplyStacksSlice[1]

	supplyStacksGraphRows := strings.Split(supplyStacksGraph, "\n")
	// build map of stacks
	var mapStack = make(map[int][]string)
	for i := 0; i < len(supplyStacksGraphRows)-1; i++ {
		for j := 0; j < len(supplyStacksGraphRows[i]); j = j + 3 {
			if supplyStacksGraphRows[i][j:j+3] != "   " {
				mapStack[j/3] = append(mapStack[j/3], supplyStacksGraphRows[i][j:j+3])
			}
			if (j + 1) < len(supplyStacksGraphRows[i]) {
				j++
			}
		}
	}

	// iterate over moves & start popping and pushing values around
	trimmedMoves := strings.TrimSpace(supplyStacksMoves)
	supplyStacksMovesRows := strings.Split(trimmedMoves, "\n")
	for i := 0; i < len(supplyStacksMovesRows); i++ {
		var moves = strings.Split(string(supplyStacksMovesRows[i]), " ")
		transfers, transferFrom, transferTo := moves[1], moves[3], moves[5]
		transfersInt, err := strconv.Atoi(transfers)
		if err != nil {
			fmt.Errorf("Connot convert value to int: %v, err")
		}

		for j := 0; j < transfersInt; j++ {
			fmt.Println(transferFrom, transferTo)
		}
	}
	fmt.Println(mapStack)
}
