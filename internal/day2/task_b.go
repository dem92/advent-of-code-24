package day2

import (
	"aoc24/internal/utils"
	"log"
)

func Day2b() {
	fileContent := utils.ReadFile("2", false)
	safeReportCount := 0

	for _, line := range fileContent {
		report := findNumbers(line)

		if isSafe := checkReport(report); isSafe {
			safeReportCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			// Make a copy of report, but without the level at index i
			redactedReport := []int{}
			redactedReport = append(redactedReport, report[:i]...)
			redactedReport = append(redactedReport, report[i+1:]...)

			if isSafe := checkReport(redactedReport); isSafe {
				safeReportCount++
				break
			}
		}
	}

	log.Printf("Total: %d", safeReportCount)
}
