package day5

import (
	"aoc24/internal/utils"
	"log"
	"sort"
	"strings"
)

func Day5b() {
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

	_, invalidUpdates := verifyUpdates(rules, updates)
	invalidMiddlePages := getinvalidMiddlePages(invalidUpdates, rules)

	for _, page := range invalidMiddlePages {
		total += utils.ConvertStringToNumber(page)
	}

	log.Printf("Total: %d", total)
}

func getinvalidMiddlePages(invalidUpdates, rules []string) []string {
	middlePages := []string{}

	for _, invalidUpdate := range invalidUpdates {
		pages := strings.Split(invalidUpdate, ",")

		sort.Slice(pages, func(i, j int) bool {
			a := pages[i]
			b := pages[j]

			requiredPages := getRequiredPages(rules, a)

			for _, reqPage := range requiredPages {
				if reqPage == b {
					return false
				}
			}

			return true
		})

		middlePage := pages[len(pages)/2]
		middlePages = append(middlePages, middlePage)
	}

	return middlePages
}
