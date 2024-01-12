package day_2

import (
	"log"
	"strings"

	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"github.com/juangw/advent-of-code/math"
)

func RunPart1(logger *log.Logger) {
	lines := files.ReadFile("./2023/day_2/day_2.txt")

	var validGames []int
	for _, game := range lines {
		s := strings.Split(game, ":")
		gameIdDetails, gameDetails := s[0], s[1]
		gameIdSplit := strings.Split(gameIdDetails, " ")
		_, gameIdString := gameIdSplit[0], gameIdSplit[1]
		gameId := cast.ToInt(gameIdString)
		gameDetailsSlice := strings.FieldsFunc(gameDetails, Split)

		isPossible := true
		for _, cubesSet := range gameDetailsSlice {
			cubesSetTrimmed := strings.Trim(cubesSet, " ")
			cubesDetails := strings.Split(cubesSetTrimmed, " ")
			countString, colorString := cubesDetails[0], cubesDetails[1]
			count := cast.ToInt(countString)
			if count > numberOfCubesTotal[colorString] {
				isPossible = false
				break
			}
		}
		if isPossible {
			validGames = append(validGames, gameId)
		}
	}

	logger.Println(math.SumSlice(validGames))
}
