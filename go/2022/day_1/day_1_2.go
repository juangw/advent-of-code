package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {

	allCalories, err := ioutil.ReadFile("day_1.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}

	var separatedCalories []string
	separatedCalories = strings.Split(string(allCalories), "\n")
	var elfIndex int = 0
	elfToCaloriesCarriedMap := make(map[int]int)

	// build map of elf to calories carried
	for i := 0; i < len(separatedCalories); i++ {
		if separatedCalories[i] != "" {
			calories, err := strconv.Atoi(separatedCalories[i])
			if err != nil {
				fmt.Errorf("Unparse-able value: %v", err)
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
	fmt.Println(calorieCarriedTop3)

	// sum up top 3 calories carried
	var top3CaloriesCarriedTotal int = 0
	for i := 0; i < len(calorieCarriedTop3); i++ {
		top3CaloriesCarriedTotal = top3CaloriesCarriedTotal + calorieCarriedTop3[i]
	}
	fmt.Println(top3CaloriesCarriedTotal)
}
