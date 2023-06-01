package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	campCleanup, err := ioutil.ReadFile("day_4.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}
	trimmedCampCleanup := strings.TrimSpace(string(campCleanup))

	var totalOverlapCount = 0
	campCleanupAssignments := strings.Split(trimmedCampCleanup, "\n")
	for i := 0; i < len(campCleanupAssignments); i++ {
		replacer := strings.NewReplacer("-", " ", ",", " ")
		campCleanupCleaned := replacer.Replace(campCleanupAssignments[i])
		campCleanupCleanedSlice := strings.Split(campCleanupCleaned, " ")

		// convert array of strings to array of ints
		var campCleanupInts []int
		for j := 0; j < len(campCleanupCleanedSlice); j++ {
			k, err := strconv.Atoi(campCleanupCleanedSlice[j])
			if err != nil {
				fmt.Errorf("Cannot convert value to int: %v", err)
			}
			campCleanupInts = append(campCleanupInts, k)
		}

		if campCleanupInts[0] <= campCleanupInts[2] && campCleanupInts[1] >= campCleanupInts[3] {
			totalOverlapCount++
		} else if campCleanupInts[2] <= campCleanupInts[0] && campCleanupInts[3] >= campCleanupInts[1] {
			totalOverlapCount++
		}
	}
	fmt.Println(totalOverlapCount)
}
