package day1

import (
	"container/ring"
	"strconv"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

//SolvePart1 returns the answer for part 1 of day 1
func SolvePart1() int {
	return SumStrings(adventInput())
}

//SolvePart1 returns the answer for part 2 of day 1
func SolvePart2() int {
	return RepeatingSum(adventInput())
}

//SumStrings returns the sum of all elements in the array in the form of '+x' '-y'
func SumStrings(strings []string) int {
	result := 0
	for _, v := range intsSlice(strings) {
		result += v
	}

	return result
}

//RepeatingSum returns the first repeating sum over a cyclic list in the form of '+x' '-y'
func RepeatingSum(strings []string) int {
	sum := 0
	reached := map[int]int{0: 1}
	cycled := sliceToRing(intsSlice(strings))

	for {
		sum += cycled.Value.(int)
		cycled = cycled.Next()
		if reached[sum] == 1 {
			return sum
		}

		reached[sum] = 1
	}
}

func sliceToRing(slice []int) *ring.Ring {
	r := ring.New(len(slice))
	for _, v := range slice {
		r.Value = v
		r = r.Next()
	}
	return r
}

func intsSlice(stringSlice []string) []int {
	converted := make([]int, len(stringSlice))
	for i, v := range stringSlice {
		convertedToInt, _ := strconv.Atoi(v)
		converted[i] = convertedToInt
	}

	return converted
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day1/input")
	return lines
}
