package day4

import (
	"aoc24/internal/utils"
	"log"
)

var masCounter = 0
var lineCountB = 0
var columnCountB = 0

func Day4b() {
	fileContent := utils.ReadFile("4", false)
	lineCountB = len(fileContent)
	columnCountB = len(fileContent[0])

	for y, line := range fileContent {
		characters := []rune(line)

		for x, char := range characters {
			if char != 'A' {
				continue
			}

			checkForMas(fileContent, x, y)
		}
	}

	log.Printf("Total: %d", masCounter)
}

func checkForMas(fileContent []string, x, y int) {
	if x == 0 || y == 0 || x == columnCountB-1 || y == lineCountB-1 {
		return
	}

	ne := []rune(fileContent[y-1])[x+1]
	se := []rune(fileContent[y+1])[x+1]
	sw := []rune(fileContent[y+1])[x-1]
	nw := []rune(fileContent[y-1])[x-1]

	if checkCharPairs(ne, sw) && checkCharPairs(se, nw) {
		masCounter++
	}
}

func checkCharPairs(dir1, dir2 rune) bool {
	if dir1 == dir2 {
		return false
	}

	if (dir1 == 'M' && dir2 == 'S') || dir1 == 'S' && dir2 == 'M' {
		return true
	}

	return false
}
