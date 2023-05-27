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

func calculateScoreForYourUser(yourRps int, opponentRps int) (int, error) {
  if yourRps == opponentRps {
    return yourRps + 3, nil
  } else if yourRps == 3 && opponentRps == 1 {
    return yourRps, nil
  } else if opponentRps == 3 && yourRps == 1 {
    return yourRps + 6, nil
  } else if yourRps > opponentRps {
    return yourRps + 6, nil
  } else if opponentRps > yourRps {
    return yourRps, nil
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
    yourSelection, opponentSelection := roundSelections[1], roundSelections[0]
    opponentRps := opponentRpsMap[opponentSelection]
    yourRps := yourRpsMap[yourSelection]
    roundScore, error := calculateScoreForYourUser(yourRps, opponentRps)
    if error == nil {
      totalScore = totalScore + roundScore
    } else {
      fmt.Println(error)
    }
  }
  fmt.Println(totalScore)
}
