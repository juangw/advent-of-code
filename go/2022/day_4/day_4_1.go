package day_4

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func RunPart1(logger *log.Logger) {
	campCleanup, err := ioutil.ReadFile("./2022/day_4/day_4.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}
	trimmedCampCleanup := strings.TrimSpace(string(campCleanup))

	totalOverlapCount := 0
	campCleanupAssignments := strings.Split(trimmedCampCleanup, "\n")
	for i := 0; i < len(campCleanupAssignments); i++ {
		replacer := strings.NewReplacer("-", " ", ",", " ")
		campCleanupCleaned := replacer.Replace(campCleanupAssignments[i])
		campCleanupCleanedSlice := strings.Split(campCleanupCleaned, " ")

		// convert array of strings to array of ints
		var campCleanupInts [4]int
		for j := 0; j < len(campCleanupCleanedSlice); j++ {
			k, err := strconv.Atoi(campCleanupCleanedSlice[j])
			if err != nil {
				logger.Fatalf("Cannot convert value to int: %v", err)
			}
			campCleanupInts[j] = k
		}

		if campCleanupInts[0] <= campCleanupInts[2] && campCleanupInts[1] >= campCleanupInts[3] {
			totalOverlapCount++
		} else if campCleanupInts[2] <= campCleanupInts[0] && campCleanupInts[3] >= campCleanupInts[1] {
			totalOverlapCount++
		}
	}

	logger.Println(totalOverlapCount)
}
