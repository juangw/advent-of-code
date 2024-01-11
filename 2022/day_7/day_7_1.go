package day_7

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func RunPart1(logger *log.Logger) {
	commands, err := ioutil.ReadFile("./2022/day_7/day_7.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	commandLines := strings.Split(string(commands), "$")
	directoryStorage := make(map[string]int)
	currentDirectory := []string{}
	for _, commandLine := range commandLines {
		commands := strings.Split(string(commandLine), "\n")
		if len(commands) > 1 {
			command := strings.TrimSpace(commands[0])
			splitCommand := strings.Split(command, " ")
			if splitCommand[0] != "cd" && splitCommand[0] != "ls" {
				continue
			}

			// determine directory structure
			if splitCommand[0] == "cd" {
				if splitCommand[1] != ".." {
					if splitCommand[1] == "/" {
						splitCommand[1] = ""
					}
					currentDirectory = append(currentDirectory, fmt.Sprintf("%s/", splitCommand[1]))
				} else {
					currentDirectory = currentDirectory[:len(currentDirectory)-1]
				}
			}

			// if ls command, calculate current directory size
			directoryFileSizeTotal := 0
			if splitCommand[0] == "ls" {
				lsOutputs := commands[1:]
				for _, lsOutput := range lsOutputs {
					parsedLsOutput := strings.Split(lsOutput, " ")
					maybeFileSize, maybeFileSizeErr := strconv.Atoi(parsedLsOutput[0])
					if maybeFileSizeErr != nil {
						continue
					}
					directoryFileSizeTotal = directoryFileSizeTotal + maybeFileSize
				}
			}

			// distribute current directory size to current & parents directories
			for i := range currentDirectory {
				partialDirectory := strings.Join(currentDirectory[:i+1], "")
				directoryStorage[partialDirectory] = directoryStorage[partialDirectory] + directoryFileSizeTotal
			}
		}
	}

	// sum directories with size < 100000
	solution := 0
	for _, val := range directoryStorage {
		if val <= 100000 {
			solution = solution + val
		}
	}

	logger.Println(solution)
}
