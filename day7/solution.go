package day7

import (
	"github.com/piotrplaneta/adventofcode2018/utils"
)

//SolvePart1 returns the answer for part 1 of day 7
func SolvePart1() string {
	return SortTopologically(adventInput())
}

//SolvePart2 returns the answer for part 2 of day 7
func SolvePart2() int {
	return TimedSortTopologically(adventInput(), 5)
}

//SortTopologically returns node names sorted topologically
func SortTopologically(input []string) string {
	edges := ParseEdges(input)
	graph := MakeGraph(26)
	graph.AddEdges(edges)

	return graph.SortTopologically()
}

//TimedSortTopologically returns time to sort topologically
func TimedSortTopologically(input []string, workerCount int) int {
	edges := ParseEdges(input)
	graph := MakeGraph(26)
	graph.AddEdges(edges)

	return graph.TimedSortTopologically(workerCount)
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day7/input")
	return lines
}
