package day_2

import (
	"log"
	"strings"

	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"github.com/juangw/advent-of-code/math"
)

func RunPart2(logger *log.Logger) {
	lines := files.ReadFile("./2023/day_2/day_2.txt")

	var validGames []int
	for _, game := range lines {
		s := strings.Split(game, ":")
		_, gameDetails := s[0], s[1]
		gameDetailsSlice := strings.FieldsFunc(gameDetails, Split)

		var cubeColorsTracker = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, cubesSet := range gameDetailsSlice {
			cubesSetTrimmed := strings.Trim(cubesSet, " ")
			cubesDetails := strings.Split(cubesSetTrimmed, " ")
			countString, colorString := cubesDetails[0], cubesDetails[1]
			count := cast.ToInt(countString)
			if count > cubeColorsTracker[colorString] {
				cubeColorsTracker[colorString] = count
			}
		}
		validGames = append(validGames, math.MultiplyMapValues(cubeColorsTracker))
	}

	logger.Println(math.SumSlice(validGames))
}
