package day3

import (
	"aoc24/internal/utils"
	"log"
	"regexp"
)

func Day3b() {
	fileContent := utils.ReadFile("3", false)
	total := 0
	do := true

	for _, line := range fileContent {
		instructions := findInstructionsB(line)

		for _, instruction := range instructions {
			switch instruction {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				if do {
					total += calculate(instruction)
				}
			}
		}
	}

	log.Printf("Total: %d", total)
}

func findInstructionsB(line string) []string {
	r := regexp.MustCompile(`(mul\(\d+,\d+\))|do\(\)|don't\(\)`)
	matches := r.FindAllString(line, -1)
	return matches
}
