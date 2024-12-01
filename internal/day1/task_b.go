package day1

import (
	"aoc24/internal/utils"
	"log"
)

func Day1b() {
	var listLeft = []int{}
	var listRight = []int{}
	fileContent := utils.ReadFile("1", false)

	for _, line := range fileContent {
		left, right := findNumbers(line)

		listLeft = append(listLeft, left)
		listRight = append(listRight, right)
	}

	total := getSimilarityscore(listLeft, listRight)
	log.Printf("Total: %d", total)
}

func getSimilarityscore(listLeft []int, listRight []int) int {
	total := 0

	for _, left := range listLeft {
		count := 0

		for _, right := range listRight {
			if left == right {
				count++
			}
		}

		total += left * count
	}

	return total
}
