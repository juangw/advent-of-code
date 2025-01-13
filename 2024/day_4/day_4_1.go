package day_4

import (
	"github.com/juangw/advent-of-code/files"
	"log"
	"strings"
)

func RunPart1(logger *log.Logger) {
	wordSearch := files.ReadFile("./2024/day_4/day_4.txt")

	totalXmas := 0

	// Check xmas by line
	for _, wordLine := range wordSearch {
		containsXmas := strings.Contains(wordLine, "XMAS") || strings.Contains(wordLine, "SAMX")
		if containsXmas {
			totalXmas++
		}
	}

	// Check xmas by column
	for i := 0; i < len(wordSearch[0]); i++ {
		wordColumn := ""
		for j := 0; j < len(wordSearch); j++ {
			wordColumn += string(wordSearch[j][i])
		}
		containsXmas := strings.Contains(wordColumn, "XMAS") || strings.Contains(wordColumn, "SAMX")
		if containsXmas {
			totalXmas++
		}
	}

	// check xmas by diagonal
	for i := 0; i < len(wordSearch); i++ {
		wordDiagonalUpRight := ""
		for j := 0; j < len(wordSearch); j++ {
			if i+j < len(wordSearch) {
				wordDiagonalUpRight += string(wordSearch[i+j][j])
			}
		}
		containsXmas := strings.Contains(wordDiagonalUpRight, "XMAS") || strings.Contains(wordDiagonalUpRight, "SAMX")
		if containsXmas {
			totalXmas++
		}

		wordDiagonalDownRight := ""
		for j := 0; j < len(wordSearch); j++ {
			if i-j >= 0 {
				wordDiagonalDownRight += string(wordSearch[i-j][j])
			}
		}
		containsXmas = strings.Contains(wordDiagonalDownRight, "XMAS") || strings.Contains(wordDiagonalDownRight, "SAMX")
		if containsXmas {
			totalXmas++
		}
	}

	logger.Println(totalXmas)
}
