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
				mapStack[(j/4)+1] = append(mapStack[(j/4)+1], supplyStacksGraphRows[i][j:j+3])
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
		transfers, transferFromStr, transferToStr := moves[1], moves[3], moves[5]
		transfersInt, transfersErr := strconv.Atoi(transfers)
		transferFromInt, transferFromErr := strconv.Atoi(transferFromStr)
		transferToInt, transferToErr := strconv.Atoi(transferToStr)
		if transfersErr == nil || transferFromErr == nil || transferToErr == nil {
			var errSlice = [...]error{transfersErr, transferFromErr, transferToErr}
			fmt.Errorf("Connot convert value to int: (%v)", errSlice)
		}

		// execute moves onto each stack
		var popSlice = mapStack[transferFromInt]
		poppedValues, popSlice := popSlice[:transfersInt], popSlice[transfersInt:]

		// we need to make a copy because its a pointer to an array and we want to
		// persist this slice as is
		var popSliceCopy = make([]string, len(popSlice))
		copy(popSliceCopy, popSlice)

		var pushSlice = mapStack[transferToInt]
		pushSlice = append(poppedValues, pushSlice...)
		mapStack[transferFromInt] = popSliceCopy
		mapStack[transferToInt] = pushSlice
	}

	// concat top of each stack together
	var solution []string
	for i := 1; i <= len(mapStack); i++ {
		solution = append(solution, string(mapStack[i][0][1]))
	}
	fmt.Println(strings.Join(solution, ""))
}
