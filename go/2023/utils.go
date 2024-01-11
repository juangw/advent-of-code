package utils

import (
	"io/ioutil"
	"log"
	"strconv"
)

func ReadFile(filename string, logger *log.Logger) []byte {
	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}
	return lines
}

func toInt(char string, logger *log.Logger) int {
	v, err := strconv.Atoi(char)
	if err != nil {
		logger.Print("unable to convert char value to int")
	}
	return v
}
