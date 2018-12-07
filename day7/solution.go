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
	edges := parseEdges(input)
	graph := utils.MakeGraph(26)

	for _, e := range edges {
		graph.AddEdge(e.U, e.V)
	}

	return graph.SortTopologically()
}

//TimedSortTopologically returns time to sort topologically
func TimedSortTopologically(input []string, workerCount int) int {
	edges := parseEdges(input)
	graph := utils.MakeGraph(26)

	for _, e := range edges {
		graph.AddEdge(e.U, e.V)
	}

	return graph.TimedSortTopologically(workerCount)
}

func parseEdges(input []string) []utils.Edge {
	edges := make([]utils.Edge, len(input))

	for i, v := range input {
		u := string(v[5])
		v := string(v[36])

		edges[i] = utils.Edge{U: u, V: v}
	}

	return edges
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day7/input")
	return lines
}
