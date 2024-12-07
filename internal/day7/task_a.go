package day7

import (
	"aoc24/internal/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func Day7a() {
	fileContent := utils.ReadFile("7", false)
	total := 0

	for _, line := range fileContent {
		splitLine := strings.Split(line, ": ")
		testValue := utils.ConvertStringToNumber(splitLine[0])
		numbers := utils.FindNumbers(splitLine[1], " ")

		// fmt.Println(line)
		total += testLine(testValue, numbers)
	}

	log.Printf("Total: %d", total)
}

func testLine(testValue int, numbers []int) int {
	operators := len(numbers) - 1
	possibleCombinations := math.Pow(2, float64(operators))

	for i := 0; i < int(possibleCombinations); i++ {
		n := int64(i)

		binaryNumber := strconv.FormatInt(n, 2)

		for j := len(binaryNumber); j < operators; j++ {
			binaryNumber = "0" + binaryNumber
		}

		// fmt.Println(binaryNumber)
		total := numbers[0]

		for i, c := range binaryNumber {
			total = calclulate(total, numbers[i+1], c)
		}

		if total == testValue {
			// fmt.Println("Total matches test value")
			return total
		}
	}

	return 0
}

func calclulate(a, b int, operator rune) int {
	switch operator {
	case '0':
		return a + b
	default:
		return a * b
	}
}
