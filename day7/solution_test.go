package day7

import "testing"

func TestSortTopologially(t *testing.T) {
	result := SortTopologically(testInput())
	if result != "CABDFE" {
		t.Errorf("Order was incorrect, got: %s, want: %s.", result, "CABDFE")
	}
}

func TestTimedSortTopologially(t *testing.T) {
	result := TimedSortTopologically(testInput(), 2)
	if result != 256 {
		t.Errorf("time was incorrect, got: %d, want: %d.", result, 256)
	}
}

func testInput() []string {
	return []string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}
}
