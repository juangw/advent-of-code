package day_2

var numberOfCubesTotal = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Split(r rune) bool {
	return r == ',' || r == ';'
}
