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

func TestRepeatingSum2Elems(t *testing.T) {
	sequence := []string{"+1", "-1"}
	sum := RepeatingSum(sequence)
	if sum != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 0)
	}
}

func TestRepeatingSumInOneGo(t *testing.T) {
	sequence := []string{"+1", "+3", "-3"}
	sum := RepeatingSum(sequence)
	if sum != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 1)
	}
}

func TestRepeatingSumIn2Cycles(t *testing.T) {
	sequence := []string{"+3", "+3", "+4", "-2", "-4"}
	sum := RepeatingSum(sequence)
	if sum != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 10)
	}
}

func TestRepeatingSumInMultipleCycles(t *testing.T) {
	sequence := []string{"+7", "+7", "-2", "-7", "-4"}
	sum := RepeatingSum(sequence)
	if sum != 14 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 14)
	}
}
