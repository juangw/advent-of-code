package files

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(filename string) []string {
	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.Split(strings.Trim(string(lines), "\n"), "\n")
}
