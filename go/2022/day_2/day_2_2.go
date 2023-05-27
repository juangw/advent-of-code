package main

import (
	"fmt"
  "errors"
	"io/ioutil"
	"strings"
)

var opponentRpsMap = map[string]int {
  "A": 1,  // rock
  "B": 2,  // papper
  "C": 3, // scissors
}

var yourRpsMap = map[string]int {
  "Y": 2,  // paper
  "X": 1,  // rock
  "Z": 3,  // scissors
}

var endingRpsMap = map[string]string {
  "Y": "D",  // draw
  "X": "L",  // lose
  "Z": "W",  // win
}

func calculateScoreForYourUser(gameResult string, opponentSelection string) (int, error) {
  opponentScore := opponentRpsMap[opponentSelection]
  if gameResult == "D" {
    return opponentScore + 3, nil
  } else if gameResult == "W" {
    if opponentScore == 3 {
      return 7, nil
    } else {
      return opponentScore + 7, nil
    }
  } else if gameResult == "L" {
    if opponentScore == 1 {
      return 3, nil
    } else {
      return opponentScore - 1, nil
    }
  }

  return -1, errors.New("Something went very wrong")
}

func main() {
	strategyBytes, err := ioutil.ReadFile("day_2.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}

  // build round of RPS games 
	var totalScore = 0
  var strategyRounds []string
	var strategyString = strings.TrimSpace(string(strategyBytes))
	strategyRounds = strings.Split(strategyString, "\n")
	for i := 0; i < len(strategyRounds); i++ {
    roundSelections := strings.Split(strategyRounds[i], " ")
    endResult, opponentSelection := roundSelections[1], roundSelections[0]
    gameResult := endingRpsMap[endResult]
    roundScore, error := calculateScoreForYourUser(gameResult, opponentSelection)
    if error == nil {
      totalScore = totalScore + roundScore
    } else {
      fmt.Println(error)
    }
  }
  fmt.Println(totalScore)
}
