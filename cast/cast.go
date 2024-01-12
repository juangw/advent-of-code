package cast

import (
	"log"
	"strconv"
)

func ToInt(char string) int {
	v, err := strconv.Atoi(char)
	if err != nil {
		log.Fatalf("unable to convert char value: %v to int", char)
	}
	return v
}
