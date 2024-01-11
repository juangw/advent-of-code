package files

import (
	"io/ioutil"
    "fmt"
)

func ReadFile(filename string) []byte {
	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("unable to read file: %v", err)
	}
	return lines
}
