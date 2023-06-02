package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	packets, err := ioutil.ReadFile("day_6.txt")
	if err != nil {
		fmt.Errorf("unable to read file: %v", err)
	}

	var packetLines = strings.Split(string(packets), "\n")
	var uniqueChars = make([]string, 4, 4)
	for i := 0; i < len(packetLines); i++ {
		fmt.Println(packetLines[i])
		for j := 0; j < len(packetLines[i]); j++ {
			uniqueChars = append(uniqueChars, string(packetLines[i][j]))
		}
	}
	fmt.Println(uniqueChars)
}
