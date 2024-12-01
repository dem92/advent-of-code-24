package day1

import (
	"aoc24/internal/utils"
	"log"
	"math"
	"sort"
	"strings"
)

func Day1a() {
	var listLeft = []int{}
	var listRight = []int{}
	fileContent := utils.ReadFile("1", false)

	for _, line := range fileContent {
		left, right := findNumbers(line)

		listLeft = append(listLeft, left)
		listRight = append(listRight, right)
	}

	sort.Ints(listLeft)
	sort.Ints(listRight)

	total := getTotalDistance(listLeft, listRight)
	log.Printf("Total: %d", total)
}

func findNumbers(line string) (int, int) {
	stringNums := strings.Split(line, "   ")
	left := utils.ConvertStringToNumber(stringNums[0])
	right := utils.ConvertStringToNumber(stringNums[1])

	return left, right
}

func getTotalDistance(listLeft []int, listRight []int) int {
	total := 0

	for i := 0; i < len(listLeft); i++ {
		diff := (listLeft[i] - listRight[i])
		total += int(math.Abs(float64(diff)))
	}

	return total
}
