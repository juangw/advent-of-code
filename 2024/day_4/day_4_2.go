package day_4

import (
	"github.com/juangw/advent-of-code/files"
	"log"
)

func RunPart2(logger *log.Logger) {
	wordSearch := files.ReadFile("./2024/day_4/day_4.txt")

	totalMas := 0

	// check xmas by diagonal
	for row := 0; row < len(wordSearch); row++ {
		for column := 0; column < len(wordSearch); column++ {

			if (row+1) < len(wordSearch) && (row-1) >= 0 && (column+1) < len(wordSearch[0]) && (column-1) >= 0 {
				rowCopy := row
				columnCopy := column

				currentChar := string(wordSearch[rowCopy][columnCopy])
				topLeft := string(wordSearch[rowCopy-1][columnCopy-1])
				topRight := string(wordSearch[rowCopy-1][columnCopy+1])
				bottomLeft := string(wordSearch[rowCopy+1][columnCopy-1])
				bottomRight := string(wordSearch[rowCopy+1][columnCopy+1])

				leftGood := false
				leftDiagonal := topLeft + currentChar + bottomRight
				if leftDiagonal == "MAS" || leftDiagonal == "SAM" {
					leftGood = true
				}

				rightGood := false
				rightDiagonal := topRight + currentChar + bottomLeft
				if rightDiagonal == "MAS" || rightDiagonal == "SAM" {
					rightGood = true
				}

				if leftGood && rightGood {
					totalMas += 1
				}
			}
		}
	}

	logger.Println(totalMas)
}
