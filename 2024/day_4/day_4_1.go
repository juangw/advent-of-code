package day_4

import (
	"github.com/juangw/advent-of-code/files"
	"log"
	"strings"
)

func RunPart1(logger *log.Logger) {
	wordSearch := files.ReadFile("./2024/day_4/day_4.txt")

	xmasLength := 4
	lineCount := 0
	columnCount := 0
	rightDiagonalCount := 0
	leftDiagonalCount := 0
	totalXmas := 0

	// Check xmas by line
	for _, wordLine := range wordSearch {
		xmasCount := strings.Count(wordLine, "XMAS")
		samxCount := strings.Count(wordLine, "SAMX")
		lineCount += xmasCount + samxCount
	}

	// Check xmas by column
	for i := 0; i < len(wordSearch[0]); i++ {
		wordColumn := ""
		for j := 0; j < len(wordSearch); j++ {
			wordColumn += string(wordSearch[j][i])
		}
		xmasCount := strings.Count(wordColumn, "XMAS")
		samxCount := strings.Count(wordColumn, "SAMX")
		columnCount += xmasCount + samxCount
	}

	// check xmas by diagonal
	for row := 0; row < len(wordSearch); row++ {
		for column := 0; column < len(wordSearch); column++ {

			wordRightDiagonal := ""
			if xmasLength <= (row+1) && (column+xmasLength) <= len(wordSearch[0]) {
				rowCopy := row
				columnCopy := column
				for k := 0; rowCopy-k >= 0 && (columnCopy+k) < len(wordSearch[0]) && k < xmasLength; k++ {
					wordRightDiagonal += string(wordSearch[rowCopy-k][columnCopy+k])
				}
				xmasCount := strings.Count(wordRightDiagonal, "XMAS")
				samxCount := strings.Count(wordRightDiagonal, "SAMX")
				rightDiagonalCount += xmasCount + samxCount
			}

			wordDiagonalUpLeft := ""
			if xmasLength <= (row+1) && ((column+1)-xmasLength) >= 0 {
				rowCopy := row
				columnCopy := column
				for k := 0; rowCopy-k >= 0 && (columnCopy-k) >= 0 && k < xmasLength; k++ {
					wordDiagonalUpLeft += string(wordSearch[rowCopy-k][columnCopy-k])
				}
				xmasCount := strings.Count(wordDiagonalUpLeft, "XMAS")
				samxCount := strings.Count(wordDiagonalUpLeft, "SAMX")
				leftDiagonalCount += xmasCount + samxCount
			}
		}
	}

	totalXmas = lineCount + columnCount + rightDiagonalCount + leftDiagonalCount
	logger.Println(totalXmas)
}
