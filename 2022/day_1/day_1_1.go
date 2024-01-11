package year_2022_day_1

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func RunPart1(logger *log.Logger) {
	allCalories, err := ioutil.ReadFile("./2022/day_1/day_1.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	separatedCalories := strings.Split(string(allCalories), "\n")
	elfToCaloriesCarriedMap := make(map[int]int)
	elfIndex := 0

	// build map of elf to calories carried
	for i := 0; i < len(separatedCalories); i++ {
		if separatedCalories[i] != "" {
			calories, err := strconv.Atoi(separatedCalories[i])
			if err != nil {
				log.Fatalf("Unparse-able value: %v", err)
			} else {
				if _, ok := elfToCaloriesCarriedMap[elfIndex]; ok {
					elfToCaloriesCarriedMap[elfIndex] = elfToCaloriesCarriedMap[elfIndex] + calories
				} else {
					elfToCaloriesCarriedMap[elfIndex] = calories
				}
			}
		} else {
			elfIndex = elfIndex + 1
		}
	}

	// grab all values from map
	calorieCarried := make([]int, 0, len(elfToCaloriesCarriedMap))
	for _, val := range elfToCaloriesCarriedMap {
		calorieCarried = append(calorieCarried, val)
	}
	sort.Ints(calorieCarried)

	logger.Println(calorieCarried[len(calorieCarried)-1])
}
