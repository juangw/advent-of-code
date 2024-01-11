package day_2

import (
	"errors"
)

var opponentRpsMap = map[string]int{
	"A": 1, // rock
	"B": 2, // papper
	"C": 3, // scissors
}

var yourRpsMap = map[string]int{
	"Y": 2, // paper
	"X": 1, // rock
	"Z": 3, // scissors
}

var endingRpsMap = map[string]string{
	"Y": "D", // draw
	"X": "L", // lose
	"Z": "W", // win
}

func calculateScoreForYourUserFromYourMove(yourRps int, opponentRps int) (int, error) {
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

func calculateScoreForYourUserFromResult(gameResult string, opponentSelection string) (int, error) {
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
