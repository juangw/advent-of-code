package day_2

import (
	"io/ioutil"
	"log"
	"strings"
)

func RunPart2(logger *log.Logger) {
	strategyBytes, err := ioutil.ReadFile("./2022/day_2/day_2.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	// build round of RPS games
	totalScore := 0
	strategyString := strings.TrimSpace(string(strategyBytes))
	strategyRounds := strings.Split(strategyString, "\n")
	for i := 0; i < len(strategyRounds); i++ {
		roundSelections := strings.Split(strategyRounds[i], " ")
		endResult, opponentSelection := roundSelections[1], roundSelections[0]
		gameResult := endingRpsMap[endResult]
		roundScore, error := calculateScoreForYourUserFromResult(gameResult, opponentSelection)
		if error == nil {
			totalScore = totalScore + roundScore
		} else {
			logger.Fatal(error)
		}
	}

	logger.Println(totalScore)
}
