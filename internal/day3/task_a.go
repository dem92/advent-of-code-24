package day3

import (
	"aoc24/internal/utils"
	"log"
	"regexp"
)

func Day3a() {
	fileContent := utils.ReadFile("3", false)
	total := 0

	for _, line := range fileContent {
		instructions := findInstructions(line)

		for _, instruction := range instructions {
			total += calculate(instruction)
		}
	}

	log.Printf("Total: %d", total)
}

func findInstructions(line string) []string {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(line, -1)
	return matches
}

func calculate(instruction string) int {
	r := regexp.MustCompile(`\d+`)
	digits := r.FindAllString(instruction, -1)
	return utils.ConvertStringToNumber(digits[0]) * utils.ConvertStringToNumber(digits[1])
}
