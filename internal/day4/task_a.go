package day4

import (
	"aoc24/internal/utils"
	"log"
)

type direction struct {
	dir string
	x   int
	y   int
}

var directions = []direction{
	{dir: "N", x: 0, y: -1},
	{dir: "NE", x: 1, y: -1},
	{dir: "E", x: 1, y: 0},
	{dir: "SE", x: 1, y: 1},
	{dir: "S", x: 0, y: 1},
	{dir: "SW", x: -1, y: 1},
	{dir: "W", x: -1, y: 0},
	{dir: "NW", x: -1, y: -1},
}
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

			for _, dir := range directions {
				checkForXmas(fileContent, 1, x, y, dir)
			}
		}
	}

	log.Printf("Total: %d", xmasCounter)
}

func checkForXmas(fileContent []string, xmasIndex, x, y int, dir direction) {
	newX := x + dir.x
	newY := y + dir.y

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
