package utils

import "strings"

func findNumbers(line string, separator string) []int {
	stringNums := strings.Split(line, separator)
	numbers := []int{}

	for _, stringNum := range stringNums {
		numbers = append(numbers, ConvertStringToNumber(stringNum))
	}

	return numbers
}
