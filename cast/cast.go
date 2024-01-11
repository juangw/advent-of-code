package cast

import (
	"log"
	"strconv"
)

func ToInt(char string) int {
	v, err := strconv.Atoi(char)
	if err != nil {
		log.Fatal("unable to convert char value to int")
	}
	return v
}
