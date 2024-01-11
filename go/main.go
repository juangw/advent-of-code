package main

import (
	"github.com/juangw/advent-of-code/go/2022/day_1"
	"github.com/juangw/advent-of-code/go/2022/day_10"
	"github.com/juangw/advent-of-code/go/2022/day_2"
	"github.com/juangw/advent-of-code/go/2022/day_3"
	"github.com/juangw/advent-of-code/go/2022/day_4"
	"github.com/juangw/advent-of-code/go/2022/day_5"
	"github.com/juangw/advent-of-code/go/2022/day_6"
	"github.com/juangw/advent-of-code/go/2022/day_7"
	"github.com/juangw/advent-of-code/go/2022/day_8"
	"github.com/juangw/advent-of-code/go/2022/day_9"
	day_1_2023 "github.com/juangw/advent-of-code/go/2023/day_1"
	"log"
	"os"
)

func main() {
	yearArg := os.Args[1]

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	if yearArg == "2022" {
		day_1.RunPart1(logger)
		day_1.RunPart2(logger)
		day_2.RunPart1(logger)
		day_2.RunPart2(logger)
		day_3.RunPart1(logger)
		day_3.RunPart2(logger)
		day_4.RunPart1(logger)
		day_4.RunPart2(logger)
		day_5.RunPart1(logger)
		day_5.RunPart2(logger)
		day_6.RunPart1(logger)
		day_6.RunPart2(logger)
		day_7.RunPart1(logger)
		day_7.RunPart2(logger)
		day_8.RunPart1(logger)
		day_8.RunPart2(logger)
		day_9.RunPart1(logger)
		day_9.RunPart2(logger)
		day_10.RunPart1(logger)
	} else if yearArg == "2023" {
		day_1_2023.RunPart1(logger)
	}
}
