package day7

import (
	"sort"
)

//Edge represents Edgeof string Graph
type Edge struct {
	u, v string
}

//Graph represents string Graph
type Graph struct {
	n             int
	edges         map[string][]string
	reversedEdges map[string][]string
}

//ParseEdges returns edges from []string
func ParseEdges(input []string) []Edge {
	edges := make([]Edge, len(input))

	for i, v := range input {
		u := string(v[5])
		v := string(v[36])

		edges[i] = Edge{u: u, v: v}
	}

	return edges
}

//MakeGraph returns a new graph
func MakeGraph(n int) *Graph {
	return &Graph{
		n:             n,
		edges:         make(map[string][]string, n),
		reversedEdges: make(map[string][]string, n),
	}
}

//AddEdges adds all edges from the slice
func (g *Graph) AddEdges(edges []Edge) {
	for _, e := range edges {
		g.AddEdge(e)
	}
}

//AddEdge adds a new edge
func (g *Graph) AddEdge(e Edge) {
	g.edges[e.u] = append(g.edges[e.u], e.v)
	g.reversedEdges[e.v] = append(g.reversedEdges[e.v], e.u)
}

//RemoveEdgesComingFromAndReturnNewStartNodes used in topo sort
func (g *Graph) RemoveEdgesComingFromAndReturnNewStartNodes(u string) []string {
	g.edges[u] = make([]string, 0)

	newStartNodes := make([]string, 0)

	for destinationNode, sourceNodes := range g.reversedEdges {
		indexOfU := index(sourceNodes, u)

		if indexOfU != -1 {
			g.reversedEdges[destinationNode] = append(sourceNodes[:indexOfU], sourceNodes[indexOfU+1:]...)
			if len(g.reversedEdges[destinationNode]) == 0 {
				newStartNodes = append(newStartNodes, destinationNode)
			}
		}
	}

	return newStartNodes
}

//StartNodes return all nodes which does not have edge leading to
func (g *Graph) StartNodes() []string {
	startNodes := make([]string, 0)
	for node := range g.edges {
		found := false

		for destinationNode := range g.reversedEdges {
			if node == destinationNode {
				found = true
			}
		}

		if !found {
			startNodes = append(startNodes, node)
		}
	}

	sort.Strings(startNodes)
	return startNodes
}

//SortTopologically returns sorted topologically nodes
func (g *Graph) SortTopologically() string {
	result := ""
	startNodes := g.StartNodes()

	for len(startNodes) > 0 {
		startNode := startNodes[0]
		startNodes = startNodes[1:]

		result += startNode
		newStartNodes := g.RemoveEdgesComingFromAndReturnNewStartNodes(startNode)
		startNodes = append(startNodes, newStartNodes...)
		sort.Strings(startNodes)
	}

	return result
}

//TimedSortTopologically returns time to sort topologically
func (g *Graph) TimedSortTopologically(workerCount int) int {
	startNodes := g.StartNodes()
	startNodesWithTime := assignTimeToNodes(startNodes)
	timePassed := 0

	for len(startNodesWithTime) > 0 {
		if anyValueEqZero(startNodesWithTime) {
			startNode := keyOfFirstZeroValue(startNodesWithTime)
			delete(startNodesWithTime, startNode)
			newStartNodes := g.RemoveEdgesComingFromAndReturnNewStartNodes(startNode)
			newStartNodesWithTime := assignTimeToNodes(newStartNodes)
			for k, v := range newStartNodesWithTime {
				startNodesWithTime[k] = v
			}
		} else {
			timePassed++

			keys := keys(startNodesWithTime)
			sort.Strings(keys)

			n := workerCount
			if len(keys) < n {
				n = len(keys)
			}
			firstNKeys := keys[:n]

			for _, k := range firstNKeys {
				startNodesWithTime[k]--
			}
		}
	}

	return timePassed
}

func assignTimeToNodes(nodes []string) map[string]int {
	nodesWithTime := make(map[string]int)

	for _, v := range nodes {
		nodesWithTime[v] = taskTime(v)
	}

	return nodesWithTime
}

func taskTime(s string) int {
	return int(s[0]) - 4
}

func anyValueEqZero(m map[string]int) bool {
	for _, v := range m {
		if v == 0 {
			return true
		}
	}

	return false
}

func keyOfFirstZeroValue(m map[string]int) string {
	for k, v := range m {
		if v == 0 {
			return k
		}
	}

	return ""
}

func keys(m map[string]int) []string {
	keys := make([]string, 0)

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func index(array []string, elem string) int {
	index := -1

	for i, v := range array {
		if v == elem {
			index = i
		}
	}

	return index
}
