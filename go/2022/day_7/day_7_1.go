package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	commands, err := ioutil.ReadFile("day_7.txt")
	if err != nil {
		err = fmt.Errorf("unable to read file: %v", err)
		fmt.Println(err)
	}

	var commandLines = strings.Split(string(commands), "$")
	var directoryStorage = make(map[string]int)
	var currentDirectory = []string{}
	for _, commandLine := range commandLines {
		var commands = strings.Split(string(commandLine), "\n")
		if len(commands) > 1 {
			var command = strings.TrimSpace(commands[0])
			var splitCommand = strings.Split(command, " ")
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

			// assume ls command, calculate current directory size
			var lsOutputs = commands[1:]
			var directoryFileSizeTotal = 0
			for _, lsOutput := range lsOutputs {
				var parsedLsOutput = strings.Split(lsOutput, " ")
				maybeFileSize, maybeFileSizeErr := strconv.Atoi(parsedLsOutput[0])
				if maybeFileSizeErr != nil {
					err = fmt.Errorf("unable to parse value to integer: %v", parsedLsOutput[0])
				}
				directoryFileSizeTotal = directoryFileSizeTotal + maybeFileSize
			}

			// distribute current directory size to current & parents directories
			for i, directory := range currentDirectory {
				if directory == "/" {
					continue
				}
				var partialDirectory = strings.Join(currentDirectory[:i+1], "")
				directoryStorage[partialDirectory] = directoryStorage[partialDirectory] + directoryFileSizeTotal
			}
		}
	}

	// sum directories with size < 100000
	var solution = 0
	for _, val := range directoryStorage {
		if val <= 100000 {
			solution = solution + val
		}
	}
	fmt.Println(solution)
}
