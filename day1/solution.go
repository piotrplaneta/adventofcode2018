package day1

import (
	"strconv"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

//SolvePart1 returns the answer for part 1 of day 1
func SolvePart1() int {
	return SumStrings(adventInput())
}

//SumStrings returns the sum of all elements in the array in the form of '+x' '-y'
func SumStrings(strings []string) int {
	result := 0
	for _, v := range strings {
		converted, _ := strconv.Atoi(v)
		result += converted
	}

	return result
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day1/input")
	return lines
}
