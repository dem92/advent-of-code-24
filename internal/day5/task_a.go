package day5

import (
	"aoc24/internal/utils"
	"log"
	"strings"
)

func Day5a() {
	fileContent := utils.ReadFile("5", false)
	total := 0

	rules := []string{}
	updates := []string{}
	breakPointReached := false

	for _, line := range fileContent {
		if !breakPointReached {
			if line == "" {
				breakPointReached = true
				continue
			}

			rules = append(rules, line)
			continue
		}

		updates = append(updates, line)
	}

	validMiddlePages, _ := verifyUpdates(rules, updates)

	for _, page := range validMiddlePages {
		total += utils.ConvertStringToNumber(page)
	}

	log.Printf("Total: %d", total)
}

func verifyUpdates(rules, updates []string) (middlePages, invalid []string) {
	validMiddlePages := []string{}
	invalidUpdates := []string{}

	for _, update := range updates {
		isValid := true
		pages := strings.Split(update, ",")

		for pi, page := range pages {
			requiredPages := getRequiredPages(rules, page)

			if !validatePage(pages, requiredPages, pi) {
				isValid = false
				break
			}
		}

		if isValid {
			middlePage := pages[len(pages)/2]
			validMiddlePages = append(validMiddlePages, middlePage)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	return validMiddlePages, invalidUpdates
}

func validatePage(pages, requiredPages []string, targetPageIndex int) bool {
	for _, requiredPage := range requiredPages {
		for i, includedPage := range pages {
			if requiredPage == includedPage && i > targetPageIndex {
				return false
			}
		}
	}

	return true
}

func getRequiredPages(rules []string, targetPage string) []string {
	required := []string{}

	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")

		if splitRule[1] == targetPage {
			required = append(required, splitRule[0])
		}
	}

	return required
}
