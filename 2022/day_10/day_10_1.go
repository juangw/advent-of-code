package day_10

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func RunPart1(logger *log.Logger) {
	program, err := ioutil.ReadFile("./2022/day_10/day_10.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	cleanedProgram := strings.TrimSpace(string(program))
	programSlice := strings.Split(cleanedProgram, "\n")

	actions := map[int]int{}
	cycleIndex := 0
	for _, commandLine := range programSlice {
		if commandLine == "noop" {
			actions[cycleIndex] += 0
			continue
		}

		commandDetail := strings.Split(commandLine, " ")
		_, commandValue := commandDetail[0], commandDetail[1]
		commandValueInt, err := strconv.Atoi(commandValue)
		if err != nil {
			logger.Fatal("unable to convert program value to int")
		}

		cycleIndex += 2
		actions[cycleIndex] += commandValueInt
	}

	xValue := 1
	logger.Println(actions)
	for i := 0; i < len(actions); i++ {
		xValue += actions[i]

		if i == 19 {
			logger.Println(xValue * 20)
		}

		if i == 60 {
			logger.Println(xValue)
			logger.Println(xValue * 19)
			break
		}
	}
}
