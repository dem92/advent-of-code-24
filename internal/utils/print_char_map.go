package utils

import "fmt"

func PrintCharMap(areaMap [][]rune) {
	for _, column := range areaMap {
		fmt.Println(string(column))
	}
	fmt.Println()
}
