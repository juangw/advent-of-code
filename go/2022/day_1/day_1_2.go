package day_1

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func RunPart2(logger *log.Logger) {
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
				logger.Fatalf("Unparse-able value: %v", err)
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

	// slice top 3 calories carried
	sort.Ints(calorieCarried)
	calorieCarriedTop3 := calorieCarried[len(calorieCarried)-3:]

	// sum up top 3 calories carried
	top3CaloriesCarriedTotal := 0
	for i := 0; i < len(calorieCarriedTop3); i++ {
		top3CaloriesCarriedTotal = top3CaloriesCarriedTotal + calorieCarriedTop3[i]
	}
	logger.Println(top3CaloriesCarriedTotal)
}
