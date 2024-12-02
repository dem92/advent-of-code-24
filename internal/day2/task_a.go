package day2

import (
	"aoc24/internal/utils"
	"log"
	"math"
	"strings"
)

func Day2a() {
	fileContent := utils.ReadFile("2", false)
	safeReportCount := 0

	for _, line := range fileContent {
		report := findNumbers(line)

		if isSafe := checkReport(report); isSafe {
			safeReportCount++
		}
	}

	log.Printf("Total: %d", safeReportCount)
}

func findNumbers(line string) []int {
	stringNums := strings.Split(line, " ")
	nums := []int{}

	for _, sNum := range stringNums {
		nums = append(nums, utils.ConvertStringToNumber(sNum))
	}

	return nums
}

func checkReport(report []int) bool {
	prevLevel := report[0]
	isIncreasing := report[0] < report[1]

	for i := 1; i < len(report); i++ {
		currentLevel := report[i]
		diff := int(math.Abs(float64(currentLevel - prevLevel)))

		if currentLevel == prevLevel {
			return false
		}
		if isIncreasing && currentLevel < prevLevel {
			return false
		}
		if !isIncreasing && currentLevel > prevLevel {
			return false
		}
		if diff > 3 {
			return false
		}

		prevLevel = currentLevel
	}

	return true
}
