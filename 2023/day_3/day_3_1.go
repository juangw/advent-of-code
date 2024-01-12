package day_3

import (
	"log"
	"unicode"

	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"github.com/juangw/advent-of-code/math"
)

type number struct {
	value       string
	row         int
	columnStart int
}

func RunPart1(logger *log.Logger) {
	lines := files.ReadFile("./2023/day_3/day_3.txt")

	var allNumbers []number
	for row, line := range lines {
		digit := ""
		for column, char := range line {
			if unicode.IsDigit(char) {
				digit += string(char)
				// check & handle if end of line
				if column == len(line)-1 {
					columnStart := column - len(digit) + 1
					allNumbers = append(allNumbers, number{digit, row, columnStart})
				}
			} else if !unicode.IsDigit(char) && digit != "" {
				columnStart := column - len(digit)
				allNumbers = append(allNumbers, number{digit, row, columnStart})
				// reset digit value
				digit = ""
			}
		}
	}

	// now that we know where all numbers are lets check if any symbols
	// are adjacent to them
	var validNumbers []int
	logger.Println(allNumbers)
	for _, num := range allNumbers {
		// scan in row above
		rowAbove := num.row - 1
		isValidAbove := false
		if rowAbove >= 0 {
			for i := num.columnStart - 1; i <= num.columnStart+len(num.value); i++ {
				if i < 0 || i >= len(lines[0]) {
					continue
				}
				if lines[rowAbove][i] != '.' {
					isValidAbove = true
					break
				}
			}
		}

		// scan in row below
		rowBelow := num.row + 1
		isValidBelow := false
		if rowBelow < len(lines) {
			for i := num.columnStart - 1; i <= num.columnStart+len(num.value); i++ {
				if i < 0 || i >= len(lines[0]) {
					continue
				}
				if lines[rowBelow][i] != '.' {
					isValidBelow = true
					break
				}
			}
		}

		// scan in column to the left
		leftColumn := num.columnStart - 1
		isValidLeft := false
		if leftColumn >= 0 {
			if lines[num.row][leftColumn] != '.' {
				isValidLeft = true
			}
		}

		// scan in column to the right
		rightColumn := num.columnStart + len(num.value)
		isValidRight := false
		if rightColumn < len(lines[0]) {
			if lines[num.row][rightColumn] != '.' {
				isValidRight = true
			}
		}

		if isValidAbove || isValidBelow || isValidLeft || isValidRight {
			validNumbers = append(validNumbers, cast.ToInt(num.value))
			logger.Printf("valid %v", num.value)
		} else {
			logger.Printf("invalid %v", num.value)
		}
	}

	logger.Println(math.SumSlice(validNumbers))
}
