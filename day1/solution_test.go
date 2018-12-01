package day1

import "testing"

func TestSumStrings0Elems(t *testing.T) {
	toSum := make([]string, 0)
	total := SumStrings(toSum)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestSumStrings1PositiveElem(t *testing.T) {
	toSum := []string{"+1"}
	total := SumStrings(toSum)
	if total != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 1)
	}
}

func TestSumStrings1NegativeElem(t *testing.T) {
	toSum := []string{"-1"}
	total := SumStrings(toSum)
	if total != -1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, -1)
	}
}

func TestSumStrings1Positive1Negative(t *testing.T) {
	toSum := []string{"-1", "+4"}
	total := SumStrings(toSum)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}
