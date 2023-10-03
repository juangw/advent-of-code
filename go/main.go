package main

import (
	"go/2022/day_1"
	"go/2022/day_10"
	"go/2022/day_2"
	"go/2022/day_3"
	"go/2022/day_4"
	"go/2022/day_5"
	"go/2022/day_6"
	"go/2022/day_7"
	"go/2022/day_8"
	"go/2022/day_9"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
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
}
