package day5

import (
	"strings"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

//SolvePart1 returns the answer for part 1 of day 5
func SolvePart1() int {
	return PolymerLength(adventInput())
}

//SolvePart2 returns the answer for part 2 of day 5
func SolvePart2() int {
	return ShortestPolymerWithoutOneUnitLength(adventInput())
}

//PolymerLength return length of polymer after all reactions
func PolymerLength(input string) int {
	polymerAfterReactions := removeReactions(input, 0)
	return len(polymerAfterReactions)
}

/*
ShortestPolymerWithoutOneUnitLength return length of the polymer
that is the shortest after removing all units of one type
*/
func ShortestPolymerWithoutOneUnitLength(input string) int {
	min := len(input)

	for _, unitToRemove := range alphabet() {
		polymerToTest := removeUnit(input, unitToRemove)
		lengthAfterReactions := len(removeReactions(polymerToTest, 0))

		if lengthAfterReactions < min {
			min = lengthAfterReactions
		}
	}

	return min
}

func removeReactions(polymer string, startFrom int) string {
	for i := startFrom; i < len(polymer)-1; i++ {
		if polymer[i] != polymer[i+1] && theSameLetter(polymer[i], polymer[i+1]) {
			newStartFrom := i - 1
			if newStartFrom < 0 {
				newStartFrom = 0
			}
			return removeReactions(polymer[:i]+polymer[i+2:], newStartFrom)
		}
	}

	return polymer
}

func theSameLetter(letter1, letter2 byte) bool {
	return strings.ToLower(string(letter1)) == strings.ToLower(string(letter2))
}

func alphabet() []string {
	return []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
}

func removeUnit(polymer, letterToRemove string) string {
	polymerWithRemovedUnit := ""

	for _, v := range polymer {
		if !theSameLetter(byte(v), letterToRemove[0]) {
			polymerWithRemovedUnit += string(v)
		}
	}

	return polymerWithRemovedUnit
}

func adventInput() string {
	lines, _ := utils.ReadLinesFromFile("day5/input")
	return lines[0]
}
