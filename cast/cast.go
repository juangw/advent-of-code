package cast 

import (
	"io/ioutil"
    "fmt"
	"strconv"
)

func ReadFile(filename string) []byte {
	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("unable to read file: %v", err)
	}
	return lines
}

func toInt(char string) int {
	v, err := strconv.Atoi(char)
	if err != nil {
		fmt.Print("unable to convert char value to int")
	}
	return v
}
