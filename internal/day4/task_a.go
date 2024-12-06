package day4

import (
	"aoc24/internal/utils"
	"log"
)

var xmasCounter = 0
var xmas = []rune{'X', 'M', 'A', 'S'}
var lineCountA = 0
var columnCountA = 0

func Day4a() {
	fileContent := utils.ReadFile("4", false)
	lineCountA = len(fileContent)
	columnCountA = len(fileContent[0])

	for y, line := range fileContent {
		characters := []rune(line)

		for x, char := range characters {
			if char != 'X' {
				continue
			}

			for _, dir := range utils.Directions {
				checkForXmas(fileContent, 1, x, y, dir)
			}
		}
	}

	log.Printf("Total: %d", xmasCounter)
}

func checkForXmas(fileContent []string, xmasIndex, x, y int, dir utils.Direction) {
	newX := x + dir.X
	newY := y + dir.Y

	if newX >= columnCountA || newX < 0 || newY >= lineCountA || newY < 0 {
		return
	}

	currentChar := []rune(fileContent[newY])[newX]

	if currentChar == xmas[xmasIndex] {
		if xmasIndex == len(xmas)-1 {
			xmasCounter++
			return
		}

		checkForXmas(fileContent, xmasIndex+1, newX, newY, dir)
	}
}
