package main

import (
	"github.com/juangw/advent-of-code/year_2022_day_1"
	day_1_2023 "github.com/juangw/advent-of-code/2023/day_1"
	"log"
	"os"
)

func main() {
	yearArg := os.Args[1]

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	if yearArg == "2022" {
		year_2022_day_1.RunPart1(logger)
		year_2022_day_1.RunPart2(logger)
	} else if yearArg == "2023" {
		day_1_2023.RunPart1(logger)
	}
}
