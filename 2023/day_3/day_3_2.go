package day_3

import (
	"log"
	"unicode"

	"github.com/juangw/advent-of-code/cast"
	"github.com/juangw/advent-of-code/files"
	"github.com/juangw/advent-of-code/math"
)

type gear struct {
	row    int
	column int
}

func RunPart2(logger *log.Logger) {
	lines := files.ReadFile("./2023/day_3/day_3.txt")

	// build list of locations of all numbers & gears
	allNumbers := make(map[int][]number)
	var allGears []gear
	for row, line := range lines {
		digit := ""
		for column, char := range line {
			if unicode.IsDigit(char) {
				digit += string(char)
				// check & handle if end of line
				if column == len(line)-1 {
					columnStart := column - len(digit) + 1
					allNumbers[row] = append(allNumbers[row], number{digit, row, columnStart})
				}
			} else if !unicode.IsDigit(char) && digit != "" {
				columnStart := column - len(digit)
				allNumbers[row] = append(allNumbers[row], number{digit, row, columnStart})
				// reset digit value
				digit = ""
			}

			if char == '*' {
				allGears = append(allGears, gear{row, column})
			}
		}
	}

	// now that we know where all numbers are lets check if any gears
	// are adjacent to two of them
	var validNumbers []int
	log.Print(allGears)
	logger.Println(allNumbers)
	for _, gear := range allGears {
		var gearRatio []int

		// scan if number in row above
		rowAbove := gear.row - 1
		_, okAbove := allNumbers[rowAbove]
		if rowAbove >= 0 && okAbove {
			for _, num := range allNumbers[rowAbove] {
				if num.columnStart <= gear.column+1 && gear.column-1 <= (num.columnStart+len(num.value)-1) {
					gearRatio = append(gearRatio, cast.ToInt(num.value))
				}
			}
		}

		// scan if number in row below
		rowBelow := gear.row + 1
		_, okBelow := allNumbers[rowBelow]
		if rowBelow >= 0 && okBelow {
			for _, num := range allNumbers[rowBelow] {
				if num.columnStart <= gear.column+1 && gear.column-1 <= (num.columnStart+len(num.value)-1) {
					gearRatio = append(gearRatio, cast.ToInt(num.value))
				}
			}
		}

		// scan if number in column to the left
		leftColumn := gear.column - 1
		if leftColumn >= 0 {
			for _, num := range allNumbers[gear.row] {
				if (num.columnStart + len(num.value)) == leftColumn {
					gearRatio = append(gearRatio, cast.ToInt(num.value))
				}
			}
		}

		// scan if number in column to the right
		rightColumn := gear.column + 1
		if rightColumn < len(lines[0]) {
			for _, num := range allNumbers[gear.row] {
				if num.columnStart == rightColumn {
					gearRatio = append(gearRatio, cast.ToInt(num.value))
				}
			}
		}

		logger.Println(gearRatio)
		if len(gearRatio) == 2 {
			validNumbers = append(validNumbers, math.MultiplySlice(gearRatio))
		}
	}

	logger.Println(math.SumSlice(validNumbers))
}
