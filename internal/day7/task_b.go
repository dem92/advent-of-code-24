package day7

import (
	"aoc24/internal/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Day7b() {
	fileContent := utils.ReadFile("7", false)
	total := 0

	for _, line := range fileContent {
		splitLine := strings.Split(line, ": ")
		testValue := utils.ConvertStringToNumber(splitLine[0])
		numbers := utils.FindNumbers(splitLine[1], " ")

		// fmt.Println(line)
		total += testLineB(testValue, numbers)
	}

	log.Printf("Total: %d", total)
}

func testLineB(testValue int, numbers []int) int {
	aCombinations := getACombinations(numbers)
	total := 0

	for _, aComb := range aCombinations {
		res := numbers[0]

		for i, c := range aComb {
			res = calclulate(res, numbers[i+1], c)
		}

		if res == testValue {
			total = res
			break
		}
	}

	if total == 0 {
		return checkForExtraOperator(aCombinations, numbers, testValue)
	}

	return total
}

func checkForExtraOperator(aCombinations []string, failedNumbers []int, testValue int) int {
	for _, combA := range aCombinations {
		for _, combB := range aCombinations {
			characters := []rune(combB)

			for i, c := range combA {
				if c == '0' {
					continue
				}

				characters[i] = '2'
			}

			total := failedNumbers[0]

			for i, c := range characters {
				if c == '2' {
					total = utils.ConvertStringToNumber(fmt.Sprintf("%d%d", total, failedNumbers[i+1]))
					continue
				}

				total = calclulateB(total, failedNumbers[i+1], c)
			}

			if total == testValue {
				// fmt.Println("Total matches test value")
				return total
			}
		}
	}

	return 0
}

func getACombinations(numbers []int) []string {
	operators := len(numbers) - 1
	possibleCombinations := math.Pow(2, float64(operators))
	combinations := []string{}

	for i := 0; i < int(possibleCombinations); i++ {
		n := int64(i)

		binaryNumber := strconv.FormatInt(n, 2)

		for j := len(binaryNumber); j < operators; j++ {
			binaryNumber = "0" + binaryNumber
		}

		combinations = append(combinations, binaryNumber)
	}

	return combinations
}

func calclulateB(a, b int, operator rune) int {
	switch operator {
	case '0':
		return a + b
	default:
		return a * b
	}
}
